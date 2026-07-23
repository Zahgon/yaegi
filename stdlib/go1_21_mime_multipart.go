//go:build go1.21 && !go1.22
// +build go1.21,!go1.22

package stdlib

import (
	"mime/multipart"
	"reflect"
)

func init() {
	Symbols["mime/multipart/multipart"] = map[string]reflect.Value{

		"ErrMessageTooLarge": reflect.ValueOf(&multipart.ErrMessageTooLarge).Elem(),
		"NewReader":          reflect.ValueOf(multipart.NewReader),
		"NewWriter":          reflect.ValueOf(multipart.NewWriter),

		"File":       reflect.ValueOf((*multipart.File)(nil)),
		"FileHeader": reflect.ValueOf((*multipart.FileHeader)(nil)),
		"Form":       reflect.ValueOf((*multipart.Form)(nil)),
		"Part":       reflect.ValueOf((*multipart.Part)(nil)),
		"Reader":     reflect.ValueOf((*multipart.Reader)(nil)),
		"Writer":     reflect.ValueOf((*multipart.Writer)(nil)),

		"_File": reflect.ValueOf((*_mime_multipart_File)(nil)),
	}
}

type _mime_multipart_File struct {
	IValue  interface{}
	WClose  func() error
	WRead   func(p []byte) (n int, err error)
	WReadAt func(p []byte, off int64) (n int, err error)
	WSeek   func(offset int64, whence int) (int64, error)
}

func (W _mime_multipart_File) Close() error { _ = "STUB: not implemented"; return nil }
func (W _mime_multipart_File) Read(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}
func (W _mime_multipart_File) ReadAt(p []byte, off int64) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (W _mime_multipart_File) Seek(offset int64, whence int) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
