package utils

import (
	"crypto/rand"
	"crypto/subtle"
	"reflect"
	"unsafe"
)

// SecureZeroMemory ensures data in RAM is overwritten before garbage collection
func SecureZeroBytes(b []byte) {
	if b == nil {
		return
	}
	// Overwrite the memory contents
	for i := range b {
		b[i] = 0
	}
	// Prevent compiler from optimizing this away
	rand.Read(b)
	
	// Pointer magic to ensure the slice header is also zeroed (advanced)
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh.Cap = 0
	sh.Len = 0
	sh.Data = 0
}