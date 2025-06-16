// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import { Test } from "forge-std/Test.sol";
import { SequencerRegistry } from "../src/non-proxy/SequencerRegistry.sol";
import { ISequencerRegistry } from "../src/ISequencerRegistry.sol";
import { MockERC20 } from "./MockERC20.sol";

contract SequencerRegistryTest is Test {
    SequencerRegistry public registry;
    MockERC20 public token;
    address public processRegistry;
    address public owner;
    address public sequencer1;
    address public sequencer2;
    address public user;

    uint256 public constant MIN_STAKE = 1000;
    uint256 public constant WITHDRAWAL_COOLDOWN = 1 days;
    uint32 public constant CHAIN_ID = 1;

    event Registered(address indexed sequencer, uint256 capacity, uint256 stake);
    event WithdrawalRequested(address indexed sequencer, uint256 timestamp);
    event Withdrawn(address indexed sequencer, uint256 amount);
    event SequencerDeleted(address indexed sequencer);
    event Deposit(address indexed sequencer, uint256 capacity, uint256 stake);
    event Slashed(address indexed sequencer, uint256 amount);
    event RewardAdded(address indexed sequencer, uint256 amount);
    event RewardClaimed(address indexed sequencer, uint256 amount);
    event SequencerDisabled(address indexed sequencer);
    event SequencerEnabled(address indexed sequencer);
    event ProcessRegistrySet(address indexed processRegistry);
    event MinStakeSet(uint256 minStake);
    event WithdrawalCooldownSet(uint256 cooldown);

    function setUp() public {
        owner = makeAddr("owner");
        processRegistry = makeAddr("processRegistry");
        sequencer1 = makeAddr("sequencer1");
        sequencer2 = makeAddr("sequencer2");
        user = makeAddr("user");

        vm.startPrank(owner);
        token = new MockERC20("Vocdoni Token", "VOC");
        registry = new SequencerRegistry(address(token), processRegistry, MIN_STAKE, WITHDRAWAL_COOLDOWN, CHAIN_ID);
        vm.stopPrank();

        // Mint tokens to sequencers and approve registry
        token.mint(sequencer1, 10000);
        token.mint(sequencer2, 10000);
        token.mint(user, 10000);
    }

    function test_Constructor() public {
        assertEq(address(registry.token()), address(token));
        assertEq(registry.processRegistry(), processRegistry);
        assertEq(registry.minStake(), MIN_STAKE);
        assertEq(registry.withdrawalCooldown(), WITHDRAWAL_COOLDOWN);
        assertEq(registry.chainID(), CHAIN_ID);
    }

    function test_Constructor_InvalidParams() public {
        vm.expectRevert(ISequencerRegistry.InvalidSequencerAddress.selector);
        new SequencerRegistry(address(0), processRegistry, MIN_STAKE, WITHDRAWAL_COOLDOWN, CHAIN_ID);

        vm.expectRevert(ISequencerRegistry.InvalidSequencerAddress.selector);
        new SequencerRegistry(address(token), address(0), MIN_STAKE, WITHDRAWAL_COOLDOWN, CHAIN_ID);

        vm.expectRevert(ISequencerRegistry.InvalidStakeAmount.selector);
        new SequencerRegistry(address(token), processRegistry, 0, WITHDRAWAL_COOLDOWN, CHAIN_ID);

        vm.expectRevert(ISequencerRegistry.InvalidCooldownPeriod.selector);
        new SequencerRegistry(address(token), processRegistry, MIN_STAKE, 0, CHAIN_ID);

        vm.expectRevert(ISequencerRegistry.InvalidChainID.selector);
        new SequencerRegistry(address(token), processRegistry, MIN_STAKE, WITHDRAWAL_COOLDOWN, 0);
    }

    function test_Register() public {
        vm.startPrank(sequencer1);
        token.approve(address(registry), MIN_STAKE);

        emit Registered(sequencer1, 100, MIN_STAKE);
        registry.register(100, MIN_STAKE);

        ISequencerRegistry.Sequencer memory seq = registry.getSequencer(sequencer1);
        assertEq(seq.stake, MIN_STAKE);
        assertEq(seq.capacity, 100);
        assertTrue(seq.active);
        assertTrue(seq.exists);
        assertEq(seq.lastWithdrawalRequest, 0);
        vm.stopPrank();
    }

    function test_Register_InvalidParams() public {
        vm.startPrank(sequencer1);

        // Test registering with zero capacity
        vm.expectRevert(ISequencerRegistry.InvalidCapacity.selector);
        registry.register(0, MIN_STAKE);

        // Test registering with insufficient stake
        vm.expectRevert(ISequencerRegistry.StakeAmountTooLow.selector);
        registry.register(100, MIN_STAKE - 1);

        // Test registering with insufficient allowance
        vm.expectRevert(ISequencerRegistry.InsufficientStake.selector);
        registry.register(100, MIN_STAKE);

        // Test registering twice
        token.approve(address(registry), MIN_STAKE);
        registry.register(100, MIN_STAKE);
        vm.expectRevert(ISequencerRegistry.SequencerAlreadyExists.selector);
        registry.register(100, MIN_STAKE);
        vm.stopPrank();
    }

    function test_WithdrawalFlow() public {
        // Register sequencer
        vm.startPrank(sequencer1);
        token.approve(address(registry), MIN_STAKE);
        registry.register(100, MIN_STAKE);

        // Try to request withdrawal while active
        vm.expectRevert(ISequencerRegistry.SequencerNotActive.selector);
        registry.requestWithdraw();

        // Set sequencer inactive
        vm.stopPrank();
        vm.prank(processRegistry);
        registry.setInactive(sequencer1);

        // Request withdrawal
        vm.startPrank(sequencer1);
        emit WithdrawalRequested(sequencer1, block.timestamp);
        registry.requestWithdraw();

        // Try to withdraw before cooldown
        vm.expectRevert(ISequencerRegistry.InvalidCooldownPeriod.selector);
        registry.withdraw(MIN_STAKE);

        // Fast forward time
        vm.warp(block.timestamp + WITHDRAWAL_COOLDOWN + 1);

        // Withdraw
        emit Withdrawn(sequencer1, MIN_STAKE);
        registry.withdraw(MIN_STAKE);

        // Verify sequencer state
        ISequencerRegistry.Sequencer memory seq = registry.getSequencer(sequencer1);
        assertEq(seq.stake, 0);
        assertFalse(seq.active);
        vm.stopPrank();
    }

    function test_Withdrawal_InvalidParams() public {
        vm.startPrank(sequencer1);
        token.approve(address(registry), MIN_STAKE);
        registry.register(100, MIN_STAKE);

        // Try to withdraw without requesting
        vm.expectRevert(ISequencerRegistry.WithdrawalRequestNotMade.selector);
        registry.withdraw(MIN_STAKE);

        // Set inactive and request withdrawal
        vm.stopPrank();
        vm.prank(processRegistry);
        registry.setInactive(sequencer1);
        vm.startPrank(sequencer1);
        registry.requestWithdraw();

        // Try to withdraw more than stake
        vm.expectRevert(ISequencerRegistry.InsufficientStake.selector);
        registry.withdraw(MIN_STAKE + 1);
        vm.stopPrank();
    }

    function test_AddStake() public {
        vm.startPrank(sequencer1);
        token.approve(address(registry), MIN_STAKE * 2);
        registry.register(100, MIN_STAKE);

        // Add more stake
        uint256 additionalStake = 500;
        emit Deposit(sequencer1, 100, MIN_STAKE + additionalStake);
        registry.addStake(additionalStake);

        ISequencerRegistry.Sequencer memory seq = registry.getSequencer(sequencer1);
        assertEq(seq.stake, MIN_STAKE + additionalStake);
        vm.stopPrank();
    }

    function test_AddStake_InvalidParams() public {
        vm.startPrank(sequencer1);

        // Try to add stake without registering
        vm.expectRevert(ISequencerRegistry.SequencerNotFound.selector);
        registry.addStake(MIN_STAKE);

        // Register and try to add stake without approval
        token.approve(address(registry), MIN_STAKE);
        registry.register(100, MIN_STAKE);
        vm.expectRevert(ISequencerRegistry.InsufficientStake.selector);
        registry.addStake(MIN_STAKE);
        vm.stopPrank();
    }

    function test_Slash() public {
        // Register sequencer
        vm.startPrank(sequencer1);
        token.approve(address(registry), MIN_STAKE);
        registry.register(100, MIN_STAKE);
        vm.stopPrank();

        // Slash sequencer
        vm.prank(processRegistry);
        emit Slashed(sequencer1, MIN_STAKE / 2);
        registry.slash(sequencer1, MIN_STAKE / 2);

        ISequencerRegistry.Sequencer memory seq = registry.getSequencer(sequencer1);
        assertEq(seq.stake, MIN_STAKE / 2);
        assertFalse(seq.active);

        // Try to slash more than stake
        vm.prank(processRegistry);
        emit Slashed(sequencer1, MIN_STAKE);
        registry.slash(sequencer1, MIN_STAKE);
        seq = registry.getSequencer(sequencer1);
        assertEq(seq.stake, 0);
    }

    function test_Slash_InvalidParams() public {
        vm.startPrank(sequencer1);
        token.approve(address(registry), MIN_STAKE);
        registry.register(100, MIN_STAKE);
        vm.stopPrank();

        // Try to slash from non-process registry
        vm.expectRevert(ISequencerRegistry.Unauthorized.selector);
        registry.slash(sequencer1, MIN_STAKE / 2);

        // Set sequencer inactive and try to slash
        vm.prank(processRegistry);
        registry.setInactive(sequencer1);
        vm.expectRevert(ISequencerRegistry.SequencerNotActive.selector);
        registry.slash(sequencer1, MIN_STAKE / 2);
    }

    function test_Rewards() public {
        // Register sequencer
        vm.startPrank(sequencer1);
        token.approve(address(registry), MIN_STAKE);
        registry.register(100, MIN_STAKE);
        vm.stopPrank();

        // Add reward
        uint256 reward = 100;
        token.mint(address(registry), reward);
        vm.prank(processRegistry);
        emit RewardAdded(sequencer1, reward);
        registry.addReward(sequencer1, reward);

        // Claim reward
        vm.startPrank(sequencer1);
        emit RewardClaimed(sequencer1, reward);
        registry.claimRewards();

        assertEq(token.balanceOf(sequencer1), 10000 + reward);
        assertEq(registry.pendingRewards(sequencer1), 0);
        vm.stopPrank();
    }

    function test_Rewards_InvalidParams() public {
        // Try to add reward to non-existent sequencer
        vm.prank(processRegistry);
        vm.expectRevert(ISequencerRegistry.SequencerNotFound.selector);
        registry.addReward(sequencer1, 100);

        // Register sequencer
        vm.startPrank(sequencer1);
        token.approve(address(registry), MIN_STAKE);
        registry.register(100, MIN_STAKE);
        vm.stopPrank();

        // Set inactive and try to add reward
        vm.prank(processRegistry);
        registry.setInactive(sequencer1);
        vm.expectRevert(ISequencerRegistry.SequencerNotActive.selector);
        registry.addReward(sequencer1, 100);

        // Try to claim non-existent reward
        vm.startPrank(sequencer1);
        vm.expectRevert(ISequencerRegistry.InvalidRewardAmount.selector);
        registry.claimRewards();
        vm.stopPrank();
    }

    function test_ActiveStatus() public {
        // Register sequencer
        vm.startPrank(sequencer1);
        token.approve(address(registry), MIN_STAKE);
        registry.register(100, MIN_STAKE);
        vm.stopPrank();

        // Set inactive
        vm.prank(processRegistry);
        emit SequencerDisabled(sequencer1);
        registry.setInactive(sequencer1);
        assertFalse(registry.isActive(sequencer1));

        // Set active
        vm.prank(processRegistry);
        emit SequencerEnabled(sequencer1);
        registry.setActive(sequencer1);
        assertTrue(registry.isActive(sequencer1));
    }

    function test_ActiveStatus_InvalidParams() public {
        // Try to set non-existent sequencer inactive
        vm.prank(processRegistry);
        vm.expectRevert(ISequencerRegistry.SequencerNotActive.selector);
        registry.setInactive(sequencer1);

        // Register sequencer
        vm.startPrank(sequencer1);
        token.approve(address(registry), MIN_STAKE);
        registry.register(100, MIN_STAKE);
        vm.stopPrank();

        // Try to set already inactive sequencer inactive
        vm.prank(processRegistry);
        registry.setInactive(sequencer1);
        vm.expectRevert(ISequencerRegistry.SequencerNotActive.selector);
        registry.setInactive(sequencer1);

        // Try to set active with insufficient stake
        vm.prank(processRegistry);
        registry.slash(sequencer1, MIN_STAKE);
        vm.expectRevert(ISequencerRegistry.StakeAmountTooLow.selector);
        registry.setActive(sequencer1);
    }

    function test_OwnerFunctions() public {
        vm.startPrank(owner);

        // Set process registry
        address newProcessRegistry = makeAddr("newProcessRegistry");
        emit ProcessRegistrySet(newProcessRegistry);
        registry.setProcessRegistry(newProcessRegistry);
        assertEq(registry.processRegistry(), newProcessRegistry);

        // Set min stake
        uint256 newMinStake = 2000;
        emit MinStakeSet(newMinStake);
        registry.setMinStake(newMinStake);
        assertEq(registry.minStake(), newMinStake);

        // Set withdrawal cooldown
        uint256 newCooldown = 2 days;
        emit WithdrawalCooldownSet(newCooldown);
        registry.setWithdrawalCooldown(newCooldown);
        assertEq(registry.withdrawalCooldown(), newCooldown);

        vm.stopPrank();
    }

    function test_OwnerFunctions_InvalidParams() public {
        vm.startPrank(owner);

        // Try to set zero address as process registry
        vm.expectRevert(ISequencerRegistry.InvalidSequencerAddress.selector);
        registry.setProcessRegistry(address(0));

        // Try to set zero min stake
        vm.expectRevert(ISequencerRegistry.InvalidStakeAmount.selector);
        registry.setMinStake(0);

        // Try to set zero withdrawal cooldown
        vm.expectRevert(ISequencerRegistry.InvalidCooldownPeriod.selector);
        registry.setWithdrawalCooldown(0);

        vm.stopPrank();
    }

    function test_UnauthorizedAccess() public {
        vm.startPrank(user);

        // Try to call owner functions
        vm.expectRevert("Ownable: caller is not the owner");
        registry.setProcessRegistry(makeAddr("newProcessRegistry"));

        vm.expectRevert("Ownable: caller is not the owner");
        registry.setMinStake(2000);

        vm.expectRevert("Ownable: caller is not the owner");
        registry.setWithdrawalCooldown(2 days);

        // Try to call process registry functions
        vm.expectRevert(ISequencerRegistry.Unauthorized.selector);
        registry.slash(sequencer1, MIN_STAKE);

        vm.expectRevert(ISequencerRegistry.Unauthorized.selector);
        registry.addReward(sequencer1, 100);

        vm.expectRevert(ISequencerRegistry.Unauthorized.selector);
        registry.setInactive(sequencer1);

        vm.expectRevert(ISequencerRegistry.Unauthorized.selector);
        registry.setActive(sequencer1);

        vm.stopPrank();
    }

    function test_DeleteSequencer() public {
        // Register sequencer
        vm.startPrank(sequencer1);
        token.approve(address(registry), MIN_STAKE);
        registry.register(100, MIN_STAKE);

        // Try to delete while active
        vm.expectRevert(ISequencerRegistry.SequencerNotActive.selector);
        registry.deleteSequencer();

        // Set inactive and delete
        vm.stopPrank();
        vm.prank(processRegistry);
        registry.setInactive(sequencer1);
        vm.startPrank(sequencer1);

        emit Withdrawn(sequencer1, MIN_STAKE);
        emit SequencerDeleted(sequencer1);
        registry.deleteSequencer();

        // Verify sequencer is deleted
        ISequencerRegistry.Sequencer memory seq = registry.getSequencer(sequencer1);
        assertFalse(seq.exists);
        assertEq(token.balanceOf(sequencer1), 10000);
        vm.stopPrank();
    }

    function test_DeleteSequencer_InvalidParams() public {
        vm.startPrank(sequencer1);

        // Try to delete non-existent sequencer
        vm.expectRevert(ISequencerRegistry.SequencerNotFound.selector);
        registry.deleteSequencer();

        // Register and try to delete while active
        token.approve(address(registry), MIN_STAKE);
        registry.register(100, MIN_STAKE);
        vm.expectRevert(ISequencerRegistry.SequencerNotActive.selector);
        registry.deleteSequencer();
        vm.stopPrank();
    }

    function test_GetAllowance() public {
        vm.startPrank(sequencer1);
        uint256 allowance = 1000;
        token.approve(address(registry), allowance);
        assertEq(registry.getAllowance(sequencer1), allowance);
        vm.stopPrank();
    }

    function test_ApproveStake() public {
        vm.startPrank(sequencer1);
        uint256 amount = 1000;
        token.approve(address(registry), amount);
        assertEq(token.allowance(sequencer1, address(registry)), amount);
        vm.stopPrank();
    }
}
