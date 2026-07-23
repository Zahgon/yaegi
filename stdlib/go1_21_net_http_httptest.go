//go:build go1.21 && !go1.22
// +build go1.21,!go1.22

package stdlib

import (
	"go/constant"
	"go/token"
	"net/http/httptest"
	"reflect"
)

func init() {
	Symbols["net/http/httptest/httptest"] = map[string]reflect.Value{

		"DefaultRemoteAddr":  reflect.ValueOf(constant.MakeFromLiteral("\"1.2.3.4\"", token.STRING, 0)),
		"NewRecorder":        reflect.ValueOf(httptest.NewRecorder),
		"NewRequest":         reflect.ValueOf(httptest.NewRequest),
		"NewServer":          reflect.ValueOf(httptest.NewServer),
		"NewTLSServer":       reflect.ValueOf(httptest.NewTLSServer),
		"NewUnstartedServer": reflect.ValueOf(httptest.NewUnstartedServer),

		"ResponseRecorder": reflect.ValueOf((*httptest.ResponseRecorder)(nil)),
		"Server":           reflect.ValueOf((*httptest.Server)(nil)),
	}
}
