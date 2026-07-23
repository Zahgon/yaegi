//go:build go1.21 && !go1.22
// +build go1.21,!go1.22

package stdlib

import (
	"image/png"
	"reflect"
)

func init() {
	Symbols["image/png/png"] = map[string]reflect.Value{

		"BestCompression":    reflect.ValueOf(png.BestCompression),
		"BestSpeed":          reflect.ValueOf(png.BestSpeed),
		"Decode":             reflect.ValueOf(png.Decode),
		"DecodeConfig":       reflect.ValueOf(png.DecodeConfig),
		"DefaultCompression": reflect.ValueOf(png.DefaultCompression),
		"Encode":             reflect.ValueOf(png.Encode),
		"NoCompression":      reflect.ValueOf(png.NoCompression),

		"CompressionLevel":  reflect.ValueOf((*png.CompressionLevel)(nil)),
		"Encoder":           reflect.ValueOf((*png.Encoder)(nil)),
		"EncoderBuffer":     reflect.ValueOf((*png.EncoderBuffer)(nil)),
		"EncoderBufferPool": reflect.ValueOf((*png.EncoderBufferPool)(nil)),
		"FormatError":       reflect.ValueOf((*png.FormatError)(nil)),
		"UnsupportedError":  reflect.ValueOf((*png.UnsupportedError)(nil)),

		"_EncoderBufferPool": reflect.ValueOf((*_image_png_EncoderBufferPool)(nil)),
	}
}

type _image_png_EncoderBufferPool struct {
	IValue interface{}
	WGet   func() *png.EncoderBuffer
	WPut   func(a0 *png.EncoderBuffer)
}

func (W _image_png_EncoderBufferPool) Get() *png.EncoderBuffer {
	_ = "STUB: not implemented"
	return nil
}
func (W _image_png_EncoderBufferPool) Put(a0 *png.EncoderBuffer) { _ = "STUB: not implemented"; return }
