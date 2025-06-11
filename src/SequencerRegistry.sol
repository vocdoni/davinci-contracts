// // SPDX-License-Identifier: AGPL-3.0-or-later
// pragma solidity ^0.8.28;

// import { Initializable } from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
// import { UUPSUpgradeable } from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
// import { OwnableUpgradeable } from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
// import { IERC20 } from "./interfaces/IERC20.sol";
// import { ISequecerRegistry } from "./ISequencerRegistry.sol";

// /**
//  * @title SequencerRegistry
//  * @dev This contract manages the registration and management of sequencers in the system.
//  * It allows sequencers to register with a stake, manage their status, and handle rewards.
//  * The contract is upgradeable and uses OpenZeppelin's UUPS pattern for upgrades.
//  * It also includes functionality for slashing sequencers, managing their active status,
//  * and handling withdrawal requests.
//  * @notice This contract is part of the Vocdoni protocol and is designed to work with the Vocdoni token.
//  * @dev The contract is initialized with the Vocdoni token address, process manager address,
//  * minimum stake amount, withdrawal cooldown period, and chain ID.
//  * The process manager is responsible for slashing sequencers and managing their active status.
//  * Sequencers can register with a specified capacity and stake, and they can request withdrawals
//  * after deactivating their status. The contract also allows for adding rewards to sequencers,
//  * claiming rewards, and managing the active/inactive status of sequencers.
//  */
// contract SequencerRegistry is Initializable, UUPSUpgradeable, OwnableUpgradeable, ISequencerRegistry {
//     /**
//      * @notice The ERC20 token used for staking and rewards in the sequencer registry.
//      * @dev This token is expected to be the Vocdoni token, which is used for staking by sequencers.
//      */
//     address public token;
//     /**
//      * @notice The address of the process manager that can slash and manage sequencers.
//      * @dev The process manager is responsible for slashing sequencers and managing their active status.
//      */
//     address public processManager;
//     /**
//      * @notice The minimum stake required for a sequencer to register.
//      * @dev This value is set during contract initialization and can be updated by the owner.
//      */
//     uint256 public minStake;
//     /**
//      * @notice The cooldown period for withdrawal requests, in seconds.
//      * @dev This is the time a sequencer must wait after requesting a withdrawal before it can be processed.
//      */
//     uint256 public withdrawalCooldown;
//     /**
//      * @notice The chain ID for which this sequencer registry is valid.
//      * @dev This is used to ensure that the contract operates on the correct blockchain network.
//      */
//     string public chainID;

//     modifier _onlyProcessManager() {
//         require(msg.sender == processManager, Unauthorized());
//         _;
//     }

//     function initialize(
//         address _token,
//         address _processManager,
//         uint256 _minStake,
//         uint256 _withdrawalCooldown,
//         string calldata _chainID
//     ) public initializer {
//         __Ownable_init(msg.sender);
//         __UUPSUpgradeable_init();
//         require(_token != address(0) && _processManager != address(0), "Invalid address");
//         require(_minStake > 0, InvalidStakeAmount());
//         require(_withdrawalCooldown > 0, "Invalid cooldown");
//         require(bytes(_chainID).length > 0, "Invalid chain ID");
//         token = _token;
//         processManager = _processManager;
//         minStake = _minStake;
//         withdrawalCooldown = _withdrawalCooldown;
//         chainID = _chainID;
//     }

//     function setProcessManager(address _processManager) external onlyOwner {
//         require(_processManager != address(0), "Zero address");
//         processManager = _processManager;
//     }

//     function setMinStake(uint256 _minStake) external onlyOwner {
//         require(_minStake > 0, "Invalid stake");
//         minStake = _minStake;
//     }

//     function setWithdrawalCooldown(uint256 _cooldown) external onlyOwner {
//         require(_cooldown > 0, "Invalid cooldown");
//         withdrawalCooldown = _cooldown;
//     }

//     function register(uint256 _capacity, uint256 _stake) external {
//         require(!sequencers[msg.sender].exists, "Already registered");
//         require(_capacity > 0, "Invalid capacity");
//         require(token.allowance(msg.sender, address(this)) >= _stake, "Allowance not strictly greater than stake");
//         require(minStake < _stake, "Stake too low");
//         require(token.transferFrom(msg.sender, address(this), _stake), "Stake failed");

//         sequencers[msg.sender] = Sequencer({
//             stake: _stake,
//             active: true,
//             capacity: _capacity,
//             registeredAt: block.timestamp,
//             lastWithdrawalRequest: 0,
//             exists: true
//         });

//         emit Registered(msg.sender, _capacity, _stake);
//     }

//     function getSequencer(address _sequencer) external view returns (Sequencer memory) {
//         return sequencers[_sequencer];
//     }

//     function requestWithdraw() external {
//         Sequencer storage s = sequencers[msg.sender];
//         require(s.exists, "Not registered");
//         require(!s.active, "Deactivate first");
//         require(s.lastWithdrawalRequest == 0, "Already requested");

//         s.lastWithdrawalRequest = block.timestamp;
//         emit WithdrawalRequested(msg.sender, s.lastWithdrawalRequest);
//     }

//     function withdraw(uint256 _amount) external {
//         Sequencer storage s = sequencers[msg.sender];
//         require(s.lastWithdrawalRequest != 0, "Withdraw not requested");
//         require(block.timestamp >= s.lastWithdrawalRequest + withdrawalCooldown, "Cooldown");
//         require(_amount <= s.stake, "Amount exceeds stake");
//         require(token.transfer(msg.sender, _amount), "Transfer failed");

//         s.stake -= _amount;

//         emit Withdrawn(msg.sender, _amount);
//     }

//     function deleteSequencer() external {
//         Sequencer storage s = sequencers[msg.sender];
//         require(s.exists, "Not registered");
//         require(!s.active, "Deactivate first");

//         delete sequencers[msg.sender];
//         require(token.transfer(msg.sender, s.stake), "Transfer failed");

//         emit Withdrawn(msg.sender, s.stake);
//         emit SequencerDeleted(msg.sender);
//     }

//     function addStake(uint256 _amount) external {
//         Sequencer storage s = sequencers[msg.sender];
//         require(s.exists, "Not registered");
//         require(token.allowance(msg.sender, address(this)) >= _amount, "Allowance too low");
//         require(token.transferFrom(msg.sender, address(this), _amount), "Transfer failed");

//         s.stake += _amount;
//         if (!s.active && s.stake >= minStake) {
//             s.active = true;
//         }

//         emit Deposit(msg.sender, s.capacity, s.stake);
//     }

//     function slash(address _sequencer, uint256 _amount) external onlyProcessManager {
//         Sequencer storage s = sequencers[_sequencer];
//         require(s.exists && s.active, "Not active");
//         if (_amount < s.stake) {
//             require(token.transfer(processManager, _amount), "Transfer failed");
//         } else {
//             require(token.transfer(processManager, s.stake), "Transfer failed");
//         }

//         s.stake -= _amount;
//         if (s.stake < minStake) {
//             s.active = false;
//         }

//         emit Slashed(_sequencer, _amount);
//     }

//     function addReward(address _sequencer, uint256 _amount) external onlyProcessManager {
//         require(sequencers[_sequencer].exists, "Invalid sequencer");
//         require(sequencers[_sequencer].active, "Sequencer not active");
//         pendingRewards[_sequencer] += _amount;
//         emit RewardAdded(_sequencer, _amount);
//     }

//     function claimRewards() external {
//         uint256 amount = pendingRewards[msg.sender];
//         require(amount > 0, "No rewards");
//         pendingRewards[msg.sender] = 0;
//         require(token.transfer(msg.sender, amount), "Claim failed");

//         emit RewardClaimed(msg.sender, amount);
//     }

//     function setInactive(address _sequencer) external onlyProcessManager {
//         Sequencer storage s = sequencers[_sequencer];
//         require(s.exists && s.active, "Invalid");

//         s.active = false;
//         emit SequencerDisabled(_sequencer);
//     }

//     function setActive(address _sequencer) external onlyProcessManager {
//         Sequencer storage s = sequencers[_sequencer];
//         require(s.exists && !s.active, "Already active");
//         require(s.stake >= minStake, "Stake too low");
//         s.active = true;
//         emit SequencerEnabled(_sequencer);
//     }

//     function isActive(address _seq) external view returns (bool) {
//         return sequencers[_seq].active;
//     }

//     function getCapacity(address _seq) external view returns (uint256) {
//         return sequencers[_seq].capacity;
//     }

//     function _authorizeUpgrade(address) internal override onlyOwner {}
// }
