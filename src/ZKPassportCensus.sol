// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity 0.8.28;

import {
    InternalLeanIMT,
    LeanIMTData
} from "zk-kit.solidity/packages/lean-imt/contracts/InternalLeanIMT.sol";

import {ICensusValidator} from "./interfaces/ICensusValidator.sol";
import {IZKPassportVerifier, ProofVerificationParams} from "./interfaces/IZKPassportVerifier.sol";
import {IZKPassportHelper} from "./interfaces/IZKPassportHelper.sol";

/// @notice Census contract for Vocdoni zkPassport voting — permissionless on-chain verification.
///
/// @dev Anyone may call register() with a valid zkPassport outer proof. There is no
///      trusted-backend restriction. Security is fully cryptographic:
///
///      1. IZKPassportVerifier.verify() validates the aggregated UltraHonk proof against
///         the on-chain circuit registry (handles any outer_evm_count_N variant).
///      2. IZKPassportHelper.getBoundData(committedInputs).senderAddress must equal `account`,
///         enforcing that the proof binds exactly to that Ethereum address (bind_evm circuit).
///      3. uniqueIdentifier (scoped nullifier) prevents the same passport from registering twice.
///
///      committedInputs MUST NOT include any disclosure circuit data beyond bind_evm —
///      removing all disclosure queries (e.g. gender) from the QR payload ensures the
///      backend can reconstruct committedInputs from just signerAddress + bindChain.
contract ZKPassportCensus is ICensusValidator {
    using InternalLeanIMT for LeanIMTData;

    // ====================================================
    // Immutable config
    // ====================================================

    /// @notice ZKPassport RootVerifier (universal, same address on all supported chains).
    IZKPassportVerifier public immutable VERIFIER;

    /// @notice Service scope string stored for off-chain tooling.
    ///         Scope enforcement is handled by the proof itself (service_scope public input).
    string public SCOPE;

    // ====================================================
    // State
    // ====================================================

    LeanIMTData private _tree;

    mapping(address => uint88) public weightOf;
    mapping(uint256 => bool) public nullifierUsed;

    uint256 private constant ROOT_HISTORY_SIZE = 100;
    uint256 private _currentRoot;
    uint256[ROOT_HISTORY_SIZE] private _historyRoots;
    uint256[ROOT_HISTORY_SIZE] private _historyLastValidBlock;
    uint256 private _historyIndex;
    mapping(uint256 => uint256) private _rootLastValidBlock;

    // ====================================================
    // Events / Errors
    // ====================================================

    event Registered(
        address indexed account,
        uint256 indexed nullifier,
        uint256 newRoot
    );

    error InvalidProof();
    error AddressBindingMismatch();
    error NullifierAlreadyUsed();
    error AlreadyRegistered();

    // ====================================================
    // Constructor
    // ====================================================

    /// @param verifier  IZKPassportVerifier at 0x1D000001000EFD9a6371f4d90bB8920D5431c0D8.
    /// @param scope     Service scope string (e.g. "vocdoni"). Must match the proof's serviceConfig.scope.
    constructor(IZKPassportVerifier verifier, string memory scope) {
        VERIFIER = verifier;
        SCOPE = scope;
        _currentRoot = _tree._root();
    }

    // ====================================================
    // Registration
    // ====================================================

    /// @notice Register a voter using an on-chain verified zkPassport outer proof.
    /// @param account  Voter's Ethereum address — must match the bind_evm committed address.
    /// @param params   Full proof verification parameters (version, proof, public inputs,
    ///                 committedInputs, service config). Constructed by the passport-prover backend.
    function register(
        address account,
        ProofVerificationParams calldata params
    ) external {
        (bool verified, bytes32 uniqueIdentifier, address helper) = VERIFIER.verify(params);
        if (!verified) revert InvalidProof();

        if (
            IZKPassportHelper(helper).getBoundData(params.committedInputs).senderAddress != account
        ) revert AddressBindingMismatch();

        uint256 nullifier = uint256(uniqueIdentifier);
        if (nullifierUsed[nullifier]) revert NullifierAlreadyUsed();
        if (weightOf[account] != 0) revert AlreadyRegistered();

        nullifierUsed[nullifier] = true;
        weightOf[account] = 1;

        uint256 newRoot = _insertAndRotateRoot(uint256(uint160(account)));

        emit WeightChanged(account, 0, 1);
        emit Registered(account, nullifier, newRoot);
    }

    // ====================================================
    // ICensusValidator
    // ====================================================

    function getCensusRoot() external view override returns (uint256) {
        return _currentRoot;
    }

    function getRootBlockNumber(
        uint256 root
    ) external view override returns (uint256) {
        if (root == 0) return 0;
        if (root == _currentRoot) return block.number;
        return _rootLastValidBlock[root];
    }

    function getTotalVotingPowerAtRoot(
        uint256 /* root */
    ) external view override returns (uint256) {
        return _tree.size;
    }

    // ====================================================
    // View helpers
    // ====================================================

    function treeSize() external view returns (uint256) {
        return _tree.size;
    }

    function treeDepth() external view returns (uint256) {
        return _tree.depth;
    }

    // ====================================================
    // Internal
    // ====================================================

    function _insertAndRotateRoot(uint256 leaf) internal returns (uint256 newRoot) {
        newRoot = _tree._insert(leaf);

        uint256 oldRoot = _currentRoot;
        if (oldRoot != 0 && oldRoot != newRoot) {
            uint256 evicted = _historyRoots[_historyIndex];
            if (evicted != 0) {
                delete _rootLastValidBlock[evicted];
            }

            _historyRoots[_historyIndex] = oldRoot;
            _historyLastValidBlock[_historyIndex] = block.number;
            _rootLastValidBlock[oldRoot] = block.number;

            _historyIndex = (_historyIndex + 1) % ROOT_HISTORY_SIZE;
        }

        _currentRoot = newRoot;
    }
}
