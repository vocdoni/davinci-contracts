// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import "./interfaces/IOrganizationRegistry.sol";
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
     * @notice Modifier to check if the caller is an administrator of the organization
     * @param orgId The organization's unique identifier
     */
    modifier onlyAdministrator(address orgId) {
        if (!_isAdmin(orgId)) revert NotAdministrator();
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
     * @notice Initializes the contract using OpenZeppelin's UUPS pattern.
     */
    function initialize() public initializer {
        __Ownable_init(msg.sender);
        __UUPSUpgradeable_init();
    }

    /// @inheritdoc IOrganizationRegistry
    function createOrganization(
        string calldata name,
        string calldata metadataURI,
        address[] calldata administrators
    ) public {
        address id = msg.sender;
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
                if (administrators[i] == address(0)) revert InvalidAddress();
                if (administrators[i] == id) continue;
                organization.administrators.push(administrators[i]);
            }
        }
        organization.administrators.push(id);
        organizationCount++;
        emit OrganizationCreated(id);
    }

    /// @inheritdoc IOrganizationRegistry
    function getOrganization(address id) public view returns (Organization memory) {
        return organizations[id];
    }

    /// @inheritdoc IOrganizationRegistry
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

    /// @inheritdoc IOrganizationRegistry
    function isAdministrator(address id, address account) public view returns (bool) {
        if (bytes(organizations[id].name).length <= 0) {
            revert OrganizationNotFound();
        }
        address[] memory admins = organizations[id].administrators;
        for (uint256 i = 0; i < admins.length; i++) {
            if (admins[i] == account) {
                return true;
            }
        }
        return false;
    }

    /// @inheritdoc IOrganizationRegistry
    function addAdministrator(address id, address administrator) public {
        // only the organization itself can add an administrator
        if (msg.sender != id) revert Unauthorized();
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

    /// @inheritdoc IOrganizationRegistry
    function removeAdministrator(address id, address administrator) public onlyAdministrator(id) {
        // if the caller is the organization itself can remove any admin
        // otherwise, the caller, if admin, can remove himself
        if (msg.sender != id) {
            administrator = msg.sender;
        }
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

    /// @inheritdoc IOrganizationRegistry
    function deleteOrganization(address id) public {
        if (bytes(organizations[id].name).length <= 0) {
            revert OrganizationNotFound();
        }
        if (msg.sender != id) revert Unauthorized();
        delete organizations[id];
        organizationCount--;
    }

    /// @inheritdoc IOrganizationRegistry
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
