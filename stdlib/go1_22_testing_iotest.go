//go:build go1.22
// +build go1.22

package stdlib

import (
	"reflect"
	"testing/iotest"
)

func init() {
	Symbols["testing/iotest/iotest"] = map[string]reflect.Value{

		"DataErrReader":  reflect.ValueOf(iotest.DataErrReader),
		"ErrReader":      reflect.ValueOf(iotest.ErrReader),
		"ErrTimeout":     reflect.ValueOf(&iotest.ErrTimeout).Elem(),
		"HalfReader":     reflect.ValueOf(iotest.HalfReader),
		"NewReadLogger":  reflect.ValueOf(iotest.NewReadLogger),
		"NewWriteLogger": reflect.ValueOf(iotest.NewWriteLogger),
		"OneByteReader":  reflect.ValueOf(iotest.OneByteReader),
		"TestReader":     reflect.ValueOf(iotest.TestReader),
		"TimeoutReader":  reflect.ValueOf(iotest.TimeoutReader),
		"TruncateWriter": reflect.ValueOf(iotest.TruncateWriter),
	}
}
