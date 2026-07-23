//go:build go1.21 && !go1.22
// +build go1.21,!go1.22

package stdlib

import (
	"encoding/gob"
	"reflect"
)

func init() {
	Symbols["encoding/gob/gob"] = map[string]reflect.Value{

		"NewDecoder":   reflect.ValueOf(gob.NewDecoder),
		"NewEncoder":   reflect.ValueOf(gob.NewEncoder),
		"Register":     reflect.ValueOf(gob.Register),
		"RegisterName": reflect.ValueOf(gob.RegisterName),

		"CommonType": reflect.ValueOf((*gob.CommonType)(nil)),
		"Decoder":    reflect.ValueOf((*gob.Decoder)(nil)),
		"Encoder":    reflect.ValueOf((*gob.Encoder)(nil)),
		"GobDecoder": reflect.ValueOf((*gob.GobDecoder)(nil)),
		"GobEncoder": reflect.ValueOf((*gob.GobEncoder)(nil)),

		"_GobDecoder": reflect.ValueOf((*_encoding_gob_GobDecoder)(nil)),
		"_GobEncoder": reflect.ValueOf((*_encoding_gob_GobEncoder)(nil)),
	}
}

type _encoding_gob_GobDecoder struct {
	IValue     interface{}
	WGobDecode func(a0 []byte) error
}

func (W _encoding_gob_GobDecoder) GobDecode(a0 []byte) error { _ = "STUB: not implemented"; return nil }

type _encoding_gob_GobEncoder struct {
	IValue     interface{}
	WGobEncode func() ([]byte, error)
}

func (W _encoding_gob_GobEncoder) GobEncode() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
