// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity 0.8.24;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

/**
 * @title ZKRegistry
 * @author Vocdoni Association
 * @notice The ZKRegistry contract is a registry of cryptographic keys that allow to delink a user from their primary signer.
 */
contract ZKRegistry is Initializable, UUPSUpgradeable, OwnableUpgradeable {
    
    function initialize() public initializer {
        __Ownable_init(msg.sender);
        __UUPSUpgradeable_init();
    }

    mapping(address => bytes32) public accountSIK;

    function SetAccountSIK(bytes32 _sik) public {
        accountSIK[msg.sender] = _sik;
    }

    function DeleteAccountSIK() public {
        delete accountSIK[msg.sender];
    }

    function _authorizeUpgrade(address) internal override onlyOwner {}
}