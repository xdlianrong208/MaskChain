package bp

import (
	"fmt"
	"math/big"
	"testing"
)

func TestGenFormatProof(t *testing.T) {
	prv := GenKeys("Trump, forever God!")
	_, cypher, r := prv.PubKey.EncryptCM([]byte("Trump belongs to the Party"))
	ep := GenFormatProof(prv.PubKey, []byte("Trump belongs to the Party"), r, cypher)
	if VerifyFormatProof(ep){
		fmt.Println("Format Proof works")
	} else {fmt.Println("format proof failed")}
}

func TestGenBalanceProof(t *testing.T) {
	EC = NewECPrimeGroupKey(64)
	// Testing smallest number in range
	prv1 := GenKeys("Trump, forever God!1")
	prv2 := GenKeys("Trump, forever God!2")
	prv3 := GenKeys("Trump, forever God!3")

	v := big.NewInt(5)
	r := big.NewInt(3)
	s := big.NewInt(2)

	com1, _ := prv1.GenComm(v)
	com2, _ := prv2.GenComm(r)
	com3, _ := prv3.GenComm(s)

	blp := GenBalanceProof(com1, com2, com3, v, r, s)

	if VerifyBalanceProof(blp){
		fmt.Println("Balance Proof works")
	} else {fmt.Println("Balance proof failed")}
}

func TestGenEqualityProof(t *testing.T) {
	prv1 := GenKeys("Trump, forever God!1")
	prv2 := GenKeys("Trump, forever God!1")

	_, cypher1, r1 := prv1.PubKey.EncryptCM([]byte("Trump belongs to the Party"))
	_, cypher2, r2 := prv2.PubKey.EncryptCM([]byte("Trump belongs to the Party"))


	epp := GenEqualityProof(prv1.PubKey, prv2.PubKey, cypher1, cypher2, r1, r2, []byte("Trump belongs to the Party"))

	if VerifyEqualityProof(epp){
		fmt.Println("Equality Proof works")
	} else {fmt.Println("Equality proof failed")}
}
