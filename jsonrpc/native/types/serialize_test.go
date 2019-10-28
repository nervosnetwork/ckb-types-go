package types

import (
	"bytes"
	tt "github.com/zeroqn/ckb-types-go/jsonrpc/types"
	"testing"
)

func TestSerializeScript(t *testing.T) {
	s := tt.Script{
		CodeHash: "0xe49352ee4984694d88eb3c1493a33d69d61c786dc5b0a32c4b3978d4fad64379",
		HashType: tt.Type,
		Args:     "0x470dcdc5e44064909650113a274b3b36aecb6dc7",
	}

	expect, err := tt.Encode(s)
	if err != nil {
		t.Errorf("Fail to encode script through ffi: %s\n", err)
		return
	}

	ss := Script{
		CodeHash: "0xe49352ee4984694d88eb3c1493a33d69d61c786dc5b0a32c4b3978d4fad64379",
		HashType: Type,
		Args:     "0x470dcdc5e44064909650113a274b3b36aecb6dc7",
	}

	got, err := ss.Serialize()
	if err != nil {
		t.Errorf("fail to serialize through native: %s\n", err)
		return
	}

	if !bytes.Equal(expect, got) {
		t.Errorf("mismatch result, expect %v, got %v", expect, got)
		return
	}
}

func TestSerializeOutPoint(t *testing.T) {
	o := tt.OutPoint{
		TxHash: "0xe49352ee4984694d88eb3c1493a33d69d61c786dc5b0a32c4b3978d4fad64379",
		Index:  "0x6",
	}

	expect, err := tt.Encode(o)
	if err != nil {
		t.Errorf("Fail to encode script through ffi: %s\n", err)
		return
	}

	oo := OutPoint{
		TxHash: "0xe49352ee4984694d88eb3c1493a33d69d61c786dc5b0a32c4b3978d4fad64379",
		Index:  "0x6",
	}

	got, err := oo.Serialize()
	if err != nil {
		t.Errorf("fail to serialize through native: %s\n", err)
		return
	}

	if !bytes.Equal(expect, got) {
		t.Errorf("mismatch result, expect %v, got %v", expect, got)
		return
	}
}

func TestSerializeCellInput(t *testing.T) {
	o := tt.OutPoint{
		TxHash: "0xe49352ee4984694d88eb3c1493a33d69d61c786dc5b0a32c4b3978d4fad64379",
		Index:  "0x6",
	}

	i := tt.CellInput{
		PreviousOutput: o,
		Since:          "0x0",
	}

	expect, err := tt.Encode(i)
	if err != nil {
		t.Errorf("Fail to encode script through ffi: %s\n", err)
		return
	}

	oo := OutPoint{
		TxHash: "0xe49352ee4984694d88eb3c1493a33d69d61c786dc5b0a32c4b3978d4fad64379",
		Index:  "0x6",
	}

	ii := CellInput{
		PreviousOutput: oo,
		Since:          "0x0",
	}

	got, err := ii.Serialize()
	if err != nil {
		t.Errorf("fail to serialize through native: %s\n", err)
		return
	}

	if !bytes.Equal(expect, got) {
		t.Errorf("mismatch result, expect %v, got %v", expect, got)
		return
	}
}

func TestSerializeCellOutput(t *testing.T) {
	// Test without type script
	s := tt.Script{
		CodeHash: "0xe49352ee4984694d88eb3c1493a33d69d61c786dc5b0a32c4b3978d4fad64379",
		HashType: tt.Type,
		Args:     "0x470dcdc5e44064909650113a274b3b36aecb6dc7",
	}

	o := tt.CellOutput{
		Capacity: "0x6666",
		Lock:     s,
		Type:     nil,
	}

	expect, err := tt.Encode(o)
	if err != nil {
		t.Errorf("Fail to encode script through ffi: %s\n", err)
		return
	}

	ss := Script{
		CodeHash: "0xe49352ee4984694d88eb3c1493a33d69d61c786dc5b0a32c4b3978d4fad64379",
		HashType: Type,
		Args:     "0x470dcdc5e44064909650113a274b3b36aecb6dc7",
	}

	oo := CellOutput{
		Capacity: "0x6666",
		Lock:     ss,
		Type:     nil,
	}

	got, err := oo.Serialize()
	if err != nil {
		t.Errorf("fail to serialize through native: %s\n", err)
		return
	}

	if !bytes.Equal(expect, got) {
		t.Errorf("mismatch result, expect %v, got %v", expect, got)
		return
	}

	// Test with type script

	o.Type = &s

	expect, err = tt.Encode(o)
	if err != nil {
		t.Errorf("Fail to encode script through ffi: %s\n", err)
		return
	}

	oo.Type = &ss

	got, err = oo.Serialize()
	if err != nil {
		t.Errorf("fail to serialize through native: %s\n", err)
		return
	}

	if !bytes.Equal(expect, got) {
		t.Errorf("mismatch result, expect %v, got %v", expect, got)
		return
	}
}

func TestSerializeCellDep(t *testing.T) {
	o := tt.OutPoint{
		TxHash: "0xe49352ee4984694d88eb3c1493a33d69d61c786dc5b0a32c4b3978d4fad64379",
		Index:  "0x6",
	}

	d := tt.CellDep{
		OutPoint: o,
		DepType:  tt.DepGroup,
	}

	expect, err := tt.Encode(d)
	if err != nil {
		t.Errorf("Fail to encode script through ffi: %s\n", err)
		return
	}

	oo := OutPoint{
		TxHash: "0xe49352ee4984694d88eb3c1493a33d69d61c786dc5b0a32c4b3978d4fad64379",
		Index:  "0x6",
	}

	dd := CellDep{
		OutPoint: oo,
		DepType:  DepGroup,
	}

	got, err := dd.Serialize()
	if err != nil {
		t.Errorf("fail to serialize through native: %s\n", err)
		return
	}

	if !bytes.Equal(expect, got) {
		t.Errorf("mismatch result, expect %v, got %v", expect, got)
		return
	}
}

func TestSerializeTransaction(t *testing.T) {
	aliceAddress := "0x470dcdc5e44064909650113a274b3b36aecb6dc7"
	bobAddress := "0xc8328aabcd9b9e8e64fbc566c4385c3bdeb219d7"
	systemCellLockCodeHash := "0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8"
	m5 := "0x1c6bf52634000"

	secpDepGroup := "0xe49352ee4984694d88eb3c1493a33d69d61c786dc5b0a32c4b3978d4fad64379"
	outpoint := tt.OutPoint{
		TxHash: secpDepGroup,
		Index:  "0x0",
	}

	cellDep := tt.CellDep{
		OutPoint: outpoint,
		DepType:  tt.DepGroup,
	}

	// Calc input
	bobPrevOutPoint := tt.OutPoint{
		TxHash: "0xe49352ee4984694d88eb3c1493a33d69d61c786dc5b0a32c4b3978d4fad64379",
		Index:  "0x6",
	}

	bobInput := tt.CellInput{
		PreviousOutput: bobPrevOutPoint,
		Since:          "0x0",
	}

	// Calc outputs
	aliceScript := tt.Script{
		Args:     aliceAddress,
		CodeHash: systemCellLockCodeHash,
		HashType: tt.Type,
	}

	aliceOutput := tt.CellOutput{
		Capacity: m5,
		Lock:     aliceScript,
		Type:     nil,
	}

	bobScript := tt.Script{
		Args:     bobAddress,
		CodeHash: systemCellLockCodeHash,
		HashType: tt.Type,
	}

	bobOutput := tt.CellOutput{
		Capacity: m5,
		Lock:     bobScript,
		Type:     nil,
	}

	// Assemble transaction
	tx := tt.Transaction{
		Version:     "0x0",
		CellDeps:    []tt.CellDep{cellDep},
		HeaderDeps:  make([]tt.H256, 0),
		Inputs:      []tt.CellInput{bobInput},
		Outputs:     []tt.CellOutput{aliceOutput, bobOutput},
		Witnesses:   make([]tt.JSONBytes, 0),
		OutputsData: []tt.JSONBytes{"0x", "0x"},
	}

	expect, err := tt.Encode(tx)
	if err != nil {
		t.Errorf("Fail to encode script through ffi: %s\n", err)
		return
	}

	oo := OutPoint{
		TxHash: secpDepGroup,
		Index:  "0x0",
	}

	cd := CellDep{
		OutPoint: oo,
		DepType:  DepGroup,
	}

	// Calc input
	po := OutPoint{
		TxHash: "0xe49352ee4984694d88eb3c1493a33d69d61c786dc5b0a32c4b3978d4fad64379",
		Index:  "0x6",
	}

	bi := CellInput{
		PreviousOutput: po,
		Since:          "0x0",
	}

	// Calc outputs
	as := Script{
		Args:     aliceAddress,
		CodeHash: systemCellLockCodeHash,
		HashType: Type,
	}

	ao := CellOutput{
		Capacity: m5,
		Lock:     as,
		Type:     nil,
	}

	bs := Script{
		Args:     bobAddress,
		CodeHash: systemCellLockCodeHash,
		HashType: Type,
	}

	bo := CellOutput{
		Capacity: m5,
		Lock:     bs,
		Type:     nil,
	}

	// Assemble transaction
	tt := Transaction{
		Version:     "0x0",
		CellDeps:    []CellDep{cd},
		HeaderDeps:  make([]Hash, 0),
		Inputs:      []CellInput{bi},
		Outputs:     []CellOutput{ao, bo},
		Witnesses:   make([]Bytes, 0),
		OutputsData: []Bytes{"0x", "0x"},
	}

	got, err := tt.Serialize()
	if err != nil {
		t.Errorf("fail to serialize through native: %s\n", err)
		return
	}

	if !bytes.Equal(expect, got) {
		t.Errorf("mismatch result\n expect %v\n got %v\n", expect, got)
		return
	}
}
