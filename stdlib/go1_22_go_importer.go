//go:build go1.22
// +build go1.22

package stdlib

import (
	"go/importer"
	"reflect"
)

func init() {
	Symbols["go/importer/importer"] = map[string]reflect.Value{

		"Default":     reflect.ValueOf(importer.Default),
		"For":         reflect.ValueOf(importer.For),
		"ForCompiler": reflect.ValueOf(importer.ForCompiler),

		"Lookup": reflect.ValueOf((*importer.Lookup)(nil)),
	}
}
