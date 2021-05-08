package bp

import (
	"crypto/elliptic"
	"encoding/binary"
	"math/big"
)

type PublicKey struct {
	G1, G2, P, H *big.Int
}
type PrivateKey struct {
	PublicKey
	X *big.Int
}
type CypherText struct {
	C1, C2 []byte
}
type Commitment struct {
	Commitment, R []byte
}
type Signature struct {
	M, M_hash, R, S []byte
}

func Encrypt(pub PublicKey, M []byte) (C CypherText) {
	p1 := ConvertPub(pub)
	_, m1 := p1.Encrypt(M)
	c1 := elliptic.Marshal(EC.C, m1.P1.X, m1.P1.Y)
	c2 := elliptic.Marshal(EC.C, m1.P2.X, m1.P2.Y)
	return CypherText{c1, c2}
}

func Decrypt(priv PrivateKey, C CypherText) (M []byte) {
	p1 := ConvertPub(priv.PublicKey)
	pr := PrivKey{p1, priv.X}
	x1, y1 := elliptic.Unmarshal(EC.C, C.C1)
	x2, y2 := elliptic.Unmarshal(EC.C, C.C2)
	return pr.Decrypt(Enc{ECPoint{x1, y1}, ECPoint{x2, y2}})
}

func GenerateKeys(info string) (pub PublicKey, priv PrivateKey, err error) {
	prv := GenKeys(info)
	pubb := RecoverPub(prv.PubKey)
	prvv := PrivateKey{pubb, prv.X}
	return pubb, prvv, nil
}

func (pub PublicKey) Commit(v *big.Int, rnd []byte) Commitment {
	pub1 := ConvertPub(pub)
	com := pub1.G1.Mult(v).Add(pub1.H.Mult(new(big.Int).SetBytes(rnd)))
	com1 := elliptic.Marshal(EC.C, com.X, com.Y)
	return Commitment{com1, rnd}
}

func (pub PublicKey) CommitByBytes(b []byte, rnd []byte) Commitment {
	pub1 := ConvertPub(pub)
	v := new(big.Int).SetBytes(b)
	com := pub1.G1.Mult(v).Add(pub1.H.Mult(new(big.Int).SetBytes(rnd)))
	com1 := elliptic.Marshal(EC.C, com.X, com.Y)
	return Commitment{com1, rnd}
}

func (pub PublicKey) CommitByUint64(v uint64, rnd []byte) Commitment {
	v_ := new(big.Int).SetUint64(v)
	return pub.Commit(v_, rnd)
}

func EncryptValue(pub PublicKey, M uint64) (C CypherText, commit Commitment, err error) {
	v := make([]byte, 8)
	binary.BigEndian.PutUint64(v, M)
	pubb := ConvertPub(pub)
	comm, cipher, r := pubb.EncryptCM(v)
	c1 := elliptic.Marshal(EC.C, cipher.P1.X, cipher.P1.Y)
	c2 := elliptic.Marshal(EC.C, cipher.P2.X, cipher.P2.Y)
	com1 := elliptic.Marshal(EC.C, comm.X, comm.Y)
	return CypherText{c1, c2}, Commitment{com1, r.Bytes()}, nil
}

func EncryptAddress(pub PublicKey, addr []byte) (C CypherText, commit Commitment, err error) {
	addr_uint64 := binary.BigEndian.Uint64(addr)
	return EncryptValue(pub, addr_uint64)
}

func (pub PublicKey) VerifyCommitment(commit Commitment) uint64 {
	x, y := elliptic.Unmarshal(EC.C, commit.Commitment)
	com := ECPoint{x, y}
	pub1 := ConvertPub(pub)
	v := big.NewInt(0)
	for {
		if pub1.G1.Mult(v).Add(pub1.G2.Mult(new(big.Int).SetBytes(commit.R))).Equal(com) {
			return 1
		}
		v = new(big.Int).Add(v, big.NewInt(1))
		if v == big.NewInt(50000) {
			break
		}
	}
	return 0
}

func Sign(priv PrivateKey, m []byte) (sig Signature) {
	prvv := PrivKey{ConvertPub(priv.PublicKey), priv.X}
	sig = Signature{}
	r, s, _, h := prvv.Sign(m)
	sig.M = m
	sig.R = r
	sig.S = s
	sig.M_hash = h
	return sig
}

func Verify(pub PublicKey, sig Signature) bool {
	pubb := ConvertPub(pub)
	return pubb.VerifySign(sig.M, sig.R, sig.S)
}

func ConvertPub(pub PublicKey) PubKey {
	re := PubKey{
		G1: EC.G.Mult(big.NewInt(0)),
		G2: EC.G.Mult(big.NewInt(0)),
	}
	re.G1.X, re.G1.Y = elliptic.Unmarshal(EC.C, pub.G1.Bytes())
	re.G2.X, re.G2.Y = elliptic.Unmarshal(EC.C, pub.G2.Bytes())
	re.H.X, re.H.Y = elliptic.Unmarshal(EC.C, pub.H.Bytes())
	return re
}
func RecoverPub(pub PubKey) PublicKey {
	re := PublicKey{
		G1: big.NewInt(0),
		G2: big.NewInt(0),
		P:  big.NewInt(0),
		H:  big.NewInt(0),
	}
	re.G1 = new(big.Int).SetBytes(elliptic.Marshal(EC.C, pub.G1.X, pub.G1.Y))
	re.G2 = new(big.Int).SetBytes(elliptic.Marshal(EC.C, pub.G2.X, pub.G2.Y))
	re.P = EC.N
	re.H = new(big.Int).SetBytes(elliptic.Marshal(EC.C, pub.H.X, pub.H.Y))
	return re
}
