Relayer Manager 将 bls 签名进行聚合，发送到一层合约进行验证，验证通过后修改提现的调整期
部署在一层
合约要做哪些

1. bls apk 注册
node 注册需要哪些步骤

// 注册需要的参数
struct PubkeyRegistrationParams {
        BN254.G1Point pubkeyRegistrationSignature;
        BN254.G1Point pubkeyG1;
        BN254.G2Point pubkeyG2;
    }
    function registerBLSPublicKey(
        address operator,
        PubkeyRegistrationParams calldata params,
        BN254.G1Point calldata pubkeyRegistrationMessageHash
    )
------


2. BN254 bls 验证签名     

参数

txHash
signature
block_number ??? 这个 block_number 是哪个网络的！！！



节点注册 解绑的时候计算 apk
还是 check 的时候再计算



### anvil 环境
mangaer
0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80

node1
0x70997970C51812dc3A010C7d01b50e0d17dc79C8
0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d

node2
0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC
0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a


-------------
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

目前报错
BLSApkRegistry.registerBLSPublicKey: either the G1 signature is wrong, or G1 and G2 private key do not match"

验证不过 不知道是 
sign/generate_params.go 生成的不对
还是 contracts/src/core/BLSApkRegistry.sol _registerBLSPublicKey方法不对
