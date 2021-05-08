package bp

import (
	"fmt"
	"math/big"
	"testing"
)


func TestRepVerify1(t *testing.T) {
	EC = NewECPrimeGroupKey(64)

	x1 := big.NewInt(1)
	x2 := big.NewInt(2)

	temp1 := big.NewInt(131421)
	temp2 := big.NewInt(421313)
	g1 := EC.G.Mult(temp1)
	g2 := EC.G.Mult(temp2)
	//y := g1.Mult(x1).Add(g2.Mult(x2))

	//v1, err := rand.Int(rand.Reader, EC.N)
	//check(err)
	//v2, err := rand.Int(rand.Reader, EC.N)
	//check(err)
	//v1 := big.NewInt(11)
	//v2 := big.NewInt(12)
	//t1 := g1.Mult(v1).Add(g2.Mult(v2))
	//
	//c := sha256.Sum256([]byte(g1.X.String()+g1.Y.String()+g2.X.String()+g2.Y.String()+y.X.String()+y.Y.String()+t1.X.String()+t1.Y.String()))
	//intc := new(big.Int).SetBytes(c[:])
	//
	//s1 := new(big.Int).Sub(v1,new(big.Int).Mul(x1,intc))
	//s2 := new(big.Int).Sub(v2,new(big.Int).Mul(x2,intc))

	//t2 := y.Mult(intc).Add(g1.Mult(s1)).Add(g2.Mult(s2))

	//fmt.Println("t1: ",t1)
	//fmt.Println("t2: ",t2)
	//fmt.Println("t1 = t2?", t1.Equal(t2))

	if RepVerify(RepProof([]ECPoint{g1,g2},[]*big.Int{x1,x2})) {
		fmt.Println("Representation Proof Verification works")
	} else {
		t.Error("*****Representation Proof FAILURE")
	}
}
