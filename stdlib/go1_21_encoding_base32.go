//go:build go1.21 && !go1.22
// +build go1.21,!go1.22

package stdlib

import (
	"encoding/base32"
	"reflect"
)

func init() {
	Symbols["encoding/base32/base32"] = map[string]reflect.Value{

		"HexEncoding": reflect.ValueOf(&base32.HexEncoding).Elem(),
		"NewDecoder":  reflect.ValueOf(base32.NewDecoder),
		"NewEncoder":  reflect.ValueOf(base32.NewEncoder),
		"NewEncoding": reflect.ValueOf(base32.NewEncoding),
		"NoPadding":   reflect.ValueOf(base32.NoPadding),
		"StdEncoding": reflect.ValueOf(&base32.StdEncoding).Elem(),
		"StdPadding":  reflect.ValueOf(base32.StdPadding),

		"CorruptInputError": reflect.ValueOf((*base32.CorruptInputError)(nil)),
		"Encoding":          reflect.ValueOf((*base32.Encoding)(nil)),
	}
}
