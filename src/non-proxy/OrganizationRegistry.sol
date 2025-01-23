// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity 0.8.28;

/**
 * @title OrganizationRegistry
 * @author Vocdoni Association
 * @notice The OrganizationRegistry contract is a registry of organizations.
 */
contract OrganizationRegistry {
    /**
     * @notice Emitted when a new organization is created
     * @param id The organization's unique identifier
     * @param creator The address of the organization's creator
     */
    event OrganizationCreated(address indexed id, address indexed creator);

    /**
     * @notice Emitted when an organization is updated
     * @param id The organization's unique identifier
     * @param updater The address of the organization's updater
     */
    event OrganizationUpdated(address indexed id, address indexed updater);

    /**
     * @notice Organization structure containing the organization's data
     * @param id The organization's unique identifier
     * @param name The organization's name
     * @param metadataURI The organization's metadata URI that can be used to store additional information
     * @param administrators The list of administrators of the organization
     */
    struct Organization {
        string name;
        string metadataURI;
        mapping(address => bool) administrators;
    }

    /**
     * @notice Modifier that checks if the sender is an administrator of the organization
     * @param id The organization's unique identifier
     */
    modifier onlyAdministrator(address id) {
        require(organizations[id].administrators[msg.sender], "OrganizationRegistry: not an administrator");
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
     * @notice Creates a new organization
     * @param id The organization's unique identifier
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
        require(id != address(0), "OrganizationRegistry: invalid id");
        require(bytes(name).length > 0, "OrganizationRegistry: invalid name");

        require(bytes(organizations[id].name).length == 0, "OrganizationRegistry: organization already exists");

        Organization storage organization = organizations[id];
        organization.name = name;
        organization.metadataURI = metadataURI;

        if (administrators.length > 0) {
            for (uint256 i = 0; i < administrators.length; i++) {
                require(administrators[i] != address(0), "OrganizationRegistry: invalid administrator address");
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
    function updateOrganization(address id, string calldata name, string calldata metadataURI)
        public
        onlyAdministrator(id)
    {
        require(bytes(name).length > 0, "OrganizationRegistry: invalid name");
        require(bytes(metadataURI).length > 0, "OrganizationRegistry: invalid metadataURI");
        require(bytes(organizations[id].name).length > 0, "OrganizationRegistry: organization does not exist");

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
        require(bytes(organizations[id].name).length > 0, "OrganizationRegistry: organization does not exist");
        require(administrator != address(0), "OrganizationRegistry: invalid administrator address");
        organizations[id].administrators[administrator] = true;
    }

    /**
     * @notice Removes an administrator from an organization
     * @param id The organization's unique identifier
     * @param administrator The address of the administrator to remove
     */
    function removeAdministrator(address id, address administrator) public onlyAdministrator(id) {
        require(bytes(organizations[id].name).length > 0, "OrganizationRegistry: organization does not exist");
        require(administrator != address(0), "OrganizationRegistry: invalid administrator address");
        organizations[id].administrators[administrator] = false;
    }

    /**
     * @notice Deletes an organization
     * @param id The ID of the organization to delete
     */
    function deleteOrganization(address id) public {
        require(bytes(organizations[id].name).length > 0, "OrganizationRegistry: organization does not exist");
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
}
