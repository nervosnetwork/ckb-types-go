package types

/*
#cgo LDFLAGS: -static -L${SRCDIR}/../ -lckb_ffi -lpthread -ldl
#include "../ckb_ffi.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import "unsafe"

func (x *Buffer) toBytes() []byte {
	if x.refb0a5a638 == nil {
		return nil
	}

	return C.GoBytes(unsafe.Pointer(x.refb0a5a638.data), C.int(x.refb0a5a638.len))
}

func newBufferFromBytes(b []byte) *Buffer {
	if b == nil || len(b) == 0 {
		return nil
	}

	buf := new(Buffer)
	length := len(b)

	buf.Len = *(*uint)(unsafe.Pointer(&length))
	buf.Data = (*byte)(unsafe.Pointer(C.CBytes(b)))

	return buf
}
