// SPDX-License-Identifier: MIT
pragma solidity ^0.8.12;

import {BN254} from "../libraries/BN254.sol";

interface IBLSApkRegistry {
    struct FinalityNodeInfo {
        BN254.G1Point pubkey;         // BLS public key
        bool isJailed;                // Jail status
        uint256 registeredTime;       // Registration timestamp
    }

    struct PubkeyRegistrationParams {
        BN254.G1Point pubkeyRegistrationSignature;
        BN254.G1Point pubkeyG1;
        BN254.G2Point pubkeyG2;
    }

    event FinalityNodeRegistered(
        address indexed operator,
        bytes32 pubkeyHash,
        uint256 registeredTime
    );

    event FinalityNodeDeregistered(
        address indexed operator,
        uint256 deregisteredTime
    );

    event FinalityNodeJailed(
        address indexed operator, 
        uint256 jailedTime
    );

    event FinalityNodeUnjailed(
        address indexed operator, 
        uint256 unjailedTime
    );

    /**
     * @notice Register an operator with their BLS public key
     * @param operator Address of the operator
     * @param params Public key registration parameters
     * @param msgHash Message hash for signature verification
     * @return Operator ID (public key hash)
     */
    function registerOperator(
        address operator,
        PubkeyRegistrationParams calldata params,
        BN254.G1Point memory msgHash
    ) external returns (bytes32);

    /**
     * @notice Deregister an operator
     * @param operator Address of the operator
     * @return Operator ID (public key hash)
     */
    function unRegisterOperator(address operator) external returns (bytes32);

    /**
     * @notice Jail an operator
     * @param operator Address of the operator
     */
    function jailOperator(address operator) external;

    /**
     * @notice Unjail an operator
     * @param operator Address of the operator
     */
    function unjailOperator(address operator) external;

    /**
     * @notice Get the registered public key of an operator
     * @param operator Address of the operator
     * @return Public key and its hash
     */
    function getRegisteredPubkey(address operator) external view returns (BN254.G1Point memory, bytes32);

    /**
     * @notice Get operator address from public key hash
     * @param pubkeyHash Hash of the public key
     * @return Address of the operator
     */
    function getOperatorFromPubkeyHash(bytes32 pubkeyHash) external view returns (address);

    /**
     * @notice Get operator ID
     * @param operator Address of the operator
     * @return Operator ID (public key hash)
     */
    function getOperatorId(address operator) external view returns (bytes32);

    /**
     * @notice Check if a node is jailed
     * @param operator Address of the operator
     * @return Jail status
     */
    function isNodeJailed(address operator) external view returns (bool);

    /**
     * @notice Get all registered operators
     * @return Array of operator addresses
     */
    function getOperators() external view returns (address[] memory);

    /**
     * @notice Get the aggregated public key
     * @return Aggregated BLS public key
     */
    function getAggregatedPubkey() external view returns (BN254.G2Point memory);

    /**
     * @notice Calculate message hash for public key registration
     * @param operator Address of the operator
     * @return Message hash point
     */
    function pubkeyRegistrationMessageHash(address operator) external view returns (BN254.G1Point memory);
}