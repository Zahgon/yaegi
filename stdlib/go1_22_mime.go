//go:build go1.22
// +build go1.22

package stdlib

import (
	"mime"
	"reflect"
)

func init() {
	Symbols["mime/mime"] = map[string]reflect.Value{

		"AddExtensionType":         reflect.ValueOf(mime.AddExtensionType),
		"BEncoding":                reflect.ValueOf(mime.BEncoding),
		"ErrInvalidMediaParameter": reflect.ValueOf(&mime.ErrInvalidMediaParameter).Elem(),
		"ExtensionsByType":         reflect.ValueOf(mime.ExtensionsByType),
		"FormatMediaType":          reflect.ValueOf(mime.FormatMediaType),
		"ParseMediaType":           reflect.ValueOf(mime.ParseMediaType),
		"QEncoding":                reflect.ValueOf(mime.QEncoding),
		"TypeByExtension":          reflect.ValueOf(mime.TypeByExtension),

		"WordDecoder": reflect.ValueOf((*mime.WordDecoder)(nil)),
		"WordEncoder": reflect.ValueOf((*mime.WordEncoder)(nil)),
	}
}
