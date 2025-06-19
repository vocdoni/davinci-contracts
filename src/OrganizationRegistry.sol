// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import "./IOrganizationRegistry.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

/**
 * @title OrganizationRegistry
 * @author Vocdoni Association
 * @notice The OrganizationRegistry contract is a registry of organizations.
 * @dev Uses OpenZeppelin's Initializable contract to manage the contract's initialization.
 */
contract OrganizationRegistry is IOrganizationRegistry, Initializable, UUPSUpgradeable, OwnableUpgradeable {
    /**
     * @notice Modifier that checks if the sender is an administrator of the organization
     * @param id The organization's unique identifier
     */
    modifier onlyAdministrator(address id) {
        if (!organizations[id].administrators[msg.sender]) {
            revert NotAdministrator();
        }
        _;
    }
    /**
     * @notice Mapping of organizations IDs to their respective organization data
     */
    mapping(address => Organization) public organizations;
    /**
     * @notice Tracks the total number of organizations
     */
    uint32 public organizationCount;

    /**
     * @notice Initializes the contract
     */
    function initialize() public initializer {
        __Ownable_init(msg.sender);
        __UUPSUpgradeable_init();
    }

    /**
     * @notice Creates a new organization
     * @param id The organization's address
     * @param name The organization's name
     * @param metadataURI The organization's metadata URI that can be used to store additional information
     * @param administrators The list of administrators of the organization
     * @dev Checks for organization existence by verifying that the organization's name is not empty
     * @dev msg.sender is added as an administrator by default
     */
    function createOrganization(
        address id,
        string calldata name,
        string calldata metadataURI,
        address[] calldata administrators
    ) public {
        if (id == address(0)) {
            revert InvalidOrganizationID();
        }
        if (bytes(name).length <= 0) {
            revert InvalidOrganizationName();
        }
        if (bytes(organizations[id].name).length != 0) {
            revert OrganizationAlreadyExists();
        }
        Organization storage organization = organizations[id];
        organization.name = name;
        organization.metadataURI = metadataURI;
        if (administrators.length > 0) {
            for (uint256 i = 0; i < administrators.length; i++) {
                if (administrators[i] == address(0)) {
                    revert InvalidAddress();
                }
                organization.administrators[administrators[i]] = true;
            }
        }
        organization.administrators[msg.sender] = true;
        organizationCount++;
        emit OrganizationCreated(id, msg.sender);
    }

    /**
     * @notice Retrieves an organization's data
     * @param id The organization's unique identifier
     * @return name The organization's name
     * @return metadataURI The organization's metadata URI that can be used to store additional information
     */
    function getOrganization(address id) public view returns (string memory, string memory) {
        Organization storage organization = organizations[id];
        return (organization.name, organization.metadataURI);
    }

    /**
     * @notice Updates an organization's data
     * @param id The organization's unique identifier
     * @param name The organization's name
     * @param metadataURI The organization's metadata URI that can be used to store additional information
     */
    function updateOrganization(
        address id,
        string calldata name,
        string calldata metadataURI
    ) public onlyAdministrator(id) {
        if (bytes(name).length <= 0) {
            revert InvalidOrganizationName();
        }
        if (bytes(metadataURI).length <= 0) {
            revert InvalidMetadataURI();
        }
        if (bytes(organizations[id].name).length <= 0) {
            revert OrganizationNotFound();
        }

        Organization storage organization = organizations[id];
        organization.name = name;
        organization.metadataURI = metadataURI;

        emit OrganizationUpdated(id, msg.sender);
    }

    /**
     * @notice Adds an administrator to an organization
     * @param id The organization's unique identifier
     * @param administrator The address of the administrator to add
     */
    function addAdministrator(address id, address administrator) public onlyAdministrator(id) {
        if (bytes(organizations[id].name).length <= 0) {
            revert OrganizationNotFound();
        }
        if (administrator == address(0)) {
            revert InvalidAddress();
        }
        address[] storage admins = organizations[id].administrators;
        for (uint256 i = 0; i < admins.length; i++) {
            if (admins[i] == administrator) {
                revert AlreadyAdministrator();
            }
        }
        admins.push(administrator);
        emit AdministratorAdded(id, administrator);
    }

    /**
     * @notice Removes an administrator from an organization
     * @param id The organization's unique identifier
     * @param administrator The address of the administrator to remove
     */
    function removeAdministrator(address id, address administrator) public onlyAdministrator(id) {
        if (bytes(organizations[id].name).length <= 0) {
            revert OrganizationNotFound();
        }
        if (administrator == address(0)) {
            revert InvalidAddress();
        }
        bool found = false;
        address[] storage admins = organizations[id].administrators;
        for (uint256 i = 0; i < admins.length; i++) {
            if (admins[i] == administrator) {
                found = true;
                admins[i] = admins[admins.length - 1];
                admins.pop();
                break;
            }
        }
        if (!found) {
            revert NotAdministrator();
        }
        emit AdministratorRemoved(id, administrator, msg.sender);
    }

    /**
     * @notice Deletes an organization
     * @param id The ID of the organization to delete
     */
    function deleteOrganization(address id) public onlyOwner {
        if (bytes(organizations[id].name).length <= 0) {
            revert OrganizationNotFound();
        }
        delete organizations[id];
        organizationCount--;
    }

    /**
     * @notice Checks if an account is an administrator of an organization
     * @param id The organization's unique identifier
     * @param account The address of the account to check
     * @return true if the account is an administrator, false otherwise
     */
    function isAdministrator(address id, address account) public view returns (bool) {
        return organizations[id].administrators[account];
    }

    /**
     * @notice Checks if an organization exists
     * @param id The organization's unique identifier
     * @return true if the organization exists, false otherwise
     */
    function exists(address id) public view returns (bool) {
        return bytes(organizations[id].name).length > 0;
    }

    // @notice Internal function to check if a user is an administrator of an organization
    // @param id The organization's unique identifier
    function _isAdmin(address id) internal view returns (bool) {
        address[] storage admins = organizations[id].administrators;
        for (uint256 i = 0; i < admins.length; i++) {
            if (admins[i] == msg.sender) return true;
        }
        return false;
    }

    function _authorizeUpgrade(address) internal override onlyOwner {}
}
