//go:build go1.21 && !go1.22
// +build go1.21,!go1.22

package stdlib

import (
	"crypto/rc4"
	"reflect"
)

func init() {
	Symbols["crypto/rc4/rc4"] = map[string]reflect.Value{

		"NewCipher": reflect.ValueOf(rc4.NewCipher),

		"Cipher":       reflect.ValueOf((*rc4.Cipher)(nil)),
		"KeySizeError": reflect.ValueOf((*rc4.KeySizeError)(nil)),
	}
}
