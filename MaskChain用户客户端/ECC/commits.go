package bp

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math/big"
)

type PubKey struct {
	G1 ECPoint
	G2 ECPoint
	H ECPoint
}

type PrivKey struct {
	PubKey
	X *big.Int
}

type Enc struct {
	P1 ECPoint
	P2 ECPoint
}

func GenKeys(s string) PrivKey {

	x1 := []byte(s)
	x := new(big.Int).SetBytes(x1[:])
	Key := PrivKey{}
	v1, err := rand.Int(rand.Reader, EC.N)
	check(err)
	//v2, err := rand.Int(rand.Reader, EC.N)
	//check(err)
	Key.G1 = EC.G.Mult(v1)
	Key.G2.X = EC.C.Params().Gx
	Key.G2.Y = EC.C.Params().Gy
	Key.X = x
	Key.H = Key.G2.Mult(x)

	return Key
}

// genetate pederson commitment: v*g + r*h, return commitment and random value r
func (pub PubKey) GenComm(v *big.Int) (ECPoint,*big.Int) {

	r, err := rand.Int(rand.Reader, EC.N)
	check(err)

	com := pub.G1.Mult(v).Add(pub.H.Mult(r))

	return com,r
}

func (pub PubKey) GenComByBytes(b []byte) (ECPoint,*big.Int) {

	v := new(big.Int).SetBytes(b[:])

	r, err := rand.Int(rand.Reader, EC.N)
	check(err)

	com := pub.G1.Mult(v).Add(pub.H.Mult(r))

	return com,r
}

func (pub PubKey) Encrypt(plainText []byte) (*big.Int, Enc){

	v := new(big.Int).SetBytes(plainText[:])

	r, err := rand.Int(rand.Reader, big.NewInt(50000))
	check(err)
	t1 := pub.G1.Mult(v).Add(pub.H.Mult(r))
	t2 := pub.G2.Mult(r)

	return  v, Enc{t1,t2}
}


// The private key and plaintext are passed in for decryption
func (priv PrivKey) Decrypt(enc Enc)(msg []byte){
	g1v := enc.P1.Add(enc.P2.Mult(priv.X).Neg())
	v := big.NewInt(0)
	for{
		if priv.G1.Mult(v).Equal(g1v){
			break
		}
		v = new(big.Int).Add(v,big.NewInt(1))
		if v == big.NewInt(50000){
			return []byte("解密了个寂寞")
			break
		}
	}
	plainText := v.Bytes()
	return plainText
}

// equal r for commitment and cyphertext, return commitment for plaitext, cyphertext, random r
func (pub PubKey) EncryptCM(plainText []byte) (ECPoint, Enc, *big.Int){
	v := new(big.Int).SetBytes(plainText[:])

	r, err := rand.Int(rand.Reader, EC.N)
	check(err)
	t1 := pub.G1.Mult(v).Add(pub.H.Mult(r))
	t2 := pub.G2.Mult(r)

	com := pub.G1.Mult(v).Add(pub.H.Mult(r))

	return  com, Enc{t1,t2}, r
}

func (priv PrivKey) DecryptCM(cyperText Enc) uint64 {
	gv := cyperText.P1.Add(cyperText.P2.Mult(priv.X).Neg())
	for i := 1; true; i++ {
		m := new(big.Int).SetInt64(int64(i))
		if(gv.Equal(priv.G1.Mult(m))){
			return uint64(i)
		}
		if i >= 262144 {
			fmt.Printf("该承诺并非价值承诺或承诺价值大于262144")
			return 0
		}
	}
	return 0
}

func (priv PrivKey) Sign(msg []byte)([]byte, []byte, error, []byte){
	Key := ecdsa.PublicKey{}
	Key.X = priv.PubKey.H.X
	Key.Y = priv.PubKey.H.Y


	Key.Curve = EC.C

	PrivKey := ecdsa.PrivateKey{}
	PrivKey.PublicKey = Key
	PrivKey.D = priv.X
	//// Convert to the private key in the ecies package in the ethereum package
	//
	myhash := sha256.New()
	resultHash := myhash.Sum(msg)
	r, s, err := ecdsa.Sign(rand.Reader, &PrivKey, resultHash)
	if err!=nil{
		return nil,nil,err,resultHash
	}
	return r.Bytes(),s.Bytes(),nil,resultHash
}

func (pub PubKey) VerifySign(msg []byte, rText, sText []byte) bool {
	Key := ecdsa.PublicKey{}
	Key.X = pub.H.X
	Key.Y = pub.H.Y


	Key.Curve = EC.C

	myhash := sha256.New()
	resultHash := myhash.Sum(msg)
	r := new(big.Int).SetBytes(rText)
	s := new(big.Int).SetBytes(sText)
	result := ecdsa.Verify(&Key, resultHash, r, s)
	return result
}

//整形转换成字节
func IntToBytes(n int) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}
//字节转换成整形
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int(x)
}