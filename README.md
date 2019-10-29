# ckb-types-go

### Note

Encode `Transaction` will strip witnesses field, so that
we can properly calculate transaction hash.

### example

#### send capacity

```go
package main

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/minio/blake2b-simd"
	"github.com/ybbus/jsonrpc"
	t "github.com/zeroqn/ckb-types-go/jsonrpc/types"
)

// from spec/dev.toml
// # issue 10M cell for random generated private key: d00c06bfd800d27397002dca6fb0993d5ba6399b4238b2f29ee9deb97593d2bc
// [[genesis.issued_cells]]
// capacity = 10_000_000_00000000
// lock.code_hash = "0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8"
// lock.args = "0xc8328aabcd9b9e8e64fbc566c4385c3bdeb219d7"
// lock.hash_type = "type"
//
// # issue 5M cell for random generated private key: 63d86723e08f0f813a36ce6aa123bb2289d90680ae1e99d4de8cdb334553f24d
// [[genesis.issued_cells]]
// capacity = 5_000_000_00000000
// lock.code_hash = "0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8"
// lock.args = "0x470dcdc5e44064909650113a274b3b36aecb6dc7"
// lock.hash_type = "type"

// CkbBlake2BHashPersonalization personal
const CkbBlake2BHashPersonalization = "ckb-default-hash"

// GenesisBlockHash genesis block hash
const GenesisBlockHash = "0xe49352ee4984694d88eb3c1493a33d69d61c786dc5b0a32c4b3978d4fad64379"

// BobAddress bob address
const BobAddress = "0xc8328aabcd9b9e8e64fbc566c4385c3bdeb219d7"

// BobSecKey bob private key
const BobSecKey = "d00c06bfd800d27397002dca6fb0993d5ba6399b4238b2f29ee9deb97593d2bc"

// AliceAddress Alice address
const AliceAddress = "0x470dcdc5e44064909650113a274b3b36aecb6dc7"

// SystemCellLockCodeHash system cell lock code hash
const SystemCellLockCodeHash = "0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8"

// M5 5 millions
const M5 = "0x1c6bf52634000"

// M49 4.9 millions
const M49 = "0x1bda703f0a000"

// M10 10 millions
const M10 = "0x38d7ea4c68000"

// RPCResultTransaction transaction response from jsonrpc
type RPCResultTransaction struct {
	Version     t.Uint32       `json:"version"`
	Hash        string         `json:"hash"`
	CellDeps    []t.CellDep    `json:"cell_deps"`
	HeaderDeps  []t.H256       `json:"header_deps"`
	Inputs      []t.CellInput  `json:"inputs"`
	Outputs     []t.CellOutput `json:"outputs"`
	Witnesses   []t.JSONBytes  `json:"witnesses"`
	OutputsData []t.JSONBytes  `json:"outputs_data"`
}

// RPCResultBlock block response from jsonrpc
type RPCResultBlock struct {
	Header       t.Header               `json:"header"`
	Proposals    []t.ProposalShortID    `json:"proposals"`
	Transactions []RPCResultTransaction `json:"transactions"`
	Uncles       []t.UncleBlock         `json:"uncles"`
}

func main() {
	rpcClient := jsonrpc.NewClient("http://127.0.0.1:8114")

	resp, err := rpcClient.Call("get_block", GenesisBlockHash)
	if err != nil {
		fmt.Printf("Err %v", err)
		return
	}

	genesisBlock := new(RPCResultBlock)

	err = resp.GetObject(&genesisBlock)
	if err != nil || genesisBlock == nil {
		fmt.Printf("Unmarsh json genesisBlock fail: %v", err)
	}

	if genesisBlock.Transactions[0].Outputs[6].Lock.Args != BobAddress {
		fmt.Println("Genesis bob address changed")
		return
	}

	if genesisBlock.Transactions[0].Outputs[7].Lock.Args != AliceAddress {
		fmt.Println("Genesis alice address changed")
		return
	}

	// We try to transfer 5M from bob to alice

	// Calc cell deps
	secpDepGroup := genesisBlock.Transactions[1].Hash
	outpoint := t.OutPoint{
		TxHash: secpDepGroup,
		Index:  "0x0",
	}

	cellDep := t.CellDep{
		OutPoint: outpoint,
		DepType:  t.DepGroup,
	}

	// Calc input
	bobPrevOutPoint := t.OutPoint{
		TxHash: genesisBlock.Transactions[0].Hash,
		Index:  "0x6",
	}

	bobInput := t.CellInput{
		PreviousOutput: bobPrevOutPoint,
		Since:          "0x0",
	}

	// Calc outputs
	aliceScript := t.Script{
		Args:     AliceAddress,
		CodeHash: SystemCellLockCodeHash,
		HashType: t.Type,
	}

	aliceOutput := t.CellOutput{
		Capacity: M5,
		Lock:     aliceScript,
		Type:     nil,
	}

	bobScript := t.Script{
		Args:     BobAddress,
		CodeHash: SystemCellLockCodeHash,
		HashType: t.Type,
	}

	bobOutput := t.CellOutput{
		Capacity: M49,
		Lock:     bobScript,
		Type:     nil,
	}

	// Assemble transaction
	tx := t.Transaction{
		Version:     "0x0",
		CellDeps:    []t.CellDep{cellDep},
		HeaderDeps:  make([]t.H256, 0),
		Inputs:      []t.CellInput{bobInput},
		Outputs:     []t.CellOutput{aliceOutput, bobOutput},
		Witnesses:   make([]t.JSONBytes, 0),
		OutputsData: []t.JSONBytes{"0x", "0x"},
	}

	// Cacl witness

	// Serialize transaction
	txBlob, err := t.Encode(tx)
	if err != nil {
		fmt.Printf("Encode failure: %s", err)
		return
	}

	// Prepare hash
	config := &blake2b.Config{
		Size:   32,
		Person: []byte(CkbBlake2BHashPersonalization),
	}
	h, err := blake2b.New(config)
	if err != nil {
		fmt.Printf("Initial blake2b hash failure, %s\n", err)
		return
	}

	// Hash tx
	h.Write(txBlob)
	txHash := h.Sum(nil)

	// Hash again
	h.Reset()
	h.Write(txHash)
	witnessMessage := h.Sum(nil)

	// Sign witness message
	bobSecKey := make([]byte, 32)
	_, err = hex.Decode(bobSecKey, []byte(BobSecKey))
	if err != nil {
		fmt.Println("Invalid bob private key")
	}

	witnessSig, err := secp256k1.Sign(witnessMessage[:], bobSecKey)
	if err != nil {
		fmt.Printf("Calc witness failure: %s\n", err)
		return
	}

	// Hex witness signature
    // NOTE: 130 is secp256k1 sig with recovery pubkey
	witness := make([]byte, 130)
	hex.Encode(witness, witnessSig[:])

	// Update transaction with witness
	tx.Witnesses = []t.JSONBytes{fmt.Sprintf("0x%s", witness)}

	resp, err = rpcClient.Call("send_transaction", []*t.Transaction{&tx})
	if err != nil {
		fmt.Printf("Err %v", err)
		return
	}

	fmt.Printf("transfer response %v", resp)
}
```
