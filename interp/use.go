package interp

import (
	"reflect"
)

func (interp *Interpreter) Symbols(importPath string) Exports {
	_ = "STUB: not implemented"
	return *new(Exports)
}

func getWrapper(n *node, t reflect.Type) reflect.Type {
	_ = "STUB: not implemented"
	return *new(reflect.Type)
}

func (interp *Interpreter) Use(values Exports) error { _ = "STUB: not implemented"; return nil }

func fixStdlib(interp *Interpreter) { _ = "STUB: not implemented"; return }
