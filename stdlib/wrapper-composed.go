package stdlib

import (
	"bufio"
	"io"
	"net"
	"net/http"
	"reflect"
)

type _netHTTPResponseWriterHijacker struct {
	IValue       interface{}
	WHeader      func() http.Header
	WWrite       func(a0 []byte) (int, error)
	WWriteHeader func(statusCode int)

	WHijack func() (net.Conn, *bufio.ReadWriter, error)
}

func (w _netHTTPResponseWriterHijacker) Header() http.Header {
	_ = "STUB: not implemented"
	return *new(http.Header)
}

func (w _netHTTPResponseWriterHijacker) Write(a0 []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (w _netHTTPResponseWriterHijacker) WriteHeader(statusCode int) {
	_ = "STUB: not implemented"
	return
}

func (w _netHTTPResponseWriterHijacker) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil, nil
}

type _ioReaderWriteTo struct {
	IValue interface{}
	WRead  func(p []byte) (n int, err error)

	WWriteTo func(w io.Writer) (n int64, err error)
}

func (w _ioReaderWriteTo) Read(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (w _ioReaderWriteTo) WriteTo(wr io.Writer) (n int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

type _ioWriterReadFrom struct {
	IValue interface{}
	WWrite func(p []byte) (n int, err error)

	WReadFrom func(r io.Reader) (n int64, err error)
}

func (w _ioWriterReadFrom) Write(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (w _ioWriterReadFrom) ReadFrom(r io.Reader) (n int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func init() {
	MapTypes[reflect.ValueOf((*_net_http_ResponseWriter)(nil))] = []reflect.Type{
		reflect.ValueOf((*_netHTTPResponseWriterHijacker)(nil)).Type().Elem(),
	}
	MapTypes[reflect.ValueOf((*_io_Reader)(nil))] = []reflect.Type{
		reflect.ValueOf((*_ioReaderWriteTo)(nil)).Type().Elem(),
	}
	MapTypes[reflect.ValueOf((*_io_Writer)(nil))] = []reflect.Type{
		reflect.ValueOf((*_ioWriterReadFrom)(nil)).Type().Elem(),
	}
}
