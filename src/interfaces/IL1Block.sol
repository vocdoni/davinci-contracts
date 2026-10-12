// SPDX-License-Identifier: AGPL-3.0-only
pragma solidity ^0.8.28;

/// @notice The OP Stack L1Block predeploy: the latest Ethereum block the L2 knows.
interface IL1Block {
    function hash() external view returns (bytes32);

    function number() external view returns (uint64);
}
