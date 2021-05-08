package utils

type NodeInfo struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Enode string `json:"enode"`
		Enr   string `json:"enr"`
		IP    string `json:"ip"`
		Ports struct {
			Discovery int `json:"discovery"`
			Listener  int `json:"listener"`
		} `json:"ports"`
		ListenAddr string `json:"listenAddr"`
		Protocols  struct {
			Eth struct {
				Network    int    `json:"network"`
				Difficulty int    `json:"difficulty"`
				Genesis    string `json:"genesis"`
				Config     struct {
					ChainID             int      `json:"chainId"`
					HomesteadBlock      int      `json:"homesteadBlock"`
					Eip150Block         int      `json:"eip150Block"`
					Eip150Hash          string   `json:"eip150Hash"`
					Eip155Block         int      `json:"eip155Block"`
					Eip158Block         int      `json:"eip158Block"`
					ByzantiumBlock      int      `json:"byzantiumBlock"`
					ConstantinopleBlock int      `json:"constantinopleBlock"`
					PetersburgBlock     int      `json:"petersburgBlock"`
					IstanbulBlock       int      `json:"istanbulBlock"`
					Ethash              struct{} `json:"ethash"`
					CryptoType          int      `json:"cryptoType"`
				} `json:"config"`
				Head string `json:"head"`
			} `json:"eth"`
		} `json:"protocols"`
	} `json:"result"`
}

type RPCbody struct {
	Jsonrpc string   `json:"jsonrpc"`
	Method  string   `json:"method"`
	Params  []string `json:"params"`
	ID      int      `json:"id"`
}

type AddPeerResult struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  bool   `json:"result"`
}

type PeersResult struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  []struct {
		Enode   string   `json:"enode"`
		ID      string   `json:"id"`
		Name    string   `json:"name"`
		Caps    []string `json:"caps"`
		Network struct {
			LocalAddress  string `json:"localAddress"`
			RemoteAddress string `json:"remoteAddress"`
			Inbound       bool   `json:"inbound"`
			Trusted       bool   `json:"trusted"`
			Static        bool   `json:"static"`
		} `json:"network"`
		Protocols struct {
			Eth struct {
				Version    int    `json:"version"`
				Difficulty int    `json:"difficulty"`
				Head       string `json:"head"`
			} `json:"eth"`
		} `json:"protocols"`
	} `json:"result"`
}

type RPCResult struct {
	ID      int    `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  string `json:"result"`
}
type AccountsResult struct {
	ID      int      `json:"id"`
	Jsonrpc string   `json:"jsonrpc"`
	Result  []string `json:"result"`
}
type Receipt struct {
	Cmv    string `json:"cmv"`
	Epkrc1 string `json:"epkrc1"`
	Epkrc2 string `json:"epkrc2"`
	Hash   string `json:"hash"` //此次购币交易的交易哈希
}
type Coin struct {
	Cmv    string `json:"cmv"`
	Vor    string `json:"vor"`
	Hash   string `json:"hash"`
	Amount string `json:"amount"`
}
type RPCtx struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		BlockHash        string `json:"blockHash"`
		BlockNumber      string `json:"blockNumber"`
		From             string `json:"from"`
		Gas              string `json:"gas"`
		GasPrice         string `json:"gasPrice"`
		Hash             string `json:"hash"`
		Input            string `json:"input"`
		Nonce            string `json:"nonce"`
		To               string `json:"to"`
		TransactionIndex string `json:"transactionIndex"`
		Value            string `json:"value"`
		V                string `json:"v"`
		R                string `json:"r"`
		S                string `json:"s"`
		ID               string `json:"ID"`
		ErpkC1           string `json:"erpkc1"`
		ErpkC2           string `json:"erpkc2"`
		EspkC1           string `json:"espkc1"`
		EspkC2           string `json:"espkc2"`
		CMRpk            string `json:"cmrpk"`
		CMSpk            string `json:"cmspk"`
		RpkEPg1          string `json:"rpkepg1"` //接收方地址公钥相等证明字段g1
		RpkEPg2      	 string `json:"rpkepg2"` //接收方地址公钥相等证明字段g2
		RpkEPy1     	 string `json:"rpkepy1"` //接收方地址公钥相等证明字段y1
		RpkEPy2    	     string `json:"rpkepy2"` //接收方地址公钥相等证明字段y2
		RpkEPt1      	 string `json:"rpkept1"` //接收方地址公钥相等证明字段t1
		RpkEPt2      	 string `json:"rpkept2"` //接收方地址公钥相等证明字段t2
		RpkEPs       	 string `json:"rpkeps"` //接收方地址公钥相等证明字段s
		RpkEPc       	 string `json:"rpkepc"` //接收方地址公钥相等证明字段c
		SpkEPg1      	 string `json:"spkepg1"` //发送方地址公钥相等证明字段g1
		SpkEPg2      	 string `json:"spkepg2"` //发送方地址公钥相等证明字段g2
		SpkEPy1      	 string `json:"spkepy1"` //发送方地址公钥相等证明字段y1
		SpkEPy2      	 string `json:"spkepy2"` //发送方地址公钥相等证明字段y2
		SpkEPt1      	 string `json:"spkept1"` //发送方地址公钥相等证明字段t1
		SpkEPt2      	 string `json:"spkept2"` //发送方地址公钥相等证明字段t2
		SpkEPs       	 string `json:"spkeps"` //发送方地址公钥相等证明字段s
		SpkEPc       	 string `json:"spkepc"` //发送方地址公钥相等证明字段c
		EvSC1            string `json:"evsc1"`
		EvSC2            string `json:"evsc2"`
		EvRC1            string `json:"evrc1"`
		EvRC2            string `json:"evrc2"`
		CmS              string `json:"cms"`
		CmR              string `json:"cmr"`
		ScmFPg1      	 string `json:"scmfpg1"` //发送金额承诺格式证明字段g1
		ScmFPg2      	 string `json:"scmfpg2"` //发送金额承诺格式证明字段g2
		ScmFPy1      	 string `json:"scmfpy1"` //发送金额承诺格式证明字段y1
		ScmFPy2      	 string `json:"scmfpy2"` //发送金额承诺格式证明字段y2
		ScmFPt1      	 string `json:"scmfpt1"` //发送金额承诺格式证明字段t1
		ScmFPt2      	 string `json:"scmfpt2"` //发送金额承诺格式证明字段t2
		ScmFPs       	 string `json:"scmfps"` //发送金额承诺格式证明字段s
		ScmFPc       	 string `json:"scmfpc"` //发送金额承诺格式证明字段c
		RcmFPg1      	 string `json:"rcmfpg1"` //接收金额承诺格式证明字段g1
		RcmFPg2      	 string `json:"rcmfpg2"` //接收金额承诺格式证明字段g2
		RcmFPy1      	 string `json:"rcmfpy1"` //接收金额承诺格式证明字段y1
		RcmFPy2      	 string `json:"rcmfpy2"` //接收金额承诺格式证明字段y2
		RcmFPt1      	 string `json:"rcmfpt1"` //接收金额承诺格式证明字段t1
		RcmFPt2      	 string `json:"rcmfpt2"` //接收金额承诺格式证明字段t2
		RcmFPs       	 string `json:"rcmfps"` //接收金额承诺格式证明字段s
		RcmFPc       	 string `json:"rcmfpc"` //接收金额承诺格式证明字段c
		EvsBsC1          string `json:"evsbsc1"`
		EvsBsC2          string `json:"evsbsc2"`
		EvOC1            string `json:"evoc1"`
		EvOC2            string `json:"evoc2"`
		CmO              string `json:"cmo"`
		VoEPg1       	 string `json:"voepg1"` //被花费承诺相等证明字段g1
		VoEPg2       	 string `json:"voepg2"` //被花费承诺相等证明字段g2
		VoEPy1       	 string `json:"voepy1"` //被花费承诺相等证明字段y1
		VoEPy2       	 string `json:"voepy2"` //被花费承诺相等证明字段y2
		VoEPt1       	 string `json:"voept1"` //被花费承诺相等证明字段t1
		VoEPt2       	 string `json:"voept2"` //被花费承诺相等证明字段t2
		VoEPs        	 string `json:"voeps"` //被花费承诺相等证明字段s
		VoEPc        	 string `json:"voepc"` //被花费承诺相等证明字段c
		BPy          	 string `json:"bpy"` //会计平衡证明字段y
		BPt          	 string `json:"bpt"` //会计平衡证明字段t
		BPsn1        	 string `json:"bpsn1"` //会计平衡证明字段sn1
		BPsn2        	 string `json:"bpsn2"` //会计平衡证明字段sn2
		BPsn3        	 string `json:"bpsn3"` //会计平衡证明字段sn3
		BPc          	 string `json:"bpc"` //会计平衡证明字段c
		EpkrC1           string `json:"epkrc1"`
		EpkrC2           string `json:"epkrc2"`
		EpkpC1           string `json:"epkpc1"`
		EpkpC2           string `json:"epkpc2"`
		SigM             string `json:"sigm"`
		SigMHash         string `json:"sigmhash"`
		SigR             string `json:"sigr"`
		SigS             string `json:"sigs"`
		CmV              string `json:"cmv"`
		CmSRC1           string `json:"cmsrc1"`
		CmSRC2           string `json:"cmsrc2"`
		CmRRC1           string `json:"cmrrc1"`
		CmRRC2           string `json:"cmrrc2"`
	} `json:"result"`
}

type SendRPCTx struct {
	Jsonrpc string            `json:"jsonrpc"`
	Method  string            `json:"method"`
	Params  []SendRPCTxParams `json:"params"`
	ID      int               `json:"id"`
}
type SendRPCTxParams struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Gas      string `json:"gas"`
	GasPrice string `json:"gasPrice"`
	Value    string `json:"value"`
	ID       string `json:"id"`
	Data     string `json:"data"`
	Spk      string `json:"spk"`
	Rpk      string `json:"rpk"`
	S        string `json:"s"`
	R        string `json:"r"`
	Vor      string `json:"vor"`
	Cmo      string `json:"cmo"`
}
