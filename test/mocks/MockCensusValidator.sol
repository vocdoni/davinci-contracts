// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import {ICensusValidator} from "../../src/interfaces/ICensusValidator.sol";

/// @notice A census root history with the OnchainCensus semantics of getRootBlockNumber:
///         block.number for the current root, the block it was replaced in for an older
///         one, 0 for a root it never held. `force` overrides the answer for one root.
contract MockCensusValidator is ICensusValidator {
    uint256 public currentRoot;
    mapping(uint256 => uint256) public replacedAt;
    mapping(uint256 => uint256) private forced;
    mapping(uint256 => bool) private isForced;

    function setRoot(uint256 root) external {
        if (currentRoot != 0) replacedAt[currentRoot] = block.number;
        currentRoot = root;
    }

    function force(uint256 root, uint256 blockNumber) external {
        forced[root] = blockNumber;
        isForced[root] = true;
    }

    function getRootBlockNumber(uint256 root) external view returns (uint256) {
        if (isForced[root]) return forced[root];
        if (root != 0 && root == currentRoot) return block.number;
        return replacedAt[root];
    }

    function getCensusRoot() external view returns (uint256) {
        return currentRoot;
    }

    function getTotalVotingPowerAtRoot(uint256) external pure returns (uint256) {
        return 0;
    }
}

/// @notice A census contract whose two root getters can revert, return short or empty data,
///         or burn all the gas they are given.
contract BadCensusValidator {
    enum Mode {
        Ok,
        Revert,
        Short,
        Empty,
        Burn
    }

    Mode public rootBlockMode;
    Mode public censusRootMode;

    function setModes(Mode rootBlock, Mode censusRoot) external {
        rootBlockMode = rootBlock;
        censusRootMode = censusRoot;
    }

    function getRootBlockNumber(uint256) external view returns (uint256) {
        _misbehave(rootBlockMode);
        return block.number;
    }

    function getCensusRoot() external view returns (uint256) {
        _misbehave(censusRootMode);
        return 0;
    }

    function _misbehave(Mode m) private view {
        if (m == Mode.Revert) revert();
        if (m == Mode.Short) {
            assembly {
                mstore(0, 1)
                return(31, 1)
            }
        }
        if (m == Mode.Empty) {
            assembly {
                return(0, 0)
            }
        }
        if (m == Mode.Burn) {
            uint256 i;
            while (gasleft() > 0) {
                i++;
            }
        }
    }
}
