// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import {stdJson} from "forge-std/StdJson.sol";
import {RegistryTestBase} from "./RegistryTestBase.sol";
import {ProcessRegistry} from "../src/ProcessRegistry.sol";
import {IProcessRegistry} from "../src/interfaces/IProcessRegistry.sol";
import {DAVINCITypes} from "../src/libraries/DAVINCITypes.sol";
import {MockCensusValidator} from "./mocks/MockCensusValidator.sol";

contract NewProcessTest is RegistryTestBase {
    using stdJson for string;

    function test_NewProcess_StoresKeyAndGenesisRoot() public {
        DAVINCITypes.EncryptionKey memory key = _encKey();
        vm.expectEmit(true, true, false, true, address(registry));
        emit IProcessRegistry.ProcessCreated(bytes31(fixture.readBytes(".process_id")), ORGANIZER);
        uint256 g = gasleft();
        bytes31 pid = _fixtureProcess();
        emit log_named_uint("newProcess gas:", g - gasleft());

        DAVINCITypes.Process memory p = registry.getProcess(pid);
        assertEq(p.encryptionKey.x, key.x);
        assertEq(p.encryptionKey.y, key.y);
        assertEq(p.latestStateRoot, fixture.readBytes32(".genesis_root"));
        assertEq(
            p.latestStateRoot,
            registry.genesisRoot(pid, _ballotMode(), key, DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1)
        );
        assertEq(p.organizationId, ORGANIZER);
        assertEq(uint256(p.census.censusOrigin), 1);
        assertEq(p.census.censusRoot, fixture.readBytes32(".census_root"));
        assertEq(p.votersCount, 0);
        assertEq(p.batchNumber, 0);
    }

    function test_NewProcess_NumFieldsBounds() public {
        DAVINCITypes.BallotMode memory m = _ballotMode();
        m.groupSize = 0;

        m.numFields = 0;
        vm.expectRevert(IProcessRegistry.InvalidMaxCount.selector);
        _newProcess(block.timestamp, DURATION, MAX_VOTERS, m, _census());

        m.numFields = 17;
        vm.expectRevert(IProcessRegistry.InvalidMaxCount.selector);
        _newProcess(block.timestamp, DURATION, MAX_VOTERS, m, _census());

        m.numFields = 1;
        _newProcess(block.timestamp, DURATION, MAX_VOTERS, m, _census());
        m.numFields = 16;
        bytes31 pid = _newProcess(block.timestamp, DURATION, MAX_VOTERS, m, _census());
        assertEq(registry.getProcess(pid).ballotMode.numFields, 16);
    }

    function test_NewProcess_RejectsUnknownCensusOrigin() public {
        DAVINCITypes.Census memory c = _census();
        c.censusOrigin = DAVINCITypes.CensusOrigin.CENSUS_UNKNOWN;
        vm.expectRevert(IProcessRegistry.InvalidCensusOrigin.selector);
        _newProcess(block.timestamp, DURATION, MAX_VOTERS, _ballotMode(), c);
    }

    /// @dev Census with the origin as a raw uint8, to encode values outside the enum.
    struct RawCensus {
        uint8 censusOrigin;
        bytes32 censusRoot;
        address contractAddress;
        string censusURI;
        bool onchainAllowAnyValidRoot;
    }

    function _rawNewProcess(uint8 origin) internal returns (bool ok) {
        return _rawNewProcess(origin, address(0));
    }

    function _rawNewProcess(uint8 origin, address censusContract) internal returns (bool ok) {
        bytes32 root = origin == 4 ? bytes32(uint256(uint160(address(0xC5C5)))) : _census().censusRoot;
        RawCensus memory c = RawCensus(origin, root, censusContract, "https://example.com/census", false);
        bytes memory data = abi.encodeWithSelector(
            registry.newProcess.selector,
            DAVINCITypes.ProcessStatus.READY,
            block.timestamp,
            DURATION,
            MAX_VOTERS,
            _ballotMode(),
            c,
            METADATA_URI,
            METADATA_HASH,
            _encKey(),
            _noDkg()
        );
        vm.prank(ORGANIZER);
        (ok,) = address(registry).call(data);
    }

    function test_NewProcess_RejectsCensusOriginFive() public {
        // 5 is outside the enum: the ABI decoder rejects the call before any check runs.
        assertFalse(_rawNewProcess(5));
        assertFalse(_rawNewProcess(0));
        assertTrue(_rawNewProcess(1));
        assertTrue(_rawNewProcess(2));
        // Origin 3 needs a census contract.
        assertFalse(_rawNewProcess(3));
        assertTrue(_rawNewProcess(3, address(new MockCensusValidator())));
        assertTrue(_rawNewProcess(4));
    }

    function test_NewProcess_CSPCensus() public {
        DAVINCITypes.Census memory c = _census();
        c.censusOrigin = DAVINCITypes.CensusOrigin.CSP_EDDSA_BABYJUBJUB_V1;
        address csp = address(uint160(uint256(keccak256("csp signer"))));
        c.censusRoot = bytes32(uint256(uint160(csp)));
        bytes31 pid = _newProcess(block.timestamp, DURATION, MAX_VOTERS, _ballotMode(), c);
        DAVINCITypes.Process memory p = registry.getProcess(pid);
        assertEq(uint256(p.census.censusOrigin), 4);
        assertEq(
            p.latestStateRoot,
            registry.genesisRoot(pid, _ballotMode(), _encKey(), DAVINCITypes.CensusOrigin.CSP_EDDSA_BABYJUBJUB_V1)
        );

        // The CSP census root is the uint160 CSP address; a left-aligned address is rejected.
        c.censusRoot = bytes32(bytes20(csp));
        vm.expectRevert(IProcessRegistry.InvalidCensusRoot.selector);
        _newProcess(block.timestamp, DURATION, MAX_VOTERS, _ballotMode(), c);
    }

    function test_NewProcess_RejectsCensusConfig() public {
        DAVINCITypes.Census memory c = _census();
        c.onchainAllowAnyValidRoot = true;
        vm.expectRevert(IProcessRegistry.InvalidCensusConfig.selector);
        _newProcess(block.timestamp, DURATION, MAX_VOTERS, _ballotMode(), c);

        c = _census();
        c.censusRoot = bytes32(0);
        vm.expectRevert(IProcessRegistry.InvalidCensusRoot.selector);
        _newProcess(block.timestamp, DURATION, MAX_VOTERS, _ballotMode(), c);

        c = _census();
        c.censusURI = "";
        vm.expectRevert(IProcessRegistry.InvalidCensusURI.selector);
        _newProcess(block.timestamp, DURATION, MAX_VOTERS, _ballotMode(), c);
    }

    /// @dev newProcess with the fixture config and key k must revert InvalidEncryptionKey.
    function _expectKeyRejected(DAVINCITypes.EncryptionKey memory k) internal {
        DAVINCITypes.BallotMode memory m = _ballotMode();
        DAVINCITypes.Census memory c = _census();
        vm.expectRevert(IProcessRegistry.InvalidEncryptionKey.selector);
        vm.prank(ORGANIZER);
        registry.newProcess(
            DAVINCITypes.ProcessStatus.READY, 0, DURATION, MAX_VOTERS, m, c, METADATA_URI, METADATA_HASH, k, _noDkg()
        );
    }

    // x + p and y + p encode the same on-curve point; only the canonical range check rejects them.
    // The results guest fails such a key with RANGE, so the election could never be tallied.
    function test_RevertWhen_KeyXPlusP() public {
        DAVINCITypes.EncryptionKey memory k = _encKey();
        k.x += BN254_P;
        _expectKeyRejected(k);
    }

    function test_RevertWhen_KeyYPlusP() public {
        DAVINCITypes.EncryptionKey memory k = _encKey();
        k.y += BN254_P;
        _expectKeyRejected(k);
    }

    function test_RevertWhen_KeyBothPlusP() public {
        DAVINCITypes.EncryptionKey memory k = _encKey();
        k.x += BN254_P;
        k.y += BN254_P;
        _expectKeyRejected(k);
    }

    function test_RevertWhen_KeyXAtLeastP() public {
        DAVINCITypes.EncryptionKey memory k = _encKey();
        k.x = BN254_P;
        _expectKeyRejected(k);
        k.x = type(uint256).max;
        _expectKeyRejected(k);
    }

    function test_RevertWhen_KeyOffCurve() public {
        DAVINCITypes.EncryptionKey memory k = _encKey();
        k.y = addmod(k.y, 1, BN254_P);
        _expectKeyRejected(k);
    }

    function test_RevertWhen_KeyInReducedForm() public {
        // gnark's reduced twisted Edwards form, x' = x * (-f), is not on the circomlib curve.
        DAVINCITypes.EncryptionKey memory k = _encKey();
        uint256 negF = BN254_P - 6360561867910373094066688120553762416144456282423235903351243436111059670888;
        k.x = mulmod(k.x, negF, BN254_P);
        _expectKeyRejected(k);
    }

    function test_RevertWhen_KeyIdentity() public {
        _expectKeyRejected(DAVINCITypes.EncryptionKey({x: 0, y: 1}));
    }

    function test_RevertWhen_KeyOrderTwo() public {
        // (0, -1) is on the curve with order 2.
        _expectKeyRejected(DAVINCITypes.EncryptionKey({x: 0, y: BN254_P - 1}));
    }

    function test_NewProcess_RejectsBallotModeOverflow() public {
        DAVINCITypes.BallotMode memory m = _ballotMode();
        m.maxValueSum = uint256(1) << 63;
        vm.expectRevert(IProcessRegistry.BallotModeMaxValueSumTooLarge.selector);
        _newProcess(block.timestamp, DURATION, MAX_VOTERS, m, _census());
    }

    /// @dev External, so vm.expectRevert has a call to expect: forge compiles a test's `new`
    ///      into a deployCode cheatcode, whose revert ends the test instead.
    function deploy(address verifier, bytes32 batchVK, bytes32 resultsVK, bytes32 rootC, bytes32 vkHash)
        external
        returns (ProcessRegistry)
    {
        return new ProcessRegistry(
            CHAIN_ID,
            verifier,
            batchVK,
            resultsVK,
            rootC,
            vkHash,
            address(0),
            address(0),
            GRACE,
            GRACE_FLOOR,
            GRACE_CEIL,
            GRACE_MAX_TOTAL,
            NOTICE_MIN
        );
    }

    function test_Constructor_RejectsZeroConfig() public {
        bytes32 vkHash = bytes32(uint256(1));
        vm.expectRevert(IProcessRegistry.InvalidVerifierConfig.selector);
        this.deploy(address(0), BATCH_VK, RESULTS_VK, ROOT_C, vkHash);
        vm.expectRevert(IProcessRegistry.InvalidVerifierConfig.selector);
        this.deploy(address(1), bytes32(0), RESULTS_VK, ROOT_C, vkHash);
        vm.expectRevert(IProcessRegistry.InvalidVerifierConfig.selector);
        this.deploy(address(1), BATCH_VK, bytes32(0), ROOT_C, vkHash);
        vm.expectRevert(IProcessRegistry.InvalidVerifierConfig.selector);
        this.deploy(address(1), BATCH_VK, RESULTS_VK, bytes32(0), vkHash);
        vm.expectRevert(IProcessRegistry.InvalidVerifierConfig.selector);
        this.deploy(address(1), BATCH_VK, RESULTS_VK, ROOT_C, bytes32(0));
        assertEq(this.deploy(address(1), BATCH_VK, RESULTS_VK, ROOT_C, vkHash).batchProgramVK(), BATCH_VK);
    }

    function test_Constructor_ExposesConfig() public view {
        assertEq(registry.batchProgramVK(), BATCH_VK);
        assertEq(registry.resultsProgramVK(), RESULTS_VK);
        assertEq(registry.rootCVadcopFinal(), ROOT_C);
        assertEq(registry.ballotVKHash(), fixture.readBytes32(".ballot_vk_hash"));
        assertEq(registry.getSTVerifierVKeyHash(), BATCH_VK);
        assertEq(registry.getRVerifierVKeyHash(), RESULTS_VK);
        assertEq(registry.chainID(), CHAIN_ID);
    }
}
