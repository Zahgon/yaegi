//go:build go1.22
// +build go1.22

package stdlib

import (
	"crypto/hmac"
	"reflect"
)

func init() {
	Symbols["crypto/hmac/hmac"] = map[string]reflect.Value{

		"Equal": reflect.ValueOf(hmac.Equal),
		"New":   reflect.ValueOf(hmac.New),
	}
}
