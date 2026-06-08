// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity 0.8.28;

import {
    InternalLeanIMT,
    LeanIMTData
} from "zk-kit.solidity/packages/lean-imt/contracts/InternalLeanIMT.sol";

import {ICensusValidator} from "./interfaces/ICensusValidator.sol";
import {IProofVerifier} from "./verifiers/IProofVerifier.sol";

/// @notice Census contract for Vocdoni zkPassport voting.
///
/// @dev The backend (passport-prover server) aggregates a zkPassport outer proof
///      and calls register() here. The contract verifies the UltraHonk proof using
///      an on-chain Barretenberg verifier (e.g. OuterCount4) rather than trusting
///      the backend to produce valid data.
///
///      Outer proof public inputs layout (outer_evm_count_N, N = number of inner circuits):
///        [0] certificate_registry_root
///        [1] circuit_registry_root
///        [2] current_date
///        [3] service_scope     — hash of the service scope string
///        [4] service_subscope  — hash of the service subscope string
///        [5..len-4] param_commitments — one per disclosure circuit
///        [len-3] nullifier_type
///        [len-2] scoped_nullifier  ← nullifier used for double-registration prevention
///        [len-1] oprf_pk_hash
///
///      For outer_evm_count_4 (9 inputs): NULLIFIER_INPUT_INDEX = 7
///      For outer_evm_count_5 (10 inputs): NULLIFIER_INPUT_INDEX = 8
///
///      Residual trust: `account` is passed by the backend and extracted from the
///      inner bind_evm disclosure proof (not directly verifiable from the outer public
///      inputs). V3 improvement: encode address(this) as service_subscope so
///      publicInputs[4] can be verified on-chain against the caller's address.
///
///      The TRUSTED_BACKEND restriction is kept to prevent address front-running:
///      only the known backend relayer may provide the account parameter.
contract ZKPassportCensus is ICensusValidator {
    using InternalLeanIMT for LeanIMTData;

    // ====================================================
    // Constants
    // ====================================================

    /// @dev Index of scoped_nullifier in the outer proof public inputs (constructor-set).
    uint256 public immutable NULLIFIER_INPUT_INDEX;

    /// @dev Expected number of public inputs (constructor-set for the deployed OuterCount variant).
    uint256 public immutable EXPECTED_PUBLIC_INPUTS;

    // ====================================================
    // State
    // ====================================================

    IProofVerifier public immutable VERIFIER;
    address public immutable TRUSTED_BACKEND;

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

    error NotTrustedBackend();
    error InvalidProof();
    error InvalidPublicInputCount();
    error NullifierAlreadyUsed();
    error AlreadyRegistered();

    // ====================================================
    // Constructor
    // ====================================================

    /// @param verifier              On-chain UltraHonk verifier for the outer zkPassport circuit.
    /// @param backend               Address of the trusted backend relayer that calls register().
    /// @param expectedPublicInputs  Total public input count for the deployed OuterCount variant
    ///                              (9 for outer_evm_count_4, 10 for outer_evm_count_5, etc.).
    /// @param nullifierInputIndex   Index of scoped_nullifier in the public inputs array
    ///                              (always expectedPublicInputs - 2).
    constructor(
        IProofVerifier verifier,
        address backend,
        uint256 expectedPublicInputs,
        uint256 nullifierInputIndex
    ) {
        VERIFIER = verifier;
        TRUSTED_BACKEND = backend;
        EXPECTED_PUBLIC_INPUTS = expectedPublicInputs;
        NULLIFIER_INPUT_INDEX = nullifierInputIndex;
        _currentRoot = _tree._root();
    }

    // ====================================================
    // Registration
    // ====================================================

    /// @notice Register a voter whose zkPassport outer proof has been verified.
    /// @param account       Voter's Ethereum address (from the bind_evm inner proof).
    /// @param proof         Outer UltraHonk proof bytes from the prover-cli.
    /// @param publicInputs  8 public inputs of the outer proof (see layout above).
    function register(
        address account,
        bytes calldata proof,
        bytes32[] calldata publicInputs
    ) external {
        if (msg.sender != TRUSTED_BACKEND) revert NotTrustedBackend();
        if (publicInputs.length != EXPECTED_PUBLIC_INPUTS) revert InvalidPublicInputCount();
        if (!VERIFIER.verify(proof, publicInputs)) revert InvalidProof();

        uint256 nullifier = uint256(publicInputs[NULLIFIER_INPUT_INDEX]);
        if (nullifierUsed[nullifier]) revert NullifierAlreadyUsed();
        if (weightOf[account] != 0) revert AlreadyRegistered();

        nullifierUsed[nullifier] = true;
        weightOf[account] = 1;

        uint256 leaf = uint256(uint160(account));
        uint256 newRoot = _insertAndRotateRoot(leaf);

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
