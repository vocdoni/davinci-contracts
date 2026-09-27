// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

/// @title ICensusValidator
/// @notice Root history of an on-chain lean-IMT census, used by MERKLE_TREE_ONCHAIN_DYNAMIC_V1
///         processes.
/// @dev The ProcessRegistry settles a batch proven against census root R only if
///      getRootBlockNumber(R) is non-zero, at most block.number and at least the process
///      creation block. It calls the getters with 100k gas and needs one 32-byte word back;
///      anything else rejects the batch. That only works with these semantics:
///      - getRootBlockNumber returns block.number for the current root, the block a replaced
///        root was replaced in (its last valid block), and 0 for a root it never held.
///      - Roots are never evicted. An evicted root answers 0, so a batch proven against it
///        stops settling and has to be proven again on a newer root.
///      - The census is append-only and a member's weight never changes, so every later root
///        holds every earlier member with the same weight. The ballot proof binds the weight
///        and the ballot slot derives from the address.
///      A contract that answers the block a root was set in (DavinciDao style) does not fit:
///      the root current at process creation would be rejected.
interface ICensusValidator {
    /// @notice Emitted when an account's weight changes in the census
    /// @dev A census used by the registry emits it only for new members (previousWeight == 0).
    /// @param account The address of the account whose weight changed
    /// @param previousWeight The previous weight of the account
    /// @param newWeight The new weight of the account
    event WeightChanged(address indexed account, uint88 previousWeight, uint88 newWeight);

    /// @notice The last block where a census root is or was valid.
    /// @dev block.number for the current root, the block it was replaced in for an older one,
    ///      0 for a root the census never held (or evicted, which the registry cannot tell apart).
    /// @param root The census Merkle root to validate
    /// @return blockNumber The last block where this root is/was valid (0 if unknown)
    function getRootBlockNumber(uint256 root) external view returns (uint256 blockNumber);

    /// @notice Current census Merkle root (Lean-IMT).
    /// @return root The latest census root
    function getCensusRoot() external view returns (uint256 root);

    /// @notice Returns total voting power associated with a recorded census root.
    /// @dev Returns 0 if the root has not been recorded, or if its total power is zero.
    ///      Pair with getRootBlockNumber(root) to disambiguate unknown roots.
    /// @param root The census root to query.
    /// @return totalVotingPower Total voting power at that root snapshot.
    function getTotalVotingPowerAtRoot(uint256 root) external view returns (uint256 totalVotingPower);
}
