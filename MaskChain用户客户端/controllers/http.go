package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
)

var (
	//Verifyurl    = "http://39.106.173.191:1423/verify"
	//Getpuburl    = "http://39.106.173.191:1423/regkey?chainID=1"
	//Ethurl       = "http://127.0.0.1:8545"
	//RegulatorURL = "http://39.106.173.191:1423/" // 监管方URL
	//ExchangeURL  = "http://127.0.0.1:1323/"

	Verifyurl    = "http://localhost:1423/verify"
	Getpuburl    = "http://localhost:1423/regkey?chainID=1"
	Ethurl       = "http://localhost:8545"
	RegulatorURL = "http://localhost:1423/" // 监管方URL
	ExchangeURL  = "http://localhost:1323/"
)

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
