## 数据结构

公钥

```
type PublicKey struct {
   G1, G2, P, H *big.Int
}
```

私钥

```
type PrivateKey struct {
   PublicKey
   X *big.Int
}
```

密文

```
type CypherText struct {
   C1, C2 []byte
}
```

承诺

```
type Commitment struct {
   commitment, r []byte
}
```

签名

```
type Signature struct {
   M, M_hash, R, S []byte
}
```

## 

## Pederson Commitment

$$ com = g^v*h^r $$

生成承诺的函数：使用公钥和一个数，返回承诺和随机数。输入金额v

```
func (pub PublicKey) Commit(v *big.Int, rnd []byte) Commitment {}
```

有需要时还可以对比特串进行承诺

```
func (pub PublicKey) CommitByBytes(b []byte, rnd []byte) Commitment {}
```

## 

## Key

$$ PubKey = (G1, G2, H), 其中H = g^v $$

公钥的数据结构：

```
type PubKey struct {
   G1 ECPoint
   G2 ECPoint
   H ECPoint
}
```

私钥的数据结构：

```
type PrivKey struct {
   PubKey
   X *big.Int
}
```

生成公私钥：输入字符串即可

```
func GenerateKeys(info string) (pub PublicKey, priv PrivateKey, err error) {}
```

## 

## Encrypt & Decrypt

加解密：公钥加密，生成t1,t2，私钥解密，使用t1,t2

```
func Encrypt(pub PublicKey, M []byte) (C CypherText) {}
func Decrypt(priv PrivateKey, C CypherText) (M []byte) {}
```

## 

## Sign & VerifySign

签名：

```
func Sign(priv PrivateKey, m []byte) (sig Signature) {}
func Verify(pub PublicKey, sig Signature) bool {}
```

## 

## Discrete Logarithm Proof

$$ (y,x) : y = g^x $$

离散对数证明

proof：提供密文，证明其知道一个数x

verify：验证prover知道x

```
func Discrete_Logarithm_Proof(x *big.Int) DLP
// 参数：一个数
func DLPVerify(dlp DLP) bool
```

## 

## Representation Proof

$$ (y,x1, ...,xl) : y = g1*x1·g2*x2...gl·xl $$

表示证明

proof：提供密文，证明其知道数组[x1, x2, ..., xn]内的所有数字

verify：验证prover知道数组

```
func RepProof(gn []ECPoint,xn []*big.Int) REP
// 参数：一堆椭圆曲线上的点，和一堆数
func RepVerify(rep REP) bool
```

## 

## Equality Proof

$$ (y1,y2,x):y1=g*x1∧y2=g*x2 $$

相等证明

proof：提供密文，证明，虽然生成元不同，但两个y的指数x相等

verify：验证这俩数相等（但不知道这俩数是多少）

```
func EPProof (g1 ECPoint, g2 ECPoint, x *big.Int) EP
// 参数：两个椭圆曲线上点，和一个数，生成EP结构密文
func EPVerify (ep EP) bool
```

## 

## Linear Equation Proof

$$ (y,b,a1, ...al,x1, ...,xl) : y = g1*x1·g2*x2...gl·xl ∧  a1x1+a2x2+...+alxl= b $$

线性证明

proof：提供密文，证明其知道x1, x2, x3 ... xl且x1, x2, x3... xl 满足自定义的线性关系（比如某几个数相等，某俩数相加等于另一个数的关系等等）

prover：验证其正确性

```
func Linear_equation_proof (gn []ECPoint, xn []*big.Int, an []*big.Int, b *big.Int) LEP
// 参数：一堆椭圆曲线的点，一堆数；数所满足的线性关系的参数一些a和b，详情请看test示例
func LepVerify(lep LEP) bool
```

## 

## Bulletproofs in Go

范围证明

proof：提供密文，证明一个数在[0, x]之间

verify：验证这个数确实在这范围里

This project implements Bulletproofs in Go. More information about Bulletproofs can be found [here](https://crypto.stanford.edu/bulletproofs/)

Paper references for the steps of the protocol:

- The inner-product argument is implemented as shown in Protocol 1 and Protocol 2.
- The range proof is implemented as described in Section 4.1.
- The multi-range proof is implemented as described in Section 4.3.
- Non-interactivity is implemented as described in Section 4.4 with SHA256.

WARNING: This is research quality code.

```
func RPProve(v *big.Int) RangeProof
func RPVerify(rp RangeProof) bool 
```

## 

## Zkp_for_ledgers

对于账本的其他几个零知识证明：

格式正确证明：输入公钥、加密原文、随机数、加密密文，输出非交互式证明，来证明密文的两部分中r1=r2

```
func GenFormatProof(pub PubKey, plainText []byte, r *big.Int, enc Enc) EP
func VerifyFormatProof(ep EP) bool
```

会计平衡证明：输入原始、转账、收款承诺，和三个对应的金额，生成证明

```
func GenBalanceProof(cmo, cms, cmr ECPoint, vo, vs, vr *big.Int) LEP
func VerifyBalanceProof(lep LEP) bool
```

相等证明：输入生成两个密文的密钥（一般是接受者和监管者）、两个密文、密文原文，来证明给两个人加密的信息相同。生成ep

```
func GenEqualityProof(pub1 PubKey, pub2 PubKey, enc1,enc2 Enc, r1,r2 *big.Int, plainText []byte) EP
func VerifyEqualityProof(ep EP) bool
```