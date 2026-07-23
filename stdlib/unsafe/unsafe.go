package unsafe

import (
	"reflect"
	"unsafe"
)

var Symbols = map[string]map[string]reflect.Value{}

func init() {
	Symbols["github.com/traefik/yaegi/stdlib/unsafe/unsafe"] = map[string]reflect.Value{
		"Symbols": reflect.ValueOf(Symbols),
	}
	Symbols["github.com/traefik/yaegi/yaegi"] = map[string]reflect.Value{
		"convert": reflect.ValueOf(convert),
	}

	Symbols["unsafe/unsafe"]["Add"] = reflect.ValueOf(add)

	Symbols["unsafe/unsafe"]["Sizeof"] = reflect.ValueOf(sizeof)
	Symbols["unsafe/unsafe"]["Alignof"] = reflect.ValueOf(alignof)

	Symbols["unsafe/unsafe"]["Offsetof"] = reflect.ValueOf(func(interface{}) uintptr { return 0 })
}

func convert(from, to reflect.Type) func(src, dest reflect.Value) {
	_ = "STUB: not implemented"
	return nil
}

func add(ptr unsafe.Pointer, l int) unsafe.Pointer {
	_ = "STUB: not implemented"
	return *new(unsafe.Pointer)
}

func sizeof(i interface{}) uintptr { _ = "STUB: not implemented"; return 0 }

func alignof(i interface{}) uintptr { _ = "STUB: not implemented"; return 0 }

//go:nocheckptr
func uintptrToUnsafePtr(src, dest reflect.Value) { _ = "STUB: not implemented"; return }

//nolint:govet

//go:generate ../../internal/cmd/extract/extract unsafe
