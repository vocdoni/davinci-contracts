// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import {Test} from "forge-std/Test.sol";
import {stdJson} from "forge-std/StdJson.sol";
import {ZiskVerifier} from "../src/verifiers/ZiskVerifier.sol";

/// @dev The vendored verifier against a batch PLONK recorded from the prover.
contract ZiskVerifierTest is Test {
    using stdJson for string;

    ZiskVerifier internal verifier;
    bytes32 internal programVK;
    bytes32 internal rootC;
    bytes internal publicValues;
    bytes internal proofBytes;

    function setUp() public {
        verifier = new ZiskVerifier();
        string memory s = vm.readFile("test/vectors/recorded_batch_snark.json");
        programVK = s.readBytes32(".program_vk");
        rootC = s.readBytes32(".root_c_vadcop_final");
        publicValues = s.readBytes(".public_values");
        proofBytes = s.readBytes(".proof_bytes");
    }

    function test_VerifiesRecordedProof() public {
        uint256 g = gasleft();
        verifier.verifySnarkProof(programVK, rootC, publicValues, proofBytes);
        emit log_named_uint("ZiskVerifier.verifySnarkProof gas", g - gasleft());
    }

    function test_RootCMatchesSetup() public view {
        assertEq(verifier.getRootCVadcopFinal(), rootC);
    }

    function test_RevertWhen_PublicValueTampered() public {
        bytes memory pv = bytes.concat(publicValues);
        pv[8 * 18] = bytes1(uint8(pv[8 * 18]) + 1);
        vm.expectRevert(ZiskVerifier.InvalidProof.selector);
        verifier.verifySnarkProof(programVK, rootC, pv, proofBytes);
    }

    function test_RevertWhen_OtherProgramVK() public {
        vm.expectRevert(ZiskVerifier.InvalidProof.selector);
        verifier.verifySnarkProof(bytes32(uint256(programVK) ^ 1), rootC, publicValues, proofBytes);
    }
}
