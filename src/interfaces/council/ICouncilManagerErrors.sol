// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

/**
 * @title ICouncilManagerErrors
 * @notice The CouncilManager errors a call through the CouncilAdapter can bubble up, declared
 *         exactly as vocdoni/davinci-dkg-council `solidity/src/CouncilTypes.sol` declares them.
 *         `ICouncilManager.sol` (verbatim upstream) carries no errors, so this interface
 *         gives clients their names: bindProcess (UnknownCeremony, WrongPhase,
 *         NotAllowedAdapter, NotAuthorizedCreator, AlreadyBound), submitRequest (UnknownCeremony,
 *         UnknownBinding, WrongPhase, AlreadyRequested, BadFieldCount, and the per-point
 *         NonCanonical, InvalidPoint, NotInSubgroup), getPlaintexts (UnknownRequest) and
 *         isDecryptionOpen (UnknownCeremony). DecryptionNotOpen is the manager's gate error
 *         (submitPartial, combine, publishPartialData); the registry reverts with the same
 *         selector (IProcessRegistry.DecryptionNotOpen) when a COUNCIL process's results are
 *         finalized before its ceremony opens decryption.
 */
interface ICouncilManagerErrors {
    error UnknownCeremony();
    error WrongPhase();
    error NotAllowedAdapter();
    error NotAuthorizedCreator();
    error AlreadyBound();
    error UnknownBinding();
    error AlreadyRequested();
    error BadFieldCount();
    error NonCanonical();
    error InvalidPoint();
    error NotInSubgroup();
    error UnknownRequest();
    error DecryptionNotOpen();
}
