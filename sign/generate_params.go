package sign

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
)

// BN254Point represents a point on the BN254 curve in JSON format
type BN254Point struct {
	X string `json:"X"`
	Y string `json:"Y"`
}

// BN254G2Point represents a G2 point on the BN254 curve in JSON format
type BN254G2Point struct {
	X []string `json:"X"`
	Y []string `json:"Y"`
}

// RegistrationParams represents the registration parameters in JSON format
type RegistrationParams struct {
	PubkeyG1                    BN254Point   `json:"pubkeyG1"`
	PubkeyG2                    BN254G2Point `json:"pubkeyG2"`
	PubkeyRegistrationSignature BN254Point   `json:"pubkeyRegistrationSignature"`
}

func makeKeyPairFromHexString(hexKey string) (*KeyPair, error) {
	// 移除 "0x" 前缀
	hexKey = strings.TrimPrefix(hexKey, "0x")

	// 解码十六进制字符串
	privateKeyBytes, err := hex.DecodeString(hexKey)
	if err != nil {
		return nil, fmt.Errorf("failed to decode hex string: %v", err)
	}

	// 转换为 big.Int
	privateKeyBig := new(big.Int).SetBytes(privateKeyBytes)

	// 对私钥进行模运算，确保在曲线阶范围内
	privateKeyBig.Mod(privateKeyBig, fr.Modulus())

	// 创建 fr.Element
	sk := new(fr.Element).SetBigInt(privateKeyBig)

	keyPair := MakeKeyPair(sk)
	if keyPair == nil {
		return nil, fmt.Errorf("failed to create key pair")
	}

	return keyPair, nil
}

// GenerateRegistrationParams generates registration parameters for given private keys
func GenerateRegistrationParams(privateKeys []string) ([]RegistrationParams, error) {
	var params []RegistrationParams

	for i, privKey := range privateKeys {
		fmt.Printf("Processing key %d: %s\n", i+1, privKey)

		keyPair, err := makeKeyPairFromHexString(privKey)
		if err != nil {
			return nil, fmt.Errorf("failed to create key pair for key %d: %v", i+1, err)
		}

		pubKeyG1 := keyPair.GetPubKeyG1()
		pubKeyG2 := keyPair.GetPubKeyG2()

		// 生成注册签名
		// 这里我们需要签名 G1 公钥
		message := pubKeyG1.Serialize()
		messageHash := [32]byte{}
		copy(messageHash[:], message)

		// 使用私钥对消息进行签名
		signature := keyPair.SignMessage(messageHash)

		// 验证签名
		isValid := signature.Verify(pubKeyG2, messageHash)
		if !isValid {
			return nil, fmt.Errorf("signature verification failed for key %d", i+1)
		}

		// 创建参数对象
		param := RegistrationParams{
			PubkeyG1: BN254Point{
				X: fmt.Sprintf("0x%x", pubKeyG1.X.BigInt(new(big.Int))),
				Y: fmt.Sprintf("0x%x", pubKeyG1.Y.BigInt(new(big.Int))),
			},
			PubkeyG2: BN254G2Point{
				X: []string{
					fmt.Sprintf("0x%x", pubKeyG2.X.A1.BigInt(new(big.Int))),
					fmt.Sprintf("0x%x", pubKeyG2.X.A0.BigInt(new(big.Int))),
				},
				Y: []string{
					fmt.Sprintf("0x%x", pubKeyG2.Y.A1.BigInt(new(big.Int))),
					fmt.Sprintf("0x%x", pubKeyG2.Y.A0.BigInt(new(big.Int))),
				},
			},
			PubkeyRegistrationSignature: BN254Point{
				X: fmt.Sprintf("0x%x", signature.X.BigInt(new(big.Int))),
				Y: fmt.Sprintf("0x%x", signature.Y.BigInt(new(big.Int))),
			},
		}

		params = append(params, param)
		fmt.Printf("Successfully processed key %d\n", i+1)
	}

	return params, nil
}

// SaveParamsToFile saves the registration parameters to a JSON file
func SaveParamsToFile(params []RegistrationParams, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	if err := encoder.Encode(params); err != nil {
		return fmt.Errorf("failed to encode JSON: %v", err)
	}

	return nil
}
