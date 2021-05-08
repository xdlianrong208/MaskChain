package bp

import (
	"fmt"
	"math/big"
	"testing"
)

func TestDLPVerify1(t *testing.T) {
	EC = NewECPrimeGroupKey(64)
	// Testing smallest number in range
	if DLPVerify(Discrete_Logarithm_Proof(big.NewInt(64))) {
		fmt.Println("Discrete Logarithm Proof Verification works")
	} else {
		t.Error("*****Discrete Logarithm Proof FAILURE")
	}
}
