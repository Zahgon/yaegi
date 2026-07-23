//go:build go1.22
// +build go1.22

package stdlib

import (
	"go/format"
	"reflect"
)

func init() {
	Symbols["go/format/format"] = map[string]reflect.Value{

		"Node":   reflect.ValueOf(format.Node),
		"Source": reflect.ValueOf(format.Source),
	}
}
