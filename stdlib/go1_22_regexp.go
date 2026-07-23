//go:build go1.22
// +build go1.22

package stdlib

import (
	"reflect"
	"regexp"
)

func init() {
	Symbols["regexp/regexp"] = map[string]reflect.Value{

		"Compile":          reflect.ValueOf(regexp.Compile),
		"CompilePOSIX":     reflect.ValueOf(regexp.CompilePOSIX),
		"Match":            reflect.ValueOf(regexp.Match),
		"MatchReader":      reflect.ValueOf(regexp.MatchReader),
		"MatchString":      reflect.ValueOf(regexp.MatchString),
		"MustCompile":      reflect.ValueOf(regexp.MustCompile),
		"MustCompilePOSIX": reflect.ValueOf(regexp.MustCompilePOSIX),
		"QuoteMeta":        reflect.ValueOf(regexp.QuoteMeta),

		"Regexp": reflect.ValueOf((*regexp.Regexp)(nil)),
	}
}
