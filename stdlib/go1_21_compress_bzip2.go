//go:build go1.21 && !go1.22
// +build go1.21,!go1.22

package stdlib

import (
	"compress/bzip2"
	"reflect"
)

func init() {
	Symbols["compress/bzip2/bzip2"] = map[string]reflect.Value{

		"NewReader": reflect.ValueOf(bzip2.NewReader),

		"StructuralError": reflect.ValueOf((*bzip2.StructuralError)(nil)),
	}
}
