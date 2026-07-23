//go:build go1.22
// +build go1.22

package stdlib

import (
	"crypto/aes"
	"go/constant"
	"go/token"
	"reflect"
)

func init() {
	Symbols["crypto/aes/aes"] = map[string]reflect.Value{

		"BlockSize": reflect.ValueOf(constant.MakeFromLiteral("16", token.INT, 0)),
		"NewCipher": reflect.ValueOf(aes.NewCipher),

		"KeySizeError": reflect.ValueOf((*aes.KeySizeError)(nil)),
	}
}
