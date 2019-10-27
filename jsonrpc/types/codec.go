package types

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func typeName(v interface{}) string {
	t := reflect.ValueOf(v)

	if t.Type().Kind() == reflect.Ptr {
		return reflect.Indirect(t).Type().Name()
	}
	return t.Type().Name()
}

// Encode given ckb type to []byte
func Encode(v interface{}) ([]byte, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	buf := new(Buffer)

	retcode := CkbEncode(buf, []byte(typeName(v)), b)
	if retcode != 0 {
		return nil, fmt.Errorf("encode failure, code %d", retcode)
	}

	return buf.toBytes(), nil
}

// Decode given []byte to specified ckb type
func Decode(b []byte, v interface{}) error {
	buf := new(Buffer)
	molbuf := newBufferFromBytes(b)
	defer molbuf.Free()

	retcode := CkbDecode(buf, []byte(typeName(v)), molbuf)
	if retcode != 0 {
		return fmt.Errorf("encode failure, code %d", retcode)
	}

	return json.Unmarshal(buf.toBytes(), v)
}
