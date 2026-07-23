//go:build go1.21 && !go1.22
// +build go1.21,!go1.22

package stdlib

import (
	"reflect"
	"testing/slogtest"
)

func init() {
	Symbols["testing/slogtest/slogtest"] = map[string]reflect.Value{

		"TestHandler": reflect.ValueOf(slogtest.TestHandler),
	}
}
