//go:build go1.22
// +build go1.22

package stdlib

import (
	"net/http/httputil"
	"reflect"
)

func init() {
	Symbols["net/http/httputil/httputil"] = map[string]reflect.Value{

		"DumpRequest":               reflect.ValueOf(httputil.DumpRequest),
		"DumpRequestOut":            reflect.ValueOf(httputil.DumpRequestOut),
		"DumpResponse":              reflect.ValueOf(httputil.DumpResponse),
		"ErrClosed":                 reflect.ValueOf(&httputil.ErrClosed).Elem(),
		"ErrLineTooLong":            reflect.ValueOf(&httputil.ErrLineTooLong).Elem(),
		"ErrPersistEOF":             reflect.ValueOf(&httputil.ErrPersistEOF).Elem(),
		"ErrPipeline":               reflect.ValueOf(&httputil.ErrPipeline).Elem(),
		"NewChunkedReader":          reflect.ValueOf(httputil.NewChunkedReader),
		"NewChunkedWriter":          reflect.ValueOf(httputil.NewChunkedWriter),
		"NewClientConn":             reflect.ValueOf(httputil.NewClientConn),
		"NewProxyClientConn":        reflect.ValueOf(httputil.NewProxyClientConn),
		"NewServerConn":             reflect.ValueOf(httputil.NewServerConn),
		"NewSingleHostReverseProxy": reflect.ValueOf(httputil.NewSingleHostReverseProxy),

		"BufferPool":   reflect.ValueOf((*httputil.BufferPool)(nil)),
		"ClientConn":   reflect.ValueOf((*httputil.ClientConn)(nil)),
		"ProxyRequest": reflect.ValueOf((*httputil.ProxyRequest)(nil)),
		"ReverseProxy": reflect.ValueOf((*httputil.ReverseProxy)(nil)),
		"ServerConn":   reflect.ValueOf((*httputil.ServerConn)(nil)),

		"_BufferPool": reflect.ValueOf((*_net_http_httputil_BufferPool)(nil)),
	}
}

type _net_http_httputil_BufferPool struct {
	IValue interface{}
	WGet   func() []byte
	WPut   func(a0 []byte)
}

func (W _net_http_httputil_BufferPool) Get() []byte   { _ = "STUB: not implemented"; return nil }
func (W _net_http_httputil_BufferPool) Put(a0 []byte) { _ = "STUB: not implemented"; return }
