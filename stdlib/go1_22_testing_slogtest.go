//go:build go1.22
// +build go1.22

package stdlib

import (
	"reflect"
	"testing/slogtest"
)

func init() {
	Symbols["testing/slogtest/slogtest"] = map[string]reflect.Value{

		"Run":         reflect.ValueOf(slogtest.Run),
		"TestHandler": reflect.ValueOf(slogtest.TestHandler),
	}
}
