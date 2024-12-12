// SPDX-License-Identifier: MIT
pragma solidity ^0.8.12;

import {BN254} from "../libraries/BN254.sol";
import {BLSApkRegistry} from "./BLSApkRegistry.sol";
import {IBLSSignatureChecker} from "../interfaces/IBLSSignatureChecker.sol";

contract BLSSignatureChecker is IBLSSignatureChecker {
    using BN254 for BN254.G1Point;
    
    // Constants
    uint256 internal constant PAIRING_EQUALITY_CHECK_GAS = 120000;

    // Immutable state variables
    BLSApkRegistry internal immutable registry;

    constructor(address registry_) {
        registry = BLSApkRegistry(registry_);
    }

    // External functions
    /**
     * @notice Verify aggregated signature
     * @param params Signature parameters
     * @return Verification result
     */
    function verifySignature(SignatureParams calldata params) 
        external 
        view 
        returns (bool) 
    {
        // require(params.blockNumber < block.number, "BLSSignatureChecker.verifySignature: Invalid block number");
        require(params.signature.length == 64, "BLSSignatureChecker.verifySignature: Invalid signature length");

        // Get all registered operators
        address[] memory operators = registry.getOperators();
        require(operators.length > 0, "BLSSignatureChecker.verifySignature: No registered operators");

        // Verify all operators are registered and not jailed
        for(uint i = 0; i < operators.length; i++) {
            require(
                registry.getOperatorId(operators[i]) != bytes32(0),
                "BLSSignatureChecker.verifySignature: Operator not registered"
            );
            require(
                !registry.isNodeJailed(operators[i]),
                "BLSSignatureChecker.verifySignature: Operator is jailed"
            );
        }

        // Extract G1 point from signature
        BN254.G1Point memory sigma = _bytesToG1Point(params.signature);

        // Get aggregated public key from registry
        BN254.G2Point memory aggregatedPubkey = registry.getAggregatedPubkey();

        // Verify signature using pairing
        return _verifySignature(params.msgHash, sigma, aggregatedPubkey);
    }

    // Internal functions
    function _bytesToG1Point(bytes memory sig) internal pure returns (BN254.G1Point memory) {
        require(sig.length == 64, "BLSSignatureChecker._bytesToG1Point: Invalid signature length");
        
        uint256 x;
        uint256 y;
        
        assembly {
            x := mload(add(sig, 32))
            y := mload(add(sig, 64))
        }
        
        return BN254.G1Point(x, y);
    }

    function _verifySignature(
        bytes32 msgHash,
        BN254.G1Point memory sigma,
        BN254.G2Point memory aggregatedPubkey
    ) internal view returns (bool) {
        return BN254.pairing(
            sigma,
            BN254.negGeneratorG2(),
            BN254.hashToG1(msgHash),
            aggregatedPubkey
        );
    }
}