package bp

import "math/big"

func GenFormatProof(pub PubKey, plainText []byte, r *big.Int, enc Enc) EP {
	hr := enc.P1.Add(pub.G1.Mult(new(big.Int).SetBytes(plainText[:])).Neg())
	return EPProof(hr, enc.P2, r)
}

func VerifyFormatProof(ep EP) bool {
	return EPVerify(ep)
}

func GenBalanceProof(cmo, cms, cmr ECPoint, vo, vs, vr *big.Int) LEP {
	return Linear_equation_proof([]ECPoint{cmo, cms, cmr}, []*big.Int{vo, vs, vr}, []*big.Int{big.NewInt(-1), big.NewInt(1), big.NewInt(1)}, big.NewInt(0))
}

func VerifyBalanceProof(lep LEP) bool {
	return LepVerify(lep)
}

func GenEqualityProof(pub1 PubKey, pub2 PubKey, enc1,enc2 Enc, r1,r2 *big.Int, plainText []byte) EP {
	return EPProof(enc1.P1.Add(pub1.H.Mult(r1).Neg()), enc2.P1.Add(pub2.H.Mult(r2).Neg()), new(big.Int).SetBytes(plainText[:]))
}

func VerifyEqualityProof(ep EP) bool {
	return EPVerify(ep)
}
