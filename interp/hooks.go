package interp

import "reflect"

type convertFn func(from, to reflect.Type) func(src, dest reflect.Value)

type hooks struct {
	convert []convertFn
}

func (h *hooks) Parse(m map[string]reflect.Value) { _ = "STUB: not implemented"; return }

func getConvertFn(v reflect.Value) (convertFn, bool) {
	_ = "STUB: not implemented"
	return *new(convertFn), false
}
