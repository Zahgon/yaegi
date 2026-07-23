//go:build go1.22
// +build go1.22

package stdlib

import (
	"encoding/ascii85"
	"reflect"
)

func init() {
	Symbols["encoding/ascii85/ascii85"] = map[string]reflect.Value{

		"Decode":        reflect.ValueOf(ascii85.Decode),
		"Encode":        reflect.ValueOf(ascii85.Encode),
		"MaxEncodedLen": reflect.ValueOf(ascii85.MaxEncodedLen),
		"NewDecoder":    reflect.ValueOf(ascii85.NewDecoder),
		"NewEncoder":    reflect.ValueOf(ascii85.NewEncoder),

		"CorruptInputError": reflect.ValueOf((*ascii85.CorruptInputError)(nil)),
	}
}
