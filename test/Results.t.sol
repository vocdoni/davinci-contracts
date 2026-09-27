// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import {stdJson} from "forge-std/StdJson.sol";
import {RegistryTestBase} from "./RegistryTestBase.sol";
import {IProcessRegistry} from "../src/interfaces/IProcessRegistry.sol";
import {DAVINCITypes} from "../src/libraries/DAVINCITypes.sol";
import {PinnedMockZiskVerifier} from "./mocks/MockZiskVerifier.sol";

/// @dev Results publication from the zkVM results guest; the PLONK verifier is mocked.
contract ResultsTest is RegistryTestBase {
    using stdJson for string;

    function _results() internal view returns (bytes memory pv, bytes memory proof) {
        pv = fixture.readBytes(".results.public_values");
        proof = fixture.readBytes(".results.proof_bytes");
    }

    function _settledProcess() internal returns (bytes31 pid) {
        pid = _fixtureProcess();
        _settleBoth(pid);
        assertEq(registry.getProcess(pid).latestStateRoot, fixture.readBytes32(".results.state_root"));
    }

    function test_SetProcessResults_AfterEnd() public {
        bytes31 pid = _settledProcess();
        vm.warp(block.timestamp + DURATION);
        (bytes memory pv, bytes memory proof) = _results();

        uint256[] memory want = fixture.readUintArray(".results.values");
        assertEq(want.length, 2);
        assertEq(want[1], (uint256(1) << 32) + 5);

        vm.expectEmit(true, false, false, true, address(registry));
        emit IProcessRegistry.ProcessStatusChanged(
            pid, DAVINCITypes.ProcessStatus.READY, DAVINCITypes.ProcessStatus.RESULTS
        );
        vm.expectEmit(true, true, false, true, address(registry));
        emit IProcessRegistry.ProcessResultsSet(pid, address(this), want);
        uint256 g = gasleft();
        registry.setProcessResults(pid, pv, proof);
        emit log_named_uint("setProcessResults gas, 2 fields (mock verifier):", g - gasleft());

        DAVINCITypes.Process memory p = registry.getProcess(pid);
        assertEq(uint256(p.status), uint256(DAVINCITypes.ProcessStatus.RESULTS));
        assertEq(p.result.length, uint256(p.ballotMode.numFields));
        assertEq(p.result, want);
    }

    function test_SetProcessResults_EndedByOrganizer() public {
        bytes31 pid = _settledProcess();
        vm.prank(ORGANIZER);
        registry.setProcessStatus(pid, DAVINCITypes.ProcessStatus.ENDED);
        (bytes memory pv, bytes memory proof) = _results();
        vm.prank(address(0xBEEF));
        registry.setProcessResults(pid, pv, proof);
        assertEq(uint256(registry.getProcess(pid).status), uint256(DAVINCITypes.ProcessStatus.RESULTS));
    }

    function test_SetProcessResults_SixteenFields() public {
        DAVINCITypes.BallotMode memory m = _ballotMode();
        m.numFields = 16;
        m.groupSize = 16;
        bytes31 pid = _newProcess(block.timestamp, DURATION, MAX_VOTERS, m, _census());
        vm.warp(block.timestamp + DURATION);

        uint64[16] memory values;
        for (uint256 i = 0; i < 16; i++) {
            values[i] = uint64(i * 1_000_000_007);
        }
        bytes32 root = registry.getProcess(pid).latestStateRoot;
        registry.setProcessResults(pid, _resultsPublics(root, values), "");

        uint256[] memory got = registry.getProcess(pid).result;
        assertEq(got.length, 16);
        for (uint256 i = 0; i < 16; i++) {
            assertEq(got[i], values[i]);
        }
    }

    function test_RevertWhen_BeforeEnd() public {
        bytes31 pid = _settledProcess();
        vm.warp(block.timestamp + DURATION - 1);
        (bytes memory pv, bytes memory proof) = _results();
        vm.expectRevert(IProcessRegistry.InvalidTimeBounds.selector);
        registry.setProcessResults(pid, pv, proof);
    }

    function test_RevertWhen_PausedBeforeEnd() public {
        bytes31 pid = _settledProcess();
        vm.prank(ORGANIZER);
        registry.setProcessStatus(pid, DAVINCITypes.ProcessStatus.PAUSED);
        (bytes memory pv, bytes memory proof) = _results();
        vm.expectRevert(IProcessRegistry.InvalidTimeBounds.selector);
        registry.setProcessResults(pid, pv, proof);
    }

    function test_RevertWhen_WrongRoot() public {
        // Results over the final root, but only the first batch settled.
        bytes31 pid = _fixtureProcess();
        _submit(pid, _transition(0));
        vm.warp(block.timestamp + DURATION);
        (bytes memory pv, bytes memory proof) = _results();
        vm.expectRevert(IProcessRegistry.InvalidStateRoot.selector);
        registry.setProcessResults(pid, pv, proof);
    }

    function test_RevertWhen_NotOk() public {
        bytes31 pid = _settledProcess();
        vm.warp(block.timestamp + DURATION);
        (bytes memory pv, bytes memory proof) = _results();
        _setWord(pv, 0, 0);
        vm.expectRevert(IProcessRegistry.CircuitFailed.selector);
        registry.setProcessResults(pid, pv, proof);
    }

    function test_RevertWhen_FailMaskSet() public {
        bytes31 pid = _settledProcess();
        vm.warp(block.timestamp + DURATION);
        (bytes memory pv, bytes memory proof) = _results();
        _setWord(pv, 1, 1 << 4);
        vm.expectRevert(IProcessRegistry.CircuitFailed.selector);
        registry.setProcessResults(pid, pv, proof);
    }

    function test_RevertWhen_PublicValuesLength() public {
        bytes31 pid = _settledProcess();
        vm.warp(block.timestamp + DURATION);
        (bytes memory pv, bytes memory proof) = _results();
        vm.expectRevert(IProcessRegistry.InvalidPublicValues.selector);
        registry.setProcessResults(pid, bytes.concat(pv, hex"00"), proof);
    }

    function test_RevertWhen_SecondCall() public {
        bytes31 pid = _settledProcess();
        vm.warp(block.timestamp + DURATION);
        (bytes memory pv, bytes memory proof) = _results();
        registry.setProcessResults(pid, pv, proof);
        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        registry.setProcessResults(pid, pv, proof);
    }

    function test_RevertWhen_Canceled() public {
        bytes31 pid = _settledProcess();
        vm.prank(ORGANIZER);
        registry.setProcessStatus(pid, DAVINCITypes.ProcessStatus.CANCELED);
        vm.warp(block.timestamp + DURATION);
        (bytes memory pv, bytes memory proof) = _results();
        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        registry.setProcessResults(pid, pv, proof);
    }

    function test_RevertWhen_UnknownProcess() public {
        (bytes memory pv, bytes memory proof) = _results();
        vm.expectRevert(IProcessRegistry.InvalidProcessId.selector);
        registry.setProcessResults(bytes31(0), pv, proof);
        vm.expectRevert(IProcessRegistry.ProcessNotFound.selector);
        registry.setProcessResults(bytes31(fixture.readBytes(".process_id")), pv, proof);
    }
}

/// @dev The registry must hand the verifier its pinned results vk.
contract ResultsPinnedVerifierTest is RegistryTestBase {
    function _verifier() internal override returns (address) {
        return address(new PinnedMockZiskVerifier(RESULTS_VK, ROOT_C));
    }

    function test_SetProcessResults_UsesResultsProgramVK() public {
        bytes31 pid = _fixtureProcess();
        vm.warp(block.timestamp + DURATION);
        uint64[16] memory values;
        registry.setProcessResults(pid, _resultsPublics(registry.getProcess(pid).latestStateRoot, values), "");
        assertEq(uint256(registry.getProcess(pid).status), uint256(DAVINCITypes.ProcessStatus.RESULTS));
    }
}

/// @dev A batch proof must not be accepted as a results proof.
contract ResultsWrongProgramTest is RegistryTestBase {
    function _verifier() internal override returns (address) {
        return address(new PinnedMockZiskVerifier(BATCH_VK, ROOT_C));
    }

    function test_RevertWhen_WrongProgramVK() public {
        bytes31 pid = _fixtureProcess();
        vm.warp(block.timestamp + DURATION);
        uint64[16] memory values;
        bytes memory pv = _resultsPublics(registry.getProcess(pid).latestStateRoot, values);
        vm.expectRevert(PinnedMockZiskVerifier.InvalidProof.selector);
        registry.setProcessResults(pid, pv, "");
    }
}
