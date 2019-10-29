package types

import (
	"encoding/hex"
	"encoding/json"
	"testing"
)

func TestSerializeScript(t *testing.T) {
	script := `{
		"code_hash": "0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8",
		"hash_type": "type",
		"args": "0xc8328aabcd9b9e8e64fbc566c4385c3bdeb219d7"
	}`

	expectHex := "490000001000000030000000310000009bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce80114000000c8328aabcd9b9e8e64fbc566c4385c3bdeb219d7"

	var s Script

	err := json.Unmarshal([]byte(script), &s)
	if err != nil {
		t.Errorf("fail to unmarshal test script json: %s\n", err)
		return
	}

	got, err := s.Serialize()
	if err != nil {
		t.Errorf("fail to serialize: %s\n", err)
		return
	}

	gotHex := hex.EncodeToString(got)

	if expectHex != gotHex {
		t.Errorf("mismatch result, expect %v, got %v", expectHex, gotHex)
		return
	}
}

func TestSerializeOutPoint(t *testing.T) {
	outpoint := `{
		"tx_hash": "0xe49352ee4984694d88eb3c1493a33d69d61c786dc5b0a32c4b3978d4fad64379",
		"index":  "0x6"
	}`

	expectHex := "e49352ee4984694d88eb3c1493a33d69d61c786dc5b0a32c4b3978d4fad6437906000000"

	var o OutPoint

	err := json.Unmarshal([]byte(outpoint), &o)
	if err != nil {
		t.Errorf("fail to unmarshal test outpoint json: %s\n", err)
		return
	}

	got, err := o.Serialize()
	if err != nil {
		t.Errorf("fail to serialize through native: %s\n", err)
		return
	}

	gotHex := hex.EncodeToString(got)

	if gotHex != expectHex {
		t.Errorf("mismatch result, expect %v, got %v", expectHex, gotHex)
		return
	}
}

func TestSerializeCellInput(t *testing.T) {
	input := `{
		"previous_output": {
			"tx_hash": "0xe49352ee4984694d88eb3c1493a33d69d61c786dc5b0a32c4b3978d4fad64379",
			"index": "0x6"
		},
		"since": "0x0"
	}`

	expectHex := "0000000000000000e49352ee4984694d88eb3c1493a33d69d61c786dc5b0a32c4b3978d4fad6437906000000"

	var i CellInput

	err := json.Unmarshal([]byte(input), &i)
	if err != nil {
		t.Errorf("fail to unmarshal test cell input json: %s\n", err)
		return
	}

	got, err := i.Serialize()
	if err != nil {
		t.Errorf("fail to serialize through native: %s\n", err)
		return
	}

	gotHex := hex.EncodeToString(got)

	if gotHex != expectHex {
		t.Errorf("mismatch result, expect %v, got %v", expectHex, gotHex)
		return
	}
}

func TestSerializeCellOutput(t *testing.T) {
	output := `{
		"capacity": "0x666",
		"lock": {
			"code_hash": "0xe49352ee4984694d88eb3c1493a33d69d61c786dc5b0a32c4b3978d4fad64379",
			"hash_type": "type",
			"args": "0x470dcdc5e44064909650113a274b3b36aecb6dc7"
		},
		"type": null
	}`

	expectHex := "61000000100000001800000061000000660600000000000049000000100000003000000031000000e49352ee4984694d88eb3c1493a33d69d61c786dc5b0a32c4b3978d4fad643790114000000470dcdc5e44064909650113a274b3b36aecb6dc7"

	var o CellOutput

	err := json.Unmarshal([]byte(output), &o)
	if err != nil {
		t.Errorf("fail to unmarshal test cell output json: %s\n", err)
		return
	}

	got, err := o.Serialize()
	if err != nil {
		t.Errorf("fail to serialize through native: %s\n", err)
		return
	}

	gotHex := hex.EncodeToString(got)

	if gotHex != expectHex {
		t.Errorf("mismatch result, expect %v, got %v", expectHex, gotHex)
		return
	}

	// Test with type script

	expectHex = "aa000000100000001800000061000000660600000000000049000000100000003000000031000000e49352ee4984694d88eb3c1493a33d69d61c786dc5b0a32c4b3978d4fad643790114000000470dcdc5e44064909650113a274b3b36aecb6dc749000000100000003000000031000000e49352ee4984694d88eb3c1493a33d69d61c786dc5b0a32c4b3978d4fad643790114000000470dcdc5e44064909650113a274b3b36aecb6dc7"

	o.Type = &o.Lock

	got, err = o.Serialize()
	if err != nil {
		t.Errorf("fail to serialize through native: %s\n", err)
		return
	}

	gotHex = hex.EncodeToString(got)

	if gotHex != expectHex {
		t.Errorf("mismatch result, expect %v, got %v", expectHex, gotHex)
		return
	}
}

func TestSerializeCellDep(t *testing.T) {
	dep := `{
		"out_point": {
			"tx_hash": "0xb815a396c5226009670e89ee514850dcde452bca746cdd6b41c104b50e559c70",
			"index": "0x0"
		},
		"dep_type": "dep_group"
	}`

	expectHex := "b815a396c5226009670e89ee514850dcde452bca746cdd6b41c104b50e559c700000000001"

	var d CellDep

	err := json.Unmarshal([]byte(dep), &d)
	if err != nil {
		t.Errorf("fail to unmarshal test cell dep json: %s\n", err)
		return
	}

	got, err := d.Serialize()
	if err != nil {
		t.Errorf("fail to serialize: %s\n", err)
		return
	}

	gotHex := hex.EncodeToString(got)

	if gotHex != expectHex {
		t.Errorf("mismatch result, expect %v, got %v", expectHex, gotHex)
		return
	}
}

func TestSerializeTransaction(t *testing.T) {
	transaction := `{
	  "cell_deps": [
		{
		  "out_point": {
			"tx_hash": "0xb815a396c5226009670e89ee514850dcde452bca746cdd6b41c104b50e559c70",
			"index": "0x0"
		  },
		  "dep_type": "dep_group"
		}
	  ],
	  "header_deps": [],
	  "inputs": [
		{
		  "previous_output": {
			"tx_hash": "0xee046ce2baeda575266d4164f394c53f66009f64759f7a9f12a014c692e79390",
			"index": "0x6"
		  },
		  "since": "0x0"
		}
	  ],
	  "outputs": [
		{
		  "capacity": "0x1c6bf52634000",
		  "lock": {
			"args": "0x470dcdc5e44064909650113a274b3b36aecb6dc7",
			"code_hash": "0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8",
			"hash_type": "type"
		  },
		  "type": null
		},
		{
		  "capacity": "0x1c6bf52634000",
		  "lock": {
			"args": "0xc8328aabcd9b9e8e64fbc566c4385c3bdeb219d7",
			"code_hash": "0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8",
			"hash_type": "type"
		  },
		  "type": null
		}
	  ],
	  "outputs_data": ["0x", "0x"],
	  "version": "0x0",
	  "witnesses": []
	}`

	expectHex := "5f0100001c00000020000000490000004d0000007d0000004b0100000000000001000000b815a396c5226009670e89ee514850dcde452bca746cdd6b41c104b50e559c70000000000100000000010000000000000000000000ee046ce2baeda575266d4164f394c53f66009f64759f7a9f12a014c692e7939006000000ce0000000c0000006d0000006100000010000000180000006100000000406352bfc60100490000001000000030000000310000009bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce80114000000470dcdc5e44064909650113a274b3b36aecb6dc76100000010000000180000006100000000406352bfc60100490000001000000030000000310000009bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce80114000000c8328aabcd9b9e8e64fbc566c4385c3bdeb219d7140000000c000000100000000000000000000000"

	var tx Transaction

	err := json.Unmarshal([]byte(transaction), &tx)
	if err != nil {
		t.Errorf("fail to unmarshal test transaction json: %s\n", err)
		return
	}

	got, err := tx.Serialize()
	if err != nil {
		t.Errorf("fail to serialize: %s\n", err)
		return
	}

	gotHex := hex.EncodeToString(got)

	if gotHex != expectHex {
		t.Errorf("mismatch result, expect %v, got %v", expectHex, gotHex)
		return
	}
}
