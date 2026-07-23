package interp

import (
	"reflect"
)

var trace = false

func traceIndent(n *node) string { _ = "STUB: not implemented"; return "" }

func tracePrintln(n *node, v ...any) { _ = "STUB: not implemented"; return }

//nolint:unused // debugging facility
func tracePrintTree(n *node, v ...any) { _ = "STUB: not implemented"; return }

func ptrAddr(v any) string { _ = "STUB: not implemented"; return "" }

//nolint:unused // debugging facility
func valString(v reflect.Value) string { _ = "STUB: not implemented"; return "" }

func (n *node) String() string { _ = "STUB: not implemented"; return "" }

func (n *node) depth() int { _ = "STUB: not implemented"; return 0 }

func (sy *symbol) String() string { _ = "STUB: not implemented"; return "" }

func (t *itype) String() string { _ = "STUB: not implemented"; return "" }
