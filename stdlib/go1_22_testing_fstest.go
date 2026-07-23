//go:build go1.22
// +build go1.22

package stdlib

import (
	"reflect"
	"testing/fstest"
)

func init() {
	Symbols["testing/fstest/fstest"] = map[string]reflect.Value{

		"TestFS": reflect.ValueOf(fstest.TestFS),

		"MapFS":   reflect.ValueOf((*fstest.MapFS)(nil)),
		"MapFile": reflect.ValueOf((*fstest.MapFile)(nil)),
	}
}
