//go:build go1.21 && !go1.22
// +build go1.21,!go1.22

package stdlib

import (
	"container/ring"
	"reflect"
)

func init() {
	Symbols["container/ring/ring"] = map[string]reflect.Value{

		"New": reflect.ValueOf(ring.New),

		"Ring": reflect.ValueOf((*ring.Ring)(nil)),
	}
}
