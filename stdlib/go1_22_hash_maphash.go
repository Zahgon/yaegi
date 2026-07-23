//go:build go1.22
// +build go1.22

package stdlib

import (
	"hash/maphash"
	"reflect"
)

func init() {
	Symbols["hash/maphash/maphash"] = map[string]reflect.Value{

		"Bytes":    reflect.ValueOf(maphash.Bytes),
		"MakeSeed": reflect.ValueOf(maphash.MakeSeed),
		"String":   reflect.ValueOf(maphash.String),

		"Hash": reflect.ValueOf((*maphash.Hash)(nil)),
		"Seed": reflect.ValueOf((*maphash.Seed)(nil)),
	}
}
