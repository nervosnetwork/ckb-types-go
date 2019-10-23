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
import (
	"sync"
	"unsafe"
)

// cgoAllocMap stores pointers to C allocated memory for future reference.
type cgoAllocMap struct {
	mux sync.RWMutex
	m   map[unsafe.Pointer]struct{}
}

var cgoAllocsUnknown = new(cgoAllocMap)

func (a *cgoAllocMap) Add(ptr unsafe.Pointer) {
	a.mux.Lock()
	if a.m == nil {
		a.m = make(map[unsafe.Pointer]struct{})
	}
	a.m[ptr] = struct{}{}
	a.mux.Unlock()
}

func (a *cgoAllocMap) IsEmpty() bool {
	a.mux.RLock()
	isEmpty := len(a.m) == 0
	a.mux.RUnlock()
	return isEmpty
}

func (a *cgoAllocMap) Borrow(b *cgoAllocMap) {
	if b == nil || b.IsEmpty() {
		return
	}
	b.mux.Lock()
	a.mux.Lock()
	for ptr := range b.m {
		if a.m == nil {
			a.m = make(map[unsafe.Pointer]struct{})
		}
		a.m[ptr] = struct{}{}
		delete(b.m, ptr)
	}
	a.mux.Unlock()
	b.mux.Unlock()
}

func (a *cgoAllocMap) Free() {
	a.mux.Lock()
	for ptr := range a.m {
		C.free(ptr)
		delete(a.m, ptr)
	}
	a.mux.Unlock()
}

// allocBufferMemory allocates memory for type C.buffer_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocBufferMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfBufferValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfBufferValue = unsafe.Sizeof([1]C.buffer_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *Buffer) Ref() *C.buffer_t {
	if x == nil {
		return nil
	}
	return x.refb0a5a638
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *Buffer) Free() {
	if x != nil && x.allocsb0a5a638 != nil {
		x.allocsb0a5a638.(*cgoAllocMap).Free()
		x.refb0a5a638 = nil
	}
}

// NewBufferRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewBufferRef(ref unsafe.Pointer) *Buffer {
	if ref == nil {
		return nil
	}
	obj := new(Buffer)
	obj.refb0a5a638 = (*C.buffer_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *Buffer) PassRef() (*C.buffer_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refb0a5a638 != nil {
		return x.refb0a5a638, nil
	}
	memb0a5a638 := allocBufferMemory(1)
	refb0a5a638 := (*C.buffer_t)(memb0a5a638)
	allocsb0a5a638 := new(cgoAllocMap)
	allocsb0a5a638.Add(memb0a5a638)

	var clen_allocs *cgoAllocMap
	refb0a5a638.len, clen_allocs = (C.uint64_t)(x.Len), cgoAllocsUnknown
	allocsb0a5a638.Borrow(clen_allocs)

	var cdata_allocs *cgoAllocMap
	refb0a5a638.data, cdata_allocs = *(**C.uint8_t)(unsafe.Pointer(&x.Data)), cgoAllocsUnknown
	allocsb0a5a638.Borrow(cdata_allocs)

	x.refb0a5a638 = refb0a5a638
	x.allocsb0a5a638 = allocsb0a5a638
	return refb0a5a638, allocsb0a5a638

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x Buffer) PassValue() (C.buffer_t, *cgoAllocMap) {
	if x.refb0a5a638 != nil {
		return *x.refb0a5a638, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *Buffer) Deref() {
	if x.refb0a5a638 == nil {
		return
	}
	x.Len = (uint)(x.refb0a5a638.len)
	x.Data = (*byte)(unsafe.Pointer(x.refb0a5a638.data))
}

type sliceHeader struct {
	Data unsafe.Pointer
	Len  int
	Cap  int
}
