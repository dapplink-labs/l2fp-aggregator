// SPDX-License-Identifier: MIT
pragma solidity ^0.8.12;

import {Script,console} from "forge-std/Script.sol";
import {BLSApkRegistry} from "../src/core/BLSApkRegistry.sol";
import {BLSSignatureChecker} from "../src/core/BLSSignatureChecker.sol";

contract Deploy is Script {
    function setUp() public {}

    function run() public {
        // Load environment variables
        // address relayerManager = vm.envAddress("RELAYER_MANAGER_ADDRESS");
        // uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address relayerManager = 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266;
        uint256 deployerPrivateKey = 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80;
        
        // Start broadcasting with the deployer's private key
        vm.startBroadcast(deployerPrivateKey);

        // Deploy BLSApkRegistry
        BLSApkRegistry registry = new BLSApkRegistry(relayerManager);

        // Deploy BLSSignatureChecker
        BLSSignatureChecker checker = new BLSSignatureChecker(address(registry));

        vm.stopBroadcast();

        // Log the deployed contract addresses
        console.log("Relayer Manager address:", relayerManager);
        console.log("BLSApkRegistry deployed at:", address(registry));
        console.log("BLSSignatureChecker deployed at:", address(checker));

        // Optional: Save the deployed addresses to a file
        string memory deploymentData = string(abi.encodePacked(
            "REGISTRY_ADDRESS=", vm.toString(address(registry)), "\n",
            "CHECKER_ADDRESS=", vm.toString(address(checker))
        ));
        vm.writeFile("./deployments/addresses.txt", deploymentData);
    }
} 