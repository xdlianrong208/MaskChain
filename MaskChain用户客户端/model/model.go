package model

type NewWallet struct {
	Name string `json:"name" form:"name"`
	Id   string `json:"id" form:"id"`
	Str  string `json:"str" form:"str"`
}

type BctoEx struct {
	G1     string `json:"g1"`
	G2     string `json:"g2"`
	P      string `json:"p"`
	H      string `json:"h"`
	X      string `json:"x"`
	Amount string `json:"amount"`
}

type ExchangeCoin struct {
	SG1    string `json:"sg1"`
	SG2    string `json:"sg2"`
	SP     string `json:"sp"`
	SH     string `json:"sh"`
	SX     string `json:"sx"`
	RG1    string `json:"rg1"`
	RG2    string `json:"rg2"`
	RP     string `json:"rp"`
	RH     string `json:"rh"`
	Amount string `json:"amount"`
	Cmv    string `json:"cmv"`
	Vor    string `json:"vor"`
	Spend  string `json:"spend"`
}
type ReceiveData struct {
	Hash string `json:"hash"`
	G1   string `json:"g1"`
	G2   string `json:"g2"`
	P    string `json:"p"`
	H    string `json:"h"`
	X    string `json:"x"`
}

type RPCbody struct {
	Jsonrpc string   `json:"jsonrpc"`
	Method  string   `json:"method"`
	Params  []string `json:"params"`
	ID      int      `json:"id"`
}
