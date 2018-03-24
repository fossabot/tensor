package core

import (
	"unsafe"
)

type Buffer struct {
	data []byte
	typ  DType
}

func (b *Buffer) DSet(i int, v interface{}) {

}

func (b *Buffer) DAt(i int, f func(p unsafe.Pointer)) {

}

func (b *Buffer) CSet(i int, v interface{}) error {

}

func (b *Buffer) CAt(i int, f func(p unsafe.Pointer)) error {
	return nil
}

func (b *Buffer) Iterate(f func(i int, p unsafe.Pointer)) {}
