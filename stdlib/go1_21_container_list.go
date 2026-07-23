//go:build go1.21 && !go1.22
// +build go1.21,!go1.22

package stdlib

import (
	"container/list"
	"reflect"
)

func init() {
	Symbols["container/list/list"] = map[string]reflect.Value{

		"New": reflect.ValueOf(list.New),

		"Element": reflect.ValueOf((*list.Element)(nil)),
		"List":    reflect.ValueOf((*list.List)(nil)),
	}
}
