//go:build go1.21 && !go1.22
// +build go1.21,!go1.22

package stdlib

import (
	"net/http/fcgi"
	"reflect"
)

func init() {
	Symbols["net/http/fcgi/fcgi"] = map[string]reflect.Value{

		"ErrConnClosed":     reflect.ValueOf(&fcgi.ErrConnClosed).Elem(),
		"ErrRequestAborted": reflect.ValueOf(&fcgi.ErrRequestAborted).Elem(),
		"ProcessEnv":        reflect.ValueOf(fcgi.ProcessEnv),
		"Serve":             reflect.ValueOf(fcgi.Serve),
	}
}
