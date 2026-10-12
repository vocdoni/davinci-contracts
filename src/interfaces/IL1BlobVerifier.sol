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
    error TxTooOld(uint256 tx);

    event OriginRecorded(uint256 indexed number, bytes32 hash);

    function l1Block() external view returns (address);

    /// Records OP's current L1 origin (number, hash) so later proofs can anchor at it.
    function record() external;

    /// The recorded hash of L1 block `number`, or zero.
    function recorded(uint256 number) external view returns (bytes32);

    /// Oldest L1 block a proof may reference, counted back from the current origin.
    function MAX_AGE() external view returns (uint256);

    /// @return vhs Versioned hashes of every transaction, in order.
    /// @return blobBlocks One entry per versioned hash (parallel to `vhs`): the L1 block of its transaction.
    function verify(L1BlobProof calldata p) external view returns (bytes32[] memory vhs, uint256[] memory blobBlocks);
}
