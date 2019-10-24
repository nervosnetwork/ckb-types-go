# ckb-types-go

### example

```go
package main

import "fmt"
import t "github.com/zeroqn/ckb-types-go/types"

func main() {
	testTx := t.Transaction{
		Version:     "0x1",
		CellDeps:    make([]t.CellDep, 0),
		HeaderDeps:  make([]t.H256, 0),
		Inputs:      make([]t.CellInput, 0),
		Outputs:     make([]t.CellOutput, 0),
		Witnesses:   make([]t.JSONBytes, 0),
		OutputsData: make([]t.JSONBytes, 0),
	}

	blob, err := t.Encode(testTx)
	if err != nil {
		fmt.Printf("Encode failure: %s", err)
		return
	}

	tx := new(t.Transaction)

	err = t.Decode(blob, tx)
	if err != nil {
		fmt.Printf("Decode failure: %s", err)
		return
	}

	fmt.Printf("tx version: %s", tx.Version)
}
```
