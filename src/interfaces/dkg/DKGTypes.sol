// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

/**
 * @title DKGTypes
 * @notice Vendored subset of davinci-dkg `src/libraries/DKGTypes.sol` (commit 2338d8a):
 *         only the types the DAVINCI adapter needs, ABI-identical to upstream.
 */
library DKGTypes {
    struct Point {
        uint256 x;
        uint256 y;
    }

    /// @notice How an application's ciphertexts get decrypted. `OrganizerLocked`
    ///         adds an organizer key (`PK_aid = P_j + PK_org`) that must be
    ///         revealed before the committee can combine; `Automatic` uses the
    ///         pool key alone (`PK_aid = P_j`).
    enum AppMode {
        OrganizerLocked,
        Automatic
    }

    /// @notice Per-application policy, fixed at registration. Submission windows
    ///         are in blocks, decryption windows in unix seconds; 0 = no bound.
    struct AppPolicy {
        AppMode mode;
        bool openSubmission; // anyone may submitCiphertext
        address[] submitters; // allow-list; empty = the registrant only
        uint16 maxCiphertexts; // 0 = unlimited (capped by MAX_CIPHERTEXT_INDEX)
        uint64 notBeforeBlock; // submitCiphertext window (blocks)
        uint64 notAfterBlock;
        uint64 decryptNotBefore; // unix seconds (0 = none)
        uint64 decryptNotAfter; // unix seconds (0 = none)
    }

    /// @notice On-chain application record, as returned by `getApplication`.
    struct Application {
        address creator;
        Point organizerPK; // (0, 1) when Automatic
        uint256 organizerSecret; // 0 while sealed and always when Automatic
        uint8 poolIndex;
        AppPolicy policy;
        uint64 createdAtBlock;
        bool exists;
    }

    struct CombinedDecryptionRecord {
        uint16 ciphertextIndex;
        bool completed;
        uint256 plaintext;
    }
}
