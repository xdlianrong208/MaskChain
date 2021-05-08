# 交易所模块

假设用户是通过向发行者购买的方式获得初始数字货币的，然后才能进行转账交易。购币的核心操作：发行者将用户购得的数字货币承诺的哈希值添加到系统的承诺池中。

其中cm形式为pederson承诺。

## 功能流程：

![](./img/process.png)

##### 其中和以前方案的冲突：

将生成签名的步骤移到这一步

## 功能说明：

#### 1 初始化

通过输入端口号初始化程序（缺省1323），发行者输入任意字符串用来初始化公私钥（该公私钥已配置好，无需再次输入），发行者输入自己在链上的账户和密码进行账户解锁。

#### 2 发行者生成公私钥

通过GenerateKeys算法生成发行者公私钥，如果本地文件中已经有了公私钥配置，则不再生成

输入：任意字符串

输出：公钥，私钥

#### 3 接收用户响应

设置监听端口，接收用户发来的publickey和amount判断publickey合法

将publickey发送至监管者服务器（:1423/verify）,若返回false跳出程序，返回ture进入下一步

#### 4 请求获得监管者公钥

调用/regkey获取监管者公钥，存储在本地文件中

#### 5 承诺生成

先生成uint64位随机数r，利用承诺生成算法CommitByUint64

输入：监管者公钥struct，amount，r

输出：CM_v，规范后的r

#### 6 加密购币交易信息

调用ELGamal加密算法

输入：监管者公钥struct，用户publickey与amount的拼接字符串作为明文

输出：购币交易信息密文(C1，C2)

#### 7 加密随机数r

调用ELGamal加密算法

输入：用户公钥struct，承诺生成过程中产生的随机数r作为明文

输出：r的密文(C1，C2)

#### 8 发行者签名

调用加密算法Sign

输入：发行者私钥，用户amount与ID的拼接字符串作为签名明文

输出：Signature

```注：购币交易中的ID为1```

#### 9 sendTranscation

将交易信息打包通过rpc方式发送上链。

如果上链成功，返回用户交易回执（CM_v，r）；

如果上链失败，返回错误。

## 监听接口

服务器端口号：缺省1323

1.路由```/buy``` [POST]暴露给用户，用户输入样例如下

```json
{
   "g1": "23021d5b6c06398e6f21a16a1b34738dcde99f330738bf02857380e824317cf5",
   "g2": "04f4e643002836bdd0480a1663a85deaff82ab88db546863f0fdf38d9afd8ae0",
   "p": "32dc3a13e86eded11e481da6c95feea5f510f9eb5a6c997c74549fccd15f74a7",
   "h": "12ef9e3bc4d032db9e9d528dfd42e7a2a2aba45949570290bd157647dae48c1e",
   "amount": "100"
}
```

2.路由```/pubpub``` [GET]暴露给用户，返回发行者公钥信息

## 启动命令

启动参数如下：

GLOBAL OPTIONS:
   --port value, -p value                    the port of this server (default: "1323")
   --generatekey value, --gk value   the string that you generate your pub/pri key
   --ethaccount value, --ea value     the eth_account of you
   --ethkey value, --ek value             the key that you unlock your eth_account
   --help, -h                                         show help

## 使用方法

```
go build 
./exchange -ea 0x75e36ea49f49d6f6619eb23904e8a8cab3a3dda2 -ek 1
```



