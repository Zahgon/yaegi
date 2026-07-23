package main

import (
	"bytes"
	"encoding/json"
	"net/netip"
	"reflect"
)

func unmarshalJSON[T any](b []byte, x *[]T) error { _ = "STUB: not implemented"; return nil }

func SliceOfViews[T ViewCloner[T, V], V StructView[T]](x []T) SliceView[T, V] {
	_ = "STUB: not implemented"
	return nil
}

type StructView[T any] interface {
	Valid() bool
	AsStruct() T
}

type SliceView[T ViewCloner[T, V], V StructView[T]] struct {
	ж []T
}

type ViewCloner[T any, V StructView[T]] interface {
	View() V
	Clone() T
}

func (v SliceView[T, V]) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (v *SliceView[T, V]) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

type Slice[T any] struct {
	ж []T
}

func (v Slice[T]) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (v *Slice[T]) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

func SliceOf[T any](x []T) Slice[T] { _ = "STUB: not implemented"; return nil }

type IPPrefixSlice struct {
	ж Slice[netip.Prefix]
}

type viewStruct struct {
	Int        int
	Strings    Slice[string]
	StringsPtr *Slice[string] `json:",omitempty"`
}

func main() {
	ss := SliceOf([]string{"bar"})
	in := viewStruct{
		Int:        1234,
		Strings:    ss,
		StringsPtr: &ss,
	}

	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetIndent("", "")
	err1 := encoder.Encode(&in)
	b := buf.Bytes()
	var got viewStruct
	err2 := json.Unmarshal(b, &got)
	println(err1 == nil, err2 == nil, reflect.DeepEqual(got, in))
}
