//go:build go1.21 && !go1.22
// +build go1.21,!go1.22

package stdlib

import (
	"go/constant"
	"go/token"
	"hash/adler32"
	"reflect"
)

func init() {
	Symbols["hash/adler32/adler32"] = map[string]reflect.Value{

		"Checksum": reflect.ValueOf(adler32.Checksum),
		"New":      reflect.ValueOf(adler32.New),
		"Size":     reflect.ValueOf(constant.MakeFromLiteral("4", token.INT, 0)),
	}
}
