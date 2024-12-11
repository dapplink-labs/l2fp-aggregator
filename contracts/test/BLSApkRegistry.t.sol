// SPDX-License-Identifier: MIT
pragma solidity ^0.8.12;

import {Test, console2} from "forge-std/Test.sol";
import {BLSApkRegistry} from "../src/core/BLSApkRegistry.sol";
import {BLSSignatureChecker} from "../src/core/BLSSignatureChecker.sol";
import {BN254} from "../src/libraries/BN254.sol";
import {IBLSApkRegistry} from "../src/interfaces/IBLSApkRegistry.sol";
import {IBLSSignatureChecker} from "../src/interfaces/IBLSSignatureChecker.sol";

// forge test -vvvv
contract BLSApkRegistryTest is Test {
    BLSApkRegistry public registry;
    BLSSignatureChecker public checker;
    address public relayerManager;
    address public operator1;
    address public operator2;

    IBLSApkRegistry.PubkeyRegistrationParams pubkeyParams1;
    IBLSApkRegistry.PubkeyRegistrationParams pubkeyParams2;

    function setUp() public {
        console2.log("=== Setting up test environment ===");
        
        // 设置账户
        relayerManager = 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266;
        operator1 = 0x70997970C51812dc3A010C7d01b50e0d17dc79C8;
        operator2 = 0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC;
        
        console2.log("Relayer Manager:", relayerManager);
        console2.log("Operator 1:", operator1);
        console2.log("Operator 2:", operator2);

        // 部署合约
        vm.startPrank(relayerManager);
        registry = new BLSApkRegistry(relayerManager);
        checker = new BLSSignatureChecker(address(registry));
        vm.stopPrank();
        
        console2.log("Registry deployed at:", address(registry));
        console2.log("Checker deployed at:", address(checker));

        // 使用配对的BLS密钥
        pubkeyParams1 = IBLSApkRegistry.PubkeyRegistrationParams({
            pubkeyG1: BN254.G1Point({
                X: 0x90aa0bbe0a82ac9a119d45e0337b8957749559961b2900046876d5c0df3466b,
                Y: 0x2c436aa1f1a8fd8e5964ecf277d3d5645483c11035b3af27f67f103492305122
            }),
            pubkeyG2: BN254.G2Point({
                X: [
                    0x8bc1f908b3df2021040970e8b4f0f1443433dc1d5b35fd6cc28e2497a340cd8,
                    0x24990c69557d7758bcc0ff5c9dfa54a5a968bad95475427367f30cdb407c1fe7
                ],
                Y: [
                    0x27c981ff755ff207b15e5cde297c1c0746e7842d775f1dc5c170b51d7dfbcbe6,
                    0x2bb3d6a222ea00e36725ec7f68b7aab8baf38a98932f398624803eb01a24c3b7
                ]
            }),
            pubkeyRegistrationSignature: BN254.G1Point({
                X: 0x7959c23705248701aebf0ec0b6c754d6dc66a6ae4c5567aa5a61606b18b4e54,
                Y: 0x684a709c321d4ea5f954c8fdafd6c625d112f0911d9c5633a13d3a3a4825855
            })
        });


        pubkeyParams2 = IBLSApkRegistry.PubkeyRegistrationParams({
            pubkeyG1: BN254.G1Point({
                X: 0x1672f853e0e6dd287eb73def27cc34e8b5534a930d8d8f83e2d3346d95fca40a,
                Y: 0x1dcbf68d03ec5e402f9536036a03d2a1829a5e03a4534137d1c0092e6bdc52ef
            }),
            pubkeyG2: BN254.G2Point({
                X: [
                    0x23093418f06aba0833a30d82df9e53cc26675bbc00f45fa0c167676925f847e0,
                    0x15a691d1e7cc2f66e68c1472f481bc6aa341a8e6bf7109f3b7c5c3e49c7c5b6b
                ],
                Y: [
                    0x21562d5cd84501476957b0651b46fdf34c9497bb15527a02fcb17fd29c618b37,
                    0x20b45f5f7403d75824fc1f3081a2f70ae111189e0372334a62801316b3af6b6e
                ]
            }),
            pubkeyRegistrationSignature: BN254.G1Point({
                X: 0x250f7e5fd717a181e3c3cae5cf9a38a8dc6a6860b8bb9fc571f2098ee13e404a,
                Y: 0x1842a76fd71918d3120658d161c8a0949d93a3714b1f20b4d399d972c46d1c4c
            })
        });



        console2.log("=== Setup complete ===\n");
    }

    function test_RegisterOperator_Success() public {
        console2.log("\n=== Testing successful operator registration ===");
        
        vm.startPrank(relayerManager);
        console2.log("Registering operator from relayerManager:", relayerManager);
        
        bytes32 operatorId = registry.registerOperator(operator1, pubkeyParams1);
        console2.log("Operator registered with ID:", uint256(operatorId));
        
        vm.stopPrank();

        bytes32 storedId = registry.getOperatorId(operator1);
        console2.log("Stored operator ID:", uint256(storedId));
        
        assertNotEq(operatorId, bytes32(0), "Operator ID should not be zero");
        assertTrue(registry.getOperatorId(operator1) == operatorId, "Operator ID mismatch");
        
        console2.log("=== Registration test passed ===\n");
    }

    // function test_UnregisterOperator_Success() public {
    //     console2.log("\n=== Testing successful operator unregistration ===");
        
    //     // 先注册
    //     vm.startPrank(relayerManager);
    //     console2.log("First registering operator...");
    //     bytes32 operatorId = registry.registerOperator(operator1, pubkeyParams1);
    //     console2.log("Operator registered with ID:", uint256(operatorId));
        
    //     // 再注销
    //     console2.log("Now unregistering operator...");
    //     registry.unRegisterOperator(operator1);
    //     vm.stopPrank();

    //     bytes32 newId = registry.getOperatorId(operator1);
    //     console2.log("Operator ID after unregistration:", uint256(newId));
        
    //     assertEq(newId, bytes32(0), "Operator should be unregistered");
    //     console2.log("=== Unregistration test passed ===\n");
    // }

    // function test_JailOperator_Success() public {
    //     console2.log("\n=== Testing successful operator jailing ===");
        
    //     // 先注册
    //     vm.startPrank(relayerManager);
    //     console2.log("First registering operator...");
    //     registry.registerOperator(operator1, pubkeyParams1);
        
    //     // 监禁
    //     console2.log("Now jailing operator...");
    //     registry.jailOperator(operator1);
    //     vm.stopPrank();

    //     bool isJailed = registry.isNodeJailed(operator1);
    //     console2.log("Is operator jailed?", isJailed);
        
    //     assertTrue(isJailed, "Operator should be jailed");
    //     console2.log("=== Jailing test passed ===\n");
    // }

    // function test_UnjailOperator_Success() public {
    //     console2.log("\n=== Testing successful operator unjailing ===");
        
    //     // 先注册
    //     vm.startPrank(relayerManager);
    //     console2.log("First registering operator...");
    //     registry.registerOperator(operator1, pubkeyParams1);
        
    //     // 先监禁
    //     console2.log("Now jailing operator...");
    //     registry.jailOperator(operator1);
        
    //     bool isJailedFirst = registry.isNodeJailed(operator1);
    //     console2.log("Is operator jailed after jailing?", isJailedFirst);
    //     assertTrue(isJailedFirst, "Operator should be jailed");

    //     // 再解除监禁
    //     console2.log("Now unjailing operator...");
    //     registry.unjailOperator(operator1);
    //     vm.stopPrank();

    //     bool isJailedAfter = registry.isNodeJailed(operator1);
    //     console2.log("Is operator jailed after unjailing?", isJailedAfter);
        
    //     assertFalse(isJailedAfter, "Operator should be unjailed");
    //     console2.log("=== Unjailing test passed ===\n");
    // }

    // function test_VerifySignature_Success() public {
    //     console2.log("\n=== Testing successful signature verification ===");
        
    //     // 先注册操作者
    //     vm.startPrank(relayerManager);
    //     console2.log("First registering operator...");
    //     registry.registerOperator(operator1, pubkeyParams1);
    //     vm.stopPrank();

    //     // 创建签名参数
    //     bytes32 msgHash = keccak256("test message");
    //     console2.log("Message hash:", uint256(msgHash));

    //     // 使用有效的BLS签名
    //     bytes memory validSignature = hex"1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef";
        
    //     IBLSSignatureChecker.SignatureParams memory params = IBLSSignatureChecker.SignatureParams({
    //         msgHash: msgHash,
    //         signature: validSignature,
    //         blockNumber: uint32(block.number)
    //     });
        
    //     console2.log("Block number used:", params.blockNumber);
    //     console2.log("Signature length:", params.signature.length);

    //     // 验证签名
    //     console2.log("Verifying signature...");
    //     bool isValid = checker.verifySignature(params);
    //     console2.log("Signature verification result:", isValid);
        
    //     assertTrue(isValid, "Signature should be valid");
    //     console2.log("=== Signature verification test passed ===\n");
    // }

    // function test_VerifySignature_Fail_WrongBlockNumber() public {
    //     console2.log("\n=== Testing signature verification with wrong block number ===");
        
    //     // 先注册操作者
    //     vm.startPrank(relayerManager);
    //     console2.log("First registering operator...");
    //     registry.registerOperator(operator1, pubkeyParams1);
    //     vm.stopPrank();

    //     uint32 currentBlock = uint32(block.number);
    //     uint32 oldBlock = currentBlock - 1000;
        
    //     console2.log("Current block:", currentBlock);
    //     console2.log("Using old block:", oldBlock);

    //     // 创建过期的签名参数
    //     bytes memory validSignature = hex"1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef";
        
    //     IBLSSignatureChecker.SignatureParams memory params = IBLSSignatureChecker.SignatureParams({
    //         msgHash: keccak256("test message"),
    //         signature: validSignature,
    //         blockNumber: oldBlock
    //     });

    //     // 验证签名
    //     console2.log("Verifying signature with old block number...");
    //     bool isValid = checker.verifySignature(params);
    //     console2.log("Signature verification result:", isValid);
        
    //     assertFalse(isValid, "Signature should be invalid due to wrong block number");
    //     console2.log("=== Wrong block number test passed ===\n");
    // }
}