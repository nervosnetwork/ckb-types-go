package types

import "testing"

func TestCodec(t *testing.T) {
	testTx := Transaction{
		Version:     "0x1",
		CellDeps:    make([]CellDep, 0),
		HeaderDeps:  make([]H256, 0),
		Inputs:      make([]CellInput, 0),
		Outputs:     make([]CellOutput, 0),
		Witnesses:   make([]JSONBytes, 0),
		OutputsData: make([]JSONBytes, 0),
	}

	blob, err := Encode(testTx)
	if err != nil {
		t.Errorf("Encode failure: %s", err)
		return
	}

	tx := new(Transaction)

	err = Decode(blob, tx)
	if err != nil {
		t.Errorf("Decode failure: %s", err)
		return
	}

	if tx.Version != "0x1" {
		t.Errorf("Decode failure, version should be 0x01, got %s", tx.Version)
	}
}
