### build
make build

### 准备三组 密钥
一组 作为 manager
两组 作为 node

## 准备配置文件
分别填入密钥

```
export MANTA_RELAYER_CONFIG=./manta-relayer.yaml
export MANTA_RELAYER_MIGRATIONS_DIR=./migrations
export MANTA_RELAYER_PRIVATE_KEY=FILE_ME_IN
```
copy 到三组 命令行中 

两个node 分别执行 
```
./manta-relayer parse-peer-id --private-key
```
得到两个 pub key

将 pubKey1, pubKey2 填入 yaml文件的
manager.node_members

填入 babylon_rpc

第一个命令行启动manger
```
./manta-relayer manager
```
启动第一个node
```
./manta-relayer node
```
启动第二个node
修改 yaml文件
node:
  level_db_folder: "node_storage2"
```
./manta-relayer node
```


# 每次重启
### build
make build

删除所有 levelDB 文件


### 主要查看日志
```
success to sign msg                      txHash=0x5b220056caa932c22a3eb6ea05253b31da4981bd855c1a8307b355b630e6c0e0 signature=0x17d8f63da9c13b2f44bceb9e30b3408965e8582ec48bf344c9b42b098736fabf0fcf23f55fdfe1b64fa7bf3ffcb0f30164fa21cfa058eeef77c2de3ac7361a7c block_number=33
```
填入 block_number 重新执行
manager/sdk/sdk_test.go

### 31
signature: [39 155 208 155 4 227 232 0 212 210 56 111 129 168 170 255 196 47 92 84 68 121 205 13 129 181 88 16 234 147 253 100 15 225 192 162 161 187 20 221 62 76 141 20 132 36 196 195 240 77 226 138 105 67 93 254 110 36 245 25 179 178 165 166]
--- PASS: TestGetBabyMsgByBlock (0.01s)

### 33
signature: [15 142 117 2 180 7 209 206 216 190 182 202 169 171 57 123 145 95 134 115 30 31 167 2 16 187 237 238 153 165 191 38 44 217 238 184 157 13 98 120 170 57 42 136 187 17 200 89 21 94 17 65 219 91 254 159 64 167 236 14 2 21 123 65]
--- PASS: TestGetBabyMsgByBlock (0.01s)

### 46
signature: [34 129 95 144 27 7 114 91 91 13 242 146 51 50 140 167 146 62 45 46 66 127 68 237 217 56 85 187 156 183 143 244 2 111 47 97 5 242 148 221 228 228 38 231 1 69 179 218 114 36 208 198 158 89 171 131 12 45 170 32 239 227 119 73]