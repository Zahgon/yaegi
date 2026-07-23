//go:build go1.22
// +build go1.22

package stdlib

import (
	"go/scanner"
	"reflect"
)

func init() {
	Symbols["go/scanner/scanner"] = map[string]reflect.Value{

		"PrintError":   reflect.ValueOf(scanner.PrintError),
		"ScanComments": reflect.ValueOf(scanner.ScanComments),

		"Error":        reflect.ValueOf((*scanner.Error)(nil)),
		"ErrorHandler": reflect.ValueOf((*scanner.ErrorHandler)(nil)),
		"ErrorList":    reflect.ValueOf((*scanner.ErrorList)(nil)),
		"Mode":         reflect.ValueOf((*scanner.Mode)(nil)),
		"Scanner":      reflect.ValueOf((*scanner.Scanner)(nil)),
	}
}
