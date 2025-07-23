// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import { Test } from "forge-std/Test.sol";
import { OrganizationRegistry } from "../src/OrganizationRegistry.sol";
import { IOrganizationRegistry } from "../src/interfaces/IOrganizationRegistry.sol";

contract OrganizationRegistryTest is Test {
    OrganizationRegistry public organizationRegistry;

    // Test addresses
    address public constant ORGANIZATION_ID = 0x0056d6eD515B2e0af39bb068F587D0de83fACD1B;
    address public constant ADMIN_1 = 0x1111111111111111111111111111111111111111;
    address public constant ADMIN_2 = 0x2222222222222222222222222222222222222222;
    address public constant ADMIN_3 = 0x3333333333333333333333333333333333333333;
    address public constant NON_ADMIN = 0x4444444444444444444444444444444444444444;

    // Test data
    string public constant ORGANIZATION_NAME = "Test Organization";
    string public constant ORGANIZATION_METADATA_URI = "https://example.com/metadata";
    string public constant UPDATED_NAME = "Updated Organization";
    string public constant UPDATED_METADATA_URI = "https://example.com/updated-metadata";

    function setUp() public {
        organizationRegistry = new OrganizationRegistry();
    }

    function createTestOrganization() internal returns (address) {
        address[] memory admins = new address[](2);
        admins[0] = ADMIN_1;
        admins[1] = ADMIN_2;

        vm.prank(ORGANIZATION_ID);
        organizationRegistry.createOrganization(ORGANIZATION_NAME, ORGANIZATION_METADATA_URI, admins);
        return ORGANIZATION_ID;
    }

    function createTestOrganizationWithoutAdmins() internal returns (address) {
        address[] memory admins = new address[](0);

        vm.prank(ORGANIZATION_ID);
        organizationRegistry.createOrganization(ORGANIZATION_NAME, ORGANIZATION_METADATA_URI, admins);
        return ORGANIZATION_ID;
    }

    // ========== Organization Creation Tests ==========

    function test_CreateOrganization_Success() public {
        address[] memory admins = new address[](2);
        admins[0] = ADMIN_1;
        admins[1] = ADMIN_2;

        vm.expectEmit(true, false, false, false);
        emit IOrganizationRegistry.OrganizationCreated(ORGANIZATION_ID);

        vm.prank(ORGANIZATION_ID);
        organizationRegistry.createOrganization(ORGANIZATION_NAME, ORGANIZATION_METADATA_URI, admins);

        // Verify organization was created
        IOrganizationRegistry.Organization memory org = organizationRegistry.getOrganization(ORGANIZATION_ID);
        assertEq(org.name, ORGANIZATION_NAME);
        assertEq(org.metadataURI, ORGANIZATION_METADATA_URI);
        assertEq(org.administrators.length, 3); // 2 provided + msg.sender
        assertEq(org.administrators[0], ADMIN_1);
        assertEq(org.administrators[1], ADMIN_2);
        assertEq(org.administrators[2], ORGANIZATION_ID); // msg.sender added last

        // Verify organization count
        assertEq(organizationRegistry.organizationCount(), 1);

        // Verify exists function
        assertTrue(organizationRegistry.exists(ORGANIZATION_ID));
    }

    function test_CreateOrganization_WithoutAdditionalAdmins() public {
        address[] memory admins = new address[](0);

        vm.expectEmit(true, false, false, false);
        emit IOrganizationRegistry.OrganizationCreated(ORGANIZATION_ID);

        vm.prank(ORGANIZATION_ID);
        organizationRegistry.createOrganization(ORGANIZATION_NAME, ORGANIZATION_METADATA_URI, admins);

        // Verify organization was created
        IOrganizationRegistry.Organization memory org = organizationRegistry.getOrganization(ORGANIZATION_ID);
        assertEq(org.name, ORGANIZATION_NAME);
        assertEq(org.metadataURI, ORGANIZATION_METADATA_URI);
        assertEq(org.administrators.length, 1); // Only msg.sender
        assertEq(org.administrators[0], ORGANIZATION_ID);

        // Verify organization count
        assertEq(organizationRegistry.organizationCount(), 1);
    }

    function test_CreateOrganization_WithEmptyMetadataURI() public {
        address[] memory admins = new address[](1);
        admins[0] = ADMIN_1;

        vm.prank(ORGANIZATION_ID);
        organizationRegistry.createOrganization(ORGANIZATION_NAME, "", admins);

        // Should succeed - empty metadata URI is allowed
        IOrganizationRegistry.Organization memory org = organizationRegistry.getOrganization(ORGANIZATION_ID);
        assertEq(org.name, ORGANIZATION_NAME);
        assertEq(org.metadataURI, "");
        assertEq(org.administrators.length, 2);
    }

    function test_CreateOrganization_WithDuplicateAdmins() public {
        address[] memory admins = new address[](3);
        admins[0] = ADMIN_1;
        admins[1] = ADMIN_1; // Duplicate
        admins[2] = ADMIN_2;

        vm.prank(ORGANIZATION_ID);
        organizationRegistry.createOrganization(ORGANIZATION_NAME, ORGANIZATION_METADATA_URI, admins);

        // Should succeed - duplicates are allowed in the input array
        IOrganizationRegistry.Organization memory org = organizationRegistry.getOrganization(ORGANIZATION_ID);
        assertEq(org.administrators.length, 4); // 3 provided + msg.sender
        assertEq(org.administrators[0], ADMIN_1);
        assertEq(org.administrators[1], ADMIN_1); // Duplicate preserved
        assertEq(org.administrators[2], ADMIN_2);
        assertEq(org.administrators[3], ORGANIZATION_ID);
    }

    function test_CreateOrganization_WithSelfInAdmins() public {
        address[] memory admins = new address[](2);
        admins[0] = ADMIN_1;
        admins[1] = ORGANIZATION_ID; // Self in admins list

        vm.prank(ORGANIZATION_ID);
        organizationRegistry.createOrganization(ORGANIZATION_NAME, ORGANIZATION_METADATA_URI, admins);

        // Self should be skipped in the loop but added at the end
        IOrganizationRegistry.Organization memory org = organizationRegistry.getOrganization(ORGANIZATION_ID);
        assertEq(org.administrators.length, 2); // ADMIN_1 + ORGANIZATION_ID (added once)
        assertEq(org.administrators[0], ADMIN_1);
        assertEq(org.administrators[1], ORGANIZATION_ID);
    }

    function test_CreateOrganization_InvalidName_Empty() public {
        address[] memory admins = new address[](1);
        admins[0] = ADMIN_1;

        vm.prank(ORGANIZATION_ID);
        vm.expectRevert(IOrganizationRegistry.InvalidOrganizationName.selector);
        organizationRegistry.createOrganization("", ORGANIZATION_METADATA_URI, admins);
    }

    function test_CreateOrganization_InvalidAddress_ZeroAddress() public {
        address[] memory admins = new address[](2);
        admins[0] = ADMIN_1;
        admins[1] = address(0); // Zero address

        vm.prank(ORGANIZATION_ID);
        vm.expectRevert(IOrganizationRegistry.InvalidAddress.selector);
        organizationRegistry.createOrganization(ORGANIZATION_NAME, ORGANIZATION_METADATA_URI, admins);
    }

    function test_CreateOrganization_AlreadyExists() public {
        createTestOrganization();

        // Try to create another organization with the same ID (msg.sender)
        address[] memory admins = new address[](1);
        admins[0] = ADMIN_1;

        vm.prank(ORGANIZATION_ID);
        vm.expectRevert(IOrganizationRegistry.OrganizationAlreadyExists.selector);
        organizationRegistry.createOrganization("Another Name", ORGANIZATION_METADATA_URI, admins);
    }

    function test_CreateOrganization_MultipleOrganizations() public {
        // Create first organization
        address org1 = ORGANIZATION_ID;
        address[] memory admins1 = new address[](1);
        admins1[0] = ADMIN_1;

        vm.prank(org1);
        organizationRegistry.createOrganization("Org 1", ORGANIZATION_METADATA_URI, admins1);

        // Create second organization
        address org2 = address(0x5555555555555555555555555555555555555555);
        address[] memory admins2 = new address[](1);
        admins2[0] = ADMIN_2;

        vm.prank(org2);
        organizationRegistry.createOrganization("Org 2", ORGANIZATION_METADATA_URI, admins2);

        // Verify both organizations exist
        assertTrue(organizationRegistry.exists(org1));
        assertTrue(organizationRegistry.exists(org2));
        assertEq(organizationRegistry.organizationCount(), 2);

        // Verify organization data
        IOrganizationRegistry.Organization memory org1Data = organizationRegistry.getOrganization(org1);
        IOrganizationRegistry.Organization memory org2Data = organizationRegistry.getOrganization(org2);

        assertEq(org1Data.name, "Org 1");
        assertEq(org2Data.name, "Org 2");
    }

    // ========== Organization Retrieval Tests ==========

    function test_GetOrganization_Success() public {
        createTestOrganization();

        IOrganizationRegistry.Organization memory org = organizationRegistry.getOrganization(ORGANIZATION_ID);
        assertEq(org.name, ORGANIZATION_NAME);
        assertEq(org.metadataURI, ORGANIZATION_METADATA_URI);
        assertEq(org.administrators.length, 3);
        assertEq(org.administrators[0], ADMIN_1);
        assertEq(org.administrators[1], ADMIN_2);
        assertEq(org.administrators[2], ORGANIZATION_ID);
    }

    function test_GetOrganization_NonExistent() public {
        // Getting non-existent organization should return empty struct
        IOrganizationRegistry.Organization memory org = organizationRegistry.getOrganization(ORGANIZATION_ID);
        assertEq(org.name, "");
        assertEq(org.metadataURI, "");
        assertEq(org.administrators.length, 0);
    }

    function test_Exists_True() public {
        createTestOrganization();
        assertTrue(organizationRegistry.exists(ORGANIZATION_ID));
    }

    function test_Exists_False() public {
        assertFalse(organizationRegistry.exists(ORGANIZATION_ID));
    }

    // ========== Administrator Check Tests ==========

    function test_IsAdministrator_True() public {
        createTestOrganization();

        assertTrue(organizationRegistry.isAdministrator(ORGANIZATION_ID, ADMIN_1));
        assertTrue(organizationRegistry.isAdministrator(ORGANIZATION_ID, ADMIN_2));
        assertTrue(organizationRegistry.isAdministrator(ORGANIZATION_ID, ORGANIZATION_ID));
    }

    function test_IsAdministrator_False() public {
        createTestOrganization();

        assertFalse(organizationRegistry.isAdministrator(ORGANIZATION_ID, NON_ADMIN));
        assertFalse(organizationRegistry.isAdministrator(ORGANIZATION_ID, address(0)));
    }

    function test_IsAdministrator_OrganizationNotFound() public {
        vm.expectRevert(IOrganizationRegistry.OrganizationNotFound.selector);
        organizationRegistry.isAdministrator(ORGANIZATION_ID, ADMIN_1);
    }

    // ========== Organization Update Tests ==========

    function test_UpdateOrganization_Success_ByCreator() public {
        createTestOrganization();

        vm.expectEmit(true, true, false, false);
        emit IOrganizationRegistry.OrganizationUpdated(ORGANIZATION_ID, ORGANIZATION_ID);

        vm.prank(ORGANIZATION_ID);
        organizationRegistry.updateOrganization(ORGANIZATION_ID, UPDATED_NAME, UPDATED_METADATA_URI);

        IOrganizationRegistry.Organization memory org = organizationRegistry.getOrganization(ORGANIZATION_ID);
        assertEq(org.name, UPDATED_NAME);
        assertEq(org.metadataURI, UPDATED_METADATA_URI);
        // Administrators should remain unchanged
        assertEq(org.administrators.length, 3);
    }

    function test_UpdateOrganization_Success_ByAdmin() public {
        createTestOrganization();

        vm.expectEmit(true, true, false, false);
        emit IOrganizationRegistry.OrganizationUpdated(ORGANIZATION_ID, ADMIN_1);

        vm.prank(ADMIN_1);
        organizationRegistry.updateOrganization(ORGANIZATION_ID, UPDATED_NAME, UPDATED_METADATA_URI);

        IOrganizationRegistry.Organization memory org = organizationRegistry.getOrganization(ORGANIZATION_ID);
        assertEq(org.name, UPDATED_NAME);
        assertEq(org.metadataURI, UPDATED_METADATA_URI);
    }

    function test_UpdateOrganization_NotAdministrator() public {
        createTestOrganization();

        vm.prank(NON_ADMIN);
        vm.expectRevert(IOrganizationRegistry.NotAdministrator.selector);
        organizationRegistry.updateOrganization(ORGANIZATION_ID, UPDATED_NAME, UPDATED_METADATA_URI);
    }

    function test_UpdateOrganization_OrganizationNotFound() public {
        vm.prank(ORGANIZATION_ID);
        vm.expectRevert(IOrganizationRegistry.NotAdministrator.selector);
        organizationRegistry.updateOrganization(ORGANIZATION_ID, UPDATED_NAME, UPDATED_METADATA_URI);
    }

    function test_UpdateOrganization_InvalidName_Empty() public {
        createTestOrganization();

        vm.prank(ORGANIZATION_ID);
        vm.expectRevert(IOrganizationRegistry.InvalidOrganizationName.selector);
        organizationRegistry.updateOrganization(ORGANIZATION_ID, "", UPDATED_METADATA_URI);
    }

    function test_UpdateOrganization_InvalidMetadataURI_Empty() public {
        createTestOrganization();

        vm.prank(ORGANIZATION_ID);
        vm.expectRevert(IOrganizationRegistry.InvalidMetadataURI.selector);
        organizationRegistry.updateOrganization(ORGANIZATION_ID, UPDATED_NAME, "");
    }

    // ========== Add Administrator Tests ==========

    function test_AddAdministrator_Success() public {
        createTestOrganization();

        vm.expectEmit(true, true, false, false);
        emit IOrganizationRegistry.AdministratorAdded(ORGANIZATION_ID, ADMIN_3);

        vm.prank(ORGANIZATION_ID);
        organizationRegistry.addAdministrator(ORGANIZATION_ID, ADMIN_3);

        // Verify administrator was added
        assertTrue(organizationRegistry.isAdministrator(ORGANIZATION_ID, ADMIN_3));

        IOrganizationRegistry.Organization memory org = organizationRegistry.getOrganization(ORGANIZATION_ID);
        assertEq(org.administrators.length, 4);
        assertEq(org.administrators[3], ADMIN_3);
    }

    function test_AddAdministrator_Unauthorized_NotOrganization() public {
        createTestOrganization();

        // Only the organization itself can add administrators, not even existing admins
        vm.prank(ADMIN_1);
        vm.expectRevert(IOrganizationRegistry.Unauthorized.selector);
        organizationRegistry.addAdministrator(ORGANIZATION_ID, ADMIN_3);
    }

    function test_AddAdministrator_OrganizationNotFound() public {
        vm.prank(ORGANIZATION_ID);
        vm.expectRevert(IOrganizationRegistry.OrganizationNotFound.selector);
        organizationRegistry.addAdministrator(ORGANIZATION_ID, ADMIN_1);
    }

    function test_AddAdministrator_InvalidAddress_ZeroAddress() public {
        createTestOrganization();

        vm.prank(ORGANIZATION_ID);
        vm.expectRevert(IOrganizationRegistry.InvalidAddress.selector);
        organizationRegistry.addAdministrator(ORGANIZATION_ID, address(0));
    }

    function test_AddAdministrator_AlreadyAdministrator() public {
        createTestOrganization();

        vm.prank(ORGANIZATION_ID);
        vm.expectRevert(IOrganizationRegistry.AlreadyAdministrator.selector);
        organizationRegistry.addAdministrator(ORGANIZATION_ID, ADMIN_1);
    }

    function test_AddAdministrator_AlreadyAdministrator_Self() public {
        createTestOrganization();

        vm.prank(ORGANIZATION_ID);
        vm.expectRevert(IOrganizationRegistry.AlreadyAdministrator.selector);
        organizationRegistry.addAdministrator(ORGANIZATION_ID, ORGANIZATION_ID);
    }

    // ========== Remove Administrator Tests ==========

    function test_RemoveAdministrator_Success_ByOrganization() public {
        createTestOrganization();

        vm.expectEmit(true, true, true, false);
        emit IOrganizationRegistry.AdministratorRemoved(ORGANIZATION_ID, ADMIN_1, ORGANIZATION_ID);

        vm.prank(ORGANIZATION_ID);
        organizationRegistry.removeAdministrator(ORGANIZATION_ID, ADMIN_1);

        // Verify administrator was removed
        assertFalse(organizationRegistry.isAdministrator(ORGANIZATION_ID, ADMIN_1));

        IOrganizationRegistry.Organization memory org = organizationRegistry.getOrganization(ORGANIZATION_ID);
        assertEq(org.administrators.length, 2);
        // ADMIN_1 should be removed, ADMIN_2 and ORGANIZATION_ID should remain
        // Due to swap-and-pop, the last element takes ADMIN_1's place
        assertEq(org.administrators[0], ORGANIZATION_ID); // Last element moved to first position
        assertEq(org.administrators[1], ADMIN_2);
    }

    function test_RemoveAdministrator_Success_SelfRemoval() public {
        createTestOrganization();

        vm.expectEmit(true, true, true, false);
        emit IOrganizationRegistry.AdministratorRemoved(ORGANIZATION_ID, ADMIN_1, ADMIN_1);

        vm.prank(ADMIN_1);
        organizationRegistry.removeAdministrator(ORGANIZATION_ID, ADMIN_2); // Parameter ignored for self-removal

        // Verify ADMIN_1 was removed (not ADMIN_2)
        assertFalse(organizationRegistry.isAdministrator(ORGANIZATION_ID, ADMIN_1));
        assertTrue(organizationRegistry.isAdministrator(ORGANIZATION_ID, ADMIN_2));

        IOrganizationRegistry.Organization memory org = organizationRegistry.getOrganization(ORGANIZATION_ID);
        assertEq(org.administrators.length, 2);
    }

    function test_RemoveAdministrator_NotAdministrator_Caller() public {
        createTestOrganization();

        vm.prank(NON_ADMIN);
        vm.expectRevert(IOrganizationRegistry.NotAdministrator.selector);
        organizationRegistry.removeAdministrator(ORGANIZATION_ID, ADMIN_1);
    }

    function test_RemoveAdministrator_NotAdministrator_Target() public {
        createTestOrganization();

        vm.prank(ORGANIZATION_ID);
        vm.expectRevert(IOrganizationRegistry.NotAdministrator.selector);
        organizationRegistry.removeAdministrator(ORGANIZATION_ID, NON_ADMIN);
    }

    function test_RemoveAdministrator_OrganizationNotFound() public {
        vm.prank(ORGANIZATION_ID);
        vm.expectRevert(IOrganizationRegistry.NotAdministrator.selector);
        organizationRegistry.removeAdministrator(ORGANIZATION_ID, ADMIN_1);
    }

    function test_RemoveAdministrator_InvalidAddress_ZeroAddress() public {
        createTestOrganization();

        vm.prank(ORGANIZATION_ID);
        vm.expectRevert(IOrganizationRegistry.InvalidAddress.selector);
        organizationRegistry.removeAdministrator(ORGANIZATION_ID, address(0));
    }

    function test_RemoveAdministrator_LastAdmin() public {
        createTestOrganizationWithoutAdmins();

        // Try to remove the only administrator (the organization itself)
        vm.prank(ORGANIZATION_ID);
        organizationRegistry.removeAdministrator(ORGANIZATION_ID, ORGANIZATION_ID);

        // Should succeed - organization can remove itself as admin
        assertFalse(organizationRegistry.isAdministrator(ORGANIZATION_ID, ORGANIZATION_ID));

        IOrganizationRegistry.Organization memory org = organizationRegistry.getOrganization(ORGANIZATION_ID);
        assertEq(org.administrators.length, 0);
    }

    function test_RemoveAdministrator_SelfRemoval_AsOnlyAdmin() public {
        createTestOrganizationWithoutAdmins();

        vm.prank(ORGANIZATION_ID);
        organizationRegistry.removeAdministrator(ORGANIZATION_ID, ORGANIZATION_ID); // Remove self

        // Verify self-removal worked
        assertFalse(organizationRegistry.isAdministrator(ORGANIZATION_ID, ORGANIZATION_ID));

        IOrganizationRegistry.Organization memory org = organizationRegistry.getOrganization(ORGANIZATION_ID);
        assertEq(org.administrators.length, 0);
    }

    // ========== Delete Organization Tests ==========

    function test_DeleteOrganization_Success() public {
        createTestOrganization();
        uint32 initialCount = organizationRegistry.organizationCount();

        vm.prank(ORGANIZATION_ID);
        organizationRegistry.deleteOrganization(ORGANIZATION_ID);

        // Verify organization was deleted
        assertFalse(organizationRegistry.exists(ORGANIZATION_ID));
        assertEq(organizationRegistry.organizationCount(), initialCount - 1);

        // Verify organization data is cleared
        IOrganizationRegistry.Organization memory org = organizationRegistry.getOrganization(ORGANIZATION_ID);
        assertEq(org.name, "");
        assertEq(org.metadataURI, "");
        assertEq(org.administrators.length, 0);
    }

    function test_DeleteOrganization_Unauthorized_NotOrganization() public {
        createTestOrganization();

        vm.prank(ADMIN_1);
        vm.expectRevert(IOrganizationRegistry.Unauthorized.selector);
        organizationRegistry.deleteOrganization(ORGANIZATION_ID);
    }

    function test_DeleteOrganization_Unauthorized_NonAdmin() public {
        createTestOrganization();

        vm.prank(NON_ADMIN);
        vm.expectRevert(IOrganizationRegistry.Unauthorized.selector);
        organizationRegistry.deleteOrganization(ORGANIZATION_ID);
    }

    function test_DeleteOrganization_OrganizationNotFound() public {
        vm.prank(ORGANIZATION_ID);
        vm.expectRevert(IOrganizationRegistry.OrganizationNotFound.selector);
        organizationRegistry.deleteOrganization(ORGANIZATION_ID);
    }

    function test_DeleteOrganization_MultipleOrganizations() public {
        // Create two organizations
        address org1 = ORGANIZATION_ID;
        address org2 = address(0x5555555555555555555555555555555555555555);

        address[] memory admins = new address[](0);

        vm.prank(org1);
        organizationRegistry.createOrganization("Org 1", ORGANIZATION_METADATA_URI, admins);

        vm.prank(org2);
        organizationRegistry.createOrganization("Org 2", ORGANIZATION_METADATA_URI, admins);

        assertEq(organizationRegistry.organizationCount(), 2);

        // Delete first organization
        vm.prank(org1);
        organizationRegistry.deleteOrganization(org1);

        // Verify only first organization was deleted
        assertFalse(organizationRegistry.exists(org1));
        assertTrue(organizationRegistry.exists(org2));
        assertEq(organizationRegistry.organizationCount(), 1);
    }

    // ========== Edge Cases and Complex Scenarios ==========

    function test_ComplexAdministratorManagement() public {
        createTestOrganization();

        // Add a new administrator
        vm.prank(ORGANIZATION_ID);
        organizationRegistry.addAdministrator(ORGANIZATION_ID, ADMIN_3);

        // Verify all administrators
        assertTrue(organizationRegistry.isAdministrator(ORGANIZATION_ID, ADMIN_1));
        assertTrue(organizationRegistry.isAdministrator(ORGANIZATION_ID, ADMIN_2));
        assertTrue(organizationRegistry.isAdministrator(ORGANIZATION_ID, ADMIN_3));
        assertTrue(organizationRegistry.isAdministrator(ORGANIZATION_ID, ORGANIZATION_ID));

        // Remove one administrator
        vm.prank(ORGANIZATION_ID);
        organizationRegistry.removeAdministrator(ORGANIZATION_ID, ADMIN_2);

        // Verify removal
        assertTrue(organizationRegistry.isAdministrator(ORGANIZATION_ID, ADMIN_1));
        assertFalse(organizationRegistry.isAdministrator(ORGANIZATION_ID, ADMIN_2));
        assertTrue(organizationRegistry.isAdministrator(ORGANIZATION_ID, ADMIN_3));
        assertTrue(organizationRegistry.isAdministrator(ORGANIZATION_ID, ORGANIZATION_ID));

        // Admin can update organization
        vm.prank(ADMIN_1);
        organizationRegistry.updateOrganization(ORGANIZATION_ID, UPDATED_NAME, UPDATED_METADATA_URI);

        IOrganizationRegistry.Organization memory org = organizationRegistry.getOrganization(ORGANIZATION_ID);
        assertEq(org.name, UPDATED_NAME);
        assertEq(org.metadataURI, UPDATED_METADATA_URI);
    }

    function test_OrganizationLifecycle() public {
        // Create organization
        createTestOrganization();
        assertTrue(organizationRegistry.exists(ORGANIZATION_ID));
        assertEq(organizationRegistry.organizationCount(), 1);

        // Update organization
        vm.prank(ORGANIZATION_ID);
        organizationRegistry.updateOrganization(ORGANIZATION_ID, UPDATED_NAME, UPDATED_METADATA_URI);

        // Add administrator
        vm.prank(ORGANIZATION_ID);
        organizationRegistry.addAdministrator(ORGANIZATION_ID, ADMIN_3);

        // Remove administrator
        vm.prank(ORGANIZATION_ID);
        organizationRegistry.removeAdministrator(ORGANIZATION_ID, ADMIN_1);

        // Delete organization
        vm.prank(ORGANIZATION_ID);
        organizationRegistry.deleteOrganization(ORGANIZATION_ID);

        // Verify deletion
        assertFalse(organizationRegistry.exists(ORGANIZATION_ID));
        assertEq(organizationRegistry.organizationCount(), 0);
    }

    function test_MaxAdministrators() public {
        createTestOrganizationWithoutAdmins();

        // Add many administrators
        for (uint160 i = 1; i <= 100; i++) {
            address newAdmin = address(i);
            vm.prank(ORGANIZATION_ID);
            organizationRegistry.addAdministrator(ORGANIZATION_ID, newAdmin);
        }

        IOrganizationRegistry.Organization memory org = organizationRegistry.getOrganization(ORGANIZATION_ID);
        assertEq(org.administrators.length, 101); // 100 added + 1 original (ORGANIZATION_ID)

        // Verify all administrators
        for (uint160 i = 1; i <= 100; i++) {
            assertTrue(organizationRegistry.isAdministrator(ORGANIZATION_ID, address(i)));
        }
        assertTrue(organizationRegistry.isAdministrator(ORGANIZATION_ID, ORGANIZATION_ID));
    }

    function test_RemoveAdministrator_OrderPreservation() public {
        createTestOrganizationWithoutAdmins();

        // Add administrators in specific order
        address[] memory adminsToAdd = new address[](5);
        adminsToAdd[0] = address(0x1001);
        adminsToAdd[1] = address(0x1002);
        adminsToAdd[2] = address(0x1003);
        adminsToAdd[3] = address(0x1004);
        adminsToAdd[4] = address(0x1005);

        for (uint i = 0; i < adminsToAdd.length; i++) {
            vm.prank(ORGANIZATION_ID);
            organizationRegistry.addAdministrator(ORGANIZATION_ID, adminsToAdd[i]);
        }

        // Remove middle administrator
        vm.prank(ORGANIZATION_ID);
        organizationRegistry.removeAdministrator(ORGANIZATION_ID, address(0x1003));

        // Verify the last element took the removed element's place
        IOrganizationRegistry.Organization memory org = organizationRegistry.getOrganization(ORGANIZATION_ID);
        assertEq(org.administrators.length, 5); // 4 added + 1 original - 1 removed

        // The removed admin should not be found
        assertFalse(organizationRegistry.isAdministrator(ORGANIZATION_ID, address(0x1003)));

        // All other admins should still be present
        assertTrue(organizationRegistry.isAdministrator(ORGANIZATION_ID, address(0x1001)));
        assertTrue(organizationRegistry.isAdministrator(ORGANIZATION_ID, address(0x1002)));
        assertTrue(organizationRegistry.isAdministrator(ORGANIZATION_ID, address(0x1004)));
        assertTrue(organizationRegistry.isAdministrator(ORGANIZATION_ID, address(0x1005)));
        assertTrue(organizationRegistry.isAdministrator(ORGANIZATION_ID, ORGANIZATION_ID));
    }

    function test_CreateOrganization_LongName() public {
        string
            memory longName = "This is a very long organization name that contains many characters to test the limits of string handling in the contract and ensure it works properly with extended names";
        address[] memory admins = new address[](1);
        admins[0] = ADMIN_1;

        vm.prank(ORGANIZATION_ID);
        organizationRegistry.createOrganization(longName, ORGANIZATION_METADATA_URI, admins);

        IOrganizationRegistry.Organization memory org = organizationRegistry.getOrganization(ORGANIZATION_ID);
        assertEq(org.name, longName);
    }

    function test_CreateOrganization_LongMetadataURI() public {
        string
            memory longURI = "https://example.com/very/long/path/to/metadata/that/contains/many/segments/and/parameters?param1=value1&param2=value2&param3=value3&param4=value4&param5=value5";
        address[] memory admins = new address[](1);
        admins[0] = ADMIN_1;

        vm.prank(ORGANIZATION_ID);
        organizationRegistry.createOrganization(ORGANIZATION_NAME, longURI, admins);

        IOrganizationRegistry.Organization memory org = organizationRegistry.getOrganization(ORGANIZATION_ID);
        assertEq(org.metadataURI, longURI);
    }

    function test_UpdateOrganization_SameName() public {
        createTestOrganization();

        // Update with the same name should work
        vm.prank(ORGANIZATION_ID);
        organizationRegistry.updateOrganization(ORGANIZATION_ID, ORGANIZATION_NAME, UPDATED_METADATA_URI);

        IOrganizationRegistry.Organization memory org = organizationRegistry.getOrganization(ORGANIZATION_ID);
        assertEq(org.name, ORGANIZATION_NAME);
        assertEq(org.metadataURI, UPDATED_METADATA_URI);
    }

    function test_UpdateOrganization_SameMetadataURI() public {
        createTestOrganization();

        // Update with the same metadata URI should work
        vm.prank(ORGANIZATION_ID);
        organizationRegistry.updateOrganization(ORGANIZATION_ID, UPDATED_NAME, ORGANIZATION_METADATA_URI);

        IOrganizationRegistry.Organization memory org = organizationRegistry.getOrganization(ORGANIZATION_ID);
        assertEq(org.name, UPDATED_NAME);
        assertEq(org.metadataURI, ORGANIZATION_METADATA_URI);
    }

    function test_AddAdministrator_MultipleSequential() public {
        createTestOrganizationWithoutAdmins();

        // Add multiple administrators sequentially
        address[] memory newAdmins = new address[](3);
        newAdmins[0] = ADMIN_1;
        newAdmins[1] = ADMIN_2;
        newAdmins[2] = ADMIN_3;

        for (uint i = 0; i < newAdmins.length; i++) {
            vm.prank(ORGANIZATION_ID);
            organizationRegistry.addAdministrator(ORGANIZATION_ID, newAdmins[i]);
        }

        // Verify all were added
        for (uint i = 0; i < newAdmins.length; i++) {
            assertTrue(organizationRegistry.isAdministrator(ORGANIZATION_ID, newAdmins[i]));
        }

        IOrganizationRegistry.Organization memory org = organizationRegistry.getOrganization(ORGANIZATION_ID);
        assertEq(org.administrators.length, 4); // 3 added + 1 original
    }

    function test_RemoveAdministrator_AllButOne() public {
        createTestOrganization();

        // Remove all administrators except the organization itself
        vm.prank(ORGANIZATION_ID);
        organizationRegistry.removeAdministrator(ORGANIZATION_ID, ADMIN_1);

        vm.prank(ORGANIZATION_ID);
        organizationRegistry.removeAdministrator(ORGANIZATION_ID, ADMIN_2);

        // Verify only organization itself remains
        assertFalse(organizationRegistry.isAdministrator(ORGANIZATION_ID, ADMIN_1));
        assertFalse(organizationRegistry.isAdministrator(ORGANIZATION_ID, ADMIN_2));
        assertTrue(organizationRegistry.isAdministrator(ORGANIZATION_ID, ORGANIZATION_ID));

        IOrganizationRegistry.Organization memory org = organizationRegistry.getOrganization(ORGANIZATION_ID);
        assertEq(org.administrators.length, 1);
        assertEq(org.administrators[0], ORGANIZATION_ID);
    }

    function test_RemoveAdministrator_RemoveAll() public {
        createTestOrganizationWithoutAdmins();

        // Remove the only administrator (organization itself)
        vm.prank(ORGANIZATION_ID);
        organizationRegistry.removeAdministrator(ORGANIZATION_ID, ORGANIZATION_ID);

        // Verify no administrators remain
        assertFalse(organizationRegistry.isAdministrator(ORGANIZATION_ID, ORGANIZATION_ID));

        IOrganizationRegistry.Organization memory org = organizationRegistry.getOrganization(ORGANIZATION_ID);
        assertEq(org.administrators.length, 0);
    }

    function test_IsAdministrator_AfterDeletion() public {
        createTestOrganization();

        // Delete organization
        vm.prank(ORGANIZATION_ID);
        organizationRegistry.deleteOrganization(ORGANIZATION_ID);

        // isAdministrator should revert for deleted organization
        vm.expectRevert(IOrganizationRegistry.OrganizationNotFound.selector);
        organizationRegistry.isAdministrator(ORGANIZATION_ID, ADMIN_1);
    }

    function test_UpdateOrganization_AfterDeletion() public {
        createTestOrganization();

        // Delete organization
        vm.prank(ORGANIZATION_ID);
        organizationRegistry.deleteOrganization(ORGANIZATION_ID);

        // Update should fail for deleted organization - modifier checks admin first
        vm.prank(ORGANIZATION_ID);
        vm.expectRevert(IOrganizationRegistry.NotAdministrator.selector);
        organizationRegistry.updateOrganization(ORGANIZATION_ID, UPDATED_NAME, UPDATED_METADATA_URI);
    }

    function test_AddAdministrator_AfterDeletion() public {
        createTestOrganization();

        // Delete organization
        vm.prank(ORGANIZATION_ID);
        organizationRegistry.deleteOrganization(ORGANIZATION_ID);

        // Add administrator should fail for deleted organization
        vm.prank(ORGANIZATION_ID);
        vm.expectRevert(IOrganizationRegistry.OrganizationNotFound.selector);
        organizationRegistry.addAdministrator(ORGANIZATION_ID, ADMIN_3);
    }

    function test_RemoveAdministrator_AfterDeletion() public {
        createTestOrganization();

        // Delete organization
        vm.prank(ORGANIZATION_ID);
        organizationRegistry.deleteOrganization(ORGANIZATION_ID);

        // Remove administrator should fail for deleted organization - modifier checks admin first
        vm.prank(ORGANIZATION_ID);
        vm.expectRevert(IOrganizationRegistry.NotAdministrator.selector);
        organizationRegistry.removeAdministrator(ORGANIZATION_ID, ADMIN_1);
    }

    // ========== Gas Optimization Tests ==========

    function test_GasUsage_CreateOrganization() public {
        address[] memory admins = new address[](2);
        admins[0] = ADMIN_1;
        admins[1] = ADMIN_2;

        uint256 gasBefore = gasleft();
        vm.prank(ORGANIZATION_ID);
        organizationRegistry.createOrganization(ORGANIZATION_NAME, ORGANIZATION_METADATA_URI, admins);
        uint256 gasUsed = gasBefore - gasleft();

        // Gas usage should be reasonable (this is more of a benchmark than a strict test)
        assertTrue(gasUsed > 0);
        assertTrue(gasUsed < 500000); // Reasonable upper bound
    }

    function test_GasUsage_AddManyAdministrators() public {
        createTestOrganizationWithoutAdmins();

        // Add 10 administrators and measure gas usage
        uint256 totalGas = 0;
        for (uint160 i = 1; i <= 10; i++) {
            uint256 gasBefore = gasleft();
            vm.prank(ORGANIZATION_ID);
            organizationRegistry.addAdministrator(ORGANIZATION_ID, address(i));
            uint256 gasUsed = gasBefore - gasleft();
            totalGas += gasUsed;
        }

        // Gas usage should scale reasonably
        assertTrue(totalGas > 0);
        assertTrue(totalGas < 2000000); // Reasonable upper bound for 10 operations
    }

    // ========== Fuzz Tests ==========

    function testFuzz_CreateOrganization_ValidName(string calldata name) public {
        vm.assume(bytes(name).length > 0);
        vm.assume(bytes(name).length < 1000); // Reasonable limit

        address[] memory admins = new address[](0);
        address orgId = address(uint160(uint256(keccak256(abi.encode(name)))));

        vm.prank(orgId);
        organizationRegistry.createOrganization(name, ORGANIZATION_METADATA_URI, admins);

        assertTrue(organizationRegistry.exists(orgId));
        IOrganizationRegistry.Organization memory org = organizationRegistry.getOrganization(orgId);
        assertEq(org.name, name);
    }

    function testFuzz_CreateOrganization_ValidMetadataURI(string calldata metadataURI) public {
        vm.assume(bytes(metadataURI).length < 1000); // Reasonable limit

        address[] memory admins = new address[](0);
        address orgId = address(uint160(uint256(keccak256(abi.encode(metadataURI)))));

        vm.prank(orgId);
        organizationRegistry.createOrganization(ORGANIZATION_NAME, metadataURI, admins);

        assertTrue(organizationRegistry.exists(orgId));
        IOrganizationRegistry.Organization memory org = organizationRegistry.getOrganization(orgId);
        assertEq(org.metadataURI, metadataURI);
    }

    function testFuzz_AddAdministrator_ValidAddress(address admin) public {
        vm.assume(admin != address(0));
        vm.assume(admin != ORGANIZATION_ID);

        createTestOrganizationWithoutAdmins();

        vm.prank(ORGANIZATION_ID);
        organizationRegistry.addAdministrator(ORGANIZATION_ID, admin);

        assertTrue(organizationRegistry.isAdministrator(ORGANIZATION_ID, admin));
    }

    function testFuzz_IsAdministrator_NonExistentOrganization(address orgId, address account) public {
        vm.assume(orgId != ORGANIZATION_ID); // Avoid collision with test organization

        // Should revert for non-existent organization
        vm.expectRevert(IOrganizationRegistry.OrganizationNotFound.selector);
        organizationRegistry.isAdministrator(orgId, account);
    }
}
