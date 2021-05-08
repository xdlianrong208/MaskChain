package utils

import (
	ecc "exchange/crypto/ECC"
	"math/rand"
	"strconv"
)

func CreateDE_CM(regpub ecc.PublicKey, amount string) (C ecc.CypherText,CM ecc.Commitment) {
	amounts, _ := strconv.Atoi(amount)
	C, CM, _ = ecc.EncryptValue(regpub, uint64(amounts))
	return
}

// create commit_v
func CreateCM_v(regpub ecc.PublicKey, amount string) (CM ecc.Commitment) {
	amounts, _ := strconv.Atoi(amount)
	rF := rand.Uint64()
	r1 := strconv.FormatUint(rF, 16)
	CM = regpub.CommitByUint64(uint64(amounts), []byte(r1))
	return
}

// create elgamal result
func CreateElgamalInfo(regpub ecc.PublicKey, amount string, publickey string) (C ecc.CypherText) {
	M := publickey + amount
	C = ecc.Encrypt(regpub, []byte(M))
	return
}

func CreateElgamalR(regpub ecc.PublicKey, r []byte) (C ecc.CypherText) {
	C = ecc.Encrypt(regpub, r)
	return
}

// create sign result
func CreateSign(privpub ecc.PrivateKey, amount string) (sig ecc.Signature) {
	ID := "1"
	sig = ecc.Sign(privpub, []byte(ID+amount))
	return
}

func CreateUsrPub(g1 string, g2 string, p string, h string) (usrpub ecc.PublicKey) {
	usrpub.G1 = stringtobig(g1, 16)
	usrpub.G2 = stringtobig(g2, 16)
	usrpub.P = stringtobig(p, 16)
	usrpub.H = stringtobig(h, 16)
	return
}
