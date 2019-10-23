// MIT

// WARNING: This file has automatically been generated on Thu, 24 Oct 2019 15:00:44 CST.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package types

/*
#cgo LDFLAGS: -static -L${SRCDIR}/../ -lckb_ffi -lpthread -ldl
#include "../ckb_ffi.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import "unsafe"

// CkbEncode function as declared in ckb-types-go/ckb_ffi.h:21
func CkbEncode(OutputTx *Buffer, TypeName []byte, JsonTx []byte) int32 {
	cOutputTx, _ := OutputTx.PassRef()
	cTypeName, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&TypeName)).Data)), cgoAllocsUnknown
	cJsonTx, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&JsonTx)).Data)), cgoAllocsUnknown
	__ret := C.ckb_encode(cOutputTx, cTypeName, cJsonTx)
	__v := (int32)(__ret)
	return __v
}

// CkbDecode function as declared in ckb-types-go/ckb_ffi.h:22
func CkbDecode(OutputTx *Buffer, TypeName []byte, MolTx *Buffer) int32 {
	cOutputTx, _ := OutputTx.PassRef()
	cTypeName, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&TypeName)).Data)), cgoAllocsUnknown
	cMolTx, _ := MolTx.PassRef()
	__ret := C.ckb_decode(cOutputTx, cTypeName, cMolTx)
	__v := (int32)(__ret)
	return __v
}
