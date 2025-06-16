// SPDX-License-Identifier: AGPL-3.0-only
pragma solidity ^0.8.28;

import { Ownable } from "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/math/Math.sol";

contract ProcessCalc is Ownable {
    using Math for uint256;

    /// @notice The process cost params are the parameters for the cost of a process.
    ProcessCostParams public processCostParams;

    /**
     * @notice Parameters for the cost of a process.
     * @param fixedCost The fixed cost.
     * @param costPerVote The cost per vote.
     * @param durationCost The cost per duration.
     * @param sequencerMultiplier The multiplier for the number of sequencers.
     */
    struct ProcessCostParams {
        uint256 fixedCost;
        uint256 costPerVote;
        uint256 durationCost;
        uint256 sequencerMultiplier;
    }

    constructor() Ownable(msg.sender) {}

    /**
     * @notice Calculates the cost of a process.
     * It is calculated as the sum of the fixed cost, the cost per vote, the number of sequencers and the duration.
     * @param maxVotes The maximum number of votes.
     * @param duration The duration of the process.
     * @param numSequencers The number of sequencers.
     * @return The cost of the process.
     */
    function calculateProcessCost(
        uint256 maxVotes,
        uint256 duration,
        uint256 numSequencers
    ) external view returns (uint256) {
        return
            processCostParams.fixedCost +
            (processCostParams.costPerVote * maxVotes) +
            (numSequencers * processCostParams.sequencerMultiplier) +
            (duration * processCostParams.durationCost);
    }

    function setProcessCostParams(ProcessCostParams memory _processCostParams) external onlyOwner {
        processCostParams = _processCostParams;
    }
}
