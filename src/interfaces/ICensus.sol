// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

interface ICensus {
    /** 
     * @notice Gets a census root.
     * @return root The census root.
    */
    function getRoot() external view returns (bytes32 root);
    
    /** 
     * @notice Validates a given census root and returns any random data that fits in 32 bytes.
     * @dev Example: Returns 0x0...0 if the root has never been set or has been evicted from history.
     * @param root The census merkle tree root to validate.
     * @return ok Indicates if the given root was considered valid or not.
     * @return data Optional return data from the external call.
    */
    function checkRoot(bytes32 root) external view returns (bool ok, bytes32 data);
}
