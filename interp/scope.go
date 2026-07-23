package interp

import (
	"reflect"
)

type sKind uint

const (
	undefSym sKind = iota
	binSym
	bltnSym
	constSym
	funcSym
	labelSym
	pkgSym
	typeSym
	varTypeSym
	varSym
)

var symKinds = [...]string{
	undefSym:   "undefSym",
	binSym:     "binSym",
	bltnSym:    "bltnSym",
	constSym:   "constSym",
	funcSym:    "funcSym",
	labelSym:   "labelSym",
	pkgSym:     "pkgSym",
	typeSym:    "typeSym",
	varTypeSym: "varTypeSym",
	varSym:     "varSym",
}

func (k sKind) String() string { _ = "STUB: not implemented"; return "" }

type symbol struct {
	kind    sKind
	typ     *itype
	node    *node
	from    []*node
	recv    *receiver
	index   int
	rval    reflect.Value
	builtin bltnGenerator
	global  bool
}

type scope struct {
	anc         *scope
	child       []*scope
	def         *node
	loop        *node
	loopRestart *node
	pkgID       string
	pkgName     string
	types       []reflect.Type
	level       int
	sym         map[string]*symbol
	global      bool
	iota        int
}

func (s *scope) push(indirect bool) *scope { _ = "STUB: not implemented"; return nil }

func (s *scope) pushBloc() *scope { _ = "STUB: not implemented"; return nil }
func (s *scope) pushFunc() *scope { _ = "STUB: not implemented"; return nil }

func (s *scope) pop() *scope { _ = "STUB: not implemented"; return nil }

func (s *scope) upperLevel() *scope { _ = "STUB: not implemented"; return nil }

func (s *scope) lookup(ident string) (*symbol, int, bool) {
	_ = "STUB: not implemented"
	return nil, 0, false
}

func (s *scope) isRedeclared(n *node) bool { _ = "STUB: not implemented"; return false }

func (s *scope) rangeChanType(n *node) *itype { _ = "STUB: not implemented"; return nil }

func (s *scope) fixType(t *itype) *itype { _ = "STUB: not implemented"; return nil }

func (s *scope) getType(ident string) *itype { _ = "STUB: not implemented"; return nil }

func (s *scope) add(typ *itype) (index int) { _ = "STUB: not implemented"; return 0 }

func (interp *Interpreter) initScopePkg(pkgID, pkgName string) *scope {
	_ = "STUB: not implemented"
	return nil
}

func (interp *Interpreter) Globals() map[string]reflect.Value {
	_ = "STUB: not implemented"
	return nil
}
