package slices

import (
	"reflect"
	"unsafe"
)

// Header returns the header (reflect.SliceHeader) of a slice
func Header[T any](s []T) *reflect.SliceHeader {
	return (*reflect.SliceHeader)((unsafe.Pointer)(&s))
}

// FromHeader returns a slice given its header (reflect.SliceHeader).
func FromHeader[T any](hdr *reflect.SliceHeader) []T {
	return *(*[]T)(unsafe.Pointer(hdr))
}
