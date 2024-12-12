// SPDX-License-Identifier: MIT
pragma solidity ^0.8.12;

import {BLSApkRegistryStorage} from "./BLSApkRegistryStorage.sol";
import {BN254} from "../libraries/BN254.sol";
import {EIP712} from "@openzeppelin/contracts/utils/cryptography/EIP712.sol";
import {OwnableUpgradeable} from "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import {Initializable} from "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import {console2} from "forge-std/console2.sol";

contract BLSApkRegistry is Initializable, EIP712, OwnableUpgradeable, BLSApkRegistryStorage {
    using BN254 for BN254.G1Point;

    modifier onlyRelayerManager() {
        require(msg.sender == relayerManager(), "BLSApkRegistry.onlyRelayerManager: Only RelayerManager can call");
        _;
    }

    constructor(
        address relayerManager_
    ) BLSApkRegistryStorage(relayerManager_) EIP712("BLSApkRegistry", "1") {
        _disableInitializers();
    }

    /*******************************************************************************
                      EXTERNAL FUNCTIONS - REGISTRY COORDINATOR
    *******************************************************************************/
    function registerOperator(
        address operator,
        PubkeyRegistrationParams calldata params,
        BN254.G1Point memory msgHash
    ) external onlyRelayerManager returns (bytes32) {
        bytes32 operatorId = getOperatorId(operator);
        if (operatorId == 0) {
            operatorId = _registerBLSPublicKey(operator, params, msgHash);
            operators.push(operator);
            _updateAggregatedPubkey();
        }
        return operatorId;
    }

    function unRegisterOperator(address operator) external onlyRelayerManager returns (bytes32) {
        bytes32 operatorId = getOperatorId(operator);
        require(operatorId != bytes32(0), "Operator not registered");

        for (uint i = 0; i < operators.length; i++) {
            if (operators[i] == operator) {
                operators[i] = operators[operators.length - 1];
                operators.pop();
                break;
            }
        }

        delete operatorToPubkey[operator];
        delete operatorToG2Pubkey[operator];
        delete operatorToPubkeyHash[operator];
        delete pubkeyHashToOperator[operatorId];
        delete finalityNodes[operator];

        _updateAggregatedPubkey();

        emit FinalityNodeDeregistered(operator, block.timestamp);
        return operatorId;
    }

    function jailOperator(address operator) external onlyRelayerManager {
        require(finalityNodes[operator].registeredTime != 0, "Operator not registered");
        finalityNodes[operator].isJailed = true;
        emit FinalityNodeJailed(operator, block.timestamp);
    }

    function unjailOperator(address operator) external onlyRelayerManager {
        require(finalityNodes[operator].registeredTime != 0, "Operator not registered");
        finalityNodes[operator].isJailed = false;
        emit FinalityNodeUnjailed(operator, block.timestamp);
    }

    function pubkeyRegistrationMessageHash(address operator) external view returns (BN254.G1Point memory) {
        return BN254.hashToG1(
            _hashTypedDataV4(
                keccak256(abi.encode(PUBKEY_REGISTRATION_TYPEHASH, operator))
            )
        );
    }
    /*******************************************************************************
                            INTERNAL FUNCTIONS
    *******************************************************************************/
    function _registerBLSPublicKey(
        address operator,
        PubkeyRegistrationParams calldata params,
        BN254.G1Point memory msgHash
    ) internal returns (bytes32) {
        bytes32 pubkeyHash = BN254.hashG1Point(params.pubkeyG1);
        require(
            pubkeyHash != ZERO_PK_HASH, "BLSApkRegistry.registerBLSPublicKey: cannot register zero pubkey"
        );
        require(
            operatorToPubkeyHash[operator] == bytes32(0),
            "BLSApkRegistry.registerBLSPublicKey: operator already registered pubkey"
        );
        require(
            pubkeyHashToOperator[pubkeyHash] == address(0),
            "BLSApkRegistry.registerBLSPublicKey: public key already registered"
        );

        uint256 gamma = uint256(keccak256(abi.encodePacked(
            params.pubkeyRegistrationSignature.X, 
            params.pubkeyRegistrationSignature.Y, 
            params.pubkeyG1.X, 
            params.pubkeyG1.Y, 
            params.pubkeyG2.X, 
            params.pubkeyG2.Y, 
            msgHash.X, 
            msgHash.Y
        ))) % BN254.FR_MODULUS;
        
        require(BN254.pairing(
            params.pubkeyRegistrationSignature.plus(params.pubkeyG1.scalar_mul(gamma)),
            BN254.negGeneratorG2(),
            msgHash.plus(BN254.generatorG1().scalar_mul(gamma)),
            params.pubkeyG2
        ), "BLSApkRegistry.registerBLSPublicKey: either the G1 signature is wrong, or G1 and G2 private key do not match");

        operatorToPubkey[operator] = params.pubkeyG1;
        operatorToG2Pubkey[operator] = params.pubkeyG2;
        operatorToPubkeyHash[operator] = pubkeyHash;
        pubkeyHashToOperator[pubkeyHash] = operator;

        finalityNodes[operator] = FinalityNodeInfo({
            pubkey: params.pubkeyG1,
            isJailed: false,
            registeredTime: block.timestamp
        });

        emit FinalityNodeRegistered(operator, pubkeyHash, block.timestamp);

        return pubkeyHash;
    }

    function _updateAggregatedPubkey() internal {
        if (operators.length == 0) {
            _aggregatedPubkey = BN254.G2Point(
                [uint256(0), uint256(0)],
                [uint256(0), uint256(0)]
            );
            console2.log("\nNo operators, setting zero pubkey");
            return;
        }

        // 获取第一个操作者的公钥
        address firstOperator = operators[0];
        _aggregatedPubkey = operatorToG2Pubkey[firstOperator];

        // 累加其他操作者的公钥
        for (uint256 i = 1; i < operators.length; i++) {
            address currentOperator = operators[i];
            BN254.G2Point memory nextPubkey = operatorToG2Pubkey[currentOperator];
            _aggregatedPubkey = BN254.plusG2(_aggregatedPubkey, nextPubkey);
        }

        console2.log("\nFinal aggregated pubkey:");
        console2.log(" - X[0]:", uint256(_aggregatedPubkey.X[0]));
        console2.log(" - X[1]:", uint256(_aggregatedPubkey.X[1]));
        console2.log(" - Y[0]:", uint256(_aggregatedPubkey.Y[0]));
        console2.log(" - Y[1]:", uint256(_aggregatedPubkey.Y[1]));
    }
    /*******************************************************************************
                            VIEW FUNCTIONS
    *******************************************************************************/
    function getRegisteredPubkey(address operator) public view returns (BN254.G1Point memory, bytes32) {
        BN254.G1Point memory pubkey = operatorToPubkey[operator];
        bytes32 pubkeyHash = operatorToPubkeyHash[operator];

        require(
            pubkeyHash != bytes32(0),
            "BLSApkRegistry.getRegisteredPubkey: operator is not registered"
        );
        
        return (pubkey, pubkeyHash);
    }

    function getOperatorFromPubkeyHash(bytes32 pubkeyHash) public view returns (address) {
        return pubkeyHashToOperator[pubkeyHash];
    }

    function getOperatorId(address operator) public view returns (bytes32) {
        return operatorToPubkeyHash[operator];
    }

    /**
     * @notice 检查 node 是否监禁状态
     * @param operator 操作者地址
     */
    function isNodeJailed(address operator) external view returns (bool) {
        return finalityNodes[operator].isJailed;
    }

    function getOperators() external view returns (address[] memory) {
        return operators;
    }
}