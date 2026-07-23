//go:build go1.21 && !go1.22
// +build go1.21,!go1.22

package stdlib

import (
	"net/http/httptrace"
	"reflect"
)

func init() {
	Symbols["net/http/httptrace/httptrace"] = map[string]reflect.Value{

		"ContextClientTrace": reflect.ValueOf(httptrace.ContextClientTrace),
		"WithClientTrace":    reflect.ValueOf(httptrace.WithClientTrace),

		"ClientTrace":      reflect.ValueOf((*httptrace.ClientTrace)(nil)),
		"DNSDoneInfo":      reflect.ValueOf((*httptrace.DNSDoneInfo)(nil)),
		"DNSStartInfo":     reflect.ValueOf((*httptrace.DNSStartInfo)(nil)),
		"GotConnInfo":      reflect.ValueOf((*httptrace.GotConnInfo)(nil)),
		"WroteRequestInfo": reflect.ValueOf((*httptrace.WroteRequestInfo)(nil)),
	}
}
