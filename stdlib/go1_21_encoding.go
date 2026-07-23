//go:build go1.21 && !go1.22
// +build go1.21,!go1.22

package stdlib

import (
	"encoding"
	"reflect"
)

func init() {
	Symbols["encoding/encoding"] = map[string]reflect.Value{

		"BinaryMarshaler":   reflect.ValueOf((*encoding.BinaryMarshaler)(nil)),
		"BinaryUnmarshaler": reflect.ValueOf((*encoding.BinaryUnmarshaler)(nil)),
		"TextMarshaler":     reflect.ValueOf((*encoding.TextMarshaler)(nil)),
		"TextUnmarshaler":   reflect.ValueOf((*encoding.TextUnmarshaler)(nil)),

		"_BinaryMarshaler":   reflect.ValueOf((*_encoding_BinaryMarshaler)(nil)),
		"_BinaryUnmarshaler": reflect.ValueOf((*_encoding_BinaryUnmarshaler)(nil)),
		"_TextMarshaler":     reflect.ValueOf((*_encoding_TextMarshaler)(nil)),
		"_TextUnmarshaler":   reflect.ValueOf((*_encoding_TextUnmarshaler)(nil)),
	}
}

type _encoding_BinaryMarshaler struct {
	IValue         interface{}
	WMarshalBinary func() (data []byte, err error)
}

func (W _encoding_BinaryMarshaler) MarshalBinary() (data []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type _encoding_BinaryUnmarshaler struct {
	IValue           interface{}
	WUnmarshalBinary func(data []byte) error
}

func (W _encoding_BinaryUnmarshaler) UnmarshalBinary(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

type _encoding_TextMarshaler struct {
	IValue       interface{}
	WMarshalText func() (text []byte, err error)
}

func (W _encoding_TextMarshaler) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type _encoding_TextUnmarshaler struct {
	IValue         interface{}
	WUnmarshalText func(text []byte) error
}

func (W _encoding_TextUnmarshaler) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}
