// SPDX-License-Identifier: AGPL-3.0-only
pragma solidity ^0.8.28;

import {Test} from "forge-std/Test.sol";
import {stdJson} from "forge-std/StdJson.sol";
import {L1BlobVerifier} from "../src/L1BlobVerifier.sol";
import {IL1BlobVerifier} from "../src/interfaces/IL1BlobVerifier.sol";
import {RLPReader} from "../src/libraries/op/RLPReader.sol";
import {MerkleTrie} from "../src/libraries/op/MerkleTrie.sol";
import {MockL1Block} from "./mocks/MockL1Block.sol";

contract L1BlobVerifierHarness is L1BlobVerifier {
    constructor(address l1) L1BlobVerifier(l1) {}

    function txAt(bytes memory header, bytes memory key, bytes[] memory proof) external pure returns (bytes memory) {
        RLPReader.RLPItem[] memory f = RLPReader.readList(header);
        bytes32 root = bytes32(RLPReader.readBytes(f[4]));
        return MerkleTrie.get(key, proof, root);
    }

    function blobHashes(bytes memory env, uint256 t) external pure returns (bytes32[] memory) {
        return _blobHashes(env, t);
    }
}

contract L1BlobVerifierTest is Test {
    using stdJson for string;

    MockL1Block l1;
    L1BlobVerifierHarness v;

    struct Fx {
        IL1BlobVerifier.L1BlobProof p;
        bytes32[] hashes;
        bytes32[] vhs;
        uint256[] blocks;
        uint256 txType;
    }

    function setUp() public {
        l1 = new MockL1Block();
        v = new L1BlobVerifierHarness(address(l1));
    }

    function _load(string memory name) internal view returns (Fx memory fx) {
        string memory j = vm.readFile(string.concat("test/vectors/l1blob/", name, ".json"));
        fx.p.headers = j.readBytesArray(".headers");
        fx.hashes = j.readBytes32Array(".header_hashes");
        fx.vhs = j.readBytes32Array(".expected_vhs");
        fx.blocks = j.readUintArray(".expected_tx_blocks");
        fx.txType = j.readUint(".tx_type");
        uint256 nt = abi.decode(vm.parseJson(j, ".txs"), (bytes[])).length;
        fx.p.txs = new IL1BlobVerifier.L1TxProof[](nt);
        for (uint256 i = 0; i < nt; ++i) {
            string memory b = string.concat(".txs[", vm.toString(i), "]");
            fx.p.txs[i].headerIndex = j.readUint(string.concat(b, ".header_index"));
            fx.p.txs[i].key = j.readBytes(string.concat(b, ".key"));
            fx.p.txs[i].proof = j.readBytesArray(string.concat(b, ".proof"));
        }
    }

    function _check(Fx memory fx, uint256 anchorIdx) internal {
        l1.set(1, fx.hashes[anchorIdx]);
        (bytes32[] memory vhs, uint256[] memory blocks) = v.verify(fx.p);
        assertEq(vhs, fx.vhs);
        // every fixture has one transaction: its block repeated once per blob
        uint256[] memory want = new uint256[](fx.vhs.length);
        for (uint256 i = 0; i < want.length; ++i) {
            want[i] = fx.blocks[0];
        }
        assertEq(blocks, want);
    }

    /// A proof whose headers start at `from`, with the tx index shifted to match.
    function _dropHeaders(Fx memory fx, uint256 from) internal pure {
        bytes[] memory hs = new bytes[](fx.p.headers.length - from);
        for (uint256 i = 0; i < hs.length; ++i) {
            hs[i] = fx.p.headers[from + i];
        }
        fx.p.headers = hs;
        for (uint256 t = 0; t < fx.p.txs.length; ++t) {
            fx.p.txs[t].headerIndex -= from;
        }
    }

    function _number(bytes memory header) internal pure returns (uint64) {
        RLPReader.RLPItem[] memory f = RLPReader.readList(header);
        bytes memory b = RLPReader.readBytes(f[8]);
        uint256 n;
        for (uint256 i = 0; i < b.length; ++i) {
            n = (n << 8) | uint8(b[i]);
        }
        return uint64(n);
    }

    function test_Recent_AnchorAtTop() public {
        Fx memory fx = _load("recent");
        _check(fx, 0);
    }

    function test_Recent_AnchorInMiddle() public {
        Fx memory fx = _load("recent");
        _check(fx, 1);
    }

    function test_Recent_AnchorAtBlobBlock() public {
        Fx memory fx = _load("recent");
        _check(fx, fx.hashes.length - 1);
    }

    function test_Cancun_DecodesShorterHeader() public {
        Fx memory fx = _load("cancun");
        _check(fx, 0);
    }

    function test_Index128Plus() public {
        Fx memory fx = _load("index128plus");
        _check(fx, 0);
    }

    function test_KeyIndex0() public view {
        Fx memory fx = _load("index0");
        bytes memory env = v.txAt(fx.p.headers[fx.p.txs[0].headerIndex], fx.p.txs[0].key, fx.p.txs[0].proof);
        assertGt(env.length, 0);
        assertEq(uint8(env[0]), fx.txType);
    }

    function test_BigBlockIndex257() public view {
        Fx memory fx = _load("big_block");
        assertEq(fx.p.txs[0].key, hex"820101");
        bytes memory env = v.txAt(fx.p.headers[fx.p.txs[0].headerIndex], fx.p.txs[0].key, fx.p.txs[0].proof);
        assertGt(env.length, 0);
        assertEq(uint8(env[0]), fx.txType);
    }

    function test_RevertWhen_NoAnchor() public {
        Fx memory fx = _load("recent");
        l1.set(1, keccak256("x"));
        vm.expectRevert(IL1BlobVerifier.L1OriginMismatch.selector);
        v.verify(fx.p);
    }

    function test_RevertWhen_BrokenChain() public {
        Fx memory fx = _load("recent");
        (fx.p.headers[1], fx.p.headers[2]) = (fx.p.headers[2], fx.p.headers[1]);
        l1.set(1, fx.hashes[0]);
        vm.expectRevert(abi.encodeWithSelector(IL1BlobVerifier.BrokenHeaderChain.selector, 1));
        v.verify(fx.p);
    }

    function test_RevertWhen_TxAboveOrigin() public {
        Fx memory fx = _load("recent");
        l1.set(1, fx.hashes[fx.hashes.length - 1]);
        fx.p.txs[0].headerIndex = 0;
        vm.expectRevert(abi.encodeWithSelector(IL1BlobVerifier.TxAboveOrigin.selector, 0));
        v.verify(fx.p);
    }

    function test_RevertWhen_NotBlobTx() public {
        Fx memory fx = _load("not_blob");
        l1.set(1, fx.hashes[fx.p.txs[0].headerIndex]);
        vm.expectRevert(abi.encodeWithSelector(IL1BlobVerifier.NotBlobTx.selector, 0));
        v.verify(fx.p);
    }

    function test_RevertWhen_BadProof() public {
        Fx memory fx = _load("recent");
        l1.set(1, fx.hashes[0]);
        fx.p.txs[0].proof[0][5] = bytes1(uint8(fx.p.txs[0].proof[0][5]) ^ 0x01);
        vm.expectRevert();
        v.verify(fx.p);
    }

    function test_RevertWhen_ShortHeader() public {
        Fx memory fx = _load("recent");
        fx.p.headers = new bytes[](1);
        fx.p.headers[0] = hex"c3010203";
        l1.set(1, keccak256(fx.p.headers[0]));
        vm.expectRevert(abi.encodeWithSelector(IL1BlobVerifier.BadHeader.selector, 0));
        v.verify(fx.p);
    }

    function test_RevertWhen_EmptyHeaders() public {
        Fx memory fx = _load("recent");
        fx.p.headers = new bytes[](0);
        fx.p.txs = new IL1BlobVerifier.L1TxProof[](0);
        l1.set(1, fx.hashes[0]);
        vm.expectRevert(IL1BlobVerifier.L1OriginMismatch.selector);
        v.verify(fx.p);
    }

    function test_RecordedAnchor() public {
        Fx memory fx = _load("recent");
        uint64 n = _number(fx.p.headers[2]);
        uint64 txBlock = uint64(fx.blocks[0]);
        l1.set(n, fx.hashes[2]);
        vm.expectEmit(true, false, false, true, address(v));
        emit IL1BlobVerifier.OriginRecorded(n, fx.hashes[2]);
        v.record();
        assertEq(v.recorded(n), fx.hashes[2]);

        l1.set(txBlock + 10, keccak256("later"));
        _dropHeaders(fx, 2);
        (bytes32[] memory vhs, uint256[] memory blocks) = v.verify(fx.p);
        assertEq(vhs, fx.vhs);
        assertEq(blocks.length, vhs.length);
        assertEq(blocks[0], txBlock);
    }

    function test_RevertWhen_NotRecorded() public {
        Fx memory fx = _load("recent");
        l1.set(uint64(fx.blocks[0]) + 10, keccak256("later"));
        _dropHeaders(fx, 2);
        vm.expectRevert(IL1BlobVerifier.L1OriginMismatch.selector);
        v.verify(fx.p);
    }

    function test_RevertWhen_TxTooOld() public {
        Fx memory fx = _load("recent");
        uint64 txBlock = uint64(fx.blocks[0]);
        uint64 maxAge = uint64(v.MAX_AGE());
        l1.set(txBlock + maxAge, fx.hashes[0]);
        v.verify(fx.p);
        l1.set(txBlock + maxAge + 1, fx.hashes[0]);
        vm.expectRevert(abi.encodeWithSelector(IL1BlobVerifier.TxTooOld.selector, 0));
        v.verify(fx.p);
    }

    function test_TwoTxs() public {
        Fx memory fx = _load("recent");
        l1.set(1, fx.hashes[0]);
        IL1BlobVerifier.L1TxProof[] memory two = new IL1BlobVerifier.L1TxProof[](2);
        two[0] = fx.p.txs[0];
        two[1] = fx.p.txs[0];
        fx.p.txs = two;
        (bytes32[] memory vhs, uint256[] memory blocks) = v.verify(fx.p);
        assertEq(vhs.length, 2 * fx.vhs.length);
        assertEq(blocks.length, vhs.length);
        for (uint256 i = 0; i < vhs.length; ++i) {
            assertEq(vhs[i], fx.vhs[i % fx.vhs.length]);
            assertEq(blocks[i], fx.blocks[0]);
        }
    }

    function test_RevertWhen_HeaderIndexOutOfRange() public {
        Fx memory fx = _load("recent");
        l1.set(1, fx.hashes[0]);
        fx.p.txs[0].headerIndex = fx.p.headers.length;
        vm.expectRevert(abi.encodeWithSelector(IL1BlobVerifier.TxAboveOrigin.selector, 0));
        v.verify(fx.p);
    }

    /// 0x03 || rlp([10 empty strings, [vh]]) with vh = `first` || 31 bytes of 0xaa.
    function _envelope(bytes1 first) internal pure returns (bytes memory) {
        return abi.encodePacked(
            hex"03ec",
            hex"80808080808080808080",
            hex"e1a0",
            first,
            bytes31(0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa)
        );
    }

    function test_RevertWhen_BadVersionedHash() public {
        vm.expectRevert(abi.encodeWithSelector(IL1BlobVerifier.BadVersionedHash.selector, 0));
        v.blobHashes(_envelope(0x02), 0);
        bytes32[] memory out = v.blobHashes(_envelope(0x01), 0);
        assertEq(out.length, 1);
        assertEq(
            out[0],
            bytes32(
                abi.encodePacked(
                    bytes1(0x01), bytes31(0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa)
                )
            )
        );
    }
}
