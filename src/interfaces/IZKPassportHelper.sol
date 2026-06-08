// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity 0.8.28;

struct BoundData {
    address senderAddress;
    uint256 chainId;
    string customData;
}

/// @notice Interface for the ZKPassport VerifierHelper returned by IZKPassportVerifier.verify().
interface IZKPassportHelper {
    /// @notice Decodes the address and chain binding from the committedInputs of a bind_evm proof.
    function getBoundData(bytes calldata committedInputs)
        external
        pure
        returns (BoundData memory boundData);
}
