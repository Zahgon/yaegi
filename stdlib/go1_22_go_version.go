//go:build go1.22
// +build go1.22

package stdlib

import (
	"go/version"
	"reflect"
)

func init() {
	Symbols["go/version/version"] = map[string]reflect.Value{

		"Compare": reflect.ValueOf(version.Compare),
		"IsValid": reflect.ValueOf(version.IsValid),
		"Lang":    reflect.ValueOf(version.Lang),
	}
}
