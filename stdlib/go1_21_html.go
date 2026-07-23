//go:build go1.21 && !go1.22
// +build go1.21,!go1.22

package stdlib

import (
	"html"
	"reflect"
)

func init() {
	Symbols["html/html"] = map[string]reflect.Value{

		"EscapeString":   reflect.ValueOf(html.EscapeString),
		"UnescapeString": reflect.ValueOf(html.UnescapeString),
	}
}
