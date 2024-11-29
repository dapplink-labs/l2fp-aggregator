// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;


interface IMantaServiceManager {
    struct BatchHeader {
        bytes32 finalityRoot;
        bytes quorumNumbers;
        bytes signedStakeForQuorums;
        uint32 referenceBlockNumber;
        // state root
        bytes32 outputRoot;
        uint256 l2BlockNumber;
        bytes32 l1BlockHash;
        uint256 l1BlockNumber;
    }

    event FinalityVerified(
        address indexed proposer,
        bytes32 outputRoot,
        uint256 l2BlockNumber,
        bytes32 l1BlockHash,
        uint256 l1BlockNumber
    );

    function verifyFinality(BatchHeader calldata batchHeader) external;
}
