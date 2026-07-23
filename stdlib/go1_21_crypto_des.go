//go:build go1.21 && !go1.22
// +build go1.21,!go1.22

package stdlib

import (
	"crypto/des"
	"go/constant"
	"go/token"
	"reflect"
)

func init() {
	Symbols["crypto/des/des"] = map[string]reflect.Value{

		"BlockSize":          reflect.ValueOf(constant.MakeFromLiteral("8", token.INT, 0)),
		"NewCipher":          reflect.ValueOf(des.NewCipher),
		"NewTripleDESCipher": reflect.ValueOf(des.NewTripleDESCipher),

		"KeySizeError": reflect.ValueOf((*des.KeySizeError)(nil)),
	}
}
