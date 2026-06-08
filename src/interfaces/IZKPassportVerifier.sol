// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity 0.8.28;

struct ProofVerificationData {
    bytes32 vkeyHash;
    bytes proof;
    bytes32[] publicInputs;
}

/// @dev oprfPubKeyHash temporarily disabled in the deployed RootVerifier to match its on-chain ABI.
struct ServiceConfig {
    uint256 validityPeriodInSeconds;
    string domain;
    string scope;
    bool devMode;
}

struct ProofVerificationParams {
    bytes32 version;
    ProofVerificationData proofVerificationData;
    bytes committedInputs;
    ServiceConfig serviceConfig;
}

/// @notice Interface for the ZKPassport RootVerifier at 0x1D000001000EFD9a6371f4d90bB8920D5431c0D8.
///         Handles all outer_evm_count_N circuit variants without hardcoding a circuit count.
interface IZKPassportVerifier {
    /// @return valid            True if the proof is cryptographically valid.
    /// @return uniqueIdentifier Scoped nullifier derived from the passport identity.
    /// @return helper           VerifierHelper address for this version; use to decode committedInputs.
    function verify(ProofVerificationParams calldata params)
        external
        view
        returns (bool valid, bytes32 uniqueIdentifier, address helper);
}
