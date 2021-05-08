package utils

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/big"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"time"
	ecc "wallet/ECC"
	"wallet/model"
)

func stringtobig(s string, base int) (b *big.Int) {
	b = new(big.Int)
	b, _ = b.SetString(s, base)
	return
}

func CreatePriKey(g1 string, g2 string, p string, h string, x string) (usrpub ecc.PrivateKey) {
	usrpub.G1 = stringtobig(g1, 16)
	usrpub.G2 = stringtobig(g2, 16)
	usrpub.P = stringtobig(p, 16)
	usrpub.H = stringtobig(h, 16)
	usrpub.X = stringtobig(x, 16)
	return
}

func CreatePubKey(g1 string, g2 string, p string, h string) (usrpub ecc.PublicKey) {
	usrpub.G1 = stringtobig(g1, 16)
	usrpub.G2 = stringtobig(g2, 16)
	usrpub.P = stringtobig(p, 16)
	usrpub.H = stringtobig(h, 16)
	return
}
func EthSendTransaction(senderRPCPort int, senderGethAccount string, receiverGethAccount string, senderAccount ecc.PrivateKey, receiverAccount ecc.PublicKey, coin Coin, total int, amount int) string {
	if !personalUnlockAccount(senderRPCPort, senderGethAccount, "1") {
		Fatalf("发送方账户解锁失败")
	}
	txs := PerpareTX(senderGethAccount, receiverGethAccount, senderAccount, receiverAccount, coin, total, amount)
	data := txs
	body := ethRPCPost(data, model.Ethurl)
	var result RPCResult
	json.Unmarshal(body, &result)
	if result.Result != "" {
		fmt.Println("转账交易发送成功，待打包共识")
	} else {
		Fatalf("转账交易发送失败")
	}
	return result.Result
}
func PerpareTX(senderGethAccount string, receiverGethAccount string, senderAccount ecc.PrivateKey, receiverAccount ecc.PublicKey, coin Coin, total int, amount int) SendRPCTx {
	param := SendRPCTxParams{
		From:     senderGethAccount,
		To:       receiverGethAccount,
		Gas:      "0x76c0",
		GasPrice: "0x9184e72a000",
		Value:    "0x1",
		ID:       "0x0",
		Data:     "0x00",
		Spk:      fmt.Sprintf("%0*x%0*x%0*x%0*x", 64, senderAccount.P, 129, senderAccount.G1, 129, senderAccount.G2, 129, senderAccount.H),
		Rpk:      fmt.Sprintf("%0*x%0*x%0*x%0*x", 64, receiverAccount.P, 129, receiverAccount.G1, 129, receiverAccount.G2, 129, receiverAccount.H),
		S:        fmt.Sprintf("0x%x", amount),
		R:        fmt.Sprintf("0x%x", total-amount),
		Vor:      coin.Vor,
		Cmo:      coin.Cmv,
	}
	var params []SendRPCTxParams
	params = append(params, param)
	tx := SendRPCTx{
		Jsonrpc: "2.0",
		Method:  "eth_sendTransaction",
		Params:  params,
		ID:      67,
	}
	return tx
}
func personalUnlockAccount(rpcPort int, account string, passphrase string) bool {
	data := model.RPCbody{
		Jsonrpc: "2.0",
		Method:  "personal_unlockAccount",
		Params:  []string{account, passphrase},
		ID:      67,
	}
	body := ethRPCPost(data, model.Ethurl)
	var result AddPeerResult
	json.Unmarshal(body, &result)
	return result.Result
}
func ethRPCPost(data interface{}, url string) []byte {
	jsonStr, _ := json.Marshal(data)
	resp, err := http.Post(url,
		"application/json",
		bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println(err)
	}
	if resp == nil {
		Fatalf("tcp连接失败,url:" + url)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}
func EthGetTransactionByHash(rpcPort int, txHash string) RPCtx {
	data := model.RPCbody{
		Jsonrpc: "2.0",
		Method:  "eth_getTransactionByHash",
		Params:  []string{txHash},
		ID:      67,
	}
	body := ethRPCPost(data, model.Ethurl)
	var result RPCtx
	json.Unmarshal(body, &result)
	return result
}
func MineTx(rpcPort int, TxHash string) bool {
	fmt.Println("打包共识使交易", TxHash, "生效")
	minerStart(rpcPort)
	for {
		if res := EthGetTransactionByHash(rpcPort, TxHash); res.Result.BlockHash != "" {
			break
		}
		time.Sleep(time.Duration(1) * time.Second) //等一秒
	}
	//多挖几个块，不然不好共识
	blockNum := ethBlockNumber(rpcPort)
	for {
		if ethBlockNumber(rpcPort)-blockNum >= 10 {
			break
		}
		time.Sleep(time.Duration(1) * time.Second) //等一秒
	}
	fmt.Println("交易", TxHash, "已被打包")
	minerStop(rpcPort)
	return true
}
func EthAccounts(rpcPort int) []string {
	data := RPCbody{
		Jsonrpc: "2.0",
		Method:  "eth_accounts",
		Params:  nil,
		ID:      67,
	}
	body := ethRPCPost(data, model.Ethurl)
	var accounts AccountsResult
	json.Unmarshal(body, &accounts)
	return accounts.Result
}
func minerStart(rpcPort int) bool {
	personalUnlockAccount(rpcPort, EthAccounts(8545)[0], "123456")
	data := RPCbody{
		Jsonrpc: "2.0",
		Method:  "miner_start",
		Params:  nil,
		ID:      67,
	}
	body := ethRPCPost(data, model.Ethurl)
	var result AddPeerResult
	json.Unmarshal(body, &result)
	return result.Result
}
func minerStop(rpcPort int) bool {
	data := RPCbody{
		Jsonrpc: "2.0",
		Method:  "miner_stop",
		Params:  nil,
		ID:      67,
	}
	body := ethRPCPost(data, model.Ethurl)
	var result AddPeerResult
	json.Unmarshal(body, &result)
	return result.Result
}
func ethBlockNumber(rpcPort int) int {
	data := RPCbody{
		Jsonrpc: "2.0",
		Method:  "eth_blockNumber",
		Params:  nil,
		ID:      67,
	}
	body := ethRPCPost(data, model.Ethurl)
	var result RPCResult
	json.Unmarshal(body, &result)
	num, _ := strconv.ParseUint(result.Result[2:], 16, 64)
	return int(num)
}

// Fatalf formats a message to standard error and exits the program.
// The message is also printed to standard output if standard error
// is redirected to a different file.
func Fatalf(format string, args ...interface{}) {
	w := io.MultiWriter(os.Stdout, os.Stderr)
	if runtime.GOOS == "windows" {
		// The SameFile check below doesn't work on Windows.
		// stdout is unlikely to get redirected though, so just print there.
		w = os.Stdout
	} else {
		outf, _ := os.Stdout.Stat()
		errf, _ := os.Stderr.Stat()
		if outf != nil && errf != nil && os.SameFile(outf, errf) {
			w = os.Stderr
		}
	}
	fmt.Fprintf(w, "Fatal: "+format+"\n", args...)
	os.Exit(1)
}
func pause(enable bool) {
	if enable {
		fmt.Print("请输入回车继续...")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			break
		}
	}
}
