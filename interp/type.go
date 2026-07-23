package interp

import (
	"go/constant"
	"reflect"
)

type tcat uint

const (
	nilT tcat = iota
	arrayT
	binT
	binPkgT
	boolT
	builtinT
	chanT
	chanSendT
	chanRecvT
	comparableT
	complex64T
	complex128T
	constraintT
	errorT
	float32T
	float64T
	funcT
	genericT
	interfaceT
	intT
	int8T
	int16T
	int32T
	int64T
	linkedT
	mapT
	ptrT
	sliceT
	srcPkgT
	stringT
	structT
	uintT
	uint8T
	uint16T
	uint32T
	uint64T
	uintptrT
	valueT
	variadicT
	maxT
)

var cats = [...]string{
	nilT:        "nilT",
	arrayT:      "arrayT",
	binT:        "binT",
	binPkgT:     "binPkgT",
	boolT:       "boolT",
	builtinT:    "builtinT",
	chanT:       "chanT",
	comparableT: "comparableT",
	complex64T:  "complex64T",
	complex128T: "complex128T",
	constraintT: "constraintT",
	errorT:      "errorT",
	float32T:    "float32",
	float64T:    "float64T",
	funcT:       "funcT",
	genericT:    "genericT",
	interfaceT:  "interfaceT",
	intT:        "intT",
	int8T:       "int8T",
	int16T:      "int16T",
	int32T:      "int32T",
	int64T:      "int64T",
	linkedT:     "linkedT",
	mapT:        "mapT",
	ptrT:        "ptrT",
	sliceT:      "sliceT",
	srcPkgT:     "srcPkgT",
	stringT:     "stringT",
	structT:     "structT",
	uintT:       "uintT",
	uint8T:      "uint8T",
	uint16T:     "uint16T",
	uint32T:     "uint32T",
	uint64T:     "uint64T",
	uintptrT:    "uintptrT",
	valueT:      "valueT",
	variadicT:   "variadicT",
}

func (c tcat) String() string { _ = "STUB: not implemented"; return "" }

type structField struct {
	name  string
	tag   string
	embed bool
	typ   *itype
}

type itype struct {
	cat          tcat
	field        []structField
	key          *itype
	val          *itype
	recv         *itype
	arg          []*itype
	ret          []*itype
	ptr          *itype
	method       []*node
	constraint   []*itype
	ulconstraint []*itype
	instance     []*itype
	name         string
	path         string
	length       int
	rtype        reflect.Type
	node         *node
	scope        *scope
	str          string
	incomplete   bool
	untyped      bool
	isBinMethod  bool
}

type generic struct{}

func untypedBool(n *node) *itype { _ = "STUB: not implemented"; return nil }

func untypedString(n *node) *itype { _ = "STUB: not implemented"; return nil }

func untypedRune(n *node) *itype { _ = "STUB: not implemented"; return nil }

func untypedInt(n *node) *itype { _ = "STUB: not implemented"; return nil }

func untypedFloat(n *node) *itype { _ = "STUB: not implemented"; return nil }

func untypedComplex(n *node) *itype { _ = "STUB: not implemented"; return nil }

func errorMethodType(sc *scope) *itype { _ = "STUB: not implemented"; return nil }

type itypeOption func(*itype)

func isBinMethod() itypeOption { _ = "STUB: not implemented"; return *new(itypeOption) }

func withRecv(typ *itype) itypeOption { _ = "STUB: not implemented"; return *new(itypeOption) }

func withNode(n *node) itypeOption { _ = "STUB: not implemented"; return *new(itypeOption) }

func withScope(sc *scope) itypeOption { _ = "STUB: not implemented"; return *new(itypeOption) }

func withUntyped(b bool) itypeOption { _ = "STUB: not implemented"; return *new(itypeOption) }

func valueTOf(rtype reflect.Type, opts ...itypeOption) *itype {
	_ = "STUB: not implemented"
	return nil
}

func wrapperValueTOf(rtype reflect.Type, val *itype, opts ...itypeOption) *itype {
	_ = "STUB: not implemented"
	return nil
}

func variadicOf(val *itype, opts ...itypeOption) *itype { _ = "STUB: not implemented"; return nil }

func ptrOf(val *itype, opts ...itypeOption) *itype { _ = "STUB: not implemented"; return nil }

func namedOf(val *itype, path, name string, opts ...itypeOption) *itype {
	_ = "STUB: not implemented"
	return nil
}

func funcOf(args []*itype, ret []*itype, opts ...itypeOption) *itype {
	_ = "STUB: not implemented"
	return nil
}

type chanDir uint8

const (
	chanSendRecv chanDir = iota
	chanSend
	chanRecv
)

func chanOf(val *itype, dir chanDir, opts ...itypeOption) *itype {
	_ = "STUB: not implemented"
	return nil
}

func arrayOf(val *itype, l int, opts ...itypeOption) *itype { _ = "STUB: not implemented"; return nil }

func sliceOf(val *itype, opts ...itypeOption) *itype { _ = "STUB: not implemented"; return nil }

func mapOf(key, val *itype, opts ...itypeOption) *itype { _ = "STUB: not implemented"; return nil }

func interfaceOf(t *itype, fields []structField, constraint, ulconstraint []*itype, opts ...itypeOption) *itype {
	_ = "STUB: not implemented"
	return nil
}

func structOf(t *itype, fields []structField, opts ...itypeOption) *itype {
	_ = "STUB: not implemented"
	return nil
}

func genericOf(val *itype, name, path string, opts ...itypeOption) *itype {
	_ = "STUB: not implemented"
	return nil
}

func seenNode(ns []*node, n *node) bool { _ = "STUB: not implemented"; return false }

func nodeType(interp *Interpreter, sc *scope, n *node) (*itype, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func nodeType2(interp *Interpreter, sc *scope, n *node, seen []*node) (t *itype, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func genType(interp *Interpreter, sc *scope, name string, lt *itype, types []*itype, seen []*node) (t *itype, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func genMethod(interp *Interpreter, sc *scope, t *itype, nod *node, types []*itype) error {
	_ = "STUB: not implemented"
	return nil
}

func findPackageType(interp *Interpreter, sc *scope, n *node) *itype {
	_ = "STUB: not implemented"
	return nil
}

func isBuiltinCall(n *node, sc *scope) bool { _ = "STUB: not implemented"; return false }

func typeName(n *node) string { _ = "STUB: not implemented"; return "" }

func structName(n *node) string { _ = "STUB: not implemented"; return "" }

func fieldName(n *node) string { _ = "STUB: not implemented"; return "" }

var zeroValues [maxT]reflect.Value

func init() {
	zeroValues[boolT] = reflect.ValueOf(false)
	zeroValues[complex64T] = reflect.ValueOf(complex64(0))
	zeroValues[complex128T] = reflect.ValueOf(complex128(0))
	zeroValues[errorT] = reflect.ValueOf(new(error)).Elem()
	zeroValues[float32T] = reflect.ValueOf(float32(0))
	zeroValues[float64T] = reflect.ValueOf(float64(0))
	zeroValues[intT] = reflect.ValueOf(int(0))
	zeroValues[int8T] = reflect.ValueOf(int8(0))
	zeroValues[int16T] = reflect.ValueOf(int16(0))
	zeroValues[int32T] = reflect.ValueOf(int32(0))
	zeroValues[int64T] = reflect.ValueOf(int64(0))
	zeroValues[stringT] = reflect.ValueOf("")
	zeroValues[uintT] = reflect.ValueOf(uint(0))
	zeroValues[uint8T] = reflect.ValueOf(uint8(0))
	zeroValues[uint16T] = reflect.ValueOf(uint16(0))
	zeroValues[uint32T] = reflect.ValueOf(uint32(0))
	zeroValues[uint64T] = reflect.ValueOf(uint64(0))
	zeroValues[uintptrT] = reflect.ValueOf(uintptr(0))
}

func (t *itype) finalize() (*itype, error) { _ = "STUB: not implemented"; return nil, nil }

func (t *itype) addMethod(n *node) { _ = "STUB: not implemented"; return }

func (t *itype) numIn() int { _ = "STUB: not implemented"; return 0 }

func (t *itype) in(i int) *itype { _ = "STUB: not implemented"; return nil }

func (t *itype) numOut() int { _ = "STUB: not implemented"; return 0 }

func (t *itype) out(i int) *itype { _ = "STUB: not implemented"; return nil }

func (t *itype) concrete() *itype { _ = "STUB: not implemented"; return nil }

func (t *itype) underlying() *itype { _ = "STUB: not implemented"; return nil }

func typeDefined(t1, t2 *itype) bool { _ = "STUB: not implemented"; return false }

func (t *itype) isVariadic() bool { _ = "STUB: not implemented"; return false }

func (t *itype) isComplete() bool { _ = "STUB: not implemented"; return false }

func isComplete(t *itype, visited map[string]bool) bool { _ = "STUB: not implemented"; return false }

func (t *itype) comparable() bool { _ = "STUB: not implemented"; return false }

func (t *itype) assignableTo(o *itype) bool { _ = "STUB: not implemented"; return false }

func (t *itype) convertibleTo(o *itype) bool { _ = "STUB: not implemented"; return false }

func (t *itype) ordered() bool { _ = "STUB: not implemented"; return false }

func (t *itype) equals(o *itype) bool { _ = "STUB: not implemented"; return false }

func (t *itype) matchDefault(o *itype) bool { _ = "STUB: not implemented"; return false }

type methodSet map[string]string

func (m methodSet) contains(n methodSet) bool { _ = "STUB: not implemented"; return false }

func (m methodSet) equals(n methodSet) bool { _ = "STUB: not implemented"; return false }

func (t *itype) methods() methodSet { _ = "STUB: not implemented"; return *new(methodSet) }

func (t *itype) id() (res string) { _ = "STUB: not implemented"; return "" }

func fixPossibleConstType(t reflect.Type) (r reflect.Type) {
	_ = "STUB: not implemented"
	return *new(reflect.Type)
}

func (t *itype) zero() (v reflect.Value, err error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}

func (t *itype) fieldIndex(name string) int { _ = "STUB: not implemented"; return 0 }

func (t *itype) fieldSeq(seq []int) *itype { _ = "STUB: not implemented"; return nil }

func (t *itype) lookupField(name string) []int { _ = "STUB: not implemented"; return nil }

func (t *itype) lookupBinField(name string) (s reflect.StructField, index []int, ok bool) {
	_ = "STUB: not implemented"
	return *new(reflect.StructField), nil, false
}

func (t *itype) methodCallType() reflect.Type { _ = "STUB: not implemented"; return *new(reflect.Type) }

func (t *itype) resolveAlias() *itype { _ = "STUB: not implemented"; return nil }

func (t *itype) getMethod(name string) *node { _ = "STUB: not implemented"; return nil }

func (t *itype) lookupMethod(name string) (*node, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *itype) lookupMethod2(name string, seen map[*itype]bool) (*node, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *itype) interfaceMethod(name string) *itype { _ = "STUB: not implemented"; return nil }

func (t *itype) interfaceMethod2(name string, seen map[*itype]bool) *itype {
	_ = "STUB: not implemented"
	return nil
}

func (t *itype) methodDepth(name string) int { _ = "STUB: not implemented"; return 0 }

func (t *itype) lookupBinMethod(name string) (m reflect.Method, index []int, isPtr, ok bool) {
	_ = "STUB: not implemented"
	return *new(reflect.Method), nil, false, false
}

func (t *itype) lookupBinMethod2(name string, seen map[*itype]bool) (m reflect.Method, index []int, isPtr, ok bool) {
	_ = "STUB: not implemented"
	return *new(reflect.Method), nil, false, false
}

func lookupFieldOrMethod(t *itype, name string) *itype { _ = "STUB: not implemented"; return nil }

func exportName(s string) string { _ = "STUB: not implemented"; return "" }

var (
	emptyInterfaceType = reflect.TypeOf((*interface{})(nil)).Elem()
	valueInterfaceType = reflect.TypeOf((*valueInterface)(nil)).Elem()
	constVal           = reflect.TypeOf((*constant.Value)(nil)).Elem()
)

type refTypeContext struct {
	defined map[string]*itype

	refs map[string][]*itype

	rect       *itype
	rebuilding bool
	slevel     int
}

func (c *refTypeContext) Clone() *refTypeContext { _ = "STUB: not implemented"; return nil }

func (c *refTypeContext) isComplete() bool { _ = "STUB: not implemented"; return false }

func (t *itype) fixDummy(typ reflect.Type) reflect.Type {
	_ = "STUB: not implemented"
	return *new(reflect.Type)
}

func (t *itype) refType(ctx *refTypeContext) reflect.Type {
	_ = "STUB: not implemented"
	return *new(reflect.Type)
}

func (t *itype) TypeOf() reflect.Type { _ = "STUB: not implemented"; return *new(reflect.Type) }

func (t *itype) frameType() (r reflect.Type) { _ = "STUB: not implemented"; return *new(reflect.Type) }

func (t *itype) implements(it *itype) bool { _ = "STUB: not implemented"; return false }

func (t *itype) defaultType(v reflect.Value, sc *scope) *itype {
	_ = "STUB: not implemented"
	return nil
}

func (t *itype) isNil() bool { _ = "STUB: not implemented"; return false }

func (t *itype) hasNil() bool { _ = "STUB: not implemented"; return false }

func (t *itype) elem() *itype { _ = "STUB: not implemented"; return nil }

func hasElem(t reflect.Type) bool { _ = "STUB: not implemented"; return false }

func constToInt(c constant.Value) int { _ = "STUB: not implemented"; return 0 }

func constToString(v reflect.Value) string { _ = "STUB: not implemented"; return "" }

func wrappedType(n *node) *itype { _ = "STUB: not implemented"; return nil }

func isShiftNode(n *node) bool { _ = "STUB: not implemented"; return false }

func chanElement(t *itype) *itype { _ = "STUB: not implemented"; return nil }

func isBool(t *itype) bool { _ = "STUB: not implemented"; return false }
func isChan(t *itype) bool { _ = "STUB: not implemented"; return false }
func isFunc(t *itype) bool { _ = "STUB: not implemented"; return false }
func isMap(t *itype) bool  { _ = "STUB: not implemented"; return false }
func isPtr(t *itype) bool  { _ = "STUB: not implemented"; return false }

func isEmptyInterface(t *itype) bool { _ = "STUB: not implemented"; return false }

func isGeneric(t *itype) bool { _ = "STUB: not implemented"; return false }

func isNamedFuncSrc(t *itype) bool { _ = "STUB: not implemented"; return false }

func isFuncSrc(t *itype) bool { _ = "STUB: not implemented"; return false }

func isPtrSrc(t *itype) bool { _ = "STUB: not implemented"; return false }

func isSendChan(t *itype) bool { _ = "STUB: not implemented"; return false }

func isArray(t *itype) bool { _ = "STUB: not implemented"; return false }

func isInterfaceSrc(t *itype) bool { _ = "STUB: not implemented"; return false }

func isInterfaceBin(t *itype) bool { _ = "STUB: not implemented"; return false }

func isInterface(t *itype) bool { _ = "STUB: not implemented"; return false }

func isBin(t *itype) bool { _ = "STUB: not implemented"; return false }

func isStruct(t *itype) bool { _ = "STUB: not implemented"; return false }

func isConstType(t *itype) bool { _ = "STUB: not implemented"; return false }

func isInt(t reflect.Type) bool { _ = "STUB: not implemented"; return false }

func isUint(t reflect.Type) bool { _ = "STUB: not implemented"; return false }

func isComplex(t reflect.Type) bool { _ = "STUB: not implemented"; return false }

func isFloat(t reflect.Type) bool { _ = "STUB: not implemented"; return false }

func isByteArray(t reflect.Type) bool { _ = "STUB: not implemented"; return false }

func isFloat32(t reflect.Type) bool { _ = "STUB: not implemented"; return false }
func isFloat64(t reflect.Type) bool { _ = "STUB: not implemented"; return false }
func isNumber(t reflect.Type) bool  { _ = "STUB: not implemented"; return false }

func isBoolean(t reflect.Type) bool       { _ = "STUB: not implemented"; return false }
func isString(t reflect.Type) bool        { _ = "STUB: not implemented"; return false }
func isConstantValue(t reflect.Type) bool { _ = "STUB: not implemented"; return false }
