// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import { ERC20 } from "@openzeppelin/contracts/token/ERC20/ERC20.sol";

/**
 * @title MockERC20
 * @dev A mock ERC20 token contract for testing purposes.
 * Extends OpenZeppelin's ERC20 with minting functionality.
 */
contract MockERC20 is ERC20 {
    /**
     * @dev Constructor that gives msg.sender all of existing tokens.
     * @param name The name of the token
     * @param symbol The symbol of the token
     */
    constructor(string memory name, string memory symbol) ERC20(name, symbol) {}

    /**
     * @dev Creates `amount` new tokens for `to`.
     * @param to The address that will receive the minted tokens
     * @param amount The amount of tokens to mint
     */
    function mint(address to, uint256 amount) public {
        _mint(to, amount);
    }

    /**
     * @dev Destroys `amount` tokens from the caller's account.
     * @param amount The amount of tokens to burn
     */
    function burn(uint256 amount) public {
        _burn(msg.sender, amount);
    }

    /**
     * @dev Destroys `amount` tokens from `account`, deducting from the caller's allowance.
     * @param account The address that will have tokens burned
     * @param amount The amount of tokens to burn
     */
    function burnFrom(address account, uint256 amount) public {
        _spendAllowance(account, msg.sender, amount);
        _burn(account, amount);
    }
}
