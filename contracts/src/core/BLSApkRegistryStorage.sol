// SPDX-License-Identifier: MIT
pragma solidity ^0.8.12;

import {IBLSApkRegistry} from "../interfaces/IBLSApkRegistry.sol";
import {Initializable} from "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import {BN254} from "../libraries/BN254.sol";

abstract contract BLSApkRegistryStorage is Initializable, IBLSApkRegistry {
    bytes32 internal constant ZERO_PK_HASH = hex"ad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5";
    bytes32 public constant PUBKEY_REGISTRATION_TYPEHASH = keccak256("BN254PubkeyRegistration(address operator)");
    address private immutable _relayerManager;

    // 状态变量
    mapping(address => FinalityNodeInfo) public finalityNodes;          
    mapping(address => bytes32) public operatorToPubkeyHash;
    mapping(bytes32 => address) public pubkeyHashToOperator;
    mapping(address => BN254.G1Point) public operatorToPubkey;
    address[] public operators;

    // 聚合公钥存储
    BN254.G2Point internal _aggregatedPubkey;

    constructor(address relayerManager_) {
        _relayerManager = relayerManager_;
    }

    function relayerManager() public view returns (address) {
        return _relayerManager;
    }

    function getAggregatedPubkey() public view returns (BN254.G2Point memory) {
        return _aggregatedPubkey;
    }

    uint256[100] private __GAP;
}