# RPC API的更改日志和使用说明

## 概述

目前被更改API如下：

+ eth_sendTransaction
+ eth_getTransactionByHash

新增API如下：

+ eth_getCMState

## 用法

## eth_sendTransaction

创建一个新的消息调用交易，如果数据字段中包含代码，则创建一个合约（请不要这么使用）。

#### 参数

`Object` - 交易对象，结果如下：

- from: DATA, 20字节 - 发送交易的源地址
- to: DATA, 20字节 - 交易的目标地址，当创建新合约时可选
- gas: QUANTITY - 交易执行可用gas量，可选整数，默认值90000，未用gas将返还。
- gasPrice: QUANTITY - gas价格，可选，默认值：待定(To-Be-Determined)
- value: QUANTITY - 交易发送的金额，可选整数
- data: DATA - 合约的编译带啊或被调用方法的签名及编码参数
- id: QUANTITY - 交易类型，可选0或1，0代表是转账交易，1代表是购币交易
- nonce: QUANTITY - nonce，可选。可以使用同一个nonce来实现挂起的交易的重写
- id==0时
  - spk: DATA - 发送方公钥
  - rpk: DATA - 接收方公钥
  - s: QUANTITY - 发送金额
  - r: QUANTITY - 返还（找零）金额
  - vor: QUANTITY - 被花费货币的承诺随机数
  - cmo: QUANTITY - 被花费货币的承诺
- id==1时
  - epkrc1: DATA - 用户公钥加密随机数r后的字段C1
  - epkrc2: DATA - 用户公钥加密随机数r后的字段C2
  - epkpc1: DATA - 利用监管者公钥加密publickey+amount的结果C1
  - epkpc2: DATA - 利用监管者公钥加密publickey+amount的结果C2
  - sigm: DATA - 发行者签名的明文信息
  - sigmhash: DATA - 发行者签名明文的hash值
  - sigr: DATA - 发行者签名的密文r
  - sigs: DATA - 发行者签名的密文s
  - cmv: DATA - 监管者公钥生成的本次购币的承诺

```json
{
    "jsonrpc": "2.0",
    "method": "eth_sendTransaction",
    "params": [
        {
            "from": "0x362de6cfc9ed13bbf207d8a243a95451883a1af2",
            "to": "0x8203599e641af59593e7dbf576dfd195eb86ff28",
            "gas": "0x76c0",
            "gasPrice": "0x9184e72a000",
            "value": "0x1",
            "id":"0x0",
            "data": "0xd46e8dd67c5d32be8d46e8dd67c5d32be8058bb8eb970870f072445675058bb8eb970870f072445675",
            "spk": "234b7f8dcdec50b47127a9ba7f03d629bd751b571ff07ac8879c4ca0a91b146205e72bd1ac5e39bcf34cbbbcf48a13edc865f862a85ce69866be24e078a3942a33333f914834ced561c145797d9b5782719dbd1b43a668d4b01151f9c0e67d9f1569899100a4ce41de3c549b649ff72d5d7c9fe8983c244cc28f2ce84b2a758c",
            "rpk": "234b7f8dcdec50b47127a9ba7f03d629bd751b571ff07ac8879c4ca0a91b146205e72bd1ac5e39bcf34cbbbcf48a13edc865f862a85ce69866be24e078a3942a33333f914834ced561c145797d9b5782719dbd1b43a668d4b01151f9c0e67d9f1569899100a4ce41de3c549b649ff72d5d7c9fe8983c244cc28f2ce84b2a758c",
            "s": "0x19",
            "r": "0x2",
            "vor":"0x0c21ccfaaa23f4562094fa71c16bbfeb1db461c2f96dc72c3a70b8cd266bd37c",
            "cmo":"0x145efb9d48584450198d2fb30a1ba7e9396eb08e0b5c662dd9414d9d8fa1abe4"
        }
    ],
    "id": 67
}
```

### 返回值

`DATA`, 32字节 - 交易哈希，如果交易还未生效则返回0值哈希

当创建合约时，在交易生效后，使用`eth_getTransactionReceipt`调用获取合约地址。

### 示例代码

请求：

```json
curl -X POST --data '{"jsonrpc":"2.0","method":"eth_sendTransaction","params":[{see above}],"id":1}'
```

响应：

```json
{
  "id":1,
  "jsonrpc": "2.0",
  "result": "0xe670ec64341771606e55d6b4ca35a1a6b75ee3d5145a99d05921026d1527331"
}
```

## eth_getTransactionByHash

返回指定哈希对应的交易。

### 参数

`DATA`, 32 字节 - 交易哈希

```json
params: [
   "0xb903239f8543d04b5dc1ba6579132b143087c68db1b2168786408fcbce568238"
]
```

### 返回值

`Object` - 交易对象，如果没有找到匹配的交易则返回null。结构如下：

- hash: DATA, 32字节 - 交易哈希
- nonce: QUANTITY - 本次交易之前发送方已经生成的交易数量
- blockHash: DATA, 32字节 - 交易所在块的哈希，对于挂起块，该值为null
- blockNumber: QUANTITY - 交易所在块的编号，对于挂起块，该值为null
- transactionIndex: QUANTITY - 交易在块中的索引位置，挂起块该值为null
- from: DATA, 20字节 - 交易发送方地址
- to: DATA, 20字节 - 交易接收方地址，对于合约创建交易，该值为null
- value: QUANTITY - 发送的以太数量，单位：wei
- gasPrice: QUANTITY - 发送方提供的gas价格，单位：wei
- gas: QUANTITY - 发送方提供的gas可用量
- input: DATA - 随交易发送的数据
- SnO: QUANTITY - 要花费代币的序列号
- Rr1: QUANTITY - 随机数，交易时对交易金额v_r进行加密
- CmSpk: QUANTITY - 发送方公钥的承诺
- CmRpk: QUANTITY - 接收方公钥的承诺
- CmO: QUANTITY - 原始金额承诺
- CmS: QUANTITY - 消费金额承诺
- CmR: QUANTITY - 找零金额承诺
- EvR: QUANTITY - E(v_r) = (v_r * G1_R + r_r2 * H_R, r_r2 * G2_R)
- EvR0: QUANTITY - EvR 的后64位
- EvR_: QUANTITY - E(v_r)’ = (v_r * G1 + r_r3 * H, r_r3 * G2；S_pk * G1 + r_spk * H，r_spk * G2；R_pk * G1 + r_rpk * H，r_rpk * G2)
- EvR_0: QUANTITY - EvR_ 的后64位
- PI: QUANTITY - 零知识证明Π
- CmV: QUANTITY - 购币承诺
- EpkV: QUANTITY - 监管者公钥对购币用户公钥和购币金额的加密
- ID: QUANTITY - 购币标识
- Sig: QUANTITY - 发行者签名

### 示例代码

请求：

```json
curl -X POST --data '{"jsonrpc":"2.0","method":"eth_getTransactionByHash","params":["0xb903239f8543d04b5dc1ba6579132b143087c68db1b2168786408fcbce568238"],"id":1}'
```

响应：

```json
{
    "jsonrpc": "2.0",
    "id": 67,
    "result": {
        "blockHash": null,
        "blockNumber": null,
        "from": "0x362de6cfc9ed13bbf207d8a243a95451883a1af2",
        "gas": "0x76c0",
        "gasPrice": "0x9184e72a000",
        "hash": "0xc0fb07973e7333a6497d71c4736f2c0a4bef94cdcea20a26f0e8710e5e026dee",
        "input": "0xd46e8dd67c5d32be8d46e8dd67c5d32be8058bb8eb970870f072445675058bb8eb970870f072445675",
        "nonce": "0x0",
        "to": "0x8203599e641af59593e7dbf576dfd195eb86ff28",
        "transactionIndex": null,
        "value": "0x1",
        "v": "0x558",
        "r": "0xb60ce212244aeb11985741812ba2833a98dbc42406c57c1750cad58f6ce65206",
        "s": "0x5bea2e4d3ab1ee82c344ac1c9f6ec9447ed4eb993215f013357bf9925ac3e8aa",
        "SnO": "0x0",
        "Rr1": "0x0",
        "CmSpk": "0x0",
        "CmRpk": "0x0",
        "CmO": "0x0",
        "CmS": "0x0",
        "CmR": "0x0",
        "EvR": "0x0",
        "EvR0": "0x0",
        "EvR_": "0x0",
        "EvR_0": "0x0",
        "PI": "0x0",
        "ID": "0xffffffffffffffff",
        "Sig": "0x922838d835e8a4d7",
        "CmV": "0xffffffffffffffff",
        "EpkV": "0x0"
    }
}
```

## eth_getCMState

返回当前承诺池内有效承诺和无效承诺的个数。

### 参数

无

### 返回值

`Object` - 承诺池内承诺个数对象

+ invalid: QUANTITY - 无效的承诺的个数
+ valid: QUANTITY - 有效的承诺的个数

### 示例代码

请求：

```json
curl -X POST --data '{"jsonrpc":"2.0","method":"eth_getCMState","id":1}'
```

响应：

```json
{
    "jsonrpc": "2.0",
    "id": 67,
    "result": {
        "invalid": "0x1",
        "valid": "0x0"
    }
}
```

