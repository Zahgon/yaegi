//go:build go1.22
// +build go1.22

package stdlib

import (
	"debug/buildinfo"
	"reflect"
)

func init() {
	Symbols["debug/buildinfo/buildinfo"] = map[string]reflect.Value{

		"Read":     reflect.ValueOf(buildinfo.Read),
		"ReadFile": reflect.ValueOf(buildinfo.ReadFile),

		"BuildInfo": reflect.ValueOf((*buildinfo.BuildInfo)(nil)),
	}
}
