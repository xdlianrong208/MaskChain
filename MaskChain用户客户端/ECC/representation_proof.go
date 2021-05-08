package bp

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
)

type REP struct {
	N int
	Sn []*big.Int
	Gn []ECPoint
	Y ECPoint
	T ECPoint
	C Hash
}

func RepProof(gn []ECPoint,xn []*big.Int) REP{

	repResult := REP{}

	repResult.Gn = gn

	n := len(gn)
	repResult.N = n
	if len(gn)!= len(xn) {
		panic("x and y are not the same length")
	}

	tempy := gn[0].Mult(xn[0])

	for i:=1;i<n;i++{
		tempy = tempy.Add(gn[i].Mult(xn[i]))
	}
	//tempy = AddP(n, gn, xn)
	repResult.Y = tempy


	vn := []*big.Int{}
	for i:=0;i<n;i++{
		v, err := rand.Int(rand.Reader, EC.N)
		check(err)
		vn = append(vn, v)
	}
	vn = append(vn, big.NewInt(11))
	vn = append(vn, big.NewInt(12))
	t := gn[0].Mult(vn[0])

	for i:=1;i<n;i++{
		t = t.Add(gn[i].Mult(vn[i]))
	}
	repResult.T = t

	var gnString string
	for i:=0;i<n;i++{
		gnString = gnString + gn[i].X.String() + gn[i].Y.String()
	}
	c := sha256.Sum256([]byte(gnString+tempy.X.String()+tempy.Y.String()+t.X.String()+t.Y.String()))
	repResult.C = c
	intc := new(big.Int).SetBytes(c[:])

	sn := []*big.Int{}
	for i:=0;i<n;i++{
		cx :=  new(big.Int).Mul(intc, xn[i])
		sn = append(sn, new(big.Int).Sub(vn[i], cx))
	}
	repResult.Sn = sn

	return repResult
}

func RepVerify(rep REP) bool {
	var gnString string
	for i:=0;i<rep.N;i++{
		gnString = gnString + rep.Gn[i].X.String() + rep.Gn[i].Y.String()
	}
	c := sha256.Sum256([]byte(gnString+rep.Y.X.String()+rep.Y.Y.String()+rep.T.X.String()+rep.T.Y.String()))
	if c!= rep.C{
		fmt.Println("REP failed: c != rep.C")
		return false
	}
	intc := new(big.Int).SetBytes(c[:])

	tempt := rep.Y.Mult(intc)
	for i:=0;i<rep.N;i++{
		tempt = tempt.Add(rep.Gn[i].Mult(rep.Sn[i]))
	}
	if !tempt.Equal(rep.T){
		fmt.Println("REP failed: t wrong")
		return false
	}
	return true
}