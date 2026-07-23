//go:build go1.22
// +build go1.22

package stdlib

import (
	"encoding/binary"
	"go/constant"
	"go/token"
	"reflect"
)

func init() {
	Symbols["encoding/binary/binary"] = map[string]reflect.Value{

		"AppendUvarint":  reflect.ValueOf(binary.AppendUvarint),
		"AppendVarint":   reflect.ValueOf(binary.AppendVarint),
		"BigEndian":      reflect.ValueOf(&binary.BigEndian).Elem(),
		"LittleEndian":   reflect.ValueOf(&binary.LittleEndian).Elem(),
		"MaxVarintLen16": reflect.ValueOf(constant.MakeFromLiteral("3", token.INT, 0)),
		"MaxVarintLen32": reflect.ValueOf(constant.MakeFromLiteral("5", token.INT, 0)),
		"MaxVarintLen64": reflect.ValueOf(constant.MakeFromLiteral("10", token.INT, 0)),
		"NativeEndian":   reflect.ValueOf(&binary.NativeEndian).Elem(),
		"PutUvarint":     reflect.ValueOf(binary.PutUvarint),
		"PutVarint":      reflect.ValueOf(binary.PutVarint),
		"Read":           reflect.ValueOf(binary.Read),
		"ReadUvarint":    reflect.ValueOf(binary.ReadUvarint),
		"ReadVarint":     reflect.ValueOf(binary.ReadVarint),
		"Size":           reflect.ValueOf(binary.Size),
		"Uvarint":        reflect.ValueOf(binary.Uvarint),
		"Varint":         reflect.ValueOf(binary.Varint),
		"Write":          reflect.ValueOf(binary.Write),

		"AppendByteOrder": reflect.ValueOf((*binary.AppendByteOrder)(nil)),
		"ByteOrder":       reflect.ValueOf((*binary.ByteOrder)(nil)),

		"_AppendByteOrder": reflect.ValueOf((*_encoding_binary_AppendByteOrder)(nil)),
		"_ByteOrder":       reflect.ValueOf((*_encoding_binary_ByteOrder)(nil)),
	}
}

type _encoding_binary_AppendByteOrder struct {
	IValue        interface{}
	WAppendUint16 func(a0 []byte, a1 uint16) []byte
	WAppendUint32 func(a0 []byte, a1 uint32) []byte
	WAppendUint64 func(a0 []byte, a1 uint64) []byte
	WString       func() string
}

func (W _encoding_binary_AppendByteOrder) AppendUint16(a0 []byte, a1 uint16) []byte {
	_ = "STUB: not implemented"
	return nil
}

func (W _encoding_binary_AppendByteOrder) AppendUint32(a0 []byte, a1 uint32) []byte {
	_ = "STUB: not implemented"
	return nil
}

func (W _encoding_binary_AppendByteOrder) AppendUint64(a0 []byte, a1 uint64) []byte {
	_ = "STUB: not implemented"
	return nil
}

func (W _encoding_binary_AppendByteOrder) String() string { _ = "STUB: not implemented"; return "" }

type _encoding_binary_ByteOrder struct {
	IValue     interface{}
	WPutUint16 func(a0 []byte, a1 uint16)
	WPutUint32 func(a0 []byte, a1 uint32)
	WPutUint64 func(a0 []byte, a1 uint64)
	WString    func() string
	WUint16    func(a0 []byte) uint16
	WUint32    func(a0 []byte) uint32
	WUint64    func(a0 []byte) uint64
}

func (W _encoding_binary_ByteOrder) PutUint16(a0 []byte, a1 uint16) {
	_ = "STUB: not implemented"
	return
}
func (W _encoding_binary_ByteOrder) PutUint32(a0 []byte, a1 uint32) {
	_ = "STUB: not implemented"
	return
}
func (W _encoding_binary_ByteOrder) PutUint64(a0 []byte, a1 uint64) {
	_ = "STUB: not implemented"
	return
}
func (W _encoding_binary_ByteOrder) String() string { _ = "STUB: not implemented"; return "" }

func (W _encoding_binary_ByteOrder) Uint16(a0 []byte) uint16 { _ = "STUB: not implemented"; return 0 }
func (W _encoding_binary_ByteOrder) Uint32(a0 []byte) uint32 { _ = "STUB: not implemented"; return 0 }
func (W _encoding_binary_ByteOrder) Uint64(a0 []byte) uint64 { _ = "STUB: not implemented"; return 0 }
