// Copyright 2018 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package core

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

type ValidationInfo struct {
	Typ     string `json:"type"`
	Message string `json:"message"`
}
type ValidationMessages struct {
	Messages []ValidationInfo
}

const (
	WARN = "WARNING"
	CRIT = "CRITICAL"
	INFO = "Info"
)

func (vs *ValidationMessages) Crit(msg string) {
	vs.Messages = append(vs.Messages, ValidationInfo{CRIT, msg})
}
func (vs *ValidationMessages) Warn(msg string) {
	vs.Messages = append(vs.Messages, ValidationInfo{WARN, msg})
}
func (vs *ValidationMessages) Info(msg string) {
	vs.Messages = append(vs.Messages, ValidationInfo{INFO, msg})
}

/// getWarnings returns an error with all messages of type WARN of above, or nil if no warnings were present
func (v *ValidationMessages) getWarnings() error {
	var messages []string
	for _, msg := range v.Messages {
		if msg.Typ == WARN || msg.Typ == CRIT {
			messages = append(messages, msg.Message)
		}
	}
	if len(messages) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(messages, ","))
	}
	return nil
}

// SendTxArgs represents the arguments to submit a transaction
type SendTxArgs struct {
	From     common.MixedcaseAddress  `json:"from"`
	To       *common.MixedcaseAddress `json:"to"`
	Gas      hexutil.Uint64           `json:"gas"`
	GasPrice hexutil.Big              `json:"gasPrice"`
	Value    hexutil.Big              `json:"value"`
	Nonce    hexutil.Uint64           `json:"nonce"`
	// We accept "data" and "input" for backwards-compatibility reasons.
	Data     *hexutil.Bytes  `json:"data"`
	Input    *hexutil.Bytes  `json:"input,omitempty"`
	ID       *hexutil.Uint64 `json:"ID"`
	ErpkC1   *hexutil.Bytes  `json:"erpkc1"`
	ErpkC2   *hexutil.Bytes  `json:"erpkc2"`
	EspkC1   *hexutil.Bytes  `json:"espkc1"`
	EspkC2   *hexutil.Bytes  `json:"espkc2"`
	CMRpk    *hexutil.Bytes  `json:"cmrpk"`
	CMSpk    *hexutil.Bytes  `json:"cmspk"`

	RpkEPg1          *hexutil.Bytes  `json:"rpkepg1"` //接收方地址公钥相等证明字段g1
	RpkEPg2      	 *hexutil.Bytes  `json:"rpkepg2"` //接收方地址公钥相等证明字段g2
	RpkEPy1     	 *hexutil.Bytes  `json:"rpkepy1"` //接收方地址公钥相等证明字段y1
	RpkEPy2    	     *hexutil.Bytes  `json:"rpkepy2"` //接收方地址公钥相等证明字段y2
	RpkEPt1      	 *hexutil.Bytes  `json:"rpkept1"` //接收方地址公钥相等证明字段t1
	RpkEPt2      	 *hexutil.Bytes  `json:"rpkept2"` //接收方地址公钥相等证明字段t2
	RpkEPs       	 *hexutil.Bytes  `json:"rpkeps"` //接收方地址公钥相等证明字段s
	RpkEPc       	 *hexutil.Bytes  `json:"rpkepc"` //接收方地址公钥相等证明字段c
	SpkEPg1      	 *hexutil.Bytes  `json:"spkepg1"` //发送方地址公钥相等证明字段g1
	SpkEPg2      	 *hexutil.Bytes  `json:"spkepg2"` //发送方地址公钥相等证明字段g2
	SpkEPy1      	 *hexutil.Bytes  `json:"spkepy1"` //发送方地址公钥相等证明字段y1
	SpkEPy2      	 *hexutil.Bytes  `json:"spkepy2"` //发送方地址公钥相等证明字段y2
	SpkEPt1      	 *hexutil.Bytes  `json:"spkept1"` //发送方地址公钥相等证明字段t1
	SpkEPt2      	 *hexutil.Bytes  `json:"spkept2"` //发送方地址公钥相等证明字段t2
	SpkEPs       	 *hexutil.Bytes  `json:"spkeps"` //发送方地址公钥相等证明字段s
	SpkEPc       	 *hexutil.Bytes  `json:"spkepc"` //发送方地址公钥相等证明字段c

	EvSC1    *hexutil.Bytes  `json:"evsc1"`
	EvSC2    *hexutil.Bytes  `json:"evsc2"`
	EvRC1    *hexutil.Bytes  `json:"evrc1"`
	EvRC2    *hexutil.Bytes  `json:"evrc2"`
	CmS      *hexutil.Bytes  `json:"cms"`
	CmR      *hexutil.Bytes  `json:"cmr"`

	ScmFPg1      	 *hexutil.Bytes  `json:"scmfpg1"` //发送金额承诺格式证明字段g1
	ScmFPg2      	 *hexutil.Bytes  `json:"scmfpg2"` //发送金额承诺格式证明字段g2
	ScmFPy1      	 *hexutil.Bytes  `json:"scmfpy1"` //发送金额承诺格式证明字段y1
	ScmFPy2      	 *hexutil.Bytes  `json:"scmfpy2"` //发送金额承诺格式证明字段y2
	ScmFPt1      	 *hexutil.Bytes  `json:"scmfpt1"` //发送金额承诺格式证明字段t1
	ScmFPt2      	 *hexutil.Bytes  `json:"scmfpt2"` //发送金额承诺格式证明字段t2
	ScmFPs       	 *hexutil.Bytes  `json:"scmfps"` //发送金额承诺格式证明字段s
	ScmFPc       	 *hexutil.Bytes  `json:"scmfpc"` //发送金额承诺格式证明字段c
	RcmFPg1      	 *hexutil.Bytes  `json:"rcmfpg1"` //接收金额承诺格式证明字段g1
	RcmFPg2      	 *hexutil.Bytes  `json:"rcmfpg2"` //接收金额承诺格式证明字段g2
	RcmFPy1      	 *hexutil.Bytes  `json:"rcmfpy1"` //接收金额承诺格式证明字段y1
	RcmFPy2      	 *hexutil.Bytes  `json:"rcmfpy2"` //接收金额承诺格式证明字段y2
	RcmFPt1      	 *hexutil.Bytes  `json:"rcmfpt1"` //接收金额承诺格式证明字段t1
	RcmFPt2      	 *hexutil.Bytes  `json:"rcmfpt2"` //接收金额承诺格式证明字段t2
	RcmFPs       	 *hexutil.Bytes  `json:"rcmfps"` //接收金额承诺格式证明字段s
	RcmFPc       	 *hexutil.Bytes  `json:"rcmfpc"` //接收金额承诺格式证明字段c

	EvsBsC1  *hexutil.Bytes  `json:"evsbsc1"`
	EvsBsC2  *hexutil.Bytes  `json:"evsbsc2"`
	EvOC1    *hexutil.Bytes  `json:"evoc1"`
	EvOC2    *hexutil.Bytes  `json:"evoc2"`
	CmO      *hexutil.Bytes  `json:"cmo"`

	VoEPg1       	 *hexutil.Bytes  `json:"voepg1"` //被花费承诺相等证明字段g1
	VoEPg2       	 *hexutil.Bytes  `json:"voepg2"` //被花费承诺相等证明字段g2
	VoEPy1       	 *hexutil.Bytes  `json:"voepy1"` //被花费承诺相等证明字段y1
	VoEPy2       	 *hexutil.Bytes  `json:"voepy2"` //被花费承诺相等证明字段y2
	VoEPt1       	 *hexutil.Bytes  `json:"voept1"` //被花费承诺相等证明字段t1
	VoEPt2       	 *hexutil.Bytes  `json:"voept2"` //被花费承诺相等证明字段t2
	VoEPs        	 *hexutil.Bytes  `json:"voeps"` //被花费承诺相等证明字段s
	VoEPc        	 *hexutil.Bytes  `json:"voepc"` //被花费承诺相等证明字段c
	BPy          	 *hexutil.Bytes  `json:"bpy"` //会计平衡证明字段y
	BPt          	 *hexutil.Bytes  `json:"bpt"` //会计平衡证明字段t
	BPsn1        	 *hexutil.Bytes  `json:"bpsn1"` //会计平衡证明字段sn1
	BPsn2        	 *hexutil.Bytes  `json:"bpsn2"` //会计平衡证明字段sn2
	BPsn3        	 *hexutil.Bytes  `json:"bpsn3"` //会计平衡证明字段sn3
	BPc          	 *hexutil.Bytes  `json:"bpc"` //会计平衡证明字段c

	EpkrC1   *hexutil.Bytes  `json:"epkrc1"`
	EpkrC2   *hexutil.Bytes  `json:"epkrc2"`
	EpkpC1   *hexutil.Bytes  `json:"epkpc1"`
	EpkpC2   *hexutil.Bytes  `json:"epkpc2"`
	SigM     *hexutil.Bytes  `json:"sigm"`
	SigMHash *hexutil.Bytes  `json:"sigmhash"`
	SigR     *hexutil.Bytes  `json:"sigr"`
	SigS     *hexutil.Bytes  `json:"sigs"`
	CmV      *hexutil.Bytes  `json:"cmv"`
	CmSRC1   *hexutil.Bytes  `json:"cmsrc1"`
	CmSRC2   *hexutil.Bytes  `json:" cmsrc2"`
	CmRRC1   *hexutil.Bytes  `json:" cmrrc1"`
	CmRRC2   *hexutil.Bytes  `json:" cmrrc2"`
}

func (args SendTxArgs) String() string {
	s, err := json.Marshal(args)
	if err == nil {
		return string(s)
	}
	return err.Error()
}

func (args *SendTxArgs) toTransaction() *types.Transaction {
	var input []byte
	if args.Data != nil {
		input = *args.Data
	} else if args.Input != nil {
		input = *args.Input
	}
	if args.To == nil {
		return types.NewContractCreation(uint64(args.Nonce), (*big.Int)(&args.Value), uint64(args.Gas), (*big.Int)(&args.GasPrice), input, uint64(*args.ID), args.ErpkC1, args.ErpkC2, args.EspkC1, args.EspkC2, args.CMRpk, args.CMSpk, args.RpkEPg1, args.RpkEPg2, args.RpkEPy1, args.RpkEPy2, args.RpkEPt1, args.RpkEPt2, args.RpkEPs, args.RpkEPc, args.SpkEPg1, args.SpkEPg2, args.SpkEPy1, args.SpkEPy2, args.SpkEPt1, args.SpkEPt2, args.SpkEPs, args.SpkEPc, args.EvSC1, args.EvSC2, args.EvRC1, args.EvRC2, args.CmS, args.CmR, args.ScmFPg1, args.ScmFPg2, args.ScmFPy1, args.ScmFPy2, args.ScmFPt1, args.ScmFPt2, args.ScmFPs, args.ScmFPc, args.RcmFPg1, args.RcmFPg2, args.RcmFPy1, args.RcmFPy2, args.RcmFPt1, args.RcmFPt2, args.RcmFPs, args.RcmFPc, args.EvsBsC1, args.EvsBsC2, args.EvOC1, args.EvOC2, args.CmO,  args.VoEPg1, args.VoEPg2, args.VoEPy1, args.VoEPy2, args.VoEPt1, args.VoEPt2, args.VoEPs, args.VoEPc, args.BPy, args.BPt, args.BPsn1, args.BPsn2, args.BPsn3, args.BPc, args.EpkrC1, args.EpkrC2, args.EpkpC1, args.EpkpC2, args.SigM, args.SigMHash, args.SigR, args.SigS, args.CmV, args.CmSRC1, args.CmSRC2, args.CmRRC1, args.CmRRC2)
	}
	return types.NewTransaction(uint64(args.Nonce), args.To.Address(), (*big.Int)(&args.Value), (uint64)(args.Gas), (*big.Int)(&args.GasPrice), input, uint64(*args.ID), args.ErpkC1, args.ErpkC2, args.EspkC1, args.EspkC2, args.CMRpk, args.CMSpk, args.RpkEPg1, args.RpkEPg2, args.RpkEPy1, args.RpkEPy2, args.RpkEPt1, args.RpkEPt2, args.RpkEPs, args.RpkEPc, args.SpkEPg1, args.SpkEPg2, args.SpkEPy1, args.SpkEPy2, args.SpkEPt1, args.SpkEPt2, args.SpkEPs, args.SpkEPc, args.EvSC1, args.EvSC2, args.EvRC1, args.EvRC2, args.CmS, args.CmR, args.ScmFPg1, args.ScmFPg2, args.ScmFPy1, args.ScmFPy2, args.ScmFPt1, args.ScmFPt2, args.ScmFPs, args.ScmFPc, args.RcmFPg1, args.RcmFPg2, args.RcmFPy1, args.RcmFPy2, args.RcmFPt1, args.RcmFPt2, args.RcmFPs, args.RcmFPc, args.EvsBsC1, args.EvsBsC2, args.EvOC1, args.EvOC2, args.CmO,  args.VoEPg1, args.VoEPg2, args.VoEPy1, args.VoEPy2, args.VoEPt1, args.VoEPt2, args.VoEPs, args.VoEPc, args.BPy, args.BPt, args.BPsn1, args.BPsn2, args.BPsn3, args.BPc, args.EpkrC1, args.EpkrC2, args.EpkpC1, args.EpkpC2, args.SigM, args.SigMHash, args.SigR, args.SigS, args.CmV, args.CmSRC1, args.CmSRC2, args.CmRRC1, args.CmRRC2)
}
