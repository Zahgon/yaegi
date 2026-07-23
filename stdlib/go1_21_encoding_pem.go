//go:build go1.21 && !go1.22
// +build go1.21,!go1.22

package stdlib

import (
	"encoding/pem"
	"reflect"
)

func init() {
	Symbols["encoding/pem/pem"] = map[string]reflect.Value{

		"Decode":         reflect.ValueOf(pem.Decode),
		"Encode":         reflect.ValueOf(pem.Encode),
		"EncodeToMemory": reflect.ValueOf(pem.EncodeToMemory),

		"Block": reflect.ValueOf((*pem.Block)(nil)),
	}
}
