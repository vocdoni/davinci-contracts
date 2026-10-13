package main

import (
	"math/big"
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/trie"
	"github.com/ethereum/go-ethereum/triedb"
	"github.com/holiman/uint256"
)

// --- L1-blob mode: one transition in several Ethereum transactions ------------

// l1TxJSON and l1ProofJSON are the shape of test/vectors/l1blob/*.json.
type l1TxJSON struct {
	HeaderIndex int      `json:"header_index"`
	Key         string   `json:"key"`
	Proof       []string `json:"proof"`
}

type l1ProofJSON struct {
	ExpectedTxBlocks []uint64   `json:"expected_tx_blocks"`
	ExpectedVhs      []string   `json:"expected_vhs"`
	HeaderHashes     []string   `json:"header_hashes"`
	HeaderNumbers    []uint64   `json:"header_numbers"`
	Headers          []string   `json:"headers"`
	Name             string     `json:"name"`
	TxType           int        `json:"tx_type"`
	Txs              []l1TxJSON `json:"txs"`
}

// l1Post is one blob transaction of the proof: the versioned hashes it carries
// and how many blocks below the top header it sits.
type l1Post struct {
	vhs   []common.Hash
	depth int
}

// keyProof records the trie nodes Trie.Prove writes, root first.
type keyProof struct{ nodes [][]byte }

func (k *keyProof) Put(_, v []byte) error {
	k.nodes = append(k.nodes, common.CopyBytes(v))
	return nil
}

func (k *keyProof) Delete([]byte) error { return nil }

// l1Envelope is a type-3 transaction from a fixed sender to itself carrying vhs
// (unsigned: the verifier reads only the versioned hashes).
func l1Envelope(nonce uint64, vhs []common.Hash) []byte {
	to := common.HexToAddress("0x00000000000000000000000000000000000000b1")
	tx := types.NewTx(&types.BlobTx{
		ChainID:    uint256.NewInt(1),
		Nonce:      nonce,
		GasTipCap:  uint256.NewInt(1_000_000),
		GasFeeCap:  uint256.NewInt(2_000_000_000),
		Gas:        21_000,
		To:         to,
		Value:      uint256.NewInt(0),
		BlobFeeCap: uint256.NewInt(2),
		BlobHashes: vhs,
		V:          uint256.NewInt(0),
		R:          uint256.NewInt(1),
		S:          uint256.NewInt(1),
	})
	b, err := tx.MarshalBinary()
	check(err, "blob envelope")
	return b
}

// l1Block is a header over envs, with the transactions-trie proof of each one.
// It carries every field of a Prague header, so it is as long as a mainnet one.
func l1Block(parent common.Hash, number uint64, envs [][]byte) (*types.Header, [][][]byte) {
	tr := trie.NewEmpty(triedb.NewDatabase(rawdb.NewMemoryDatabase(), nil))
	for i, e := range envs {
		k, err := rlp.EncodeToBytes(uint64(i))
		check(err, "tx key")
		check(tr.Update(k, e), "trie update")
	}
	proofs := make([][][]byte, len(envs))
	for i := range envs {
		k, _ := rlp.EncodeToBytes(uint64(i))
		var p keyProof
		check(tr.Prove(k, &p), "trie proof")
		proofs[i] = p.nodes
	}
	withdrawals, requests := types.EmptyWithdrawalsHash, types.EmptyRequestsHash
	beacon := crypto.Keccak256Hash([]byte("beacon"), new(big.Int).SetUint64(number).Bytes())
	blobGas, excess := uint64(131_072*len(envs)), uint64(0)
	h := &types.Header{
		ParentHash:       parent,
		UncleHash:        types.EmptyUncleHash,
		Coinbase:         common.HexToAddress("0x00000000000000000000000000000000000000c0"),
		Root:             crypto.Keccak256Hash([]byte("state"), new(big.Int).SetUint64(number).Bytes()),
		TxHash:           tr.Hash(),
		ReceiptHash:      crypto.Keccak256Hash([]byte("receipts"), new(big.Int).SetUint64(number).Bytes()),
		Difficulty:       big.NewInt(0),
		Number:           new(big.Int).SetUint64(number),
		GasLimit:         45_000_000,
		GasUsed:          21_000 * uint64(len(envs)),
		Time:             1_760_000_000 + 12*number,
		Extra:            []byte("davinci l1blob vector"),
		BaseFee:          big.NewInt(1_000_000_000),
		WithdrawalsHash:  &withdrawals,
		BlobGasUsed:      &blobGas,
		ExcessBlobGas:    &excess,
		ParentBeaconRoot: &beacon,
		RequestsHash:     &requests,
	}
	return h, proofs
}

// l1ProofFixture builds `top - bottom + 1` chained headers, newest first, with
// each post's transaction (after one filler blob transaction in its block) at
// its depth, and the proof listing the posts in the given order.
func l1ProofFixture(name string, bottom, top uint64, posts []l1Post) l1ProofJSON {
	type placed struct{ depth, index int }
	blocks := make(map[uint64][][]byte)
	where := make([]placed, len(posts))
	for i, p := range posts {
		n := top - uint64(p.depth)
		if blocks[n] == nil {
			// A filler, so the trie has a branch and the key is not 0x80.
			blocks[n] = [][]byte{l1Envelope(100+n, []common.Hash{{0x01, 0xf1}})}
		}
		where[i] = placed{p.depth, len(blocks[n])}
		blocks[n] = append(blocks[n], l1Envelope(uint64(i), p.vhs))
	}
	var (
		parent  common.Hash
		headers = make([]*types.Header, 0, top-bottom+1)
		tproofs = make(map[uint64][][][]byte)
	)
	for n := bottom; n <= top; n++ {
		h, proofs := l1Block(parent, n, blocks[n])
		headers = append(headers, h)
		tproofs[n] = proofs
		parent = h.Hash()
	}
	out := l1ProofJSON{Name: name, TxType: 3}
	for i := len(headers) - 1; i >= 0; i-- {
		h := headers[i]
		enc, err := rlp.EncodeToBytes(h)
		check(err, "header rlp")
		if crypto.Keccak256Hash(enc) != h.Hash() {
			panic("header hash")
		}
		out.Headers = append(out.Headers, hex0x(enc))
		out.HeaderHashes = append(out.HeaderHashes, h.Hash().Hex())
		out.HeaderNumbers = append(out.HeaderNumbers, h.Number.Uint64())
	}
	for i, p := range posts {
		n := top - uint64(where[i].depth)
		k, _ := rlp.EncodeToBytes(uint64(where[i].index))
		tx := l1TxJSON{HeaderIndex: where[i].depth, Key: hex0x(k)}
		for _, node := range tproofs[n][where[i].index] {
			tx.Proof = append(tx.Proof, hex0x(node))
		}
		out.Txs = append(out.Txs, tx)
		for _, vh := range p.vhs {
			out.ExpectedVhs = append(out.ExpectedVhs, vh.Hex())
			out.ExpectedTxBlocks = append(out.ExpectedTxBlocks, n)
		}
	}
	return out
}

// syntheticVhs are n versioned hashes (0x01 || keccak(tag, i)[1:]).
func syntheticVhs(tag string, n int) []common.Hash {
	out := make([]common.Hash, n)
	for i := range out {
		out[i] = crypto.Keccak256Hash([]byte(tag), []byte{byte(i)})
		out[i][0] = 0x01
	}
	return out
}

// l1BlobFixtures writes l1blob/two_tx.json and two_tx_reposted.json: the
// fixture's second transition (two blobs) posted in two transactions, two
// blocks apart, under four headers. In two_tx the first blob's transaction is
// the older; in two_tx_reposted it was posted again and is the newer, while
// the proof still lists it first. one_tx.json proves the first transition's
// blob, so a registry can settle both in order. worst_case.json is the
// largest proof a node sends, for the size and gas bound.
func l1BlobFixtures(dir string, fx fixtureFile) {
	vhs := func(t int) []common.Hash {
		var out []common.Hash
		for _, s := range fx.Transitions[t].VersionedHashes {
			out = append(out, common.HexToHash(s))
		}
		return out
	}
	t0, t1 := vhs(0), vhs(1)
	if len(t0) != 1 || len(t1) != 2 {
		panic("the fixture transitions are no longer one and two blobs")
	}
	dir = filepath.Join(dir, "l1blob")
	check(os.MkdirAll(dir, 0o755), "create l1blob directory")
	writeJSON(dir, "one_tx.json", l1ProofFixture("one_tx", 1000, 1002,
		[]l1Post{{vhs: t0, depth: 2}}))
	writeJSON(dir, "two_tx.json", l1ProofFixture("two_tx", 2000, 2003,
		[]l1Post{{vhs: t1[:1], depth: 3}, {vhs: t1[1:], depth: 1}}))
	writeJSON(dir, "two_tx_reposted.json", l1ProofFixture("two_tx_reposted", 2000, 2003,
		[]l1Post{{vhs: t1[:1], depth: 1}, {vhs: t1[1:], depth: 3}}))
	// The largest proof a node sends: 64 headers, three transactions of a
	// 13-blob transition (6 + 6 + 1), the oldest at the bottom.
	vh := syntheticVhs("worst case", 13)
	writeJSON(dir, "worst_case.json", l1ProofFixture("worst_case", 3000, 3063,
		[]l1Post{{vhs: vh[:6], depth: 63}, {vhs: vh[6:12], depth: 40}, {vhs: vh[12:], depth: 2}}))
}
