// Copyright 2015 The go-ethereum Authors
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

package bind

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/event"
)

// SignerFn is a signer function callback when a contract requires a method to
// sign the transaction before submission.
type SignerFn func(types.Signer, common.Address, *types.Transaction) (*types.Transaction, error)

// CallOpts is the collection of options to fine tune a contract call request.
type CallOpts struct {
	Pending     bool            // Whether to operate on the pending state or the last known one
	From        common.Address  // Optional the sender address, otherwise the first account is used
	BlockNumber *big.Int        // Optional the block number on which the call should be performed
	Context     context.Context // Network context to support cancellation and timeouts (nil = no timeout)
}

// TransactOpts is the collection of authorization data required to create a
// valid Ethereum transaction.
type TransactOpts struct {
	From   common.Address // Ethereum account to send the transaction from
	Nonce  *big.Int       // Nonce to use for the transaction execution (nil = use pending state)
	Signer SignerFn       // Method to use for signing the transaction (mandatory)

	Value    *big.Int // Funds to transfer along along the transaction (nil = 0 = no funds)
	GasPrice *big.Int // Gas price to use for the transaction execution (nil = gas price oracle)
	GasLimit uint64   // Gas limit to set for the transaction execution (0 = estimate)

	Context context.Context // Network context to support cancellation and timeouts (nil = no timeout)

	//新增交易字段
	ID       uint64
	ErpkC1   *hexutil.Bytes
	ErpkC2   *hexutil.Bytes
	EspkC1   *hexutil.Bytes
	EspkC2   *hexutil.Bytes
	CMRpk    *hexutil.Bytes
	CMSpk    *hexutil.Bytes

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

	EvSC1    *hexutil.Bytes
	EvSC2    *hexutil.Bytes
	EvRC1    *hexutil.Bytes
	EvRC2    *hexutil.Bytes
	CmS      *hexutil.Bytes
	CmR      *hexutil.Bytes

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

	EvsBsC1  *hexutil.Bytes
	EvsBsC2  *hexutil.Bytes
	EvOC1    *hexutil.Bytes
	EvOC2    *hexutil.Bytes
	CmO      *hexutil.Bytes

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

	EpkrC1   *hexutil.Bytes
	EpkrC2   *hexutil.Bytes
	EpkpC1   *hexutil.Bytes
	EpkpC2   *hexutil.Bytes
	SigM     *hexutil.Bytes
	SigMHash *hexutil.Bytes
	SigR     *hexutil.Bytes
	SigS     *hexutil.Bytes
	CmV      *hexutil.Bytes
	CmSRC1   *hexutil.Bytes
	CmSRC2   *hexutil.Bytes
	CmRRC1   *hexutil.Bytes
	CmRRC2   *hexutil.Bytes
}

// FilterOpts is the collection of options to fine tune filtering for events
// within a bound contract.
type FilterOpts struct {
	Start uint64  // Start of the queried range
	End   *uint64 // End of the range (nil = latest)

	Context context.Context // Network context to support cancellation and timeouts (nil = no timeout)
}

// WatchOpts is the collection of options to fine tune subscribing for events
// within a bound contract.
type WatchOpts struct {
	Start   *uint64         // Start of the queried range (nil = latest)
	Context context.Context // Network context to support cancellation and timeouts (nil = no timeout)
}

// BoundContract is the base wrapper object that reflects a contract on the
// Ethereum network. It contains a collection of methods that are used by the
// higher level contract bindings to operate.
type BoundContract struct {
	address    common.Address     // Deployment address of the contract on the Ethereum blockchain
	abi        abi.ABI            // Reflect based ABI to access the correct Ethereum methods
	caller     ContractCaller     // Read interface to interact with the blockchain
	transactor ContractTransactor // Write interface to interact with the blockchain
	filterer   ContractFilterer   // Event filtering to interact with the blockchain
}

// NewBoundContract creates a low level contract interface through which calls
// and transactions may be made through.
func NewBoundContract(address common.Address, abi abi.ABI, caller ContractCaller, transactor ContractTransactor, filterer ContractFilterer) *BoundContract {
	return &BoundContract{
		address:    address,
		abi:        abi,
		caller:     caller,
		transactor: transactor,
		filterer:   filterer,
	}
}

// DeployContract deploys a contract onto the Ethereum blockchain and binds the
// deployment address with a Go wrapper.
func DeployContract(opts *TransactOpts, abi abi.ABI, bytecode []byte, backend ContractBackend, params ...interface{}) (common.Address, *types.Transaction, *BoundContract, error) {
	// Otherwise try to deploy the contract
	c := NewBoundContract(common.Address{}, abi, backend, backend, backend)

	input, err := c.abi.Pack("", params...)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	tx, err := c.transact(opts, nil, append(bytecode, input...))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	c.address = crypto.CreateAddress(opts.From, tx.Nonce())
	return c.address, tx, c, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (c *BoundContract) Call(opts *CallOpts, result interface{}, method string, params ...interface{}) error {
	// Don't crash on a lazy user
	if opts == nil {
		opts = new(CallOpts)
	}
	// Pack the input, call and unpack the results
	input, err := c.abi.Pack(method, params...)
	if err != nil {
		return err
	}
	var (
		msg    = ethereum.CallMsg{From: opts.From, To: &c.address, Data: input}
		ctx    = ensureContext(opts.Context)
		code   []byte
		output []byte
	)
	if opts.Pending {
		pb, ok := c.caller.(PendingContractCaller)
		if !ok {
			return ErrNoPendingState
		}
		output, err = pb.PendingCallContract(ctx, msg)
		if err == nil && len(output) == 0 {
			// Make sure we have a contract to operate on, and bail out otherwise.
			if code, err = pb.PendingCodeAt(ctx, c.address); err != nil {
				return err
			} else if len(code) == 0 {
				return ErrNoCode
			}
		}
	} else {
		output, err = c.caller.CallContract(ctx, msg, opts.BlockNumber)
		if err == nil && len(output) == 0 {
			// Make sure we have a contract to operate on, and bail out otherwise.
			if code, err = c.caller.CodeAt(ctx, c.address, opts.BlockNumber); err != nil {
				return err
			} else if len(code) == 0 {
				return ErrNoCode
			}
		}
	}
	if err != nil {
		return err
	}
	return c.abi.Unpack(result, method, output)
}

// Transact invokes the (paid) contract method with params as input values.
func (c *BoundContract) Transact(opts *TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	// Otherwise pack up the parameters and invoke the contract
	input, err := c.abi.Pack(method, params...)
	if err != nil {
		return nil, err
	}
	return c.transact(opts, &c.address, input)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (c *BoundContract) Transfer(opts *TransactOpts) (*types.Transaction, error) {
	return c.transact(opts, &c.address, nil)
}

// transact executes an actual transaction invocation, first deriving any missing
// authorization fields, and then scheduling the transaction for execution.
func (c *BoundContract) transact(opts *TransactOpts, contract *common.Address, input []byte) (*types.Transaction, error) {
	var err error

	// Ensure a valid value field and resolve the account nonce
	value := opts.Value
	if value == nil {
		value = new(big.Int)
	}
	var nonce uint64
	if opts.Nonce == nil {
		nonce, err = c.transactor.PendingNonceAt(ensureContext(opts.Context), opts.From)
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve account nonce: %v", err)
		}
	} else {
		nonce = opts.Nonce.Uint64()
	}
	// Figure out the gas allowance and gas price values
	gasPrice := opts.GasPrice
	if gasPrice == nil {
		gasPrice, err = c.transactor.SuggestGasPrice(ensureContext(opts.Context))
		if err != nil {
			return nil, fmt.Errorf("failed to suggest gas price: %v", err)
		}
	}
	gasLimit := opts.GasLimit
	if gasLimit == 0 {
		// Gas estimation cannot succeed without code for method invocations
		if contract != nil {
			if code, err := c.transactor.PendingCodeAt(ensureContext(opts.Context), c.address); err != nil {
				return nil, err
			} else if len(code) == 0 {
				return nil, ErrNoCode
			}
		}
		// If the contract surely has code (or code is not needed), estimate the transaction
		msg := ethereum.CallMsg{From: opts.From, To: contract, GasPrice: gasPrice, Value: value, Data: input}
		gasLimit, err = c.transactor.EstimateGas(ensureContext(opts.Context), msg)
		if err != nil {
			return nil, fmt.Errorf("failed to estimate gas needed: %v", err)
		}
	}
	// Create the transaction, sign it and schedule it for execution
	var rawTx *types.Transaction
	if contract == nil {
		rawTx = types.NewContractCreation(nonce, value, gasLimit, gasPrice, input, opts.ID, opts.ErpkC1, opts.ErpkC2, opts.EspkC1, opts.EspkC2, opts.CMRpk, opts.CMSpk, opts.RpkEPg1, opts.RpkEPg2, opts.RpkEPy1, opts.RpkEPy2, opts.RpkEPt1, opts.RpkEPt2, opts.RpkEPs, opts.RpkEPc, opts.SpkEPg1, opts.SpkEPg2, opts.SpkEPy1, opts.SpkEPy2, opts.SpkEPt1, opts.SpkEPt2, opts.SpkEPs, opts.SpkEPc, opts.EvSC1, opts.EvSC2, opts.EvRC1, opts.EvRC2, opts.CmS, opts.CmR, opts.ScmFPg1, opts.ScmFPg2, opts.ScmFPy1, opts.ScmFPy2, opts.ScmFPt1, opts.ScmFPt2, opts.ScmFPs, opts.ScmFPc, opts.RcmFPg1, opts.RcmFPg2, opts.RcmFPy1, opts.RcmFPy2, opts.RcmFPt1, opts.RcmFPt2, opts.RcmFPs, opts.RcmFPc, opts.EvsBsC1, opts.EvsBsC2, opts.EvOC1, opts.EvOC2, opts.CmO,  opts.VoEPg1, opts.VoEPg2, opts.VoEPy1, opts.VoEPy2, opts.VoEPt1, opts.VoEPt2, opts.VoEPs, opts.VoEPc, opts.BPy, opts.BPt, opts.BPsn1, opts.BPsn2, opts.BPsn3, opts.BPc, opts.EpkrC1, opts.EpkrC2, opts.EpkpC1, opts.EpkpC2, opts.SigM, opts.SigMHash, opts.SigR, opts.SigS, opts.CmV, opts.CmSRC1, opts.CmSRC2, opts.CmRRC1, opts.CmRRC2)
	} else {
		rawTx = types.NewTransaction(nonce, c.address, value, gasLimit, gasPrice, input, opts.ID, opts.ErpkC1, opts.ErpkC2, opts.EspkC1, opts.EspkC2, opts.CMRpk, opts.CMSpk, opts.RpkEPg1, opts.RpkEPg2, opts.RpkEPy1, opts.RpkEPy2, opts.RpkEPt1, opts.RpkEPt2, opts.RpkEPs, opts.RpkEPc, opts.SpkEPg1, opts.SpkEPg2, opts.SpkEPy1, opts.SpkEPy2, opts.SpkEPt1, opts.SpkEPt2, opts.SpkEPs, opts.SpkEPc, opts.EvSC1, opts.EvSC2, opts.EvRC1, opts.EvRC2, opts.CmS, opts.CmR, opts.ScmFPg1, opts.ScmFPg2, opts.ScmFPy1, opts.ScmFPy2, opts.ScmFPt1, opts.ScmFPt2, opts.ScmFPs, opts.ScmFPc, opts.RcmFPg1, opts.RcmFPg2, opts.RcmFPy1, opts.RcmFPy2, opts.RcmFPt1, opts.RcmFPt2, opts.RcmFPs, opts.RcmFPc, opts.EvsBsC1, opts.EvsBsC2, opts.EvOC1, opts.EvOC2, opts.CmO,  opts.VoEPg1, opts.VoEPg2, opts.VoEPy1, opts.VoEPy2, opts.VoEPt1, opts.VoEPt2, opts.VoEPs, opts.VoEPc, opts.BPy, opts.BPt, opts.BPsn1, opts.BPsn2, opts.BPsn3, opts.BPc, opts.EpkrC1, opts.EpkrC2, opts.EpkpC1, opts.EpkpC2, opts.SigM, opts.SigMHash, opts.SigR, opts.SigS, opts.CmV, opts.CmSRC1, opts.CmSRC2, opts.CmRRC1, opts.CmRRC2)
	}
	if opts.Signer == nil {
		return nil, errors.New("no signer to authorize the transaction with")
	}
	signedTx, err := opts.Signer(types.HomesteadSigner{}, opts.From, rawTx)
	if err != nil {
		return nil, err
	}
	if err := c.transactor.SendTransaction(ensureContext(opts.Context), signedTx); err != nil {
		return nil, err
	}
	return signedTx, nil
}

// FilterLogs filters contract logs for past blocks, returning the necessary
// channels to construct a strongly typed bound iterator on top of them.
func (c *BoundContract) FilterLogs(opts *FilterOpts, name string, query ...[]interface{}) (chan types.Log, event.Subscription, error) {
	// Don't crash on a lazy user
	if opts == nil {
		opts = new(FilterOpts)
	}
	// Append the event selector to the query parameters and construct the topic set
	query = append([][]interface{}{{c.abi.Events[name].ID()}}, query...)

	topics, err := makeTopics(query...)
	if err != nil {
		return nil, nil, err
	}
	// Start the background filtering
	logs := make(chan types.Log, 128)

	config := ethereum.FilterQuery{
		Addresses: []common.Address{c.address},
		Topics:    topics,
		FromBlock: new(big.Int).SetUint64(opts.Start),
	}
	if opts.End != nil {
		config.ToBlock = new(big.Int).SetUint64(*opts.End)
	}
	/* TODO(karalabe): Replace the rest of the method below with this when supported
	sub, err := c.filterer.SubscribeFilterLogs(ensureContext(opts.Context), config, logs)
	*/
	buff, err := c.filterer.FilterLogs(ensureContext(opts.Context), config)
	if err != nil {
		return nil, nil, err
	}
	sub, err := event.NewSubscription(func(quit <-chan struct{}) error {
		for _, log := range buff {
			select {
			case logs <- log:
			case <-quit:
				return nil
			}
		}
		return nil
	}), nil

	if err != nil {
		return nil, nil, err
	}
	return logs, sub, nil
}

// WatchLogs filters subscribes to contract logs for future blocks, returning a
// subscription object that can be used to tear down the watcher.
func (c *BoundContract) WatchLogs(opts *WatchOpts, name string, query ...[]interface{}) (chan types.Log, event.Subscription, error) {
	// Don't crash on a lazy user
	if opts == nil {
		opts = new(WatchOpts)
	}
	// Append the event selector to the query parameters and construct the topic set
	query = append([][]interface{}{{c.abi.Events[name].ID()}}, query...)

	topics, err := makeTopics(query...)
	if err != nil {
		return nil, nil, err
	}
	// Start the background filtering
	logs := make(chan types.Log, 128)

	config := ethereum.FilterQuery{
		Addresses: []common.Address{c.address},
		Topics:    topics,
	}
	if opts.Start != nil {
		config.FromBlock = new(big.Int).SetUint64(*opts.Start)
	}
	sub, err := c.filterer.SubscribeFilterLogs(ensureContext(opts.Context), config, logs)
	if err != nil {
		return nil, nil, err
	}
	return logs, sub, nil
}

// UnpackLog unpacks a retrieved log into the provided output structure.
func (c *BoundContract) UnpackLog(out interface{}, event string, log types.Log) error {
	if len(log.Data) > 0 {
		if err := c.abi.Unpack(out, event, log.Data); err != nil {
			return err
		}
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	return parseTopics(out, indexed, log.Topics[1:])
}

// UnpackLogIntoMap unpacks a retrieved log into the provided map.
func (c *BoundContract) UnpackLogIntoMap(out map[string]interface{}, event string, log types.Log) error {
	if len(log.Data) > 0 {
		if err := c.abi.UnpackIntoMap(out, event, log.Data); err != nil {
			return err
		}
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	return parseTopicsIntoMap(out, indexed, log.Topics[1:])
}

// ensureContext is a helper method to ensure a context is not nil, even if the
// user specified it as such.
func ensureContext(ctx context.Context) context.Context {
	if ctx == nil {
		return context.TODO()
	}
	return ctx
}
