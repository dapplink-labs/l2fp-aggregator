// SPDX-License-Identifier: MIT
pragma solidity ^0.8.12;

import {BN254} from "../libraries/BN254.sol";

interface IBLSApkRegistry {
    struct FinalityNodeInfo {
        BN254.G1Point pubkey;         // BLS公钥
        bool isJailed;                // 是否监禁状态
        uint256 registeredTime;       // 注册时间
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
     * @notice 注册操作者
     * @param operator 操作者地址
     * @param params 公钥注册参数
     */
    function registerOperator(
        address operator,
        PubkeyRegistrationParams calldata params
    ) external returns (bytes32);

    /**
     * @notice 注销操作者
     * @param operator 操作者地址
     */
    function unRegisterOperator(
        address operator
    ) external returns (bytes32);

    /**
     * @notice 监禁操作者
     * @param operator 操作者地址
     */
    function jailOperator(address operator) external;

    /**
     * @notice 解除操作者监禁
     * @param operator 操作者地址
     */
    function unjailOperator(address operator) external;

    /**
     * @notice 获取操作者的注册公钥
     * @param operator 操作者地址
     */
    function getRegisteredPubkey(address operator) external view returns (BN254.G1Point memory, bytes32);

    /**
     * @notice 从公钥哈希获取操作者地址
     * @param pubkeyHash 公钥哈希
     */
    function getOperatorFromPubkeyHash(bytes32 pubkeyHash) external view returns (address);

    /**
     * @notice 获取操作者ID
     * @param operator 操作者地址
     */
    function getOperatorId(address operator) external view returns (bytes32);

    /**
     * @notice 检查节点是否被监禁
     * @param operator 操作者地址
     */
    function isNodeJailed(address operator) external view returns (bool);

    /**
     * @notice 获取所有操作者地址
     */
    function getOperators() external view returns (address[] memory);

    /**
     * @notice 获取聚合公钥
     */
    function getAggregatedPubkey() external view returns (BN254.G2Point memory);
}