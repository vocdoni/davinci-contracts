// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {Test, console} from "forge-std/Test.sol";
import {OrganizationRegistry} from "../src/OrganizationRegistry.sol";

contract OrganizationRegistryTest is Test {
    OrganizationRegistry public organizationRegistry;

    function setUp() public {
        organizationRegistry = new OrganizationRegistry();
    }
}
