package bp

import "C"
import (
	"crypto/elliptic"
	"encoding/binary"
	"math/big"
)

type FormatProof struct {
	G1, G2 []byte
	Y1, Y2 []byte
	T1, T2 []byte
	S  []byte
	C  []byte
}

type BalanceProof struct {
	Y, T []byte
	Sn_1, Sn_2, Sn_3 []byte
	C []byte
}

type EqualityProof struct {
	FormatProof
}

func GenerateFormatProof(pub PublicKey, v uint64, r []byte, enc CypherText) (fp FormatProof) {
	pubb := ConvertPub(pub)
	rr := new(big.Int).SetBytes(r)
	x1, y1 := elliptic.Unmarshal(EC.C, enc.C1)
	enc_1 := ECPoint{x1,y1}
	x2, y2 := elliptic.Unmarshal(EC.C, enc.C2)
	enc_2 := ECPoint{x2,y2}
	hr := enc_1.Add(pubb.G1.Mult(new(big.Int).SetUint64(v)).Neg())
	formatproof := EPProof(hr, enc_2, rr)
	fp.G1 = elliptic.Marshal(EC.C, formatproof.G1.X, formatproof.G1.Y)
	fp.G2 = elliptic.Marshal(EC.C, formatproof.G2.X, formatproof.G2.Y)
	fp.Y1 = elliptic.Marshal(EC.C, formatproof.Y1.X, formatproof.Y1.Y)
	fp.Y2 = elliptic.Marshal(EC.C, formatproof.Y2.X, formatproof.Y2.Y)
	fp.T1 = elliptic.Marshal(EC.C, formatproof.T1.X, formatproof.T1.Y)
	fp.T2 = elliptic.Marshal(EC.C, formatproof.T2.X, formatproof.T2.Y)
	fp.S = formatproof.S.Bytes()
	fp.C = formatproof.C[:]
	return
}

func VerifyFormatProof(Ct CypherText, fp FormatProof) bool {
	formatproof := EP{}
	formatproof.G1.X, formatproof.G1.Y = elliptic.Unmarshal(EC.C, fp.G1)
	formatproof.G2.X, formatproof.G2.Y = elliptic.Unmarshal(EC.C, fp.G2)
	formatproof.Y1.X, formatproof.Y1.Y = elliptic.Unmarshal(EC.C, fp.Y1)
	formatproof.Y2.X, formatproof.Y2.Y = elliptic.Unmarshal(EC.C, fp.Y2)
	formatproof.T1.X, formatproof.T1.Y = elliptic.Unmarshal(EC.C, fp.T1)
	formatproof.T2.X, formatproof.T2.Y = elliptic.Unmarshal(EC.C, fp.T2)
	formatproof.S = new(big.Int).SetBytes(fp.S)
	formatproof.C = BytesToHash(fp.C)
	return EPVerify(formatproof)
}

func GenerateBalanceProof(vR, vS, vO uint64, cmr, cms, cmo []byte) BalanceProof {
	R := new(big.Int).SetUint64(vR)
	S := new(big.Int).SetUint64(vS)
	O := new(big.Int).SetUint64(vO)
	rx, ry := elliptic.Unmarshal(EC.C, cmr)
	commr := ECPoint{rx,ry}
	sx, sy := elliptic.Unmarshal(EC.C, cms)
	comms := ECPoint{sx,sy}
	ox, oy := elliptic.Unmarshal(EC.C, cmo)
	commo := ECPoint{ox,oy}
	linearproof := Linear_equation_proof_tx([]ECPoint{commo, comms, commr}, []*big.Int{O, S, R}, []*big.Int{big.NewInt(-1), big.NewInt(1), big.NewInt(1)})
	bp := BalanceProof{}
	bp.Y = elliptic.Marshal(EC.C, linearproof.Y.X, linearproof.Y.Y)
	bp.T = elliptic.Marshal(EC.C, linearproof.T.X, linearproof.T.Y)
	bp.Sn_1 = linearproof.Sn[0].Bytes()
	bp.Sn_2 = linearproof.Sn[1].Bytes()
	bp.Sn_3 = linearproof.Sn[2].Bytes()
	bp.C = linearproof.C[:]
	return bp
}

func VerifyBalanceProof(CM_r, CM_s, CM_o []byte, bp BalanceProof) bool {
	linearproof := LEP_tx{}
	linearproof.Y.X, linearproof.Y.Y = elliptic.Unmarshal(EC.C, bp.Y)
	linearproof.T.X, linearproof.T.Y = elliptic.Unmarshal(EC.C, bp.T)
	linearproof.Sn = append(linearproof.Sn, new(big.Int).SetBytes(bp.Sn_1))
	linearproof.Sn = append(linearproof.Sn, new(big.Int).SetBytes(bp.Sn_2))
	linearproof.Sn = append(linearproof.Sn, new(big.Int).SetBytes(bp.Sn_3))
	linearproof.C = BytesToHash(bp.C)

	rx, ry := elliptic.Unmarshal(EC.C, CM_r)
	commr := ECPoint{rx,ry}
	sx, sy := elliptic.Unmarshal(EC.C, CM_s)
	comms := ECPoint{sx,sy}
	ox, oy := elliptic.Unmarshal(EC.C, CM_o)
	commo := ECPoint{ox,oy}
	return LepVerify_tx(linearproof,[]ECPoint{commo,comms,commr})
}

func GenerateEqualityProof(pub1, pub2 PublicKey, C1, C2 Commitment, v uint) (ep EqualityProof) {
	pubb1 := ConvertPub(pub1)
	pubb2 := ConvertPub(pub2)
	c1x, c1y := elliptic.Unmarshal(EC.C, C1.Commitment)
	c1comm := ECPoint{c1x,c1y}
	c2x, c2y := elliptic.Unmarshal(EC.C, C2.Commitment)
	c2comm := ECPoint{c2x,c2y}
	r1 := new(big.Int).SetBytes(C1.R)
	r2 := new(big.Int).SetBytes(C2.R)
	equalityproof := EPProof(c1comm.Add(pubb1.H.Mult(r1).Neg()), c2comm.Add(pubb2.H.Mult(r2).Neg()), big.NewInt(int64(v)))
	ep.G1 = elliptic.Marshal(EC.C, equalityproof.G1.X, equalityproof.G1.Y)
	ep.G2 = elliptic.Marshal(EC.C, equalityproof.G2.X, equalityproof.G2.Y)
	ep.Y1 = elliptic.Marshal(EC.C, equalityproof.Y1.X, equalityproof.Y1.Y)
	ep.Y2 = elliptic.Marshal(EC.C, equalityproof.Y2.X, equalityproof.Y2.Y)
	ep.T1 = elliptic.Marshal(EC.C, equalityproof.T1.X, equalityproof.T1.Y)
	ep.T2 = elliptic.Marshal(EC.C, equalityproof.T2.X, equalityproof.T2.Y)
	ep.S = equalityproof.S.Bytes()
	ep.C = equalityproof.C[:]
	return
}

func VerifyEqualityProof(ep EqualityProof) bool {
	equalityproof := EP{}
	equalityproof.G1.X, equalityproof.G1.Y = elliptic.Unmarshal(EC.C, ep.G1)
	equalityproof.G2.X, equalityproof.G2.Y = elliptic.Unmarshal(EC.C, ep.G2)
	equalityproof.Y1.X, equalityproof.Y1.Y = elliptic.Unmarshal(EC.C, ep.Y1)
	equalityproof.Y2.X, equalityproof.Y2.Y = elliptic.Unmarshal(EC.C, ep.Y2)
	equalityproof.T1.X, equalityproof.T1.Y = elliptic.Unmarshal(EC.C, ep.T1)
	equalityproof.T2.X, equalityproof.T2.Y = elliptic.Unmarshal(EC.C, ep.T2)
	equalityproof.S = new(big.Int).SetBytes(ep.S)
	equalityproof.C = BytesToHash(ep.C)
	return EPVerify(equalityproof)
}

func GenerateAddressEqualityProof(pub1, pub2 PublicKey, C1, C2 Commitment, addr []byte) (ep EqualityProof) {
	return GenerateEqualityProof(pub1, pub2, C1, C2, uint(binary.BigEndian.Uint64(addr)))
}
