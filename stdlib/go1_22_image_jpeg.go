//go:build go1.22
// +build go1.22

package stdlib

import (
	"go/constant"
	"go/token"
	"image/jpeg"
	"reflect"
)

func init() {
	Symbols["image/jpeg/jpeg"] = map[string]reflect.Value{

		"Decode":         reflect.ValueOf(jpeg.Decode),
		"DecodeConfig":   reflect.ValueOf(jpeg.DecodeConfig),
		"DefaultQuality": reflect.ValueOf(constant.MakeFromLiteral("75", token.INT, 0)),
		"Encode":         reflect.ValueOf(jpeg.Encode),

		"FormatError":      reflect.ValueOf((*jpeg.FormatError)(nil)),
		"Options":          reflect.ValueOf((*jpeg.Options)(nil)),
		"Reader":           reflect.ValueOf((*jpeg.Reader)(nil)),
		"UnsupportedError": reflect.ValueOf((*jpeg.UnsupportedError)(nil)),

		"_Reader": reflect.ValueOf((*_image_jpeg_Reader)(nil)),
	}
}

type _image_jpeg_Reader struct {
	IValue    interface{}
	WRead     func(p []byte) (n int, err error)
	WReadByte func() (byte, error)
}

func (W _image_jpeg_Reader) Read(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}
func (W _image_jpeg_Reader) ReadByte() (byte, error) { _ = "STUB: not implemented"; return 0, nil }
