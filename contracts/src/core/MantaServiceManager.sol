// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Initializable} from "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/access/AccessControlUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/utils/ReentrancyGuardUpgradeable.sol";

import { MantaServiceManagerStorage } from "./MantaServiceManagerStorage.sol";


contract MantaServiceManager is Initializable, AccessControlUpgradeable, ReentrancyGuardUpgradeable, OwnableUpgradeable,  MantaServiceManagerStorage {
    constructor(address _finalityAddress, address _l2OutputOracle) MantaServiceManagerStorage(_finalityAddress, _l2OutputOracle) {
        _disableInitializers();
    }

    function initialize(address initialOwner) public initializer {
        _transferOwnership(initialOwner);
    }

    modifier onlyFinalityManager() {
        require(msg.sender == address(finalityAddress), "MantaServiceManager.only finality manager can call this function");
        _;
    }

    function verifyFinality(BatchHeader calldata batchHeader) external onlyFinalityManager {
        require(tx.origin == msg.sender, "MantaServiceManager.verifyFinality: header and nonsigner data must be in calldata");
        require(
            batchHeader.referenceBlockNumber < block.number, "MantaServiceManager.verifyFinality: specified referenceBlockNumber is in future"
        );

        require(
            (batchHeader.referenceBlockNumber + BLOCK_STALE_MEASURE) >= uint32(block.number),
            "MantaServiceManager.verifyFinality: specified referenceBlockNumber is too far in past"
        );

        require(
            batchHeader.quorumNumbers.length == batchHeader.signedStakeForQuorums.length,
            "MantaServiceManager.verifyFinality: quorumNumbers and signedStakeForQuorums must be of the same length"
        );
        // todo check sign and verify zk proof
        // call l2 output oracle propose state root
        (bool success, ) = l2OutputOracle.call(
            abi.encodeWithSignature(
                "proposeL2Output(bytes32,uint256,bytes32,uint256)",
                batchHeader.outputRoot,
                batchHeader.l2BlockNumber,
                batchHeader.l1BlockHash,
                batchHeader.l1BlockNumber
            )
        );
        require(success, "call propose L2Output fail");

        // emit FinalityVerified event
        emit FinalityVerified(msg.sender, batchHeader.outputRoot, batchHeader.l2BlockNumber, batchHeader.l1BlockHash, batchHeader.l1BlockNumber);
    }

}