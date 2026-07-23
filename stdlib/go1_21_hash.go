//go:build go1.21 && !go1.22
// +build go1.21,!go1.22

package stdlib

import (
	"hash"
	"reflect"
)

func init() {
	Symbols["hash/hash"] = map[string]reflect.Value{

		"Hash":   reflect.ValueOf((*hash.Hash)(nil)),
		"Hash32": reflect.ValueOf((*hash.Hash32)(nil)),
		"Hash64": reflect.ValueOf((*hash.Hash64)(nil)),

		"_Hash":   reflect.ValueOf((*_hash_Hash)(nil)),
		"_Hash32": reflect.ValueOf((*_hash_Hash32)(nil)),
		"_Hash64": reflect.ValueOf((*_hash_Hash64)(nil)),
	}
}

type _hash_Hash struct {
	IValue     interface{}
	WBlockSize func() int
	WReset     func()
	WSize      func() int
	WSum       func(b []byte) []byte
	WWrite     func(p []byte) (n int, err error)
}

func (W _hash_Hash) BlockSize() int                    { _ = "STUB: not implemented"; return 0 }
func (W _hash_Hash) Reset()                            { _ = "STUB: not implemented"; return }
func (W _hash_Hash) Size() int                         { _ = "STUB: not implemented"; return 0 }
func (W _hash_Hash) Sum(b []byte) []byte               { _ = "STUB: not implemented"; return nil }
func (W _hash_Hash) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

type _hash_Hash32 struct {
	IValue     interface{}
	WBlockSize func() int
	WReset     func()
	WSize      func() int
	WSum       func(b []byte) []byte
	WSum32     func() uint32
	WWrite     func(p []byte) (n int, err error)
}

func (W _hash_Hash32) BlockSize() int                    { _ = "STUB: not implemented"; return 0 }
func (W _hash_Hash32) Reset()                            { _ = "STUB: not implemented"; return }
func (W _hash_Hash32) Size() int                         { _ = "STUB: not implemented"; return 0 }
func (W _hash_Hash32) Sum(b []byte) []byte               { _ = "STUB: not implemented"; return nil }
func (W _hash_Hash32) Sum32() uint32                     { _ = "STUB: not implemented"; return 0 }
func (W _hash_Hash32) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

type _hash_Hash64 struct {
	IValue     interface{}
	WBlockSize func() int
	WReset     func()
	WSize      func() int
	WSum       func(b []byte) []byte
	WSum64     func() uint64
	WWrite     func(p []byte) (n int, err error)
}

func (W _hash_Hash64) BlockSize() int                    { _ = "STUB: not implemented"; return 0 }
func (W _hash_Hash64) Reset()                            { _ = "STUB: not implemented"; return }
func (W _hash_Hash64) Size() int                         { _ = "STUB: not implemented"; return 0 }
func (W _hash_Hash64) Sum(b []byte) []byte               { _ = "STUB: not implemented"; return nil }
func (W _hash_Hash64) Sum64() uint64                     { _ = "STUB: not implemented"; return 0 }
func (W _hash_Hash64) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }
