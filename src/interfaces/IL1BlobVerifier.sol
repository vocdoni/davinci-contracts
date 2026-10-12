// SPDX-License-Identifier: AGPL-3.0-only
pragma solidity ^0.8.28;

/// @notice Verifies that blob transactions are in Ethereum's canonical chain, as seen from an OP Stack L2.
interface IL1BlobVerifier {
    struct L1TxProof {
        uint256 headerIndex;
        bytes key;
        bytes[] proof;
    }

    struct L1BlobProof {
        bytes[] headers;
        L1TxProof[] txs;
    }

    error L1OriginMismatch();
    error BrokenHeaderChain(uint256 index);
    error BadHeader(uint256 index);
    error TxAboveOrigin(uint256 tx);
    error NotBlobTx(uint256 tx);
    error BadVersionedHash(uint256 tx);

    function l1Block() external view returns (address);

    function verify(L1BlobProof calldata p) external view returns (bytes32[] memory vhs, uint256[] memory txBlocks);
}
