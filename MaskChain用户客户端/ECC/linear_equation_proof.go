package bp

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"math/rand"
)

type LEP struct {
	Y ECPoint
	T ECPoint
	An []*big.Int
	B *big.Int
	Sn []*big.Int
	C Hash
}

// TODO: 现在的随机数生成限制在了-500 - 500 之间，因为要保证aivi相加等于0，生成匹配的随机数太耗时，所以数组a也最好不能太大，不知此处有何优化办法
func Linear_equation_proof (gn []ECPoint, xn []*big.Int, an []*big.Int, b *big.Int) LEP{

	lepResult := LEP{}

	if len(gn)!=len(xn) || len(gn)!= len(an) || len(xn)!= len(an) {
		panic("Generate LEP Failed: Array length error")
	}

	lepResult.An = an
	lepResult.B = b
	n := len(gn)


	tempy := gn[0].Mult(xn[0])
	for i:=1;i<n;i++{
		tempy = tempy.Add(gn[i].Mult(xn[i]))
	}
	lepResult.Y = tempy

	t := ECPoint{}
	vn := []*big.Int{}

	for{
		viai := big.NewInt(0)
		tmpvn := []*big.Int{}
		for i:=0;i<n;i++{
			tempv := rand.Intn(1000) - 500
			v := big.NewInt(int64(tempv))
			tmpvn = append(tmpvn,v)
			viai = new(big.Int).Add(viai, new(big.Int).Mul(v,an[i]))
		}
		if viai.CmpAbs(big.NewInt(0))==0{
			vn = tmpvn
			break
		}
	}

	t = gn[0].Mult(vn[0])
	for i:=1;i<n;i++{
		t = t.Add(gn[i].Mult(vn[i]))
	}
	lepResult.T = t

	var gnString string
	for i:=0;i<n;i++{
		gnString = gnString + gn[i].X.String() + gn[i].Y.String()
	}
	c := sha256.Sum256([]byte(gnString+tempy.X.String()+tempy.Y.String()+t.X.String()+t.Y.String()))
	lepResult.C = c
	intc := new(big.Int).SetBytes(c[:])

	sn := []*big.Int{}
	for i:=0;i<n;i++{
		cx :=  new(big.Int).Mul(intc, xn[i])
		sn = append(sn, new(big.Int).Sub(vn[i], cx))
	}
	lepResult.Sn = sn

	return  lepResult
}



func LepVerify(lep LEP, Gn []ECPoint) bool{

	var gnString string
	for i:=0;i<len(Gn);i++{
		gnString = gnString + Gn[i].X.String() + Gn[i].Y.String()
	}
	c := sha256.Sum256([]byte(gnString+lep.Y.X.String()+lep.Y.Y.String()+lep.T.X.String()+lep.T.Y.String()))
	if c!= lep.C{
		fmt.Println("lep failed: c wrong")
		return  false
	}

	intc := new(big.Int).SetBytes(c[:])
	n := len(Gn)
	gisi := Gn[0].Mult(lep.Sn[0])
	for i:=1;i<n;i++{
		gisi = gisi.Add(Gn[i].Mult(lep.Sn[i]))
	}
	tempT := lep.Y.Mult(intc).Add(gisi)
	if !tempT.Equal(lep.T){
		fmt.Println("lep failed: t wrong")
		return false
	}

	aisi := new(big.Int).Mul(lep.An[0], lep.Sn[0])
	for i:=1;i<n;i++{
		aisi = new(big.Int).Add(aisi,new(big.Int).Mul(lep.An[i], lep.Sn[i]))
	}
	if aisi.Cmp(new(big.Int).Mul(new(big.Int).Mul(intc, lep.B),big.NewInt(-1)))!=0{
		fmt.Println("lep failed: -cb wrong")
		return false
	}

	return true
}

type LEP_tx struct {
	Y ECPoint
	T ECPoint
	Sn []*big.Int
	C Hash
}

func Linear_equation_proof_tx (gn []ECPoint, xn []*big.Int, an []*big.Int) LEP_tx{

	lepResult := LEP_tx{}

	if len(gn)!=len(xn) || len(gn)!= len(an) || len(xn)!= len(an) {
		panic("Generate LEP Failed: Array length error")
	}

	n := len(gn)

	tempy := gn[0].Mult(xn[0])
	for i:=1;i<n;i++{
		tempy = tempy.Add(gn[i].Mult(xn[i]))
	}
	lepResult.Y = tempy

	t := ECPoint{}
	vn := []*big.Int{}

	for{
		viai := big.NewInt(0)
		tmpvn := []*big.Int{}
		for i:=0;i<n;i++{
			tempv := rand.Intn(1000) - 500
			v := big.NewInt(int64(tempv))
			tmpvn = append(tmpvn,v)
			viai = new(big.Int).Add(viai, new(big.Int).Mul(v,an[i]))
		}
		if viai.CmpAbs(big.NewInt(0))==0{
			vn = tmpvn
			break
		}
	}

	t = gn[0].Mult(vn[0])
	for i:=1;i<n;i++{
		t = t.Add(gn[i].Mult(vn[i]))
	}
	lepResult.T = t

	var gnString string
	for i:=0;i<n;i++{
		gnString = gnString + gn[i].X.String() + gn[i].Y.String()
	}
	c := sha256.Sum256([]byte(gnString+tempy.X.String()+tempy.Y.String()+t.X.String()+t.Y.String()))
	lepResult.C = c
	intc := new(big.Int).SetBytes(c[:])

	sn := []*big.Int{}
	for i:=0;i<n;i++{
		cx :=  new(big.Int).Mul(intc, xn[i])
		sni := new(big.Int).Sub(vn[i], cx)
		sni.Sub(EC.N, sni)
		sn = append(sn, sni)
	}
	lepResult.Sn = sn

	return  lepResult
}

func LepVerify_tx(lep LEP_tx, Gn []ECPoint) bool{

	var gnString string
	for i:=0;i<len(Gn);i++{
		gnString = gnString + Gn[i].X.String() + Gn[i].Y.String()
	}
	c := sha256.Sum256([]byte(gnString+lep.Y.X.String()+lep.Y.Y.String()+lep.T.X.String()+lep.T.Y.String()))
	if c!= lep.C{
		fmt.Println("lep failed: c wrong")
		return  false
	}

	intc := new(big.Int).SetBytes(c[:])
	n := len(Gn)
	lep.Sn[0].Sub(EC.N, lep.Sn[0])
	gisi := Gn[0].Mult(lep.Sn[0])
	for i:=1;i<n;i++{
		lep.Sn[i].Sub(EC.N, lep.Sn[i])
		gisi = gisi.Add(Gn[i].Mult(lep.Sn[i]))
	}
	tempT := lep.Y.Mult(intc).Add(gisi)
	if !tempT.Equal(lep.T){
		fmt.Println("lep failed: t wrong")
		return false
	}

	aisi := new(big.Int).Mul(big.NewInt(-1), lep.Sn[0])
	for i:=1;i<n;i++{
		aisi = new(big.Int).Add(aisi,new(big.Int).Mul(big.NewInt(1), lep.Sn[i]))
	}
	if aisi.Cmp(new(big.Int).Mul(new(big.Int).Mul(intc, big.NewInt(0)),big.NewInt(-1)))!=0{
		fmt.Println("lep failed: -cb wrong")
		return false
	}

	return true
}