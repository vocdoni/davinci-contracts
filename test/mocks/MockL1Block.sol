// SPDX-License-Identifier: AGPL-3.0-only
pragma solidity ^0.8.28;

import {IL1Block} from "../../src/interfaces/IL1Block.sol";

/// Stands in for OP's L1Block predeploy: the test or the e2e harness sets the origin.
contract MockL1Block is IL1Block {
    bytes32 public hash;
    uint64 public number;

    function set(uint64 n, bytes32 h) external {
        number = n;
        hash = h;
    }
}
