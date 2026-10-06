// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity >=0.8.4 <0.9.0;

/// @title ICouncilManager
/// @notice The CouncilManager surface a DAVINCI `CouncilAdapter` uses (architecture §3.1), protocol
///         v2. Self-contained on purpose: copy this file verbatim into another repository.
///
///         Flow: the organizer allows the adapter and authorizes the process creator for a `Live`
///         ceremony; the adapter `bindProcess`es a process (receiving the request id and the
///         ceremony key `P`), later `submitRequest`s the process's final accumulator (admitted
///         while the decryption gate is still closed), and polls `getPlaintexts` until `ready`.
///         `isDecryptionOpen` is the ceremony-wide decryption gate (protocol §8.7): no partial or
///         combine is accepted before it, and the DAVINCI registry publishes no result of a bound
///         process before it either, not even an all-zero vector that never needed a request.
///
///         Points are circomlib twisted Edwards (TE) `(x, y)` pairs, each `< p` (BN254 scalar
///         field); `cts[k] = [C1.x, C1.y, C2.x, C2.y]`.
interface ICouncilManager {
    /// @notice Bind `processId` to ceremony `cid`. `msg.sender` must be an allowed adapter of the
    ///         ceremony and `creator` an authorized creator; the ceremony must be `Live`; one
    ///         binding per `(msg.sender, processId)`, ever.
    /// @return requestId keccak256(abi.encode(keccak256("davinci-dkg-council/v1/request"), chainId, manager,
    ///         cid, adapter, processId)) (protocol §4.5).
    /// @return pkX ceremony public key `P`, TE x.
    /// @return pkY ceremony public key `P`, TE y.
    function bindProcess(bytes12 cid, bytes31 processId, address creator)
        external
        returns (bytes32 requestId, uint256 pkX, uint256 pkY);

    /// @notice Submit the decryption request of a bound process, once. Called by the same adapter
    ///         that bound it; `cid` is cross-checked against the binding record. 1..16 fields,
    ///         every C1/C2 canonical, on curve, in the prime subgroup and not the identity.
    ///         Admitted whether or not the decryption gate is open.
    function submitRequest(bytes12 cid, bytes31 processId, uint256[4][] calldata cts)
        external
        returns (bytes32 requestId);

    /// @notice Combined plaintexts of a submitted request. `values.length == fieldCount`;
    ///         `ready` is true only once every field is combined (uncombined fields read 0).
    function getPlaintexts(bytes32 requestId) external view returns (bool ready, uint256[] memory values);

    /// @notice Request record. `fieldCount == 0` means bound but not yet submitted.
    function getRequestMeta(bytes32 requestId)
        external
        view
        returns (bytes12 cid, uint8 fieldCount, uint16 completedBitmap, uint16 partialBitmap);

    /// @notice Binding record of `(adapter, processId)`.
    function getBinding(address adapter, bytes31 processId)
        external
        view
        returns (bytes12 cid, bytes32 requestId, bool requested);

    /// @notice Ceremony public key `P` in TE; reverts unless the ceremony is `Live`.
    function getPublicKey(bytes12 cid) external view returns (uint256 x, uint256 y);

    /// @notice The decryption gate of ceremony `cid` at `block.timestamp` (protocol §8.7): `Live`
    ///         and either Scheduled past `decryptionOpenAt`, or Manual and opened by the organizer
    ///         or past a nonzero fallback date. Monotonic: once true it stays true. Reverts
    ///         `UnknownCeremony()` for a missing ceremony.
    function isDecryptionOpen(bytes12 cid) external view returns (bool);
}
