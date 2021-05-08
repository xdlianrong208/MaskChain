package bp

import "C"
import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
)

type EP struct {
	G1 ECPoint
	G2 ECPoint
	Y1 ECPoint
	Y2 ECPoint
	T1 ECPoint
	T2 ECPoint
	S *big.Int
	C Hash
}

//TODO : 这个证明需不需要加两个生成元g1 g2 的阶参数
func EPProof (g1 ECPoint, g2 ECPoint, x *big.Int) EP {

	epResult := EP{}
	epResult.G1 = g1
	epResult.G2 = g2

	y1 := g1.Mult(x)
	y2 := g2.Mult(x)
	epResult.Y1 = y1
	epResult.Y2 = y2

	v, err := rand.Int(rand.Reader, EC.N)
	check(err)
	t1 := g1.Mult(v)
	t2 := g2.Mult(v)
	epResult.T1 = t1
	epResult.T2 = t2

	c := sha256.Sum256([]byte(g1.X.String()+g1.Y.String()+g2.X.String()+g2.Y.String()+y1.X.String()+y1.Y.String()+y2.X.String()+y2.Y.String()+t1.X.String()+t1.Y.String()+t2.X.String()+t2.Y.String()))
	epResult.C = c

	intc := new(big.Int).SetBytes(c[:])
	cx :=  new(big.Int).Mul(intc, x)
	s := new(big.Int).Sub(v, cx)
	s.Sub(EC.N, s)
	epResult.S = s

	return  epResult
}

func EPVerify (ep EP) bool{
	c := sha256.Sum256([]byte(ep.G1.X.String()+ep.G1.Y.String()+ep.G2.X.String()+ep.G2.Y.String()+ep.Y1.X.String()+ep.Y1.Y.String()+ep.Y2.X.String()+ep.Y2.Y.String()+ep.T1.X.String()+ep.T1.Y.String()+ep.T2.X.String()+ep.T2.Y.String()))
	intc := new(big.Int).SetBytes(c[:])

	if c!=ep.C{
		fmt.Println("Equality proof failed: c wrong")
		return false
	}
	ep.S.Sub(EC.N, ep.S)
	tempT1 := ep.G1.Mult(ep.S).Add(ep.Y1.Mult(intc))
	if !tempT1.Equal(ep.T1){
		fmt.Println("Equality proof failed: t1 wrong")
		return false
	}

	tempT2 := ep.G2.Mult(ep.S).Add(ep.Y2.Mult(intc))
	if !tempT2.Equal(ep.T2){
		fmt.Println("Equality proof failed: t2 wrong")
		return false
	}

	return true
}