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
        assertEq(blocks, fx.blocks);
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
}
