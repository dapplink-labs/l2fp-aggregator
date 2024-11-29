// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import { Script, console } from "forge-std/Script.sol";
import "../src/core/MantaServiceManager.sol";


contract MantaManagerServiceScript is Script {
    ProxyAdmin public mantaProxyAdmin;
    MantaServiceManager public mantaServiceManager;


    function run() public {
        // env
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address finalityAddress =  vm.envAddress("FINALITY_ADDRESS");
        address l2OutputOracle =  vm.envAddress("L2OUTPUT_ORACLE");

        // deployerAddress
        address deployerAddress = vm.addr(deployerPrivateKey);

        vm.startBroadcast(deployerPrivateKey);

        mantaProxyAdmin = new ProxyAdmin(deployerAddress);
        console.log("deploy mantaProxyAdmin:", address(mantaProxyAdmin));

        mantaServiceManager = new MantaServiceManager(finalityAddress, l2OutputOracle);

        TransparentUpgradeableProxy proxyMantaServiceManager = new TransparentUpgradeableProxy(
            address(mantaServiceManager),
            address(mantaProxyAdmin),
            abi.encodeWithSelector(MantaServiceManager.initialize.selector, deployerAddress)
        );
        console.log("address:", address(deployerAddress));
        console.log("deploy proxyMantaServiceManager:", address(proxyMantaServiceManager));

        vm.stopBroadcast();
    }
}
