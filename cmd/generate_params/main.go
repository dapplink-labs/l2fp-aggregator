package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/eniac-x-labs/manta-relayer/sign"
)

const paramTemplate = `        pubkeyParams{{.Index}} = IBLSApkRegistry.PubkeyRegistrationParams({
            pubkeyG1: BN254.G1Point({
                X: 0x{{.PubkeyG1.X}},
                Y: 0x{{.PubkeyG1.Y}}
            }),
            pubkeyG2: BN254.G2Point({
                X: [
                    0x{{index .PubkeyG2.X 0}},
                    0x{{index .PubkeyG2.X 1}}
                ],
                Y: [
                    0x{{index .PubkeyG2.Y 0}},
                    0x{{index .PubkeyG2.Y 1}}
                ]
            }),
            pubkeyRegistrationSignature: BN254.G1Point({
                X: 0x{{.PubkeyRegistrationSignature.X}},
                Y: 0x{{.PubkeyRegistrationSignature.Y}}
            })
        });
`

type TemplateData struct {
	Index int
	sign.RegistrationParams
}

func main() {
	// 定义私钥
	privateKeys := []string{
		"0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d", // node1
		"0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a", // node2
	}

	// 生成参数
	params, err := sign.GenerateRegistrationParams(privateKeys)
	if err != nil {
		panic(fmt.Sprintf("Failed to generate parameters: %v", err))
	}

	// 创建模板
	tmpl, err := template.New("params").Parse(paramTemplate)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse template: %v", err))
	}

	// 创建输出文件
	file, err := os.Create("registration_params.txt")
	if err != nil {
		panic(fmt.Sprintf("Failed to create file: %v", err))
	}
	defer file.Close()

	// 写入每个参数
	for i, param := range params {
		// 处理参数格式：移除所有的 "0x" 前缀，因为模板中会添加
		processedParam := sign.RegistrationParams{
			PubkeyG1: sign.BN254Point{
				X: strings.TrimPrefix(param.PubkeyG1.X, "0x"),
				Y: strings.TrimPrefix(param.PubkeyG1.Y, "0x"),
			},
			PubkeyG2: sign.BN254G2Point{
				X: []string{
					strings.TrimPrefix(param.PubkeyG2.X[0], "0x"),
					strings.TrimPrefix(param.PubkeyG2.X[1], "0x"),
				},
				Y: []string{
					strings.TrimPrefix(param.PubkeyG2.Y[0], "0x"),
					strings.TrimPrefix(param.PubkeyG2.Y[1], "0x"),
				},
			},
			PubkeyRegistrationSignature: sign.BN254Point{
				X: strings.TrimPrefix(param.PubkeyRegistrationSignature.X, "0x"),
				Y: strings.TrimPrefix(param.PubkeyRegistrationSignature.Y, "0x"),
			},
		}

		data := TemplateData{
			Index:              i + 1,
			RegistrationParams: processedParam,
		}

		if err := tmpl.Execute(file, data); err != nil {
			panic(fmt.Sprintf("Failed to execute template: %v", err))
		}

		// 添加换行，除非是最后一个参数
		if i < len(params)-1 {
			file.WriteString("\n\n")
		}
	}

	fmt.Println("Parameters successfully generated and saved to registration_params.txt")
}
