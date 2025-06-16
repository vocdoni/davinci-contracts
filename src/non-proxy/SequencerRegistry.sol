// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import { IERC20 } from "../interfaces/IERC20.sol";
import { ISequencerRegistry } from "../ISequencerRegistry.sol";
import { Ownable } from "@openzeppelin/contracts/access/Ownable.sol";
import "../IProcessRegistry.sol";

/**
 * @title SequencerRegistry
 * @dev This contract manages the registration and management of sequencers in the system.
 * It allows sequencers to register with a stake, manage their status, and handle rewards.
 * It also includes functionality for slashing sequencers, managing their active status,
 * and handling withdrawal requests.
 * @dev The contract is initialized with the Vocdoni token address, process registry address,
 * minimum stake amount, withdrawal cooldown period, and chain ID.
 * The process registry is responsible for slashing sequencers and managing their active status.
 * Sequencers can register with a specified capacity and stake, and they can request withdrawals
 * after deactivating their status. The contract also allows for adding rewards to sequencers,
 * claiming rewards, and managing the active/inactive status of sequencers.
 */
contract SequencerRegistry is ISequencerRegistry, Ownable {
    /**
     * @notice The ERC20 token used for staking and rewards in the sequencer registry.
     * @dev This token is expected to be the Vocdoni token, which is used for staking by sequencers.
     */
    address public token;
    /**
     * @notice The address of the process registry that can slash and manage sequencers.
     * @dev The process registry is responsible for slashing sequencers and managing their active status.
     */
    address public processRegistry;
    /**
     * @notice The minimum stake required for a sequencer to register.
     * @dev This value is set during contract initialization and can be updated by the owner.
     */
    uint256 public minStake;
    /**
     * @notice The cooldown period for withdrawal requests, in seconds.
     * @dev This is the time a sequencer must wait after requesting a withdrawal before it can be processed.
     */
    uint256 public withdrawalCooldown;
    /**
     * @notice The chain ID for which this sequencer registry is valid.
     * @dev This is used to ensure that the contract operates on the correct blockchain network.
     */
    uint32 public chainID;
    /**
     * @notice The sequencers registered in the registry.
     * @dev This mapping stores the sequencers registered in the registry.
     */
    mapping(address => Sequencer) public sequencers;
    /**
     * @notice The rewards for each sequencer in a process.
     * @dev This mapping stores the rewards for each sequencer in a process.
     */
    mapping(bytes32 => mapping(address => ProcessReward)) public rewards;
    /**
     * @notice The pending rewards for each sequencer.
     * @dev This mapping stores the pending rewards for each sequencer.
     */
    mapping(address => uint256) public pendingRewards;
    /**
     * @notice The cost factor for the votes processed.
     * @dev This is the cost factor for the votes processed.
     */
    uint256 public R;
    /**
     * @notice The cost factor for the rewrites processed.
     * @dev This is the cost factor for the rewrites processed.
     */
    uint256 public W;

    modifier onlyProcessRegistry() {
        if (msg.sender != processRegistry) revert Unauthorized();
        _;
    }

    constructor(
        address _token,
        address _processRegistry,
        uint256 _minStake,
        uint256 _withdrawalCooldown,
        uint32 _chainID
    ) Ownable(msg.sender) {
        if (_token == address(0) || _processRegistry == address(0)) revert InvalidSequencerAddress();
        if (_minStake == 0) revert InvalidStakeAmount();
        if (_withdrawalCooldown == 0) revert InvalidCooldownPeriod();
        if (_chainID == 0) revert InvalidChainID();
        token = _token;
        processRegistry = _processRegistry;
        minStake = _minStake;
        withdrawalCooldown = _withdrawalCooldown;
        chainID = _chainID;
    }

    function setProcessRegistry(address _processRegistry) external override onlyOwner {
        if (_processRegistry == address(0)) revert InvalidSequencerAddress();
        processRegistry = _processRegistry;

        emit ProcessRegistrySet(_processRegistry);
    }

    function setMinStake(uint256 _minStake) external override onlyOwner {
        if (_minStake == 0) revert InvalidStakeAmount();
        minStake = _minStake;

        emit MinStakeSet(_minStake);
    }

    function setWithdrawalCooldown(uint256 _cooldown) external override onlyOwner {
        if (_cooldown == 0) revert InvalidCooldownPeriod();
        withdrawalCooldown = _cooldown;

        emit WithdrawalCooldownSet(_cooldown);
    }

    function register(uint256 _capacity, uint256 _stake) external override {
        if (sequencers[msg.sender].exists) revert SequencerAlreadyExists();
        if (_capacity == 0) revert InvalidCapacity();
        if (_stake == 0) _stake = minStake;
        if (minStake > _stake) revert StakeAmountTooLow();
        if (IERC20(token).allowance(msg.sender, address(this)) < _stake) revert InsufficientStake();

        if (!IERC20(token).transferFrom(msg.sender, address(this), _stake)) revert InsufficientStake();

        sequencers[msg.sender] = Sequencer({
            stake: _stake,
            capacity: _capacity,
            active: true,
            registeredAt: block.timestamp,
            lastWithdrawalRequest: 0,
            exists: true
        });

        emit Registered(msg.sender, _capacity, _stake);
    }

    function getSequencer(address _sequencer) external view override returns (Sequencer memory) {
        return sequencers[_sequencer];
    }

    function requestWithdraw() external override {
        Sequencer storage s = sequencers[msg.sender];
        if (!s.exists) revert SequencerNotFound();
        if (s.active) revert SequencerNotActive();
        if (s.lastWithdrawalRequest != 0) revert WithdrawalRequestAlreadyMade();

        s.lastWithdrawalRequest = block.timestamp;
        emit WithdrawalRequested(msg.sender, s.lastWithdrawalRequest);
    }

    function withdraw(uint256 _amount) external override {
        Sequencer storage s = sequencers[msg.sender];
        if (s.lastWithdrawalRequest == 0) revert WithdrawalRequestNotMade();
        if (block.timestamp < s.lastWithdrawalRequest + withdrawalCooldown) revert InvalidCooldownPeriod();
        if (_amount > s.stake) revert InsufficientStake();
        if (!IERC20(token).transfer(msg.sender, _amount)) revert InsufficientStake();

        s.stake -= _amount;

        emit Withdrawn(msg.sender, _amount);
    }

    function deleteSequencer() external override {
        Sequencer storage s = sequencers[msg.sender];
        if (!s.exists) revert SequencerNotFound();
        if (s.active) revert SequencerNotActive();

        delete sequencers[msg.sender];
        if (!IERC20(token).transfer(msg.sender, s.stake)) revert InsufficientStake();

        emit Withdrawn(msg.sender, s.stake);
        emit SequencerDeleted(msg.sender);
    }

    function addStake(uint256 _amount) external override {
        Sequencer storage s = sequencers[msg.sender];
        if (!s.exists) revert SequencerNotFound();
        if (IERC20(token).allowance(msg.sender, address(this)) < _amount) revert InsufficientStake();
        if (!IERC20(token).transferFrom(msg.sender, address(this), _amount)) revert InsufficientStake();

        s.stake += _amount;
        if (!s.active && s.stake >= minStake) {
            s.active = true;
        }

        emit Deposit(msg.sender, s.capacity, s.stake);
    }

    function slash(address _sequencer, uint256 _amount) external override onlyProcessRegistry {
        Sequencer storage s = sequencers[_sequencer];
        if (!s.exists || !s.active) revert SequencerNotActive();
        if (_amount < s.stake) {
            if (!IERC20(token).transfer(processRegistry, _amount)) revert InsufficientStake();
        } else {
            if (!IERC20(token).transfer(processRegistry, s.stake)) revert InsufficientStake();
        }

        s.stake -= _amount;
        if (s.stake < minStake) {
            s.active = false;
        }

        emit Slashed(_sequencer, _amount);
    }

    function addReward(
        bytes32 _processId,
        address _sequencer,
        uint256 _votes,
        uint256 _rewrites
    ) external override onlyProcessRegistry {
        if (!sequencers[_sequencer].exists) revert SequencerNotFound();
        if (!sequencers[_sequencer].active) revert SequencerNotActive();
        uint256 newVotes = rewards[_processId][_sequencer].votes + _votes;
        uint256 newRewrites = rewards[_processId][_sequencer].rewrites + _rewrites;

        rewards[_processId][_sequencer].votes = newVotes;
        rewards[_processId][_sequencer].rewrites = newRewrites;

        emit RewardAdded(_processId, _sequencer, newVotes, newRewrites);
    }

    function claimRewards() external override {
        uint256 reward = pendingRewards[msg.sender];
        if (reward == 0) revert InvalidRewardAmount();

        pendingRewards[msg.sender] = 0;
        if (!IERC20(token).transfer(msg.sender, reward)) revert InsufficientStake();

        emit RewardClaimed(msg.sender, reward);
    }

    function finalizeProcessRewards(bytes32 processId) external onlyProcessRegistry {
        IProcessRegistry.Process memory p = IProcessRegistry(processRegistry).getProcess(processId);
        if (p.status != IProcessRegistry.ProcessStatus.RESULTS) revert("Process not finalized");

        uint256 rewardFromVotes = (p.cost * R) / (R + W);
        uint256 rewardFromRewrites = p.cost - rewardFromVotes;

        for (uint256 i = 0; i < p.sequencers.length; i++) {
            address seq = p.sequencers[i];
            uint256 votePart = (p.voteCount > 0) ? (rewardFromVotes * rewards[processId][seq].votes) / p.voteCount : 0;
            uint256 rewritePart = (p.voteOverwriteCount > 0)
                ? (rewardFromRewrites * rewards[processId][seq].rewrites) / p.voteOverwriteCount
                : 0;
            uint256 reward = votePart + rewritePart;

            pendingRewards[seq] += reward;
        }
    }

    function setInactive(address _sequencer) external override onlyProcessRegistry {
        Sequencer storage s = sequencers[_sequencer];
        if (!s.exists || !s.active) revert SequencerNotActive();

        s.active = false;
        emit SequencerDisabled(_sequencer);
    }

    function setActive(address _sequencer) external override onlyProcessRegistry {
        Sequencer storage s = sequencers[_sequencer];
        if (!s.exists || s.active) revert SequencerNotActive();
        if (s.stake < minStake) revert StakeAmountTooLow();
        s.active = true;
        emit SequencerEnabled(_sequencer);
    }

    function isActive(address _seq) external view override returns (bool) {
        return sequencers[_seq].active;
    }

    function getCapacity(address _seq) external view returns (uint256) {
        return sequencers[_seq].capacity;
    }
}
