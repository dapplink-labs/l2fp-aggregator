// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "../interfaces/IMantaServiceManager.sol";


abstract contract MantaServiceManagerStorage is IMantaServiceManager {
    uint256 public constant THRESHOLD_DENOMINATOR = 100;

    bytes public constant quorumAdversaryThresholdPercentages = hex"21";

    bytes public constant quorumConfirmationThresholdPercentages = hex"37";

    bytes public constant quorumNumbersRequired = hex"00";

    uint32 public constant BLOCK_STALE_MEASURE = 300;

    address public immutable finalityAddress;

    address public immutable l2OutputOracle;

    uint32 public batchId;

    mapping(uint32 => bytes32) public batchIdToBatchMetadataHash;

    mapping(address => bool) public isBatchConfirmer;

    constructor(address _finalityAddress,  address _l2OutputOracle) {
        finalityAddress = _finalityAddress;
        l2OutputOracle = _l2OutputOracle;
    }

    uint256[100] private __GAP;
}

