package types

import (
	"encoding/json"
	"errors"
	"reflect"
)

func typeName(v interface{}) string {
	if t := reflect.ValueOf(v); t.Type().Kind() == reflect.Ptr {
		return reflect.Indirect(t).Type().Name()
	} else {
		return t.Type().Name()
	}
}

func Encode(v interface{}) ([]byte, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	buf := new(Buffer)

	retcode := CkbEncode(buf, []byte(typeName(v)), b)
	if retcode != 0 {
		return nil, errors.New("encode failure, invalid or unsupport type")
	}

	return buf.toBytes(), nil
}

func Decode(b []byte, v interface{}) error {
	buf := new(Buffer)
	mol_buf := newBufferFromBytes(b)
	defer mol_buf.Free()

	retcode := CkbDecode(buf, []byte(typeName(v)), mol_buf)
	if retcode != 0 {
		return errors.New("decode failure, invalid or unsupport type")
	}

	return json.Unmarshal(buf.toBytes(), v)
}
