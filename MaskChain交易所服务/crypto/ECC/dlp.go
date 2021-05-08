package bp

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	_ "github.com/btcsuite/btcd/btcec"
	"math/big"
)

type DLP struct {
	Y ECPoint
	T ECPoint
	C Hash
	S *big.Int
}
// 此证明没带有G 参数，所用G为secp256k1默认的G生成元 G= 04 79BE667E F9DCBBAC 55A06295 CE870B07 029BFCDB 2DCE28D9 59F2815B 16F81798 483ADA77 26A3C465 5DA4FBFC 0E1108A8 FD17B448 A6855419 9C47D08F FB10D4B8
// 证明中G的阶为 n = FFFFFFFF FFFFFFFF FFFFFFFF FFFFFFFE BAAEDCE6 AF48A03B BFD25E8C D0364141
func Discrete_Logarithm_Proof(x *big.Int) DLP{

	dlpResult := DLP{}

	y := EC.G.Mult(x)
	dlpResult.Y = y

	v, err := rand.Int(rand.Reader, EC.N)
	check(err)

	t := EC.G.Mult(v)
	dlpResult.T = t
	//fmt.Println("t: ",t)
	c := sha256.Sum256([]byte(EC.G.X.String() + EC.G.Y.String()+y.X.String()+y.Y.String()+t.X.String()+t.Y.String()))
	dlpResult.C = c

	intc := new(big.Int).SetBytes(c[:])
	cx :=  new(big.Int).Mul(intc, x)
	s := new(big.Int).Sub(v, cx)
	dlpResult.S = s

	return  dlpResult
}

func DLPVerify(dlp DLP) bool{

	tempC := sha256.Sum256([]byte(EC.G.X.String() + EC.G.Y.String()+dlp.Y.X.String()+dlp.Y.Y.String()+dlp.T.X.String()+dlp.T.Y.String()))
	if tempC != dlp.C {
		fmt.Println("DLP failed: tem[C != dlp.C")
		return false
	}
	tempintc := new(big.Int).SetBytes(tempC[:])

	tempT :=  EC.G.Mult(dlp.S).Add(dlp.Y.Mult(tempintc))
	//fmt.Println("dlp.S: ",dlp.S,"dlp.Y.Mult(tempintc): ",dlp.Y.Mult(tempintc))
	//fmt.Println("tempT: ",tempT)
	if !tempT.Equal(dlp.T){
		fmt.Println("DLP failed: tempT != dlp.T")
		return false
	}

	return true
}