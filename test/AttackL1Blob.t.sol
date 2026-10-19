// SPDX-License-Identifier: AGPL-3.0-only
pragma solidity ^0.8.28;

import {Test} from "forge-std/Test.sol";
import {stdJson} from "forge-std/StdJson.sol";
import {IL1BlobVerifier} from "../src/interfaces/IL1BlobVerifier.sol";
import {L1BlobVerifierHarness} from "./L1BlobVerifier.t.sol";
import {MockL1Block} from "./mocks/MockL1Block.sol";

/// @dev Mutation fuzzer over the honest L1 blob proofs (test/vectors/l1blob). With OP's origin
///      at any authenticated header of the proof, every mutation of the headers, the
///      transaction proofs, the keys or the header indices must either revert or return
///      exactly the honest versioned hashes and blocks: nothing may make the verifier return a
///      hash that is not in the proven transactions.
contract AttackL1BlobTest is Test {
    using stdJson for string;

    struct Fx {
        IL1BlobVerifier.L1BlobProof p;
        bytes32[] hashes;
        uint256[] numbers;
        bytes32[] vhs;
        uint256[] blocks;
    }

    MockL1Block internal l1;
    L1BlobVerifierHarness internal v;

    function setUp() public {
        l1 = new MockL1Block();
        v = new L1BlobVerifierHarness(address(l1));
    }

    function _name(uint256 i) internal pure returns (string memory) {
        string[6] memory names = ["recent", "one_tx", "two_tx", "two_tx_reposted", "index128plus", "cancun"];
        return names[i % names.length];
    }

    function _load(string memory name) internal view returns (Fx memory fx) {
        string memory j = vm.readFile(string.concat("test/vectors/l1blob/", name, ".json"));
        fx.p.headers = j.readBytesArray(".headers");
        fx.hashes = j.readBytes32Array(".header_hashes");
        fx.numbers = j.readUintArray(".header_numbers");
        fx.vhs = j.readBytes32Array(".expected_vhs");
        fx.blocks = j.readUintArray(".expected_tx_blocks");
        uint256 nt = abi.decode(vm.parseJson(j, ".txs"), (bytes[])).length;
        fx.p.txs = new IL1BlobVerifier.L1TxProof[](nt);
        for (uint256 i = 0; i < nt; ++i) {
            string memory b = string.concat(".txs[", vm.toString(i), "]");
            fx.p.txs[i].headerIndex = j.readUint(string.concat(b, ".header_index"));
            fx.p.txs[i].key = j.readBytes(string.concat(b, ".key"));
            fx.p.txs[i].proof = j.readBytesArray(string.concat(b, ".proof"));
        }
    }

    /// @dev The lowest header index any transaction sits at: OP's origin may be anywhere above.
    function _minTxHeader(Fx memory fx) internal pure returns (uint256 m) {
        m = type(uint256).max;
        for (uint256 t = 0; t < fx.p.txs.length; ++t) {
            if (fx.p.txs[t].headerIndex < m) m = fx.p.txs[t].headerIndex;
        }
    }

    function _copyBytes(bytes memory b) internal pure returns (bytes memory) {
        return bytes.concat(b);
    }

    function _without(bytes[] memory a, uint256 i) internal pure returns (bytes[] memory r) {
        r = new bytes[](a.length - 1);
        for (uint256 k = 0; k < r.length; ++k) {
            r[k] = a[k < i ? k : k + 1];
        }
    }

    function _withCopy(bytes[] memory a, uint256 i) internal pure returns (bytes[] memory r) {
        r = new bytes[](a.length + 1);
        for (uint256 k = 0; k < r.length; ++k) {
            r[k] = a[k <= i ? k : k - 1];
        }
    }

    function _flip(bytes memory b, uint256 pos, uint8 x) internal pure returns (bytes memory r) {
        r = _copyBytes(b);
        if (r.length == 0) return hex"00";
        pos = pos % r.length;
        r[pos] = bytes1(uint8(r[pos]) ^ (x == 0 ? 1 : x));
    }

    function _truncate(bytes memory b, uint256 n) internal pure returns (bytes memory r) {
        n = b.length == 0 ? 0 : 1 + (n % b.length);
        r = new bytes(b.length - n);
        for (uint256 i = 0; i < r.length; ++i) {
            r[i] = b[i];
        }
    }

    function _mutate(Fx memory fx, Fx memory other, uint256 kind, uint256 a, uint256 b, uint8 x) internal pure {
        IL1BlobVerifier.L1BlobProof memory p = fx.p;
        uint256 nh = p.headers.length;
        uint256 t = a % p.txs.length;
        bytes[] memory proof = p.txs[t].proof;
        uint256 k = b % proof.length;
        if (kind == 0) {
            p.headers[a % nh] = _flip(p.headers[a % nh], b, x);
        } else if (kind == 1 && nh > 1) {
            p.headers = _without(p.headers, a % nh);
        } else if (kind == 2) {
            (p.headers[a % nh], p.headers[b % nh]) = (p.headers[b % nh], p.headers[a % nh]);
        } else if (kind == 3) {
            p.headers = _withCopy(p.headers, a % nh);
        } else if (kind == 4) {
            proof[k] = _flip(proof[k], uint256(keccak256(abi.encode(b))), x);
        } else if (kind == 5 && proof.length > 1) {
            p.txs[t].proof = _without(proof, k);
        } else if (kind == 6) {
            p.txs[t].proof = _withCopy(proof, k);
        } else if (kind == 7) {
            uint256 k2 = uint256(keccak256(abi.encode(b))) % proof.length;
            (proof[k], proof[k2]) = (proof[k2], proof[k]);
        } else if (kind == 8) {
            // A node of another transaction proof (another fixture, or the other tx here).
            bytes[] memory src = other.p.txs[b % other.p.txs.length].proof;
            proof[k] = src[uint256(keccak256(abi.encode(a))) % src.length];
        } else if (kind == 9) {
            p.txs[t].headerIndex = b % (nh + 2);
        } else if (kind == 10) {
            uint256 m = b % 3;
            if (m == 0) p.txs[t].key = _flip(p.txs[t].key, a, x);
            else if (m == 1) p.txs[t].key = bytes.concat(p.txs[t].key, bytes1(x));
            else p.txs[t].key = _truncate(p.txs[t].key, a);
        } else if (kind == 11) {
            p.headers[a % nh] = _truncate(p.headers[a % nh], b);
        } else if (kind == 12) {
            p.headers[a % nh] = bytes.concat(p.headers[a % nh], bytes1(x));
        } else if (kind == 13) {
            proof[k] = b % 2 == 0 ? _truncate(proof[k], a) : bytes.concat(proof[k], bytes1(x));
        } else if (kind == 14) {
            p.headers[a % nh] = other.p.headers[b % other.p.headers.length];
        } else if (kind == 15) {
            // Another transaction's whole proof and key under this transaction's header.
            IL1BlobVerifier.L1TxProof memory o = other.p.txs[b % other.p.txs.length];
            p.txs[t].key = o.key;
            p.txs[t].proof = o.proof;
        }
    }

    /// The core property, over every fixture, origin position and mutation kind.
    function testFuzz_MutatedProofNeverForges(
        uint256 fxi,
        uint256 oxi,
        uint256 anchor,
        uint256 kind,
        uint256 a,
        uint256 b,
        uint8 x
    ) public {
        Fx memory fx = _load(_name(fxi));
        Fx memory other = _load(_name(oxi));
        anchor = bound(anchor, 0, _minTxHeader(fx));
        l1.set(uint64(fx.numbers[anchor]), fx.hashes[anchor]);
        (bytes32[] memory hv, uint256[] memory hb) = v.verify(fx.p);
        assertEq(hv, fx.vhs);

        kind = bound(kind, 0, 15);
        _mutate(fx, other, kind, a, b, x);
        try v.verify(fx.p) returns (bytes32[] memory vhs, uint256[] memory blocks) {
            if (kind == 8 || kind == 14 || kind == 15) {
                // Parts of another proof: whatever is returned must still be real.
                _assertGenuine(vhs, blocks, hv, hb, other);
            } else {
                assertEq(vhs, hv, "mutated proof changed the versioned hashes");
                assertEq(blocks, hb, "mutated proof changed the blocks");
            }
        } catch {}
    }

    /// Every returned (hash, block) is one of a fixture's honest pairs.
    function _assertGenuine(
        bytes32[] memory vhs,
        uint256[] memory blocks,
        bytes32[] memory hv,
        uint256[] memory hb,
        Fx memory other
    ) internal pure {
        assertEq(vhs.length, blocks.length);
        for (uint256 i = 0; i < vhs.length; ++i) {
            bool found;
            for (uint256 j = 0; j < hv.length && !found; ++j) {
                found = vhs[i] == hv[j] && blocks[i] == hb[j];
            }
            for (uint256 j = 0; j < other.vhs.length && !found; ++j) {
                found = vhs[i] == other.vhs[j] && blocks[i] == other.blocks[j % other.blocks.length];
            }
            assertTrue(found, "verifier returned a hash no proven transaction carries");
        }
    }

    /// Two mutations at once.
    function testFuzz_DoubleMutation(uint256 fxi, uint256 kind1, uint256 kind2, uint256 a, uint256 b, uint8 x) public {
        Fx memory fx = _load(_name(fxi));
        Fx memory other = _load(_name(fxi % 6 + 1));
        l1.set(uint64(fx.numbers[0]), fx.hashes[0]);
        (bytes32[] memory hv, uint256[] memory hb) = v.verify(fx.p);
        _mutate(fx, other, bound(kind1, 0, 15), a, b, x);
        _mutate(fx, other, bound(kind2, 0, 15), uint256(keccak256(abi.encode(a))), uint256(keccak256(abi.encode(b))), x);
        try v.verify(fx.p) returns (bytes32[] memory vhs, uint256[] memory blocks) {
            _assertGenuine(vhs, blocks, hv, hb, other);
        } catch {}
    }

    /// Arbitrary proof nodes against a real transactions root never resolve.
    function testFuzz_RandomProofNodes(bytes[] memory nodes, bytes memory key) public {
        Fx memory fx = _load("recent");
        l1.set(uint64(fx.numbers[0]), fx.hashes[0]);
        if (key.length == 0) key = hex"80";
        fx.p.txs[0].key = key;
        fx.p.txs[0].proof = nodes;
        vm.expectRevert();
        v.verify(fx.p);
    }

    /// Headers above the anchor are unauthenticated: a forged child of the origin (any content
    /// with the right parentHash) is accepted as padding but no transaction may point at it.
    function test_ForgedChildOfTheOriginIsUnusable() public {
        Fx memory fx = _load("recent");
        l1.set(uint64(fx.numbers[0]), fx.hashes[0]);
        // A forged header: parentHash = the real top header, transactionsRoot = keccak of the
        // proof's root node (the real one), number 1 above the top.
        bytes memory forged = _forgedHeader(fx.hashes[0], keccak256(fx.p.txs[0].proof[0]), fx.numbers[0] + 1);
        bytes[] memory hs = new bytes[](fx.p.headers.length + 1);
        hs[0] = forged;
        for (uint256 i = 0; i < fx.p.headers.length; ++i) {
            hs[i + 1] = fx.p.headers[i];
        }
        fx.p.headers = hs;
        fx.p.txs[0].headerIndex += 1;
        (bytes32[] memory vhs,) = v.verify(fx.p);
        assertEq(vhs, fx.vhs);

        // Pointing the transaction at the forged header is refused even though its root
        // would match the proof.
        fx.p.txs[0].headerIndex = 0;
        vm.expectRevert(abi.encodeWithSelector(IL1BlobVerifier.TxAboveOrigin.selector, 0));
        v.verify(fx.p);
    }

    /// The same, when the anchor is a recorded origin and the current origin moved on.
    function test_RecordedAnchorDoesNotAuthenticateChildren() public {
        Fx memory fx = _load("recent");
        l1.set(uint64(fx.numbers[1]), fx.hashes[1]);
        v.record();
        l1.set(uint64(fx.numbers[0]) + 5, keccak256("later origin"));
        // headers[0] is the real child of the recorded block but not authenticated by anything.
        fx.p.txs[0].headerIndex = 0;
        vm.expectRevert(abi.encodeWithSelector(IL1BlobVerifier.TxAboveOrigin.selector, 0));
        v.verify(fx.p);
    }

    /// A header whose number field lies about its height cannot be the recorded anchor: the
    /// record is keyed by the real number.
    function test_RecordKeyedByNumber() public {
        Fx memory fx = _load("recent");
        l1.set(uint64(fx.numbers[1]), fx.hashes[1]);
        v.record();
        assertEq(v.recorded(fx.numbers[1]), fx.hashes[1]);
        assertEq(v.recorded(fx.numbers[0]), bytes32(0));
        // A later record of another origin does not disturb it.
        l1.set(uint64(fx.numbers[0]), fx.hashes[0]);
        v.record();
        assertEq(v.recorded(fx.numbers[1]), fx.hashes[1]);
    }

    /// The age bound counts from the current origin even when a recorded anchor is used.
    function testFuzz_AgeBoundWithRecordedAnchor(uint256 later) public {
        Fx memory fx = _load("recent");
        uint256 txBlock = fx.blocks[0];
        l1.set(uint64(fx.numbers[2]), fx.hashes[2]);
        v.record();
        later = bound(later, fx.numbers[0], txBlock + 3 * v.MAX_AGE());
        l1.set(uint64(later), keccak256(abi.encode("origin", later)));
        // Proof anchored at the recorded header (drop the two above it).
        bytes[] memory hs = new bytes[](fx.p.headers.length - 2);
        for (uint256 i = 0; i < hs.length; ++i) {
            hs[i] = fx.p.headers[i + 2];
        }
        fx.p.headers = hs;
        fx.p.txs[0].headerIndex -= 2;
        if (later - txBlock > v.MAX_AGE()) {
            vm.expectRevert(abi.encodeWithSelector(IL1BlobVerifier.TxTooOld.selector, 0));
            v.verify(fx.p);
        } else {
            (bytes32[] memory vhs,) = v.verify(fx.p);
            assertEq(vhs, fx.vhs);
        }
    }

    /// A non-blob transaction (type 2) at index 0 and index 257 of real mainnet blocks.
    function test_TypeTwoTransactionsRefused() public {
        string[2] memory names = ["index0", "big_block"];
        for (uint256 i = 0; i < 2; ++i) {
            Fx memory fx = _load(names[i]);
            l1.set(uint64(fx.numbers[0]), fx.hashes[0]);
            vm.expectRevert(abi.encodeWithSelector(IL1BlobVerifier.NotBlobTx.selector, 0));
            v.verify(fx.p);
        }
    }

    /// The type-2 transaction's proof under the blob transaction's key path, and the reverse:
    /// a proof only opens the key it was built for.
    function test_KeyAndProofMustMatch() public {
        Fx memory a = _load("index0");
        Fx memory b = _load("big_block");
        l1.set(uint64(b.numbers[0]), b.hashes[0]);
        b.p.txs[0].key = hex"80";
        vm.expectRevert();
        v.verify(b.p);
        b = _load("big_block");
        b.p.txs[0].proof = a.p.txs[0].proof;
        vm.expectRevert();
        v.verify(b.p);
    }

    /// RLP encoding of a minimal header list: parentHash, 3 fillers, transactionsRoot, 3
    /// fillers, number.
    function _forgedHeader(bytes32 parent, bytes32 txRoot, uint256 number) internal pure returns (bytes memory) {
        bytes memory num = _rlpUint(number);
        bytes memory body = bytes.concat(hex"a0", parent, hex"808080", hex"a0", txRoot, hex"808080", num);
        require(body.length < 56 + 200, "len");
        return bytes.concat(hex"f8", bytes1(uint8(body.length)), body);
    }

    function _rlpUint(uint256 n) internal pure returns (bytes memory) {
        if (n == 0) return hex"80";
        if (n < 0x80) return abi.encodePacked(uint8(n));
        uint256 len;
        for (uint256 t = n; t > 0; t >>= 8) {
            len++;
        }
        bytes memory out = new bytes(len);
        for (uint256 i = 0; i < len; ++i) {
            out[len - 1 - i] = bytes1(uint8(n >> (8 * i)));
        }
        return bytes.concat(bytes1(uint8(0x80 + len)), out);
    }
}
