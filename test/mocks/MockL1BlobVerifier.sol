// SPDX-License-Identifier: AGPL-3.0-only
pragma solidity ^0.8.28;

import {IL1BlobVerifier} from "../../src/interfaces/IL1BlobVerifier.sol";

/// @dev Returns canned hashes and blocks, or reverts L1OriginMismatch when `fail` is set.
contract MockL1BlobVerifier is IL1BlobVerifier {
    bytes32[] public vhs;
    uint256[] public blocks;
    bool public fail;

    function set(bytes32[] memory _vhs, uint256[] memory _blocks, bool _fail) external {
        vhs = _vhs;
        blocks = _blocks;
        fail = _fail;
    }

    function l1Block() external pure returns (address) {
        return address(0);
    }

    function record() external {}

    function recorded(uint256) external pure returns (bytes32) {
        return bytes32(0);
    }

    function MAX_AGE() external pure returns (uint256) {
        return 7200;
    }

    function verify(L1BlobProof calldata) external view returns (bytes32[] memory, uint256[] memory) {
        if (fail) revert L1OriginMismatch();
        return (vhs, blocks);
    }
}
