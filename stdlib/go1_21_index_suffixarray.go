//go:build go1.21 && !go1.22
// +build go1.21,!go1.22

package stdlib

import (
	"index/suffixarray"
	"reflect"
)

func init() {
	Symbols["index/suffixarray/suffixarray"] = map[string]reflect.Value{

		"New": reflect.ValueOf(suffixarray.New),

		"Index": reflect.ValueOf((*suffixarray.Index)(nil)),
	}
}
