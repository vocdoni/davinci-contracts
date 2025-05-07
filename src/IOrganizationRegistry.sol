// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

/**
 * @title IOrganizationRegistry
 * @author Vocdoni Association
 * @notice The Organization Registry interface contract.
 */
interface IOrganizationRegistry {
    /// EVENTS ///

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

    /// ERRORS ///

    /**
     * @notice NotAdministrator error is emitted when the caller is not an administrator of the organization.
     */
    error NotAdministrator();
    /**
     * @notice InvalidOrganizationID error is emitted when an organization ID is not valid.
     */
    error InvalidOrganizationID();
    /**
     * InvalidOrganizationName error is emitted when an organization name is not valid.
     */
    error InvalidOrganizationName();
    /**
     * OrganizationAlreadyExists error is emitted when an organization already exists.
     */
    error OrganizationAlreadyExists();
    /**
     * InvalidAddress error is emitted if an address is invalid.
     */
    error InvalidAddress();
    /**
     * OrganizationNotFound error is emitted when an organization is not found in the organization registry
     */
    error OrganizationNotFound();
    /**
     * InvalidMetadataURI error is emitted when the metadata URI is invalid
     */
    error InvalidMetadataURI();

    /// STRUCTS ///

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

    /// GETTERS ///

    /**
     * @notice Retrieves an organization's data
     * @param id The organization's unique identifier
     * @return name The organization's name
     * @return metadataURI The organization's metadata URI that can be used to store additional information
     */
    function getOrganization(address id) external view returns (string memory, string memory);

    /**
     * @notice Checks if an account is an administrator of an organization
     * @param id The organization's unique identifier
     * @param account The address of the account to check
     * @return true if the account is an administrator, false otherwise
     */
    function isAdministrator(address id, address account) external view returns (bool);

    /**
     * @notice Checks if an organization exists
     * @param id The organization's unique identifier
     * @return true if the organization exists, false otherwise
     */
    function exists(address id) external view returns (bool);

    /// SETTERS ///

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
    ) external;

    /**
     * @notice Updates an organization's data
     * @param id The organization's unique identifier
     * @param name The organization's name
     * @param metadataURI The organization's metadata URI that can be used to store additional information
     */
    function updateOrganization(address id, string calldata name, string calldata metadataURI) external;

    /**
     * @notice Adds an administrator to an organization
     * @param id The organization's unique identifier
     * @param administrator The address of the administrator to add
     */
    function addAdministrator(address id, address administrator) external;

    /**
     * @notice Removes an administrator from an organization
     * @param id The organization's unique identifier
     * @param administrator The address of the administrator to remove
     */
    function removeAdministrator(address id, address administrator) external;

    /**
     * @notice Deletes an organization
     * @param id The ID of the organization to delete
     */
    function deleteOrganization(address id) external;
}
