package interp

import (
	"reflect"
)

type cfgError struct {
	*node
	error
}

func (c *cfgError) Error() string { _ = "STUB: not implemented"; return "" }

var constOp = map[action]func(*node){
	aAdd:    addConst,
	aSub:    subConst,
	aMul:    mulConst,
	aQuo:    quoConst,
	aRem:    remConst,
	aAnd:    andConst,
	aOr:     orConst,
	aShl:    shlConst,
	aShr:    shrConst,
	aAndNot: andNotConst,
	aXor:    xorConst,
	aNot:    notConst,
	aBitNot: bitNotConst,
	aNeg:    negConst,
	aPos:    posConst,
}

var constBltn = map[string]func(*node){
	bltnComplex: complexConst,
	bltnImag:    imagConst,
	bltnReal:    realConst,
}

const nilIdent = "nil"

func init() {

	constBltn[bltnAlignof] = alignof
	constBltn[bltnOffsetof] = offsetof
	constBltn[bltnSizeof] = sizeof
}

func (interp *Interpreter) cfg(root *node, sc *scope, importPath, pkgName string) ([]*node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func fixUntyped(nod *node, sc *scope) { _ = "STUB: not implemented"; return }

func compDefineX(sc *scope, n *node) error { _ = "STUB: not implemented"; return nil }

func childPos(n *node) int { _ = "STUB: not implemented"; return 0 }

func (n *node) cfgErrorf(format string, a ...interface{}) *cfgError {
	_ = "STUB: not implemented"
	return nil
}

func genRun(nod *node) error { _ = "STUB: not implemented"; return nil }

func genGlobalVars(roots []*node, sc *scope) (*node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getVars(n *node) (vars []*node) { _ = "STUB: not implemented"; return nil }

func genGlobalVarDecl(nodes []*node, sc *scope) (*node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getVarDependencies(nod *node, sc *scope) (deps []*node) { _ = "STUB: not implemented"; return nil }

func setFNext(cond, next *node) { _ = "STUB: not implemented"; return }

func getDefault(n *node) int { _ = "STUB: not implemented"; return 0 }

func isBinType(v reflect.Value) bool { _ = "STUB: not implemented"; return false }

func (n *node) isType(sc *scope) bool { _ = "STUB: not implemented"; return false }

func wireChild(n *node, exclude ...nkind) { _ = "STUB: not implemented"; return }

func excludeNodeKind(child []*node, kinds []nkind) []*node { _ = "STUB: not implemented"; return nil }

func (n *node) name() (s string) { _ = "STUB: not implemented"; return "" }

func (n *node) isNatural() bool { _ = "STUB: not implemented"; return false }

func (n *node) isNil() bool { _ = "STUB: not implemented"; return false }

func (n *node) fieldType(m int) *node { _ = "STUB: not implemented"; return nil }

func (n *node) lastChild() *node { _ = "STUB: not implemented"; return nil }

func (n *node) hasAnc(nod *node) bool { _ = "STUB: not implemented"; return false }

func isKey(n *node) bool { _ = "STUB: not implemented"; return false }

func isField(n *node) bool { _ = "STUB: not implemented"; return false }

func isInInterfaceType(n *node) bool { _ = "STUB: not implemented"; return false }

func isInConstOrTypeDecl(n *node) bool { _ = "STUB: not implemented"; return false }

func isNewDefine(n *node, sc *scope) bool { _ = "STUB: not implemented"; return false }

func isMethod(n *node) bool { _ = "STUB: not implemented"; return false }

func isFuncField(n *node) bool { _ = "STUB: not implemented"; return false }

func isMapEntry(n *node) bool { _ = "STUB: not implemented"; return false }

func isCall(n *node) bool { _ = "STUB: not implemented"; return false }

func isBinCall(n *node, sc *scope) bool { _ = "STUB: not implemented"; return false }

func mustReturnValue(n *node) bool { _ = "STUB: not implemented"; return false }

func isRegularCall(n *node) bool { _ = "STUB: not implemented"; return false }

func variadicPos(n *node) int { _ = "STUB: not implemented"; return 0 }

func canExport(name string) bool { _ = "STUB: not implemented"; return false }

func getExec(n *node) bltn { _ = "STUB: not implemented"; return *new(bltn) }

func setExec(n *node) { _ = "STUB: not implemented"; return }

func typeSwichAssign(n *node) bool { _ = "STUB: not implemented"; return false }

func compositeGenerator(n *node, typ *itype, rtyp reflect.Type) (gen bltnGenerator) {
	_ = "STUB: not implemented"
	return *new(bltnGenerator)
}

func matchSelectorMethod(sc *scope, n *node) (err error) { _ = "STUB: not implemented"; return nil }

func arrayTypeLen(n *node, sc *scope) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func isValueUntyped(v reflect.Value) bool { _ = "STUB: not implemented"; return false }

func isArithmeticAction(n *node) bool { _ = "STUB: not implemented"; return false }

func isBoolAction(n *node) bool { _ = "STUB: not implemented"; return false }

func isBlank(n *node) bool { _ = "STUB: not implemented"; return false }

func alignof(n *node) { _ = "STUB: not implemented"; return }

func offsetof(n *node) { _ = "STUB: not implemented"; return }

func sizeof(n *node) { _ = "STUB: not implemented"; return }
