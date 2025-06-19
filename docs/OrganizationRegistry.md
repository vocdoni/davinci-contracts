# OrganizationRegistry

## Summary

The OrganizationRegistry is a decentralized registry contract that manages organizations within the Vocdoni protocol. It provides a simple yet powerful system for creating and managing organizations with multiple administrators. Each organization is identified by a unique address and contains metadata that can be used to store additional information about the organization. The contract implements a role-based access control system where only administrators can modify organization data.

## Contract Architecture

### Core Components

1. **Organization**: Each organization has a name, metadata URI, and a list of administrators
2. **Organization Administrator System**: Multiple administrators per organizations with optional access control

### State Variables

- `organizations`: Main mapping storing all organization data (address → Organization struct)
- `organizationCount`: Tracks the total number of active organizations
- `token`: Token used for covering process costs

### Organization Structure

```solidity
struct Organization {
    string name; // Organization's display name
    string metadataURI; // URI for additional metadata (e.g., IPFS hash)
    address[] administrators; // Administrator addresses with permissions
}
```

## Functions

### Organization Management Functions

#### `createOrganization(address id, string calldata name, string calldata metadataURI, address[] calldata administrators)`

- **Access**: Public (anyone can create an organization)
- **Purpose**: Creates a new organization with specified parameters
- **Parameters**:
    - `name`: Display name for the organization
    - `metadataURI`: URI pointing to additional metadata
    - `administrators`: Initial list of administrator addresses
- **Process**:
    1. Validates name is not empty
    2. Checks organization doesn't already exist
    3. Creates organization with provided data
    4. Adds all specified administrators
    5. Automatically adds msg.sender as administrator
    6. Increments organization counter
- **Events**: Emits `OrganizationCreated(msg.sender)`
- **Errors**:
    - `InvalidOrganizationID`: If ID is zero address
    - `InvalidOrganizationName`: If name is empty
    - `OrganizationAlreadyExists`: If organization already exists
    - `InvalidAddress`: If any administrator address is zero

#### `updateOrganization(address id, string calldata name, string calldata metadataURI)`

- **Access**: Organization administrators only
- **Purpose**: Updates organization's name and metadata
- **Parameters**:
    - `id`: Organization's unique identifier
    - `name`: New display name
    - `metadataURI`: New metadata URI
- **Validations**:
    - Caller must be administrator
    - Name cannot be empty
    - Metadata URI cannot be empty
    - Organization must exist
- **Events**: Emits `OrganizationUpdated(id, msg.sender)`
- **Errors**:
    - `NotAdministrator`: If caller is not an administrator
    - `InvalidOrganizationName`: If name is empty
    - `InvalidMetadataURI`: If metadata URI is empty
    - `OrganizationNotFound`: If organization doesn't exist

#### `deleteOrganization(address id)`

- **Access**: Organization itself
- **Purpose**: Permanently deletes an organization
- **Parameters**:
    - `id`: Organization's unique identifier
- **Process**:
    1. Validates caller is the organization
    2. Validates organization exists
    3. Deletes all organization data
    4. Decrements organization counter
- **Note**: This is irreversible and removes the organization entirely
- **Errors**:
    - `Unauthorized`: If caller is not the organization
    - `OrganizationNotFound`: If organization doesn't exist

### Administrator Management Functions

#### `addAdministrator(address id)`

- **Access**: Organization only
- **Purpose**: Adds a new administrator to the organization
- **Parameters**:
    - `id`: Organization's unique identifier
    - `administrator`: Address to grant administrator privileges
- **Validations**:
    - Caller must be the organization
    - Organization must exist
    - Administrator address cannot be zero and not already added
- **Note**: No event emitted for administrator changes
- **Errors**:
    - `NotAdministrator`: If caller is not an administrator
    - `OrganizationNotFound`: If organization doesn't exist
    - `InvalidAddress`: If administrator address is zero

#### `removeAdministrator(address id, address administrator)`

- **Access**: Organization administrators and organization itself
- **Purpose**: Removes administrator privileges
- **Parameters**:
    - `id`: Organization's unique identifier
    - `administrator`: Address to revoke privileges from
- **Validations**:
    - Caller must be administrator
    - If the caller is the organization itself, can remove any administrator from the organization
    - Organization must exist
    - Administrator address cannot be zero
- **Errors**:
    - `NotAdministrator`: If caller is not an administrator
    - `OrganizationNotFound`: If organization doesn't exist
    - `InvalidAddress`: If administrator address is zero

### Query Functions

#### `getOrganization(address id)`

- **Access**: Public view function
- **Purpose**: Retrieves organization's basic information
- **Parameters**:
    - `id`: Organization's unique identifier
- **Returns**: Tuple of (name, metadataURI)
- **Note**: Returns empty strings if organization doesn't exist

#### `isAdministrator(address id, address account)`

- **Access**: Public view function
- **Purpose**: Checks if an address has administrator privileges
- **Parameters**:
    - `id`: Organization's unique identifier
    - `account`: Address to check
- **Returns**: Boolean indicating administrator status

#### `exists(address id)`

- **Access**: Public view function
- **Purpose**: Checks if an organization exists
- **Parameters**:
    - `id`: Organization's unique identifier
- **Returns**: Boolean indicating existence
- **Implementation**: Checks if organization name length > 0

## What to expect

### Normal Operations

1. **Organization Creation**: Anyone can create an organization with themselves as administrator
2. **Multi-Admin Management**: Organizations can have multiple administrators with equal permissions
3. **Metadata Flexibility**: Organizations can store arbitrary metadata via URI (e.g., IPFS)
4. **Dynamic Updates**: Administrators can update organization details at any time
5. **Administrator Rotation**: Administrators can be added/removed
6. **Self-Removal**: Administrators can remove their own privileges

### Access Control

1. **Public Creation**: No restrictions on who can create organizations
2. **Protected Modifications**: Only administrators can modify organization data
3. **Equal Privileges**: All administrators have the same permissions by default
4. **Automatic Admin**: Organization creator is automatically added as administrator

### Security Guarantees

1. **Duplicate Organizations**: Cannot create organization with existing ID
2. **Unauthorized Modifications**: Non-administrators cannot modify organization data
3. **Zero Address Operations**: Cannot use zero address as organization ID or administrator
4. **Empty Names**: Cannot create or update organization with empty name
5. **Cross-Organization Access**: Administrators of one organization cannot modify another
6. **Global Admin**: No super-admin or owner role can control organizations

### Data Integrity

1. **Name Preservation**: Organization names cannot be set to empty string
2. **ID Immutability**: Organization IDs cannot be changed after creation
3. **Partial Deletion**: Cannot partially delete organization data

## Usage Examples

### 1. Creating an Organization

```solidity
// Create organization with multiple administrators
string memory name = "Vocdoni DAO";
string memory metadataURI = "ipfs://QmXxx...";  // IPFS hash with extended metadata
address[] memory admins = [0xAdmin1..., 0xAdmin2...];

organizationRegistry.createOrganization(name, metadataURI, admins);
// msg.sender is automatically added as administrator
```

### 2. Updating Organization Information

```solidity
// Update organization name and metadata
address orgId = 0x123...;
string memory newName = "Vocdoni Foundation";
string memory newMetadataURI = "ipfs://QmYyy...";

// Must be called by an administrator
organizationRegistry.updateOrganization(orgId, newName, newMetadataURI);
```

### 3. Managing Administrators

```solidity
// Add a new administrator
address orgId = 0x123...;
address newAdmin = 0x456...;
organizationRegistry.addAdministrator(orgId, newAdmin);

// Remove an administrator
address adminToRemove = 0x789...;
organizationRegistry.removeAdministrator(orgId, adminToRemove);

// Check administrator status
bool isAdmin = organizationRegistry.isAdministrator(orgId, newAdmin);
```

### 4. Querying Organization Data

```solidity
// Get organization information
address orgId = 0x123...;
(string memory name, string memory metadataURI) = organizationRegistry.getOrganization(orgId);

// Get total organization count
uint32 totalOrgs = organizationRegistry.organizationCount();
```

### 5. Deleting an Organization

```solidity
// Permanently delete organization (administrator only)
address orgId = 0x123...;
organizationRegistry.deleteOrganization(orgId);
// This action is irreversible!
```

## Best Practices

### Administrator Management

1. **Maintain Multiple Admins**: Always keep at least 2-3 administrators to prevent lockout
2. **Implement Governance**: Consider implementing time delays or multi-sig for critical operations
3. **Admin Key Management**: Use hardware wallets or multi-sig for administrator accounts

### Metadata Management

1. **Use IPFS**: Store extended metadata on IPFS and reference via URI
2. **Version Metadata**: Include version information in metadata for compatibility
3. **Validate URIs**: Implement off-chain validation for metadata URI format

### Metadata Standards

```json
// Example metadata structure (stored on IPFS)
{
    "version": "1.0",
    "languages": [],
    "name": {
        "default": "Example Organization"
    },
    "description": {
        "default": "This is an example account metadata description."
    },
    "newsFeed": {
        "default": "https://example.org/news"
    },
    "media": {
        "avatar": "https://example.org/media/avatar.png",
        "header": "https://example.org/media/header.png",
        "logo": "https://example.org/media/logo.png"
    },
    "meta": {
        "createdBy": "admin@example.org",
        "tags": ["demo", "test", "organization"]
    }
}
```

## Limitations and Considerations

1. **No Built-in Recovery**: Lost administrator access cannot be recovered by the contract
2. **No Hierarchy**: All administrators have equal permissions
3. **No Time Locks**: Changes take effect immediately
4. **Storage Costs**: On-chain storage of names may be expensive for long names
5. **No Enumeration**: Cannot list all organizations efficiently
