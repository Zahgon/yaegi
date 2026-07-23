package interp

import (
	"errors"
	"go/constant"
	"reflect"
)

type opPredicates map[action]func(reflect.Type) bool

type typecheck struct {
	scope *scope
}

func (check typecheck) op(p opPredicates, a action, n, c *node, t reflect.Type) error {
	_ = "STUB: not implemented"
	return nil
}

func (check typecheck) assignment(n *node, typ *itype, context string) error {
	_ = "STUB: not implemented"
	return nil
}

func (check typecheck) assignExpr(n, dest, src *node) error { _ = "STUB: not implemented"; return nil }

func (check typecheck) addressExpr(n *node) error { _ = "STUB: not implemented"; return nil }

func (check typecheck) starExpr(n *node) error { _ = "STUB: not implemented"; return nil }

var unaryOpPredicates = opPredicates{
	aInc:    isNumber,
	aDec:    isNumber,
	aPos:    isNumber,
	aNeg:    isNumber,
	aBitNot: isInt,
	aNot:    isBoolean,
}

func (check typecheck) unaryExpr(n *node) error { _ = "STUB: not implemented"; return nil }

func (check typecheck) shift(n *node) error { _ = "STUB: not implemented"; return nil }

func (check typecheck) comparison(n *node) error { _ = "STUB: not implemented"; return nil }

var binaryOpPredicates = opPredicates{
	aAdd: func(typ reflect.Type) bool { return isNumber(typ) || isString(typ) },
	aSub: isNumber,
	aMul: isNumber,
	aQuo: isNumber,
	aRem: isInt,

	aAnd:    isInt,
	aOr:     isInt,
	aXor:    isInt,
	aAndNot: isInt,

	aLand: isBoolean,
	aLor:  isBoolean,
}

func (check typecheck) binaryExpr(n *node) error { _ = "STUB: not implemented"; return nil }

func zeroConst(n *node) bool { _ = "STUB: not implemented"; return false }

func (check typecheck) index(n *node, max int) error { _ = "STUB: not implemented"; return nil }

func (check typecheck) arrayLitExpr(child []*node, typ *itype) error {
	_ = "STUB: not implemented"
	return nil
}

func (check typecheck) mapLitExpr(child []*node, ktyp, vtyp *itype) error {
	_ = "STUB: not implemented"
	return nil
}

func (check typecheck) structLitExpr(child []*node, typ *itype) error {
	_ = "STUB: not implemented"
	return nil
}

func (check typecheck) structBinLitExpr(child []*node, typ reflect.Type) error {
	_ = "STUB: not implemented"
	return nil
}

func (check typecheck) sliceExpr(n *node) error { _ = "STUB: not implemented"; return nil }

func (check typecheck) typeAssertionExpr(n *node, typ *itype) error {
	_ = "STUB: not implemented"
	return nil
}

func (check typecheck) conversion(n *node, typ *itype) error { _ = "STUB: not implemented"; return nil }

type param struct {
	nod *node
	typ *itype
}

func (p param) Type() *itype { _ = "STUB: not implemented"; return nil }

func (check typecheck) unpackParams(child []*node) (params []param) {
	_ = "STUB: not implemented"
	return nil
}

var builtinFuncs = map[string]struct {
	args     int
	variadic bool
}{
	bltnAlignof:  {args: 1, variadic: false},
	bltnAppend:   {args: 1, variadic: true},
	bltnCap:      {args: 1, variadic: false},
	bltnClose:    {args: 1, variadic: false},
	bltnComplex:  {args: 2, variadic: false},
	bltnImag:     {args: 1, variadic: false},
	bltnCopy:     {args: 2, variadic: false},
	bltnDelete:   {args: 2, variadic: false},
	bltnLen:      {args: 1, variadic: false},
	bltnMake:     {args: 1, variadic: true},
	bltnNew:      {args: 1, variadic: false},
	bltnOffsetof: {args: 1, variadic: false},
	bltnPanic:    {args: 1, variadic: false},
	bltnPrint:    {args: 0, variadic: true},
	bltnPrintln:  {args: 0, variadic: true},
	bltnReal:     {args: 1, variadic: false},
	bltnRecover:  {args: 0, variadic: false},
	bltnSizeof:   {args: 1, variadic: false},
}

func (check typecheck) builtin(name string, n *node, child []*node, ellipsis bool) error {
	_ = "STUB: not implemented"
	return nil
}

func arrayDeref(typ *itype) *itype { _ = "STUB: not implemented"; return nil }

func (check typecheck) arguments(n *node, child []*node, fun *node, ellipsis bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (check typecheck) argument(p param, ftyp *itype, i, l int, ellipsis bool) error {
	_ = "STUB: not implemented"
	return nil
}

func getArg(ftyp *itype, i int) *itype { _ = "STUB: not implemented"; return nil }

func getArgsID(ftyp *itype) string { _ = "STUB: not implemented"; return "" }

var errCantConvert = errors.New("cannot convert")

func (check typecheck) convertUntyped(n *node, typ *itype) error {
	_ = "STUB: not implemented"
	return nil
}

func (check typecheck) representable(n *node, t reflect.Type) error {
	_ = "STUB: not implemented"
	return nil
}

func (check typecheck) convertConst(v reflect.Value, t reflect.Type) (reflect.Value, error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}

var bitlen = [...]int{
	reflect.Int:     64,
	reflect.Int8:    8,
	reflect.Int16:   16,
	reflect.Int32:   32,
	reflect.Int64:   64,
	reflect.Uint:    64,
	reflect.Uint8:   8,
	reflect.Uint16:  16,
	reflect.Uint32:  32,
	reflect.Uint64:  64,
	reflect.Uintptr: 64,
}

func representableConst(c constant.Value, t reflect.Type) bool {
	_ = "STUB: not implemented"
	return false
}

func isShiftAction(a action) bool { _ = "STUB: not implemented"; return false }

func isComparisonAction(a action) bool { _ = "STUB: not implemented"; return false }
