package bp

import (
	"fmt"
	"testing"
)

func TestBase(t *testing.T){
	fmt.Printf("\n\n========================= EXAMPLE 1 =========================\n\n")
	pub, priv, err := GenerateKeys("五点共圆")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("公钥：\nP:%x\nG1:%x\nG2:%x\nH:%x\n私钥：\nX:%x\n", pub.P, pub.G1, pub.G2, pub.H, priv.X)
	//
	//C := Encrypt(pub, []byte("1"))
	//fmt.Printf("\n加密后的密文C1为：%x\n加密后的密文C2为：%x\n", C.C1, C.C2)
	//
	//M := Decrypt(priv, C)
	//M_word := string(M)
	//fmt.Printf("\n解密后的明文为：%s\n", M_word)
	//
	//tem,_ := ecdsa.GenerateKey(btcec.S256(),rand.Reader)
	//fmt.Println(tem.PublicKey, tem.D)
	//PrivKey := GenKeys("123")
	////PrivKey.PubKey.H.X = tem.PublicKey.X
	////PrivKey.PubKey.H.Y = tem.PublicKey.Y
	////PrivKey.X = tem.D
	//r,s,_,_ := PrivKey.Sign([]byte("1"))
	//
	//
	//a := PrivKey.VerifySign([]byte("1"),r,s)
	//fmt.Println("------------sign----------:",a)


	//_, comm, _ := EncryptValue(pub, uint64(20))
	//fmt.Println(new(big.Int).SetBytes(comm.Commitment))
	//fmt.Println(new(big.Int).SetBytes(comm.R))

	t1 := ECCEncrypt(pub, []byte(string("12345, go hunt tiger in the mountain")))
	fmt.Println("encrypt message: ", "12345, go hunt tiger in the mountain")
	//fmt.Println("ecc公钥加密后：", hex.EncodeToString(t1.C1))

	fmt.Println("dcrypt message: ",string(ECCDecrypt(priv,t1)))
	//sig := Sign(priv, []byte("1"))
	//M_word := string(sig.M)
	//Mx_word := new(big.Int).SetBytes(sig.M_hash)
	//R_word := new(big.Int).SetBytes(sig.R)
	//S_word := new(big.Int).SetBytes(sig.S)
	//fmt.Printf("\n明文为：%s\n明文哈希为：%x\n签名R为：%x\n签名S为：%x\n", M_word, Mx_word, R_word, S_word)
	//
	//
	//
	//fmt.Printf("\n验证签名是否合法：\n")
	//verify := Verify(pub, sig)
	//if verify {
	//	fmt.Println("签名合法!")
	//} else {
	//	fmt.Println("签名不合法!")
	//}
	//
	//fmt.Printf("\n篡改签名后验证签名是否合法：\n")
	//sig.S[0] += 1
	//verify = Verify(pub, sig)
	//if verify {
	//	fmt.Println("签名合法!")
	//} else {
	//	fmt.Println("签名不合法!")
	//}

}
