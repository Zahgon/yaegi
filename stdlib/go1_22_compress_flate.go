//go:build go1.22
// +build go1.22

package stdlib

import (
	"compress/flate"
	"go/constant"
	"go/token"
	"io"
	"reflect"
)

func init() {
	Symbols["compress/flate/flate"] = map[string]reflect.Value{

		"BestCompression":    reflect.ValueOf(constant.MakeFromLiteral("9", token.INT, 0)),
		"BestSpeed":          reflect.ValueOf(constant.MakeFromLiteral("1", token.INT, 0)),
		"DefaultCompression": reflect.ValueOf(constant.MakeFromLiteral("-1", token.INT, 0)),
		"HuffmanOnly":        reflect.ValueOf(constant.MakeFromLiteral("-2", token.INT, 0)),
		"NewReader":          reflect.ValueOf(flate.NewReader),
		"NewReaderDict":      reflect.ValueOf(flate.NewReaderDict),
		"NewWriter":          reflect.ValueOf(flate.NewWriter),
		"NewWriterDict":      reflect.ValueOf(flate.NewWriterDict),
		"NoCompression":      reflect.ValueOf(constant.MakeFromLiteral("0", token.INT, 0)),

		"CorruptInputError": reflect.ValueOf((*flate.CorruptInputError)(nil)),
		"InternalError":     reflect.ValueOf((*flate.InternalError)(nil)),
		"ReadError":         reflect.ValueOf((*flate.ReadError)(nil)),
		"Reader":            reflect.ValueOf((*flate.Reader)(nil)),
		"Resetter":          reflect.ValueOf((*flate.Resetter)(nil)),
		"WriteError":        reflect.ValueOf((*flate.WriteError)(nil)),
		"Writer":            reflect.ValueOf((*flate.Writer)(nil)),

		"_Reader":   reflect.ValueOf((*_compress_flate_Reader)(nil)),
		"_Resetter": reflect.ValueOf((*_compress_flate_Resetter)(nil)),
	}
}

type _compress_flate_Reader struct {
	IValue    interface{}
	WRead     func(p []byte) (n int, err error)
	WReadByte func() (byte, error)
}

func (W _compress_flate_Reader) Read(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}
func (W _compress_flate_Reader) ReadByte() (byte, error) { _ = "STUB: not implemented"; return 0, nil }

type _compress_flate_Resetter struct {
	IValue interface{}
	WReset func(r io.Reader, dict []byte) error
}

func (W _compress_flate_Resetter) Reset(r io.Reader, dict []byte) error {
	_ = "STUB: not implemented"
	return nil
}
