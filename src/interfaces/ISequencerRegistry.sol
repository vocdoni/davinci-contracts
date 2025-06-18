// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

/**
 * @title ISequencerRegistry
 * @notice Interface for the SequencerRegistry contract managing sequencer registration,
 * staking, rewards, and slashing in the Vocdoni protocol.
 */
interface ISequencerRegistry {
    /// STRUCTS ///

    /**
     * @notice Represents a sequencer in the registry.
     * @dev Contains information about the sequencer's stake, capacity, registration time,
     * last withdrawal request time, and whether it is active or not.
     * @param exists Indicates if the sequencer is registered in the registry.
     * @param active Indicates if the sequencer is currently active and can process requests.
     * @param stake The amount of Vocdoni tokens staked by the sequencer.
     * @param capacity The maximum number of processes the sequencer can handle concurrently.
     * @param registeredAt The timestamp when the sequencer registered.
     * @param lastWithdrawalRequest The timestamp of the last withdrawal request made by the sequencer.
     */
    struct Sequencer {
        bool exists;
        bool active;
        uint256 stake;
        uint256 capacity;
        uint256 registeredAt;
        uint256 lastWithdrawalRequest;
    }

    /**
     * @notice Represents the rewards for a sequencer in a process.
     * @param votes The number of votes the sequencer has received.
     * @param rewrites The number of rewrites the sequencer has received.
     */
    struct ProcessReward {
        uint256 votes;
        uint256 rewrites;
    }

    /// EVENTS ///

    /**
     * @notice Emitted when a sequencer is registered.
     * @param sequencer The address of the registered sequencer.
     * @param capacity The capacity of the sequencer.
     * @param stake The stake amount of the sequencer.
     */
    event Registered(address indexed sequencer, uint256 capacity, uint256 stake);
    /**
     * @notice Emitted when a user approves tokens for staking.
     * @param approver The address of the user approving tokens.
     * @param amount The amount of tokens approved.
     */
    event StakeApproved(address indexed approver, uint256 amount);
    /**
     * @notice Emitted when a sequencer is deleted.
     * @param sequencer The address of the deleted sequencer.
     */
    event SequencerDeleted(address indexed sequencer);
    /**
     * @notice Emitted when a sequencer is set to inactive.
     * @param sequencer The address of the sequencer that was set to inactive.
     */
    event SequencerDisabled(address indexed sequencer);
    /**
     * @notice Emitted when a sequencer is set to active.
     * @param sequencer The address of the sequencer that was set to active.
     */
    event SequencerEnabled(address indexed sequencer);
    /**
     * @notice Emitted when a sequencer adds stake.
     * @param sequencer The address of the sequencer that added stake.
     * @param amount The amount of stake added.
     */
    event Slashed(address indexed sequencer, uint256 amount);
    /**
     * @notice Emitted when a sequencer adds a reward.
     * @param sequencer The address of the sequencer that received the reward.
     * @param amount The amount of the reward added.
     */
    event RewardClaimed(address indexed sequencer, uint256 amount);
    /**
     * @notice Emitted when a sequencer requests a withdrawal.
     * @param sequencer The address of the sequencer that requested the withdrawal.
     * @param time The timestamp when the withdrawal was requested.
     */
    event WithdrawalRequested(address indexed sequencer, uint256 time);
    /**
     * @notice Emitted when a sequencer withdraws funds.
     * @param sequencer The address of the sequencer that withdrew funds.
     * @param amount The amount of funds withdrawn.
     */
    event Withdrawn(address indexed sequencer, uint256 amount);
    /**
     * @notice Emitted when a sequencer deposits funds.
     * @param sequencer The address of the sequencer that deposited funds.
     * @param capacity The capacity of the sequencer after the deposit.
     * @param stake The stake amount of the sequencer after the deposit.
     */
    event Deposit(address indexed sequencer, uint256 capacity, uint256 stake);
    /**
     * @notice Emitted when a sequencer adds a reward.
     * @param processId The ID of the process.
     * @param sequencer The address of the sequencer that received the reward.
     * @param votes The number of votes the sequencer has received.
     * @param rewrites The number of rewrites the sequencer has received.
     */
    event RewardAdded(bytes32 indexed processId, address indexed sequencer, uint256 votes, uint256 rewrites);
    /**
     * @notice Emitted when the process registry is set.
     * @param processRegistry The address of the process registry.
     */
    event ProcessRegistrySet(address indexed processRegistry);
    /**
     * @notice Emitted when the minimum stake is set.
     * @param minStake The minimum stake amount.
     */
    event MinStakeSet(uint256 minStake);
    /**
     * @notice Emitted when the withdrawal cooldown is set.
     * @param cooldown The withdrawal cooldown period.
     */
    event WithdrawalCooldownSet(uint256 cooldown);

    /// ERRORS ///

    /**
     * @notice Emmitted when a sequencer already exists in the registry.
     */
    error SequencerAlreadyExists();
    /**
     * @notice Emitted when a sequencer does not exist in the registry.
     */
    error SequencerNotFound();
    /**
     * @notice Emitted when a sequencer is not active.
     */
    error SequencerNotActive();
    /**
     * @notice Emitted when a sequencer has insufficient stake to perform an operation.
     */
    error InsufficientStake();
    /**
     * @notice Emitted when a sequencer has insufficient capacity to handle a request.
     */
    error InsufficientCapacity();
    /**
     * @notice Emitted when a sequencer has already requested a withdrawal.
     */
    error WithdrawalRequestAlreadyMade();
    /**
     * @notice Emitted when a sequencer has not made a withdrawal request.
     */
    error WithdrawalRequestNotMade();
    /**
     * @notice Emmitted when the stake amount is too low.
     */
    error StakeAmountTooLow();
    /**
     * @notice Emmitted when the sequencer address is invalid.
     */
    error InvalidSequencerAddress();
    /**
     * @notice Emitted when the caller is not authorized to perform an action.
     */
    error Unauthorized();
    /**
     * @notice Emitted when the cooldown period for withdrawal is invalid.
     */
    error InvalidCooldownPeriod();
    /**
     * @notice Emitted when the stake amount is invalid.
     */
    error InvalidStakeAmount();
    /**
     * @notice Emitted when the capacity amount is invalid.
     */
    error InvalidCapacity();
    /**
     * @notice Emitted when the process registry address is invalid.
     */
    error InvalidProcessRegistryAddress();
    /**
     * @notice Emitted when the chainID is invalid.
     */
    error InvalidChainID();
    /**
     * @notice Emitted when the Vocdoni token address is invalid.
     */
    error InvalidRewardAmount();
    /**
     * @notice Thrown when the rewards are already claimed.
     */
    error RewardsAlreadyClaimed();

    /// WRITE ///

    /**
     * @notice Sets the process registry address.
     * @param _processRegistry The address of the process registry contract.
     */
    function setProcessRegistry(address _processRegistry) external;
    /**
     * @notice Sets the minimum stake required for a sequencer.
     * @param _minStake The minimum stake amount in Vocdoni tokens.
     */
    function setMinStake(uint256 _minStake) external;
    /**
     * @notice Sets the withdrawal cooldown period for sequencers.
     * @param _cooldown The cooldown period in seconds.
     */
    function setWithdrawalCooldown(uint256 _cooldown) external;
    /**
     * @notice Registers a new sequencer in the registry.
     * @dev The sequencer must provide a capacity and a stake amount.
     * @param _capacity The maximum number of processes the sequencer can handle concurrently.
     * @param _stake The amount of Vocdoni tokens to stake.
     */
    function register(uint256 _capacity, uint256 _stake) external;
    /**
     * @notice Requests a withdrawal of the sequencer's stake.
     * @dev The sequencer must wait for the withdrawal cooldown period before withdrawing.
     */
    function requestWithdraw() external;
    /**
     * @notice Withdraws the sequencer's stake after the cooldown period.
     * @param _amount The amount to withdraw in Vocdoni tokens.
     * @dev The sequencer must have requested a withdrawal before calling this function.
     */
    function withdraw(uint256 _amount) external;
    /**
     * @notice Deletes a sequencer from the registry.
     * @dev This function can only be called by the sequencer itself.
     */
    function deleteSequencer() external;
    /**
     * @notice Adds stake to the sequencer's account.
     * @param _amount The amount of Vocdoni tokens to add as stake.
     */
    function addStake(uint256 _amount) external;
    /**
     * @notice Slashes a sequencer's stake.
     * @dev This function is called to penalize a sequencer for misbehavior or failure to perform its duties.
     * @param _sequencer The address of the sequencer to be slashed.
     * @param _amount The amount of Vocdoni tokens to deposit.
     */
    function slash(address _sequencer, uint256 _amount) external;
    /**
     * @notice Adds a reward to a sequencer's account.
     * @dev This function is called to reward a sequencer for its performance.
     * @param _processId The ID of the process.
     * @param _sequencer The address of the sequencer to receive the reward.
     * @param _votes The number of votes the sequencer has received.
     * @param _rewrites The number of rewrites the sequencer has received.
     */
    function addReward(bytes32 _processId, address _sequencer, uint256 _votes, uint256 _rewrites) external;
    /**
     * @notice Claims the pending rewards for the caller.
     * @dev This function allows a sequencer to claim its accumulated rewards.
     */
    function claimRewards() external;
    /**
     * @notice Sets a sequencer to inactive status.
     * @param _sequencer The address of the sequencer to set as inactive.
     */
    function setInactive(address _sequencer) external;
    /**
     * @notice Sets a sequencer to active status.
     * @param _sequencer The address of the sequencer to set as active.
     */
    function setActive(address _sequencer) external;
    /**
     * @notice Finalizes the rewards for a process.
     * @param processId The ID of the process.
     */
    function finalizeProcessRewards(bytes32 processId) external;

    /// READ ///

    /**
     * @notice Returns the sequencer information.
     * @param _sequencer The address of the sequencer.
     * @return The sequencer information.
     */
    function getSequencer(address _sequencer) external view returns (Sequencer memory);

    /**
     * @notice Returns the active status of a sequencer.
     * @param _sequencer The address of the sequencer.
     * @return The active status of the sequencer.
     */
    function isActive(address _sequencer) external view returns (bool);

    /**
     * @notice Returns the capacity of a sequencer.
     * @param _sequencer The address of the sequencer.
     * @return The capacity of the sequencer.
     */
    function getCapacity(address _sequencer) external view returns (uint256);

    /**
     * @notice Returns the process registry address.
     * @return The process registry address.
     */
    function processRegistry() external view returns (address);

    /**
     * @notice Returns the minimum stake required for a sequencer.
     * @return The minimum stake required for a sequencer.
     */
    function minStake() external view returns (uint256);

    /**
     * @notice Returns the withdrawal cooldown period for sequencers.
     * @return The withdrawal cooldown period for sequencers.
     */
    function withdrawalCooldown() external view returns (uint256);

    /**
     * @notice Returns the chain ID.
     * @return The chain ID.
     */
    function chainID() external view returns (uint32);
}
