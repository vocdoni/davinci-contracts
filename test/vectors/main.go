// Command vectors writes the forge test vectors for the zkVM ProcessRegistry.
//
//	genesis.json     genesis roots from go-sdk chain.NewState, registry-style pids
//	smt.json         roots of arbitrary arbo SHA-256 trees (64 levels, 8-byte keys)
//	transition.json  two settlement fixtures with real KZG openings, plus a results fixture
//	publics.json     the decoded registers of recorded_batch_snark.json
//
// Run it from this directory with `go run .`. The defaults match the forge tests
// (registry = CREATE(deployer, 0), chain id 1337); the flags regenerate the
// fixture for another deployment.
package main

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/big"
	"math/rand/v2"
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	arbo "github.com/vocdoni/arbo"
	"github.com/vocdoni/arbo/memdb"
	davinci "github.com/vocdoni/davinci-zkvm/go-sdk"
	"github.com/vocdoni/davinci-zkvm/go-sdk/chain"
	bjjgnark "github.com/vocdoni/davinci-zkvm/go-sdk/vocdoni/crypto/ecc/bjj_gnark"
	"github.com/vocdoni/davinci-zkvm/go-sdk/vocdoni/crypto/ecc/format"
	"github.com/vocdoni/davinci-zkvm/go-sdk/vocdoni/spec"
)

const anvilAccount0 = "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"

// fixtureCensusRoot is a Merkle census root (a BN254 field element).
var fixtureCensusRoot, _ = new(big.Int).SetString("2bc6f255d02b18329662a71d7d66c8ce06fe984607fd4fe26ef33ab93a78f683", 16)

func main() {
	out := flag.String("out", ".", "output directory")
	chainID := flag.Uint("chain-id", 1337, "chainID passed to the ProcessRegistry constructor")
	deployer := flag.String("deployer", anvilAccount0, "registry deployer")
	deployerNonce := flag.Uint64("deployer-nonce", 0, "deployer nonce when the registry is created")
	registryFlag := flag.String("registry", "", "registry address (overrides -deployer/-deployer-nonce)")
	creatorFlag := flag.String("creator", anvilAccount0, "process creator (organizer)")
	vkPath := flag.String("ballot-vk", "../../../davinci-circom/artifacts/ballot_proof_vkey.json", "snarkjs ballot proof VK")
	sdkGenesis := flag.String("sdk-genesis", "../../../davinci-zkvm/rust-sdk/testdata/genesis.json", "rust-sdk genesis vector to cross-check and include (skipped if missing)")
	recorded := flag.String("recorded", "recorded_batch_snark.json", "recorded batch snark.json to decode")
	writeBlobs := flag.Bool("blobs", false, "also write the raw fixture blobs as transition-<t>-blob-<b>.bin")
	flag.Parse()

	registry := common.HexToAddress(*registryFlag)
	if *registryFlag == "" {
		registry = crypto.CreateAddress(common.HexToAddress(*deployer), *deployerNonce)
	}
	creator := common.HexToAddress(*creatorFlag)
	prefix := pidPrefix(uint32(*chainID), registry)
	check(os.MkdirAll(*out, 0o755), "create output directory")

	vkJSON, err := os.ReadFile(*vkPath)
	check(err, "read ballot vk")
	vkLeaf, err := davinci.BallotVKLeaf(vkJSON)
	check(err, "ballot vk leaf")
	vkSum := sha256.Sum256(vkJSON)

	genesis := genesisVectors(prefix, creator, vkLeaf, *sdkGenesis)
	genesis.Cases = append(genesis.Cases, dynamicGenesisCases(prefix, creator, vkLeaf)...)
	genesis.BallotVKFileSHA256 = hex.EncodeToString(vkSum[:])
	writeJSON(*out, "genesis.json", genesis)
	writeJSON(*out, "smt.json", map[string][]smtCase{"cases": smtVectors()})

	fx, blobs := transitionFixture(uint32(*chainID), registry, creator, prefix, vkLeaf)
	writeJSON(*out, "transition.json", fx)
	if *writeBlobs {
		for t, bs := range blobs {
			for b, blob := range bs {
				name := filepath.Join(*out, fmt.Sprintf("transition-%d-blob-%d.bin", t, b))
				check(os.WriteFile(name, blob, 0o644), "write blob")
			}
		}
	}

	writeJSON(*out, "publics.json", decodeRecorded(*recorded))
}

// --- process ids -------------------------------------------------------------

// pidPrefix mirrors ProcessIdLib.getPrefix: the low 4 bytes of
// keccak256(uint32 chainID ‖ registry).
func pidPrefix(chainID uint32, registry common.Address) uint32 {
	var buf [24]byte
	binary.BigEndian.PutUint32(buf[:4], chainID)
	copy(buf[4:], registry.Bytes())
	h := crypto.Keccak256(buf[:])
	return binary.BigEndian.Uint32(h[28:])
}

// processID mirrors ProcessIdLib.computeProcessId: creator(20) ‖ prefix(4) ‖ nonce(7).
func processID(prefix uint32, creator common.Address, nonce uint64) *big.Int {
	pid := new(big.Int).SetBytes(creator.Bytes())
	pid.Lsh(pid, 88)
	pid.Or(pid, new(big.Int).Lsh(new(big.Int).SetUint64(uint64(prefix)), 56))
	pid.Or(pid, new(big.Int).SetUint64(nonce&(1<<56-1)))
	return pid
}

// --- genesis ---------------------------------------------------------------

type ballotModeJSON struct {
	NumFields    uint8  `json:"num_fields"`
	GroupSize    uint8  `json:"group_size"`
	UniqueValues bool   `json:"unique_values"`
	CostExponent uint8  `json:"cost_exponent"`
	MaxValue     uint64 `json:"max_value"`
	MinValue     uint64 `json:"min_value"`
	MaxValueSum  uint64 `json:"max_value_sum"`
	MinValueSum  uint64 `json:"min_value_sum"`
	Packed       string `json:"packed"`
}

func (m ballotModeJSON) spec() spec.BallotMode {
	return spec.BallotMode{
		NumFields: m.NumFields, GroupSize: m.GroupSize, UniqueValues: m.UniqueValues,
		CostExponent: m.CostExponent, MaxValue: m.MaxValue, MinValue: m.MinValue,
		MaxValueSum: m.MaxValueSum, MinValueSum: m.MinValueSum,
	}
}

type pointJSON struct {
	X string `json:"x"`
	Y string `json:"y"`
}

// genesisCase uses the rust-sdk testdata schema, with 0x-prefixed hex.
type genesisCase struct {
	Source       string            `json:"source"`
	ProcessID    string            `json:"process_id"`
	BallotMode   ballotModeJSON    `json:"ballot_mode"`
	EncKey       pointJSON         `json:"enc_key"`
	CensusOrigin uint64            `json:"census_origin"`
	BallotVKHash string            `json:"ballot_vk_hash"`
	Leaves       map[string]string `json:"leaves"`
	Root         string            `json:"root"`
}

type genesisFile struct {
	BallotVKHash       string        `json:"ballot_vk_hash"`
	BallotVKFileSHA256 string        `json:"ballot_vk_file_sha256"`
	IdentityAccLeaf    string        `json:"identity_acc_leaf"`
	Cases              []genesisCase `json:"cases"`
}

var fixtureModes = []ballotModeJSON{
	{NumFields: 2, GroupSize: 1, CostExponent: 1, MaxValue: 1, MaxValueSum: 2},
	{NumFields: 6, GroupSize: 0, CostExponent: 2, MaxValue: 16, MaxValueSum: 1280, MinValueSum: 5},
	{NumFields: 16, GroupSize: 16, UniqueValues: true, CostExponent: 255,
		MaxValue: 1<<48 - 1, MinValue: 7, MaxValueSum: 1<<63 - 1, MinValueSum: 1},
}

func genesisVectors(prefix uint32, creator common.Address, vkLeaf *big.Int, sdkPath string) genesisFile {
	f := genesisFile{
		BallotVKHash:    hex0x(be32(vkLeaf)),
		IdentityAccLeaf: hex0x(be32(identityAccLeaf())),
	}
	n := uint64(0)
	for _, mode := range fixtureModes {
		for _, origin := range []uint64{1, 4} {
			pid := processID(prefix, creator, 100+n)
			key := testKey(int64(1000 + 7919*n))
			f.Cases = append(f.Cases, genesisCaseFor("registry", pid, mode, key, origin, vkLeaf))
			n++
		}
	}

	// Recompute the rust-sdk vector with chain.NewState and include it, so the
	// Solidity test and the Rust SDK are pinned to the same Go reference.
	raw, err := os.ReadFile(sdkPath)
	if err != nil {
		log.Printf("skipping rust-sdk genesis cross-check: %v", err)
		return f
	}
	var sdk []struct {
		ProcessID    string         `json:"process_id"`
		BallotMode   ballotModeJSON `json:"ballot_mode"`
		EncKey       pointJSON      `json:"enc_key"`
		CensusOrigin uint64         `json:"census_origin"`
		BallotVKHash string         `json:"ballot_vk_hash"`
		Root         string         `json:"root"`
	}
	check(json.Unmarshal(raw, &sdk), "parse rust-sdk genesis")
	for i, c := range sdk {
		pid := mustBig(c.ProcessID, 16)
		vk := mustBig(c.BallotVKHash, 16)
		if vk.Cmp(vkLeaf) != 0 {
			log.Fatalf("rust-sdk case %d: ballot vk hash %s differs from %s", i, c.BallotVKHash, hex0x(be32(vkLeaf)))
		}
		x, y := mustBig(c.EncKey.X, 10), mustBig(c.EncKey.Y, 10)
		rx, ry := format.FromTEtoRTE(x, y)
		key := bjjgnark.New().SetPoint(rx, ry).(*bjjgnark.BJJ)
		gc := genesisCaseFor("rust-sdk", pid, c.BallotMode, key, c.CensusOrigin, vkLeaf)
		if gc.Root != "0x"+c.Root {
			log.Fatalf("rust-sdk case %d: root %s, go-sdk chain.NewState gives %s", i, c.Root, gc.Root)
		}
		f.Cases = append(f.Cases, gc)
	}
	log.Printf("rust-sdk genesis vector: %d cases agree with go-sdk chain.NewState", len(sdk))
	return f
}

// dynamicGenesisCases covers the dynamic Merkle origins (2 off-chain, 3 on-chain). They go
// after the rust-sdk cases so the indices of the earlier cases do not move.
func dynamicGenesisCases(prefix uint32, creator common.Address, vkLeaf *big.Int) []genesisCase {
	var out []genesisCase
	n := uint64(0)
	for _, mode := range fixtureModes {
		for _, origin := range []uint64{2, 3} {
			pid := processID(prefix, creator, 200+n)
			key := testKey(int64(2000 + 7919*n))
			out = append(out, genesisCaseFor("registry", pid, mode, key, origin, vkLeaf))
			n++
		}
	}
	return out
}

func genesisCaseFor(source string, pid *big.Int, mode ballotModeJSON, key *bjjgnark.BJJ, origin uint64, vkLeaf *big.Int) genesisCase {
	packed, err := mode.spec().Pack()
	check(err, "pack ballot mode")
	mode.Packed = packed.String()

	st, err := chain.NewState(chain.Config{
		ProcessID:    pid,
		BallotMode:   packed,
		EncKey:       key,
		CensusOrigin: origin,
		CensusRoot:   fixtureCensusRoot,
		BallotVKHash: vkLeaf,
	})
	check(err, "chain.NewState")

	tx, ty := teKey(key)
	var xy [64]byte
	tx.FillBytes(xy[:32])
	ty.FillBytes(xy[32:])
	encLeaf := sha256.Sum256(xy[:])

	le := func(v *big.Int) string { return hex0x(arbo.BigIntToBytes(32, v)) }
	return genesisCase{
		Source:       source,
		ProcessID:    hex0x(pid.FillBytes(make([]byte, 31))),
		BallotMode:   mode,
		EncKey:       pointJSON{X: tx.String(), Y: ty.String()},
		CensusOrigin: origin,
		BallotVKHash: hex0x(be32(vkLeaf)),
		Leaves: map[string]string{
			"0x00": le(pid),
			"0x02": le(packed),
			"0x03": le(new(big.Int).SetBytes(encLeaf[:])),
			"0x04": le(identityAccLeaf()),
			"0x06": le(new(big.Int).SetUint64(origin)),
			"0x07": le(vkLeaf),
		},
		Root: st.Root(),
	}
}

// identityAccLeaf is the genesis results leaf: sha256 over the identity
// accumulator [0,1,0,1,...] as 64 BE32 words, read BE.
func identityAccLeaf() *big.Int {
	h := sha256.New()
	var w [32]byte
	for i := 0; i < davinci.BallotFields; i++ {
		w[31] = byte(i % 2)
		h.Write(w[:])
	}
	return new(big.Int).SetBytes(h.Sum(nil))
}

// testKey is s·G on the gnark (RTE) curve.
func testKey(s int64) *bjjgnark.BJJ {
	p := bjjgnark.New().(*bjjgnark.BJJ)
	p.ScalarBaseMult(big.NewInt(s))
	return p
}

// teKey returns the circomlib TE coordinates of an RTE point.
func teKey(p *bjjgnark.BJJ) (*big.Int, *big.Int) {
	rx, ry := p.Point()
	return format.FromRTEtoTE(rx, ry)
}

// --- generic SMT -------------------------------------------------------------

type smtCase struct {
	Name   string   `json:"name"`
	Keys   []uint64 `json:"keys"`
	Values []string `json:"values"` // BE32 of the leaf integer
	Root   string   `json:"root"`
}

func smtVectors() []smtCase {
	rng := rand.New(rand.NewPCG(1, 2))
	randVal := func() *big.Int {
		var b [32]byte
		for i := range b {
			b[i] = byte(rng.Uint32())
		}
		return new(big.Int).SetBytes(b[:])
	}
	cases := []struct {
		name string
		keys []uint64
	}{
		{"empty", nil},
		{"single", []uint64{5}},
		{"two_split_at_bit0", []uint64{1, 2}},
		{"shared_low_40_bits", []uint64{0x10, 0x10 + 1<<40}},
		{"genesis_keys", []uint64{0, 2, 3, 4, 6, 7}},
	}
	var mixed []uint64
	seen := map[uint64]bool{}
	for len(mixed) < 40 {
		k := 0x10 + rng.Uint64N(1<<20)
		if len(mixed)%2 == 1 {
			k = 1<<63 | rng.Uint64()
		}
		if !seen[k] {
			seen[k] = true
			mixed = append(mixed, k)
		}
	}
	cases = append(cases, struct {
		name string
		keys []uint64
	}{"mixed_40", mixed})

	var out []smtCase
	for _, c := range cases {
		tree, err := arbo.NewTree(arbo.Config{
			Database: memdb.New(), MaxLevels: 64, HashFunction: arbo.HashFunctionSha256,
		})
		check(err, "arbo.NewTree")
		sc := smtCase{Name: c.name, Keys: c.keys, Values: []string{}}
		if sc.Keys == nil {
			sc.Keys = []uint64{}
		}
		for _, k := range c.keys {
			v := randVal()
			check(tree.Add(arbo.BigIntToBytes(8, new(big.Int).SetUint64(k)), arbo.BigIntToBytes(32, v)), "arbo add")
			sc.Values = append(sc.Values, hex0x(be32(v)))
		}
		root, err := tree.Root()
		check(err, "arbo root")
		sc.Root = hex0x(pad32(root))
		out = append(out, sc)
	}
	return out
}

// --- settlement fixture -------------------------------------------------------

type transitionJSON struct {
	PublicValues    string   `json:"public_values"`
	ProofBytes      string   `json:"proof_bytes"`
	RootBefore      string   `json:"root_before"`
	RootAfter       string   `json:"root_after"`
	Voters          uint64   `json:"voters"`
	Overwrites      uint64   `json:"overwrites"`
	OccupiedBefore  uint64   `json:"occupied_before"`
	NBlobs          uint64   `json:"n_blobs"`
	Commitments     []string `json:"commitments"`
	Ys              []string `json:"ys"`
	KzgProofs       []string `json:"kzg_proofs"`
	Zs              []string `json:"zs"`
	VersionedHashes []string `json:"versioned_hashes"`
	BlobsDigest     string   `json:"blobs_digest"`
}

type resultsJSON struct {
	PublicValues string   `json:"public_values"`
	ProofBytes   string   `json:"proof_bytes"`
	StateRoot    string   `json:"state_root"`
	Values       []uint64 `json:"values"`
}

type fixtureFile struct {
	ChainID      uint32           `json:"chain_id"`
	Registry     string           `json:"registry"`
	Creator      string           `json:"creator"`
	ProcessID    string           `json:"process_id"`
	BallotMode   ballotModeJSON   `json:"ballot_mode"`
	EncKey       pointJSON        `json:"enc_key"`
	CensusOrigin uint64           `json:"census_origin"`
	CensusRoot   string           `json:"census_root"`
	BallotVKHash string           `json:"ballot_vk_hash"`
	GenesisRoot  string           `json:"genesis_root"`
	Transitions  []transitionJSON `json:"transitions"`
	Results      resultsJSON      `json:"results"`
	Dynamic      []dynamicJSON    `json:"dynamic"`
}

// dynamicJSON is the first transition of the same process created with a dynamic
// census origin: only the origin leaf of the genesis differs, which moves every root
// and the blob evaluation points.
type dynamicJSON struct {
	CensusOrigin uint64         `json:"census_origin"`
	GenesisRoot  string         `json:"genesis_root"`
	Transition   transitionJSON `json:"transition"`
}

// transitionFixture builds two transitions for the organizer's first process:
// three fresh votes on one blob, then 820 votes (three overwrites) on two
// blobs. Roots after are placeholders: the forge tests use a mock verifier.
func transitionFixture(chainID uint32, registry, creator common.Address, prefix uint32, vkLeaf *big.Int) (fixtureFile, [][][]byte) {
	mode := fixtureModes[0]
	nf := int(mode.NumFields)
	pid := processID(prefix, creator, 0)
	key := testKey(424242)
	gc := genesisCaseFor("fixture", pid, mode, key, 1, vkLeaf)
	genesis := mustBytes(gc.Root)

	fx := fixtureFile{
		ChainID: chainID, Registry: registry.Hex(), Creator: creator.Hex(),
		ProcessID: gc.ProcessID, BallotMode: gc.BallotMode, EncKey: gc.EncKey,
		CensusOrigin: 1, CensusRoot: hex0x(be32(fixtureCensusRoot)),
		BallotVKHash: gc.BallotVKHash, GenesisRoot: gc.Root,
	}

	pidBE := [32]byte(be32(pid))
	pt := pointGen{}
	var allBlobs [][][]byte

	type batch struct {
		voteIDs []uint64
		slots   []uint64
		w, occ  uint64
	}
	b1 := batch{voteIDs: []uint64{1<<63 | 11, 1<<63 | 7, 1<<63 | 23}, slots: []uint64{0x18, 0x19, 0x1a}}
	b2 := batch{w: 3, occ: 3}
	for j := uint64(0); j < 820; j++ {
		b2.voteIDs = append(b2.voteIDs, 1<<63|(1000+j))
		if j < 3 {
			b2.slots = append(b2.slots, 0x18+j)
		} else {
			b2.slots = append(b2.slots, 0x10+(1<<10|j))
		}
	}

	build := func(pt *pointGen, root []byte, b batch, t int) (transitionJSON, [][]byte, []byte) {
		updates := make([]davinci.SlotUpdate, len(b.slots))
		for i, s := range b.slots {
			updates[i] = davinci.SlotUpdate{Key: s, Ballot: pt.ballot(nf)}
		}
		rootBE := [32]byte(reverse(root))
		tb, err := davinci.BuildTransitionBlobs(nf, pidBE, rootBE, b.voteIDs, updates, pt.ballot(nf))
		check(err, "BuildTransitionBlobs")

		after := sha256.Sum256([]byte(fmt.Sprintf("davinci-contracts fixture root %d", t+1)))
		var regs [64]uint32
		regs[0] = 1
		setReg32(&regs, 2, root)
		setReg32(&regs, 10, after[:])
		regs[18] = uint32(len(b.voteIDs))
		regs[19] = uint32(b.w)
		setReg32(&regs, 20, reverse(be32(fixtureCensusRoot)))
		setReg32(&regs, 28, tb.Digest[:])
		regs[36] = uint32(len(tb.Blobs))
		regs[40], regs[41] = 1, 1
		regs[42] = uint32(b.occ)
		regs[43], regs[44] = uint32(len(b.voteIDs)), 3

		tj := transitionJSON{
			PublicValues:   hex0x(publicValues(regs)),
			ProofBytes:     hex0x(make([]byte, 768)),
			RootBefore:     hex0x(root),
			RootAfter:      hex0x(after[:]),
			Voters:         uint64(len(b.voteIDs)),
			Overwrites:     b.w,
			OccupiedBefore: b.occ,
			NBlobs:         uint64(len(tb.Blobs)),
			BlobsDigest:    hex0x(tb.Digest[:]),
		}
		var blobs [][]byte
		for i := range tb.Blobs {
			tj.Commitments = append(tj.Commitments, hex0x(tb.Commitments[i][:]))
			tj.Ys = append(tj.Ys, hex0x(tb.Ys[i][:]))
			tj.KzgProofs = append(tj.KzgProofs, hex0x(tb.Proofs[i][:]))
			tj.Zs = append(tj.Zs, hex0x(tb.Zs[i][:]))
			tj.VersionedHashes = append(tj.VersionedHashes, hex0x(tb.VersionedHashes[i][:]))
			blobs = append(blobs, tb.Blobs[i][:])
		}
		return tj, blobs, after[:]
	}

	root := genesis
	for t, b := range []batch{b1, b2} {
		tj, blobs, after := build(&pt, root, b, t)
		fx.Transitions = append(fx.Transitions, tj)
		allBlobs = append(allBlobs, blobs)
		root = after
	}

	// The first batch again, for the same process created with origin 2 and with origin 3.
	for _, origin := range []uint64{2, 3} {
		dg := genesisCaseFor("fixture", pid, mode, key, origin, vkLeaf)
		tj, _, _ := build(&pointGen{}, mustBytes(dg.Root), b1, 0)
		fx.Dynamic = append(fx.Dynamic, dynamicJSON{CensusOrigin: origin, GenesisRoot: dg.Root, Transition: tj})
	}

	// Results over the final root: field 1 needs the high register.
	values := make([]uint64, davinci.NumFields)
	values[0], values[1] = 7, 1<<32+5
	var regs [64]uint32
	regs[0] = 1
	setReg32(&regs, 2, root)
	for i, v := range values {
		regs[10+2*i] = uint32(v)
		regs[11+2*i] = uint32(v >> 32)
	}
	regs[42] = 0xFFFFFFFF
	fx.Results = resultsJSON{
		PublicValues: hex0x(publicValues(regs)),
		ProofBytes:   hex0x(make([]byte, 768)),
		StateRoot:    hex0x(root),
		Values:       values[:nf],
	}
	return fx, allBlobs
}

// pointGen hands out distinct on-curve TE points for ballot coordinates.
type pointGen struct{ n int64 }

func (g *pointGen) ballot(nf int) []string {
	out := make([]string, davinci.BallotFields)
	for f := 0; f < davinci.NumFields; f++ {
		for c := 0; c < 2; c++ {
			x, y := "0x00", "0x01"
			if f < nf {
				g.n++
				px, py := teKey(testKey(g.n*7 + 3))
				x, y = hex0x(be32(px)), hex0x(be32(py))
			}
			out[4*f+2*c], out[4*f+2*c+1] = x, y
		}
	}
	return out
}

// publicValues encodes 64 u32 registers as 8-byte LE words (ZisK 1.3 snark inputs).
func publicValues(regs [64]uint32) []byte {
	out := make([]byte, 512)
	for k, r := range regs {
		binary.LittleEndian.PutUint64(out[8*k:], uint64(r))
	}
	return out
}

// setReg32 stores 32 bytes as 8 LE u32 limbs starting at register base.
func setReg32(regs *[64]uint32, base int, b []byte) {
	for j := 0; j < 8; j++ {
		regs[base+j] = binary.LittleEndian.Uint32(b[4*j : 4*j+4])
	}
}

// --- recorded publics ---------------------------------------------------------

type publicsFile struct {
	ProgramVK      string   `json:"program_vk"`
	PublicValues   string   `json:"public_values"`
	Words          []uint64 `json:"words"`
	OK             uint64   `json:"ok"`
	FailMask       uint64   `json:"fail_mask"`
	RootBefore     string   `json:"root_before"`
	RootAfter      string   `json:"root_after"`
	Voters         uint64   `json:"voters"`
	Overwrites     uint64   `json:"overwrites"`
	CensusRootLE   string   `json:"census_root_le"`
	CensusRoot     string   `json:"census_root"`
	BlobsDigest    string   `json:"blobs_digest"`
	NBlobs         uint64   `json:"n_blobs"`
	OccupiedBefore uint64   `json:"occupied_before"`
}

func decodeRecorded(path string) publicsFile {
	raw, err := os.ReadFile(path)
	check(err, "read recorded snark")
	var s struct {
		ProgramVK    string `json:"program_vk"`
		PublicValues string `json:"public_values"`
	}
	check(json.Unmarshal(raw, &s), "parse recorded snark")
	pv := mustBytes(s.PublicValues)
	if len(pv) != 512 {
		log.Fatalf("recorded public_values: %d bytes, want 512", len(pv))
	}
	word := func(k int) uint64 { return binary.LittleEndian.Uint64(pv[8*k:]) }
	reg32 := func(k int) []byte {
		out := make([]byte, 0, 32)
		for j := 0; j < 8; j++ {
			out = append(out, pv[8*(k+j):8*(k+j)+4]...)
		}
		return out
	}
	f := publicsFile{
		ProgramVK: s.ProgramVK, PublicValues: s.PublicValues,
		OK: word(0), FailMask: word(1),
		RootBefore: hex0x(reg32(2)), RootAfter: hex0x(reg32(10)),
		Voters: word(18), Overwrites: word(19),
		CensusRootLE: hex0x(reg32(20)), CensusRoot: hex0x(reverse(reg32(20))),
		BlobsDigest: hex0x(reg32(28)), NBlobs: word(36), OccupiedBefore: word(42),
	}
	for k := 0; k < 64; k++ {
		f.Words = append(f.Words, word(k))
	}
	return f
}

// --- helpers ------------------------------------------------------------------

func be32(v *big.Int) []byte { return v.FillBytes(make([]byte, 32)) }

func pad32(b []byte) []byte {
	out := make([]byte, 32)
	copy(out[32-len(b):], b)
	return out
}

func reverse(b []byte) []byte {
	out := make([]byte, len(b))
	for i := range b {
		out[len(b)-1-i] = b[i]
	}
	return out
}

func hex0x(b []byte) string { return "0x" + hex.EncodeToString(b) }

func mustBytes(s string) []byte {
	if len(s) >= 2 && s[:2] == "0x" {
		s = s[2:]
	}
	b, err := hex.DecodeString(s)
	check(err, "hex")
	return b
}

func mustBig(s string, base int) *big.Int {
	if len(s) >= 2 && s[:2] == "0x" {
		s = s[2:]
	}
	v, ok := new(big.Int).SetString(s, base)
	if !ok {
		log.Fatalf("bad integer %q", s)
	}
	return v
}

func writeJSON(dir, name string, v any) {
	b, err := json.MarshalIndent(v, "", "  ")
	check(err, "marshal "+name)
	check(os.WriteFile(filepath.Join(dir, name), append(b, '\n'), 0o644), "write "+name)
}

func check(err error, what string) {
	if err != nil {
		log.Fatalf("%s: %v", what, err)
	}
}
