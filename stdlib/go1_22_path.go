//go:build go1.22
// +build go1.22

package stdlib

import (
	"path"
	"reflect"
)

func init() {
	Symbols["path/path"] = map[string]reflect.Value{

		"Base":          reflect.ValueOf(path.Base),
		"Clean":         reflect.ValueOf(path.Clean),
		"Dir":           reflect.ValueOf(path.Dir),
		"ErrBadPattern": reflect.ValueOf(&path.ErrBadPattern).Elem(),
		"Ext":           reflect.ValueOf(path.Ext),
		"IsAbs":         reflect.ValueOf(path.IsAbs),
		"Join":          reflect.ValueOf(path.Join),
		"Match":         reflect.ValueOf(path.Match),
		"Split":         reflect.ValueOf(path.Split),
	}
}
