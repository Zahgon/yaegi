//go:build go1.21 && !go1.22
// +build go1.21,!go1.22

package stdlib

import (
	"io/ioutil"
	"reflect"
)

func init() {
	Symbols["io/ioutil/ioutil"] = map[string]reflect.Value{

		"Discard":   reflect.ValueOf(&ioutil.Discard).Elem(),
		"NopCloser": reflect.ValueOf(ioutil.NopCloser),
		"ReadAll":   reflect.ValueOf(ioutil.ReadAll),
		"ReadDir":   reflect.ValueOf(ioutil.ReadDir),
		"ReadFile":  reflect.ValueOf(ioutil.ReadFile),
		"TempDir":   reflect.ValueOf(ioutil.TempDir),
		"TempFile":  reflect.ValueOf(ioutil.TempFile),
		"WriteFile": reflect.ValueOf(ioutil.WriteFile),
	}
}
