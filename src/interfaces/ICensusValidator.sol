// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

/// @title ICensusValidator
/// @notice Interface for validating census Merkle roots
/// @dev Implement this interface to enable external contracts to verify census roots
///      Useful for voting systems, governance contracts, and other on-chain mechanisms
///      that need to validate voting power at specific block numbers
interface ICensusValidator {
    /// @notice Validates a census root and returns the block number when it was set
    /// @dev Returns 0 if the root has never been set or has been evicted from history
    ///      The root history is maintained in a circular buffer (last 100 roots)
    /// @param root The census Merkle root to validate
    /// @return blockNumber The block number when this root was set (0 if invalid/evicted)
    function getRootBlockNumber(uint256 root) external view returns (uint256 blockNumber);
    
    /// @notice Current census Merkle root (Lean-IMT).
    /// @return root The latest census root
    function getCensusRoot() external view returns (uint256 root);
}
