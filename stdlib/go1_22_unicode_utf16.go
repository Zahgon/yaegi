//go:build go1.22
// +build go1.22

package stdlib

import (
	"reflect"
	"unicode/utf16"
)

func init() {
	Symbols["unicode/utf16/utf16"] = map[string]reflect.Value{

		"AppendRune":  reflect.ValueOf(utf16.AppendRune),
		"Decode":      reflect.ValueOf(utf16.Decode),
		"DecodeRune":  reflect.ValueOf(utf16.DecodeRune),
		"Encode":      reflect.ValueOf(utf16.Encode),
		"EncodeRune":  reflect.ValueOf(utf16.EncodeRune),
		"IsSurrogate": reflect.ValueOf(utf16.IsSurrogate),
	}
}
