package interp

import (
	"go/ast"
	"go/token"
)

type nkind uint

const (
	undefNode nkind = iota
	addressExpr
	arrayType
	assignStmt
	assignXStmt
	basicLit
	binaryExpr
	blockStmt
	branchStmt
	breakStmt
	callExpr
	caseBody
	caseClause
	chanType
	chanTypeSend
	chanTypeRecv
	commClause
	commClauseDefault
	compositeLitExpr
	constDecl
	continueStmt
	declStmt
	deferStmt
	defineStmt
	defineXStmt
	ellipsisExpr
	exprStmt
	fallthroughtStmt
	fieldExpr
	fieldList
	fileStmt
	forStmt0
	forStmt1
	forStmt2
	forStmt3
	forStmt4
	forStmt5
	forStmt6
	forStmt7
	forRangeStmt
	funcDecl
	funcLit
	funcType
	goStmt
	gotoStmt
	identExpr
	ifStmt0
	ifStmt1
	ifStmt2
	ifStmt3
	importDecl
	importSpec
	incDecStmt
	indexExpr
	indexListExpr
	interfaceType
	keyValueExpr
	labeledStmt
	landExpr
	lorExpr
	mapType
	parenExpr
	rangeStmt
	returnStmt
	selectStmt
	selectorExpr
	selectorImport
	sendStmt
	sliceExpr
	starExpr
	structType
	switchStmt
	switchIfStmt
	typeAssertExpr
	typeDecl
	typeSpec
	typeSpecAssign
	typeSwitch
	unaryExpr
	valueSpec
	varDecl
)

var kinds = [...]string{
	undefNode:         "undefNode",
	addressExpr:       "addressExpr",
	arrayType:         "arrayType",
	assignStmt:        "assignStmt",
	assignXStmt:       "assignXStmt",
	basicLit:          "basicLit",
	binaryExpr:        "binaryExpr",
	blockStmt:         "blockStmt",
	branchStmt:        "branchStmt",
	breakStmt:         "breakStmt",
	callExpr:          "callExpr",
	caseBody:          "caseBody",
	caseClause:        "caseClause",
	chanType:          "chanType",
	chanTypeSend:      "chanTypeSend",
	chanTypeRecv:      "chanTypeRecv",
	commClause:        "commClause",
	commClauseDefault: "commClauseDefault",
	compositeLitExpr:  "compositeLitExpr",
	constDecl:         "constDecl",
	continueStmt:      "continueStmt",
	declStmt:          "declStmt",
	deferStmt:         "deferStmt",
	defineStmt:        "defineStmt",
	defineXStmt:       "defineXStmt",
	ellipsisExpr:      "ellipsisExpr",
	exprStmt:          "exprStmt",
	fallthroughtStmt:  "fallthroughStmt",
	fieldExpr:         "fieldExpr",
	fieldList:         "fieldList",
	fileStmt:          "fileStmt",
	forStmt0:          "forStmt0",
	forStmt1:          "forStmt1",
	forStmt2:          "forStmt2",
	forStmt3:          "forStmt3",
	forStmt4:          "forStmt4",
	forStmt5:          "forStmt5",
	forStmt6:          "forStmt6",
	forStmt7:          "forStmt7",
	forRangeStmt:      "forRangeStmt",
	funcDecl:          "funcDecl",
	funcType:          "funcType",
	funcLit:           "funcLit",
	goStmt:            "goStmt",
	gotoStmt:          "gotoStmt",
	identExpr:         "identExpr",
	ifStmt0:           "ifStmt0",
	ifStmt1:           "ifStmt1",
	ifStmt2:           "ifStmt2",
	ifStmt3:           "ifStmt3",
	importDecl:        "importDecl",
	importSpec:        "importSpec",
	incDecStmt:        "incDecStmt",
	indexExpr:         "indexExpr",
	indexListExpr:     "indexListExpr",
	interfaceType:     "interfaceType",
	keyValueExpr:      "keyValueExpr",
	labeledStmt:       "labeledStmt",
	landExpr:          "landExpr",
	lorExpr:           "lorExpr",
	mapType:           "mapType",
	parenExpr:         "parenExpr",
	rangeStmt:         "rangeStmt",
	returnStmt:        "returnStmt",
	selectStmt:        "selectStmt",
	selectorExpr:      "selectorExpr",
	selectorImport:    "selectorImport",
	sendStmt:          "sendStmt",
	sliceExpr:         "sliceExpr",
	starExpr:          "starExpr",
	structType:        "structType",
	switchStmt:        "switchStmt",
	switchIfStmt:      "switchIfStmt",
	typeAssertExpr:    "typeAssertExpr",
	typeDecl:          "typeDecl",
	typeSpec:          "typeSpec",
	typeSpecAssign:    "typeSpecAssign",
	typeSwitch:        "typeSwitch",
	unaryExpr:         "unaryExpr",
	valueSpec:         "valueSpec",
	varDecl:           "varDecl",
}

func (k nkind) String() string { _ = "STUB: not implemented"; return "" }

type astError error

type action uint

const (
	aNop action = iota
	aAddr
	aAssign
	aAssignX
	aAdd
	aAddAssign
	aAnd
	aAndAssign
	aAndNot
	aAndNotAssign
	aBitNot
	aBranch
	aCall
	aCallSlice
	aCase
	aCompositeLit
	aConvert
	aDec
	aEqual
	aGreater
	aGreaterEqual
	aGetFunc
	aGetIndex
	aGetMethod
	aGetSym
	aInc
	aLand
	aLor
	aLower
	aLowerEqual
	aMethod
	aMul
	aMulAssign
	aNeg
	aNot
	aNotEqual
	aOr
	aOrAssign
	aPos
	aQuo
	aQuoAssign
	aRange
	aRecv
	aRem
	aRemAssign
	aReturn
	aSend
	aShl
	aShlAssign
	aShr
	aShrAssign
	aSlice
	aSlice0
	aStar
	aSub
	aSubAssign
	aTypeAssert
	aXor
	aXorAssign
)

var actions = [...]string{
	aNop:          "nop",
	aAddr:         "&",
	aAssign:       "=",
	aAssignX:      "X=",
	aAdd:          "+",
	aAddAssign:    "+=",
	aAnd:          "&",
	aAndAssign:    "&=",
	aAndNot:       "&^",
	aAndNotAssign: "&^=",
	aBitNot:       "^",
	aBranch:       "branch",
	aCall:         "call",
	aCallSlice:    "callSlice",
	aCase:         "case",
	aCompositeLit: "compositeLit",
	aConvert:      "convert",
	aDec:          "--",
	aEqual:        "==",
	aGreater:      ">",
	aGreaterEqual: ">=",
	aGetFunc:      "getFunc",
	aGetIndex:     "getIndex",
	aGetMethod:    "getMethod",
	aGetSym:       ".",
	aInc:          "++",
	aLand:         "&&",
	aLor:          "||",
	aLower:        "<",
	aLowerEqual:   "<=",
	aMethod:       "Method",
	aMul:          "*",
	aMulAssign:    "*=",
	aNeg:          "-",
	aNot:          "!",
	aNotEqual:     "!=",
	aOr:           "|",
	aOrAssign:     "|=",
	aPos:          "+",
	aQuo:          "/",
	aQuoAssign:    "/=",
	aRange:        "range",
	aRecv:         "<-",
	aRem:          "%",
	aRemAssign:    "%=",
	aReturn:       "return",
	aSend:         "<~",
	aShl:          "<<",
	aShlAssign:    "<<=",
	aShr:          ">>",
	aShrAssign:    ">>=",
	aSlice:        "slice",
	aSlice0:       "slice0",
	aStar:         "*",
	aSub:          "-",
	aSubAssign:    "-=",
	aTypeAssert:   "TypeAssert",
	aXor:          "^",
	aXorAssign:    "^=",
}

func (a action) String() string { _ = "STUB: not implemented"; return "" }

func isAssignAction(a action) bool { _ = "STUB: not implemented"; return false }

func (interp *Interpreter) firstToken(src string) token.Token {
	_ = "STUB: not implemented"
	return *new(token.Token)
}

func ignoreError(err error, src string) bool { _ = "STUB: not implemented"; return false }

func wrapInMain(src string) string { _ = "STUB: not implemented"; return "" }

func (interp *Interpreter) parse(src, name string, inc bool) (node ast.Node, err error) {
	_ = "STUB: not implemented"
	return *new(ast.Node), nil
}

func (interp *Interpreter) ast(f ast.Node) (string, *node, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

type astNode struct {
	node *node
	ast  ast.Node
}

type nodestack []astNode

func (s *nodestack) push(n *node, a ast.Node) { _ = "STUB: not implemented"; return }

func (s *nodestack) pop() astNode { _ = "STUB: not implemented"; return *new(astNode) }

func (s *nodestack) top() astNode { _ = "STUB: not implemented"; return *new(astNode) }

func (interp *Interpreter) dup(nod, anc *node) *node { _ = "STUB: not implemented"; return nil }
