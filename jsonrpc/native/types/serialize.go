package types

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

// SerializeHash serialize hash
func SerializeHash(h string) ([]byte, error) {
	if !strings.HasPrefix(h, "0x") {
		return nil, fmt.Errorf("invalid hash, should be 0x-prefix")
	}

	b, err := hex.DecodeString(h[2:])
	if err != nil {
		return nil, err
	}

	if len(b) != 32 {
		return nil, fmt.Errorf("invalid hash, should be 32 bytes")
	}

	return b, nil
}

// SerializeScriptHashType serialize script hash type
func SerializeScriptHashType(t string) ([]byte, error) {
	if strings.Compare(t, string(Data)) != 0 && strings.Compare(t, string(Type)) != 0 {
		return nil, fmt.Errorf("invalid script hash type")
	}

	if strings.Compare(t, string(Data)) == 0 {
		return []byte{00}, nil
	}

	return []byte{01}, nil
}

// SerializeDepType serialize dep type
func SerializeDepType(t string) ([]byte, error) {
	if strings.Compare(t, string(Code)) != 0 && strings.Compare(t, string(DepGroup)) != 0 {
		return nil, fmt.Errorf("invalid dep group")
	}

	if strings.Compare(t, string(Code)) == 0 {
		return []byte{00}, nil
	}

	return []byte{01}, nil
}

// SerializeBytes serialize bytes
func SerializeBytes(b string) ([]byte, error) {
	if !strings.HasPrefix(b, "0x") {
		return nil, fmt.Errorf("invalid hash, should be 0x-prefix")
	}

	// Fixvec, vector Bytes <byte>
	if len(b[2:]) == 0 {
		return []byte{00, 00, 00, 00}, nil
	}

	bs, err := hex.DecodeString(b[2:])
	if err != nil {
		return nil, err
	}

	return bs, nil
}

// SerializeStrUint32 serialize string represented uint32
func SerializeStrUint32(u string) ([]byte, error) {
	if !strings.HasPrefix(u, "0x") {
		return nil, fmt.Errorf("invalid uin32, should be 0x-prefix")
	}

	uu := u[2:]
	if len(u)%2 != 0 {
		uu = "0" + uu
	}

	n, err := strconv.ParseUint(uu, 16, 32)
	if err != nil {
		return nil, err
	}

	return SerializeUint32(uint32(n)), nil
}

// SerializeStrUint64 serialize string represented uint64
func SerializeStrUint64(u string) ([]byte, error) {
	if !strings.HasPrefix(u, "0x") {
		return nil, fmt.Errorf("invalid uint64, should be 0x-prefix")
	}

	uu := u[2:]
	if len(u)%2 != 0 {
		uu = "0" + uu
	}

	n, err := strconv.ParseUint(uu, 16, 64)
	if err != nil {
		return nil, err
	}

	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, n)

	return b, nil
}

// SerializeUint32 serialize uint32
func SerializeUint32(n uint32) []byte {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, n)

	return b
}

// SerializeFixVec serialize fixvec
func SerializeFixVec(v [][]byte) []byte {
	if len(v) == 0 {
		return []byte{00, 00, 00, 00}
	}

	l := SerializeUint32(uint32(len(v)))

	b := new(bytes.Buffer)

	b.Write(l)

	for i := 0; i < len(v); i++ {
		b.Write(v[i])
	}

	return b.Bytes()
}

// SerializeDynVec serialize dynvec
func SerializeDynVec(v [][]byte) []byte {
	size := 4
	if len(v) == 0 {
		return SerializeUint32(uint32(size))
	}

	offsets := make([]uint32, len(v))

	offsets[0] = uint32(4 + 4*len(v))
	for i := 0; i < len(v); i++ {
		size += 4 + len(v[i])

		if i != 0 {
			offsets[i] = offsets[i-1] + uint32(len(v[i-1]))
		}
	}

	b := new(bytes.Buffer)

	b.Write(SerializeUint32(uint32(size)))

	for i := 0; i < len(v); i++ {
		b.Write(SerializeUint32(uint32(offsets[i])))
	}

	for i := 0; i < len(v); i++ {
		b.Write(v[i])
	}

	return b.Bytes()
}

// Serialize script
func (s *Script) Serialize() ([]byte, error) {
	h, err := SerializeHash(s.CodeHash)
	if err != nil {
		return nil, err
	}

	t, err := SerializeScriptHashType(string(s.HashType))
	if err != nil {
		return nil, err
	}

	a, err := SerializeBytes(s.Args)
	if err != nil {
		return nil, err
	}

	size := 4 + 4*3 + len(h) + len(t) + len(a) + 4
	hOffset := 4 + 4*3
	tOffset := hOffset + len(h)
	aOffset := tOffset + len(t)

	b := new(bytes.Buffer)

	b.Write(SerializeUint32(uint32(size)))
	b.Write(SerializeUint32(uint32(hOffset)))
	b.Write(SerializeUint32(uint32(tOffset)))
	b.Write(SerializeUint32(uint32(aOffset)))
	b.Write(h)
	b.Write(t)
	b.Write(SerializeUint32(uint32(len(a))))
	b.Write(a)

	return b.Bytes(), nil
}

// Serialize outpoint
func (o *OutPoint) Serialize() ([]byte, error) {
	h, err := SerializeHash(o.TxHash)
	if err != nil {
		return nil, err
	}

	i, err := SerializeStrUint32(o.Index)
	if err != nil {
		return nil, err
	}

	b := new(bytes.Buffer)

	b.Write(h)
	b.Write(i)

	return b.Bytes(), nil
}

// Serialize cell input
func (i *CellInput) Serialize() ([]byte, error) {
	s, err := SerializeStrUint64(i.Since)
	if err != nil {
		return nil, err
	}

	o, err := i.PreviousOutput.Serialize()
	if err != nil {
		return nil, err
	}

	b := new(bytes.Buffer)

	b.Write(s)
	b.Write(o)

	return b.Bytes(), nil
}

// Serialize cell output
func (o *CellOutput) Serialize() ([]byte, error) {
	c, err := SerializeStrUint64(o.Capacity)
	if err != nil {
		return nil, err
	}

	l, err := o.Lock.Serialize()
	if err != nil {
		return nil, err
	}

	var t []byte
	if o.Type != nil {
		t, err = o.Type.Serialize()
		if err != nil {
			return nil, err
		}
	}

	size := 4 + 4*3 + len(c) + len(l)
	if len(t) != 0 {
		size += len(t)
	}
	cOffset := 4 + 4*3
	lOffset := cOffset + len(c)
	tOffset := lOffset + len(l)

	b := new(bytes.Buffer)

	b.Write(SerializeUint32(uint32(size)))
	b.Write(SerializeUint32(uint32(cOffset)))
	b.Write(SerializeUint32(uint32(lOffset)))
	b.Write(SerializeUint32(uint32(tOffset)))
	b.Write(c)
	b.Write(l)
	b.Write(t)

	return b.Bytes(), nil
}

// Serialize cell dep
func (d *CellDep) Serialize() ([]byte, error) {
	o, err := d.OutPoint.Serialize()
	if err != nil {
		return nil, err
	}

	dd, err := SerializeDepType(string(d.DepType))
	if err != nil {
		return nil, err
	}

	b := new(bytes.Buffer)

	b.Write(o)
	b.Write(dd)

	return b.Bytes(), nil
}

// Serialize transaction
func (t *Transaction) Serialize() ([]byte, error) {
	v, err := SerializeStrUint32(t.Version)
	if err != nil {
		return nil, err
	}

	cds := make([][]byte, len(t.CellDeps))
	for i := 0; i < len(t.CellDeps); i++ {
		cd, err := t.CellDeps[i].Serialize()
		if err != nil {
			return nil, err
		}

		cds[i] = cd
	}
	cdsBytes := SerializeFixVec(cds)

	hds := make([][]byte, len(t.HeaderDeps))
	for i := 0; i < len(t.HeaderDeps); i++ {
		hd, err := SerializeHash(t.HeaderDeps[i])
		if err != nil {
			return nil, err
		}

		hds[i] = hd
	}
	hdsBytes := SerializeFixVec(hds)

	ips := make([][]byte, len(t.Inputs))
	for i := 0; i < len(t.Inputs); i++ {
		ip, err := t.Inputs[i].Serialize()
		if err != nil {
			return nil, err
		}

		ips[i] = ip
	}
	ipsBytes := SerializeFixVec(ips)

	ops := make([][]byte, len(t.Outputs))
	for i := 0; i < len(t.Outputs); i++ {
		op, err := t.Outputs[i].Serialize()
		if err != nil {
			return nil, err
		}

		ops[i] = op
	}
	opsBytes := SerializeDynVec(ops)

	ods := make([][]byte, len(t.OutputsData))
	for i := 0; i < len(t.OutputsData); i++ {
		od, err := SerializeBytes(t.OutputsData[i])
		if err != nil {
			return nil, err
		}

		ods[i] = od
	}
	odsBytes := SerializeDynVec(ods)

	size := 4 + 4*6 + len(v) + len(cdsBytes) + len(hdsBytes) + len(ipsBytes) + len(opsBytes) + len(odsBytes)

	vOffset := 4 + 4*6
	cOffset := vOffset + len(v)
	hOffset := cOffset + len(cdsBytes)
	iOffset := hOffset + len(hdsBytes)
	oOffset := iOffset + len(ipsBytes)
	odOffset := oOffset + len(opsBytes)

	b := new(bytes.Buffer)

	b.Write(SerializeUint32(uint32(size)))
	b.Write(SerializeUint32(uint32(vOffset)))
	b.Write(SerializeUint32(uint32(cOffset)))
	b.Write(SerializeUint32(uint32(hOffset)))
	b.Write(SerializeUint32(uint32(iOffset)))
	b.Write(SerializeUint32(uint32(oOffset)))
	b.Write(SerializeUint32(uint32(odOffset)))
	b.Write(v)
	b.Write(cdsBytes)
	b.Write(hdsBytes)
	b.Write(ipsBytes)
	b.Write(opsBytes)
	b.Write(odsBytes)

	return b.Bytes(), nil
}
