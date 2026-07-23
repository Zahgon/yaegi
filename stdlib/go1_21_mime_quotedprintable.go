//go:build go1.21 && !go1.22
// +build go1.21,!go1.22

package stdlib

import (
	"mime/quotedprintable"
	"reflect"
)

func init() {
	Symbols["mime/quotedprintable/quotedprintable"] = map[string]reflect.Value{

		"NewReader": reflect.ValueOf(quotedprintable.NewReader),
		"NewWriter": reflect.ValueOf(quotedprintable.NewWriter),

		"Reader": reflect.ValueOf((*quotedprintable.Reader)(nil)),
		"Writer": reflect.ValueOf((*quotedprintable.Writer)(nil)),
	}
}
