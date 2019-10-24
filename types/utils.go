package types

/*
#cgo LDFLAGS: -static -L${SRCDIR}/../ -lckb_ffi -lpthread -ldl
#include "../ckb_ffi.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import "unsafe"

func (b *Buffer) toBytes() []byte {
	if b == nil {
		return nil
	}

	b.Deref()

	return C.GoBytes(unsafe.Pointer(b.Data), C.int(b.Len))
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
