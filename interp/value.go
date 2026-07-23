package interp

import (
	"go/constant"
	"reflect"
)

const (
	notInFrame  = -1
	globalFrame = -1
)

func valueGenerator(n *node, i int) func(*frame) reflect.Value {
	_ = "STUB: not implemented"
	return nil
}

func valueOf(data []reflect.Value, i int) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

func genValueRecv(n *node) func(*frame) reflect.Value { _ = "STUB: not implemented"; return nil }

func genValueAsFunctionWrapper(n *node) func(*frame) reflect.Value {
	_ = "STUB: not implemented"
	return nil
}

func genValueAs(n *node, t reflect.Type) func(*frame) reflect.Value {
	_ = "STUB: not implemented"
	return nil
}

func genValue(n *node) func(*frame) reflect.Value { _ = "STUB: not implemented"; return nil }

func genDestValue(typ *itype, n *node) func(*frame) reflect.Value {
	_ = "STUB: not implemented"
	return nil
}

func genFuncValue(n *node) func(*frame) reflect.Value { _ = "STUB: not implemented"; return nil }

func genValueArray(n *node) func(*frame) reflect.Value { _ = "STUB: not implemented"; return nil }

func genValueRangeArray(n *node) func(*frame) reflect.Value { _ = "STUB: not implemented"; return nil }

func genValueInterface(n *node) func(*frame) reflect.Value { _ = "STUB: not implemented"; return nil }

func getConcreteValue(val reflect.Value) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

func zeroInterfaceValue() reflect.Value { _ = "STUB: not implemented"; return *new(reflect.Value) }

func wantEmptyInterface(n *node) bool { _ = "STUB: not implemented"; return false }

func genValueOutput(n *node, t reflect.Type) func(*frame) reflect.Value {
	_ = "STUB: not implemented"
	return nil
}

func getBinValue(getMapType func(*itype) reflect.Type, value func(*frame) reflect.Value, f *frame) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

func valueInterfaceValue(v reflect.Value) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

func genValueInterfaceValue(n *node) func(*frame) reflect.Value {
	_ = "STUB: not implemented"
	return nil
}

func vInt(v reflect.Value) (i int64) { _ = "STUB: not implemented"; return 0 }

func vUint(v reflect.Value) (i uint64) { _ = "STUB: not implemented"; return 0 }

func vComplex(v reflect.Value) (c complex128) { _ = "STUB: not implemented"; return 0 }

func vFloat(v reflect.Value) (i float64) { _ = "STUB: not implemented"; return 0 }

func vString(v reflect.Value) (s string) { _ = "STUB: not implemented"; return "" }

func vConstantValue(v reflect.Value) (c constant.Value) {
	_ = "STUB: not implemented"
	return *new(constant.Value)
}

func genValueInt(n *node) func(*frame) (reflect.Value, int64) {
	_ = "STUB: not implemented"
	return nil
}

func genValueUint(n *node) func(*frame) (reflect.Value, uint64) {
	_ = "STUB: not implemented"
	return nil
}

func genValueFloat(n *node) func(*frame) (reflect.Value, float64) {
	_ = "STUB: not implemented"
	return nil
}

func genValueComplex(n *node) func(*frame) reflect.Value { _ = "STUB: not implemented"; return nil }

func genComplex(n *node) func(*frame) complex128 { _ = "STUB: not implemented"; return nil }

func genValueString(n *node) func(*frame) (reflect.Value, string) {
	_ = "STUB: not implemented"
	return nil
}
