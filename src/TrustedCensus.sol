// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity 0.8.28;

import {
    InternalLeanIMT,
    LeanIMTData
} from "zk-kit.solidity/packages/lean-imt/contracts/InternalLeanIMT.sol";

import {ICensusValidator} from "./interfaces/ICensusValidator.sol";

/// @notice Census contract for Vocdoni zkPassport voting.
/// @dev The trusted backend (passport-prover server) is the sole caller of register().
///      It verifies the zkPassport outer proof off-chain and pays gas on behalf of voters.
///      Enforces one-passport-per-registration via nullifier uniqueness.
///      Implements ICensusValidator so davinci-node can validate census roots.
contract TrustedCensus is ICensusValidator {
    using InternalLeanIMT for LeanIMTData;

    // ====================================================
    // State
    // ====================================================

    address public immutable TRUSTED_BACKEND;

    LeanIMTData private _tree;

    mapping(address => uint88) public weightOf;
    mapping(uint256 => bool) public nullifierUsed;

    // Root history: circular buffer of the last ROOT_HISTORY_SIZE replaced roots.
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
    error NullifierAlreadyUsed();
    error AlreadyRegistered();

    // ====================================================
    // Constructor
    // ====================================================

    constructor(address _TRUSTED_BACKEND) {
        TRUSTED_BACKEND = _TRUSTED_BACKEND;
        _currentRoot = _tree._root();
    }

    // ====================================================
    // Registration
    // ====================================================

    /// @notice Register a voter who has proven their zkPassport to the backend.
    /// @param account  Voter's Ethereum address (bound via the bind_evm circuit).
    /// @param nullifier Scoped nullifier from the zkPassport disclosure proof.
    function register(address account, uint256 nullifier) external {
        if (msg.sender != TRUSTED_BACKEND) revert NotTrustedBackend();
        if (nullifierUsed[nullifier]) revert NullifierAlreadyUsed();
        if (weightOf[account] != 0) revert AlreadyRegistered();

        nullifierUsed[nullifier] = true;
        weightOf[account] = 1;

        // Leaf is the voter's address cast to uint256 (weight always 1).
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
            // Evict oldest entry from circular buffer before overwriting.
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
