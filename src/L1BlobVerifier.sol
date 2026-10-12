// SPDX-License-Identifier: AGPL-3.0-only
pragma solidity ^0.8.28;

import {IL1Block} from "./interfaces/IL1Block.sol";
import {IL1BlobVerifier} from "./interfaces/IL1BlobVerifier.sol";
import {RLPReader} from "./libraries/op/RLPReader.sol";
import {MerkleTrie} from "./libraries/op/MerkleTrie.sol";
import {Bytes} from "./libraries/op/Bytes.sol";

/// @title L1BlobVerifier
/// @notice Proves on an OP Stack chain that blob transactions are in Ethereum's canonical chain.
///         Headers come newest first; the chain is anchored at any header whose hash is OP's
///         current L1 origin, and transactions may only sit in that header or older ones.
///         Stateless and ownerless.
contract L1BlobVerifier is IL1BlobVerifier {
    /// The L1Block predeploy (0x4200…0015 on OP Stack chains).
    address public immutable override l1Block;

    constructor(address _l1Block) {
        require(_l1Block != address(0), "L1BlobVerifier: zero L1Block");
        l1Block = _l1Block;
    }

    /// @inheritdoc IL1BlobVerifier
    function verify(L1BlobProof calldata p)
        external
        view
        override
        returns (bytes32[] memory vhs, uint256[] memory txBlocks)
    {
        uint256 nh = p.headers.length;
        bytes32[] memory txRoots = new bytes32[](nh);
        uint256[] memory numbers = new uint256[](nh);
        uint256 anchor = _walkHeaders(p, txRoots, numbers);

        uint256 nt = p.txs.length;
        bytes32[][] memory perTx = new bytes32[][](nt);
        txBlocks = new uint256[](nt);
        uint256 total;
        for (uint256 t = 0; t < nt; ++t) {
            L1TxProof calldata tx_ = p.txs[t];
            if (tx_.headerIndex < anchor || tx_.headerIndex >= nh) revert TxAboveOrigin(t);
            bytes memory env = MerkleTrie.get(tx_.key, tx_.proof, txRoots[tx_.headerIndex]);
            perTx[t] = _blobHashes(env, t);
            txBlocks[t] = numbers[tx_.headerIndex];
            total += perTx[t].length;
        }
        vhs = new bytes32[](total);
        uint256 k;
        for (uint256 t = 0; t < nt; ++t) {
            for (uint256 j = 0; j < perTx[t].length; ++j) {
                vhs[k++] = perTx[t][j];
            }
        }
    }

    /// Checks the parentHash chain, decodes roots and numbers, returns the anchor index.
    function _walkHeaders(L1BlobProof calldata p, bytes32[] memory txRoots, uint256[] memory numbers)
        private
        view
        returns (uint256 anchor)
    {
        bytes32 origin = IL1Block(l1Block).hash();
        anchor = type(uint256).max;
        bytes32 expected; // parentHash of the previous header
        for (uint256 i = 0; i < p.headers.length; ++i) {
            bytes memory h = p.headers[i];
            bytes32 hash = keccak256(h);
            if (i > 0 && hash != expected) revert BrokenHeaderChain(i);
            if (anchor == type(uint256).max && hash == origin) anchor = i;
            RLPReader.RLPItem[] memory f = RLPReader.readList(h);
            if (f.length < 9) revert BadHeader(i);
            expected = _bytes32(RLPReader.readBytes(f[0]), i);
            txRoots[i] = _bytes32(RLPReader.readBytes(f[4]), i);
            numbers[i] = _uint(RLPReader.readBytes(f[8]));
        }
        if (anchor == type(uint256).max) revert L1OriginMismatch();
    }

    /// Type-3 envelope -> blob_versioned_hashes (list item 10), non-empty, each 0x01-prefixed.
    function _blobHashes(bytes memory env, uint256 t) private pure returns (bytes32[] memory out) {
        if (env.length < 2 || env[0] != 0x03) revert NotBlobTx(t);
        RLPReader.RLPItem[] memory f = RLPReader.readList(Bytes.slice(env, 1));
        if (f.length < 11) revert NotBlobTx(t);
        RLPReader.RLPItem[] memory hs = RLPReader.readList(f[10]);
        if (hs.length == 0) revert NotBlobTx(t);
        out = new bytes32[](hs.length);
        for (uint256 j = 0; j < hs.length; ++j) {
            bytes memory b = RLPReader.readBytes(hs[j]);
            if (b.length != 32 || b[0] != 0x01) revert BadVersionedHash(t);
            out[j] = bytes32(b);
        }
    }

    function _bytes32(bytes memory b, uint256 i) private pure returns (bytes32) {
        if (b.length != 32) revert BadHeader(i);
        return bytes32(b);
    }

    function _uint(bytes memory b) private pure returns (uint256 v) {
        for (uint256 j = 0; j < b.length; ++j) {
            v = (v << 8) | uint8(b[j]);
        }
    }
}
