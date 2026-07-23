//go:build go1.22
// +build go1.22

package stdlib

import (
	"go/printer"
	"reflect"
)

func init() {
	Symbols["go/printer/printer"] = map[string]reflect.Value{

		"Fprint":    reflect.ValueOf(printer.Fprint),
		"RawFormat": reflect.ValueOf(printer.RawFormat),
		"SourcePos": reflect.ValueOf(printer.SourcePos),
		"TabIndent": reflect.ValueOf(printer.TabIndent),
		"UseSpaces": reflect.ValueOf(printer.UseSpaces),

		"CommentedNode": reflect.ValueOf((*printer.CommentedNode)(nil)),
		"Config":        reflect.ValueOf((*printer.Config)(nil)),
		"Mode":          reflect.ValueOf((*printer.Mode)(nil)),
	}
}
