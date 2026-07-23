//go:build go1.22
// +build go1.22

package stdlib

import (
	"net/http/cgi"
	"reflect"
)

func init() {
	Symbols["net/http/cgi/cgi"] = map[string]reflect.Value{

		"Request":        reflect.ValueOf(cgi.Request),
		"RequestFromMap": reflect.ValueOf(cgi.RequestFromMap),
		"Serve":          reflect.ValueOf(cgi.Serve),

		"Handler": reflect.ValueOf((*cgi.Handler)(nil)),
	}
}
