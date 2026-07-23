//go:build go1.21 && !go1.22
// +build go1.21,!go1.22

package stdlib

import (
	"net/http/cookiejar"
	"reflect"
)

func init() {
	Symbols["net/http/cookiejar/cookiejar"] = map[string]reflect.Value{

		"New": reflect.ValueOf(cookiejar.New),

		"Jar":              reflect.ValueOf((*cookiejar.Jar)(nil)),
		"Options":          reflect.ValueOf((*cookiejar.Options)(nil)),
		"PublicSuffixList": reflect.ValueOf((*cookiejar.PublicSuffixList)(nil)),

		"_PublicSuffixList": reflect.ValueOf((*_net_http_cookiejar_PublicSuffixList)(nil)),
	}
}

type _net_http_cookiejar_PublicSuffixList struct {
	IValue        interface{}
	WPublicSuffix func(domain string) string
	WString       func() string
}

func (W _net_http_cookiejar_PublicSuffixList) PublicSuffix(domain string) string {
	_ = "STUB: not implemented"
	return ""
}

func (W _net_http_cookiejar_PublicSuffixList) String() string { _ = "STUB: not implemented"; return "" }
