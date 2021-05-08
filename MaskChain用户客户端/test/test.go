package main

import (
	"fmt"
	"wallet/model"
)

// 解锁账户需要的参数
var(
	ethaccount = "0x41c060c18d1ba76971dc2d298d6e7cc64f7be57f"
	ethkey     = "1"
)
// 发送转账交易需要的参数
var(
	spk        = "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd036414141f7ac794c5f6769e689fa397ecdeaa62c7d91445cbd58890d2aab869f910d6a51a66e91b2785ff5bc219ba68ac23fb92c5366fb2f9df3880db134c9a3eee1e41479be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b846a1e9101b4e073dd71831ecc43d6afaa0f5a083fd16edb2bdb6dca8d445801721395d248d9fb0983c588e81aec372fb4e3591f5ce67680275faa1dc88c492f58"
	rpk        = "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd036414141f7ac794c5f6769e689fa397ecdeaa62c7d91445cbd58890d2aab869f910d6a51a66e91b2785ff5bc219ba68ac23fb92c5366fb2f9df3880db134c9a3eee1e41479be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b846a1e9101b4e073dd71831ecc43d6afaa0f5a083fd16edb2bdb6dca8d445801721395d248d9fb0983c588e81aec372fb4e3591f5ce67680275faa1dc88c492f58"
	r          = "0xf"
	s          = "0x5"
	vor        = "0x1cdc4d05345260bcdb04d28c8c627b2fe2315525ab7215fc2660769516659595"
	cmo        = "0x04cae9f47037baae677a25b3064db540d854df7e02eb5de5acb85a1f82ad18cef0b867f2ccc4dcc017382ded3f0f9f03124dc160079da6ced19fca1aee58ca11a2"
)
// 查看交易内容需要的参数
var(
	txhash     = "0x2c551b1b51e995216f65a1be7ce52bfe08ee21f3655d6ad53a4173e534ed2596"
)

func main(){
	//TestUnlock()
	//TestSendTransaction()
	TestGetTransaction()
}

func TestUnlock(){
	if(model.UnlockAccount(ethaccount, ethkey) == true){
		fmt.Println("unlock account " + ethaccount + " right")
	}else{
		fmt.Println("unlock account " + ethaccount + " erro")
	}
}

func TestSendTransaction(){
	if(model.SendTransaction(spk, rpk, s, r, vor, cmo) == true){
		fmt.Println("sendtx right")
	}else{
		fmt.Println("sendtx erro")
	}
}

func TestGetTransaction() {
	if(model.GetTransaction(txhash) == true){
		fmt.Println("gettx "+ txhash)
	}else{
		fmt.Println("gettx erro")
	}
}
