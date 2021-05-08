package bp

import (
	"fmt"
	"testing"
)

func TestGenFormatProof(t *testing.T) {
	pub, _, _ := GenerateKeys("Trump, forever God!")
	cypher, comm, _ := EncryptValue(pub, uint64(12))
	ep := GenerateFormatProof(pub, uint64((12)), comm.R, cypher)
	if VerifyFormatProof(cypher, ep){
		fmt.Println("Format Proof works")
	} else {fmt.Println("format proof failed")}
}

func TestGenBalanceProof(t *testing.T) {
	// Testing smallest number in range
	pub1, _, _ := GenerateKeys("Trump, forever God!1")
	pub2, _, _ := GenerateKeys("Trump, forever God!2")
	pub3, _, _ := GenerateKeys("Trump, forever God!3")

	_, commr, _ := EncryptValue(pub2, uint64(3))
	_, comms, _ := EncryptValue(pub3, uint64(2))
	_, commo, _ := EncryptValue(pub1, uint64(5))

	blp := GenerateBalanceProof(uint64(3),uint64(2),uint64(5),commr.Commitment, comms.Commitment,commo.Commitment)

	if VerifyBalanceProof(commr.Commitment, comms.Commitment,commo.Commitment,blp){
		fmt.Println("Balance Proof works")
	} else {fmt.Println("Balance proof failed")}
}

func TestGenEqualityProof(t *testing.T) {
	pub1, _, _ := GenerateKeys("Trump, forever God!1")
	pub2, _, _ := GenerateKeys("Trump, forever God!2")

	_, comm1, _ := EncryptValue(pub1, uint64(100))
	_, comm2, _ := EncryptValue(pub2, uint64(100))


	epp := GenerateEqualityProof(pub1, pub2, comm1, comm2, uint(100))

	if VerifyEqualityProof(epp){
		fmt.Println("Equality Proof works")
	} else {fmt.Println("Equality proof failed")}
}

func TestGenAddrEqualityProof(t *testing.T) {
	pub1, _, _ := GenerateKeys("Trump, forever God!")
	_, CMrpk, _ := EncryptAddress(pub1, []byte("Make USA Great Again!"))
	epp := GenerateAddressEqualityProof(pub1, pub1, CMrpk, CMrpk, []byte("Make USA Great Again!"))
	if VerifyEqualityProof(epp){
		fmt.Println("Equality Proof works")
	} else {fmt.Println("Equality proof failed")}
}
