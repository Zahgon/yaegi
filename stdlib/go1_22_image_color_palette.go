//go:build go1.22
// +build go1.22

package stdlib

import (
	"image/color/palette"
	"reflect"
)

func init() {
	Symbols["image/color/palette/palette"] = map[string]reflect.Value{

		"Plan9":   reflect.ValueOf(&palette.Plan9).Elem(),
		"WebSafe": reflect.ValueOf(&palette.WebSafe).Elem(),
	}
}
