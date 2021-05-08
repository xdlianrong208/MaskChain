package controllers

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	ecc "wallet/ECC"
	"wallet/model"
	"wallet/utils"
)

const (
	ErrorValue   = "value cannot be empty"
	RejectServer = "Server Error"
)

func Register(c echo.Context) error {
	w := new(model.NewWallet)
	// 因为 echo 的 bind 无绑定检查功能
	// echo 强制要求 post 的参数写在 body 里，写在 header 里会绑定不上
	if err := c.Bind(w); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	// 暂时只能验证是否为空
	if w.Id == "" || w.Name == "" || w.Str == "" {
		return c.JSON(http.StatusBadRequest, ErrorValue)
	}
	// 计算公私钥
	account := ecc.GenerateAccount(w.Str, w.Name, w.Id, w.Str)
	if res := register(account); res == "Successful!" {
		return c.JSON(http.StatusOK, account.KeyToString())
	} else {
		return c.JSON(http.StatusInternalServerError, RejectServer)
	}

	//_, priv, err := ELGamal.GenerateKeys(w.Str)
	//if err != nil {
	//	return c.JSON(http.StatusInternalServerError, err)
	//}
	// 取哈希  zr:默认注册时，将公钥发给监管者存入公钥池，这里不取hash
	// pub.G1 = new(big.Int)
	// HashInfoBuf := sha256.Sum256([]byte(w.Str))
	// 向监管者提交注册请求，并返回相关信息
	// [32]byte 是一个数组，要把他转换成切片
	//if resp, err := http.PostForm(RegulatorURL+"register", url.Values{"name": {w.Name}, "id": {w.Id}, "Hashky": {account.KeyToString().Publickey}}); err != nil {
	//	c.JSON(http.StatusInternalServerError, err)
	//	return c.JSON(http.StatusInternalServerError, err)
	//} else {
	//	if res, err := ioutil.ReadAll(resp.Body); err != nil {
	//		c.JSON(http.StatusInternalServerError, err)
	//		return c.JSON(http.StatusInternalServerError, err)
	//	} else {
	//		// 判断应该返回的信息
	//		if bytes.Equal(res,[]byte("Successful!")) {
	//			return c.JSON(http.StatusOK, account.KeyToString())
	//		} else {
	//			return c.JSON(http.StatusInternalServerError, RejectServer)
	//		}
	//	}
	//}
}

func register(account ecc.Account) string {
	data := account.Info
	body := ethRPCPost(data, RegulatorURL+"register")
	res := string(body)
	if res == "Successful!" {
		fmt.Println("账户" + account.Info.Name + "注册成功")
	} else if res == "Account registered!" {
		fmt.Println("账户" + account.Info.Name + "已注册")
	} else if res == "Fail!" {
		Fatalf("账户" + account.Info.Name + "注册失败")
	}
	return string(body)
}

// 39.105.58.136
func Buycoin(c echo.Context) error {
	w := new(model.BctoEx)
	if err := c.Bind(w); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	// 向交易所发出购币请求
	body := ethRPCPost(w, ExchangeURL + "buy")
	var receipt utils.Receipt
	json.Unmarshal(body, &receipt)

	if receipt.Cmv == "" || receipt.Epkrc1 == "" || receipt.Epkrc2 == "" || receipt.Hash == "" {
		return c.JSON(http.StatusBadRequest, ErrorValue)
	} else {
		// 购买成功,随机数解密
		privKey := utils.CreatePriKey(w.G1, w.G2, w.P, w.H, w.X)
		coin := decryptCoinReceipt(receipt, privKey, w.Amount)
		utils.MineTx(8545, coin.Hash)
		return c.JSON(http.StatusOK, coin)
	}
}

func ExchangeCoin(c echo.Context) error {
	w := new(model.ExchangeCoin)
	if err := c.Bind(w); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	senderPriv := utils.CreatePriKey(w.SG1, w.SG2, w.SP, w.SH, w.SX)
	reciverPub := utils.CreatePubKey(w.RG1, w.RG2, w.RP, w.RH)
	coin := utils.Coin{
		Cmv:    w.Cmv,
		Vor:    w.Vor,
		Amount: w.Amount,
	}
	amount, _ := strconv.Atoi(coin.Amount)
	spend, _ := strconv.Atoi(w.Spend)
	senderGethAccount := utils.EthAccounts(8545)[0]
	receiverGethAccount := utils.EthAccounts(8545)[0]
	txHash := utils.EthSendTransaction(8545, senderGethAccount, receiverGethAccount, senderPriv, reciverPub, coin, amount, spend)
	utils.MineTx(8545, txHash)
	rpcTx := utils.EthGetTransactionByHash(8545, txHash)
	tx := rpcTx.Result
	returnCoin := utils.Coin{
		Cmv:    tx.CmR,
		Vor:    decrypt(tx.CmRRC1, tx.CmRRC2, senderPriv),
		Hash:   txHash,
		Amount: strconv.Itoa(amount - spend),
	}
	return c.JSON(http.StatusOK, returnCoin)
}
func Receive(c echo.Context) error {
	w := new(model.ReceiveData)
	if err := c.Bind(w); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	privKey := utils.CreatePriKey(w.G1, w.G2, w.P, w.H, w.X)
	rpcTx := utils.EthGetTransactionByHash(8545, w.Hash)
	tx := rpcTx.Result
	returnCoin := utils.Coin{
		Cmv:    tx.CmO,
		Vor:    decrypt(tx.CmSRC1, tx.CmSRC2, privKey),
		Hash:   w.Hash,
		Amount: decryptValue(tx.EvsBsC1, tx.EvsBsC2, privKey),
	}
	return c.JSON(http.StatusOK, returnCoin)
}
func decryptCoinReceipt(recript utils.Receipt, priv ecc.PrivateKey, amount string) utils.Coin {
	return utils.Coin{
		Cmv:    recript.Cmv,
		Vor:    decrypt(recript.Epkrc1, recript.Epkrc2, priv),
		Hash:   recript.Hash,
		Amount: amount,
	}
}

//	解密随机数密文
func decrypt(hex0xStringC1 string, hex0xStringC2 string, priv ecc.PrivateKey) string {
	hexData1, _ := hex.DecodeString(hex0xStringC1[2:])
	hexData2, _ := hex.DecodeString(hex0xStringC2[2:])
	C := ecc.CypherText{
		C1: hexData1,
		C2: hexData2,
	}
	M := fmt.Sprintf("0x%x", ecc.Decrypt(priv, C))
	return M
}

//	解密随机数密文
func decryptValue(hex0xStringC1 string, hex0xStringC2 string, priv ecc.PrivateKey) string {
	hexData1, _ := hex.DecodeString(hex0xStringC1[2:])
	hexData2, _ := hex.DecodeString(hex0xStringC2[2:])
	C := ecc.CypherText{
		C1: hexData1,
		C2: hexData2,
	}
	M := fmt.Sprintf("0x%x", ecc.DecryptValue(priv, C))
	return M
}
