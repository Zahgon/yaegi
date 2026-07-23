package interp

//go:generate go run ../internal/cmd/genop/genop.go

import (
	"errors"
	"reflect"
	"regexp"
)

type bltn func(f *frame) bltn

type bltnGenerator func(n *node)

var builtin = [...]bltnGenerator{
	aNop:          nop,
	aAddr:         addr,
	aAssign:       assign,
	aAdd:          add,
	aAddAssign:    addAssign,
	aAnd:          and,
	aAndAssign:    andAssign,
	aAndNot:       andNot,
	aAndNotAssign: andNotAssign,
	aBitNot:       bitNot,
	aCall:         call,
	aCallSlice:    call,
	aCase:         _case,
	aCompositeLit: arrayLit,
	aDec:          dec,
	aEqual:        equal,
	aGetFunc:      getFunc,
	aGreater:      greater,
	aGreaterEqual: greaterEqual,
	aInc:          inc,
	aLand:         land,
	aLor:          lor,
	aLower:        lower,
	aLowerEqual:   lowerEqual,
	aMul:          mul,
	aMulAssign:    mulAssign,
	aNeg:          neg,
	aNot:          not,
	aNotEqual:     notEqual,
	aOr:           or,
	aOrAssign:     orAssign,
	aPos:          pos,
	aQuo:          quo,
	aQuoAssign:    quoAssign,
	aRange:        _range,
	aRecv:         recv,
	aRem:          rem,
	aRemAssign:    remAssign,
	aReturn:       _return,
	aSend:         send,
	aShl:          shl,
	aShlAssign:    shlAssign,
	aShr:          shr,
	aShrAssign:    shrAssign,
	aSlice:        slice,
	aSlice0:       slice0,
	aStar:         deref,
	aSub:          sub,
	aSubAssign:    subAssign,
	aTypeAssert:   typeAssertShort,
	aXor:          xor,
	aXorAssign:    xorAssign,
}

var receiverStripperRxp *regexp.Regexp

func init() {
	re := `func\(((.*?(, |\)))(.*))`
	var err error
	receiverStripperRxp, err = regexp.Compile(re)
	if err != nil {
		panic(err)
	}
}

type valueInterface struct {
	node  *node
	value reflect.Value
}

var floatType, complexType reflect.Type

func init() {
	floatType = reflect.ValueOf(0.0).Type()
	complexType = reflect.ValueOf(complex(0, 0)).Type()
}

func (interp *Interpreter) run(n *node, cf *frame) { _ = "STUB: not implemented"; return }

func isExecNode(n *node, exec bltn) bool { _ = "STUB: not implemented"; return false }

func originalExecNode(n *node, exec bltn) *node { _ = "STUB: not implemented"; return nil }

var errAbortHandler = errors.New("net/http: abort Handler")

func panicFunc(s *scope) string { _ = "STUB: not implemented"; return "" }

func runCfg(n *node, f *frame, funcNode, callNode *node) { _ = "STUB: not implemented"; return }

func stripReceiverFromArgs(signature string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func typeAssertShort(n *node) { _ = "STUB: not implemented"; return }

func typeAssertLong(n *node) { _ = "STUB: not implemented"; return }

func typeAssertStatus(n *node) { _ = "STUB: not implemented"; return }

func typeAssert(n *node, withResult, withOk bool) { _ = "STUB: not implemented"; return }

func canAssertTypes(src, dest reflect.Type) bool { _ = "STUB: not implemented"; return false }

func firstMissingMethod(src, dest reflect.Type) string { _ = "STUB: not implemented"; return "" }

func convert(n *node) { _ = "STUB: not implemented"; return }

func assignFromCall(n *node) { _ = "STUB: not implemented"; return }

func assign(n *node) { _ = "STUB: not implemented"; return }

func not(n *node) { _ = "STUB: not implemented"; return }

func addr(n *node) { _ = "STUB: not implemented"; return }

func deref(n *node) { _ = "STUB: not implemented"; return }

func _print(n *node) { _ = "STUB: not implemented"; return }

func _println(n *node) { _ = "STUB: not implemented"; return }

func _recover(n *node) { _ = "STUB: not implemented"; return }

func _panic(n *node) { _ = "STUB: not implemented"; return }

func genBuiltinDeferWrapper(n *node, in, out []func(*frame) reflect.Value, fn func([]reflect.Value) []reflect.Value) {
	_ = "STUB: not implemented"
	return
}

func genFunctionWrapper(n *node) func(*frame) reflect.Value { _ = "STUB: not implemented"; return nil }

func genInterfaceWrapper(n *node, typ reflect.Type) func(*frame) reflect.Value {
	_ = "STUB: not implemented"
	return nil
}

func methodByName(value reflect.Value, name string, index []int) (v reflect.Value) {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

func checkFieldIndex(typ reflect.Type, index []int) bool { _ = "STUB: not implemented"; return false }

func call(n *node) { _ = "STUB: not implemented"; return }

//nolint:gocritic

func getFrame(f *frame, l int) *frame { _ = "STUB: not implemented"; return nil }

func callBin(n *node) { _ = "STUB: not implemented"; return }

func getIndexBinMethod(n *node) { _ = "STUB: not implemented"; return }

func getIndexBinElemMethod(n *node) { _ = "STUB: not implemented"; return }

func getIndexBinPtrMethod(n *node) { _ = "STUB: not implemented"; return }

func getIndexArray(n *node) { _ = "STUB: not implemented"; return }

func getIndexMap(n *node) { _ = "STUB: not implemented"; return }

func getIndexMap2(n *node) { _ = "STUB: not implemented"; return }

func getFunc(n *node) { _ = "STUB: not implemented"; return }

func getMethod(n *node) { _ = "STUB: not implemented"; return }

func getMethodByName(n *node) { _ = "STUB: not implemented"; return }

func lookupMethodValue(val valueInterface, name string) (r reflect.Value, m *node, li []int) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil, nil
}

func getIndexSeq(n *node) { _ = "STUB: not implemented"; return }

func getPtrIndexSeq(n *node) { _ = "STUB: not implemented"; return }

func getIndexSeqField(n *node) { _ = "STUB: not implemented"; return }

func getIndexSeqPtrMethod(n *node) { _ = "STUB: not implemented"; return }

func getIndexSeqMethod(n *node) { _ = "STUB: not implemented"; return }

func neg(n *node) { _ = "STUB: not implemented"; return }

func pos(n *node) { _ = "STUB: not implemented"; return }

func bitNot(n *node) { _ = "STUB: not implemented"; return }

func land(n *node) { _ = "STUB: not implemented"; return }

func lor(n *node) { _ = "STUB: not implemented"; return }

func nop(n *node) { _ = "STUB: not implemented"; return }

func branch(n *node) { _ = "STUB: not implemented"; return }

func _return(n *node) { _ = "STUB: not implemented"; return }

func arrayLit(n *node) { _ = "STUB: not implemented"; return }

func mapLit(n *node) { _ = "STUB: not implemented"; return }

func compositeBinMap(n *node) { _ = "STUB: not implemented"; return }

func compositeBinSlice(n *node) { _ = "STUB: not implemented"; return }

func doCompositeBinStruct(n *node, hasType bool) { _ = "STUB: not implemented"; return }

func compositeBinStruct(n *node)       { _ = "STUB: not implemented"; return }
func compositeBinStructNotype(n *node) { _ = "STUB: not implemented"; return }

func destType(n *node) *itype { _ = "STUB: not implemented"; return nil }

func doComposite(n *node, hasType bool, keyed bool) { _ = "STUB: not implemented"; return }

func doCompositeLit(n *node, hasType bool) { _ = "STUB: not implemented"; return }

func compositeLit(n *node)       { _ = "STUB: not implemented"; return }
func compositeLitNotype(n *node) { _ = "STUB: not implemented"; return }

func doCompositeLitKeyed(n *node, hasType bool) { _ = "STUB: not implemented"; return }

func compositeLitKeyed(n *node)       { _ = "STUB: not implemented"; return }
func compositeLitKeyedNotype(n *node) { _ = "STUB: not implemented"; return }

func empty(n *node) { _ = "STUB: not implemented"; return }

var rat = reflect.ValueOf((*[]rune)(nil)).Type().Elem()

func _range(n *node) { _ = "STUB: not implemented"; return }

func rangeInt(n *node) { _ = "STUB: not implemented"; return }

func loopVarKey(n *node) { _ = "STUB: not implemented"; return }

func loopVarVal(n *node) { _ = "STUB: not implemented"; return }

func loopVarFor(n *node) { _ = "STUB: not implemented"; return }

func rangeChan(n *node) { _ = "STUB: not implemented"; return }

func rangeMap(n *node) { _ = "STUB: not implemented"; return }

func _case(n *node) { _ = "STUB: not implemented"; return }

func implementsInterface(v reflect.Value, t *itype) bool { _ = "STUB: not implemented"; return false }

func appendSlice(n *node) { _ = "STUB: not implemented"; return }

func _append(n *node) { _ = "STUB: not implemented"; return }

func _cap(n *node) { _ = "STUB: not implemented"; return }

func _copy(n *node) { _ = "STUB: not implemented"; return }

func _close(n *node) { _ = "STUB: not implemented"; return }

func _complex(n *node) { _ = "STUB: not implemented"; return }

func _imag(n *node) { _ = "STUB: not implemented"; return }

func _real(n *node) { _ = "STUB: not implemented"; return }

func _delete(n *node) { _ = "STUB: not implemented"; return }

func capConst(n *node) { _ = "STUB: not implemented"; return }

func lenConst(n *node) { _ = "STUB: not implemented"; return }

func _len(n *node) { _ = "STUB: not implemented"; return }

func _new(n *node) { _ = "STUB: not implemented"; return }

func _make(n *node) { _ = "STUB: not implemented"; return }

func reset(n *node) { _ = "STUB: not implemented"; return }

func recv(n *node) { _ = "STUB: not implemented"; return }

func recv2(n *node) { _ = "STUB: not implemented"; return }

func convertLiteralValue(n *node, t reflect.Type) { _ = "STUB: not implemented"; return }

func convertConstantValue(n *node) { _ = "STUB: not implemented"; return }

func send(n *node) { _ = "STUB: not implemented"; return }

func clauseChanDir(n *node) (*node, *node, *node, reflect.SelectDir) {
	_ = "STUB: not implemented"
	return nil, nil, nil, *new(reflect.SelectDir)
}

func _select(n *node) { _ = "STUB: not implemented"; return }

func slice(n *node) { _ = "STUB: not implemented"; return }

func slice0(n *node) { _ = "STUB: not implemented"; return }

func isNilChild(child int) func(n *node) { _ = "STUB: not implemented"; return nil }

func isNotNil(n *node) { _ = "STUB: not implemented"; return }

func complexConst(n *node) { _ = "STUB: not implemented"; return }

func imagConst(n *node) { _ = "STUB: not implemented"; return }

func realConst(n *node) { _ = "STUB: not implemented"; return }
