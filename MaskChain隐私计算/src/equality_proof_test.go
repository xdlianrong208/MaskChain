package bp

import (
	"fmt"
	"testing"
)


func TestEpVerify1(t *testing.T) {
	//EC = NewECPrimeGroupKey(64)
	//
	//v1, err := rand.Int(rand.Reader, EC.N)
	//check(err)
	//v2, err := rand.Int(rand.Reader, EC.N)
	//check(err)
	//
	//x1 := EC.G.Mult(v1)
	//x2 := EC.G.Mult(v2)
	//
	//a1 := big.NewInt(1)
	//Genep1 := EPProof(x1, x2, a1)
	//Genep1.ToHex()
	//
	//if EPVerify(Genep1) {
	//	fmt.Println("Equality Proof Verification works")
	//} else {
	//	t.Error("*****Equality ProofFAILURE")
	//}
}

func (e EP) ToHex() {
	//fmt.Println("g1.X:",fmt.Sprintf("%x", e.G1.X.Bytes()))
	//fmt.Println("g1.Y:",fmt.Sprintf("%x", e.G1.Y.Bytes()))
	//fmt.Println("g2.X:",fmt.Sprintf("%x", e.G2.X.Bytes()))
	//fmt.Println("g2.Y:",fmt.Sprintf("%x", e.G2.X.Bytes()))
	fmt.Println("t1.X:",fmt.Sprintf("%x", e.T1.X.Bytes()))
	fmt.Println("t1.Y:",fmt.Sprintf("%x", e.T2.X.Bytes()))
	fmt.Println("y1.X:",fmt.Sprintf("%x", e.Y1.X.Bytes()))
	fmt.Println("y1.Y:",fmt.Sprintf("%x", e.Y2.X.Bytes()))
	fmt.Println("S:",fmt.Sprintf("%x", e.S.Bytes()))
	fmt.Println("C:",fmt.Sprintf("%x", e.C))
}
