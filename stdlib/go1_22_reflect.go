//go:build go1.22
// +build go1.22

package stdlib

import (
	"reflect"
)

func init() {
	Symbols["reflect/reflect"] = map[string]reflect.Value{

		"Append":          reflect.ValueOf(reflect.Append),
		"AppendSlice":     reflect.ValueOf(reflect.AppendSlice),
		"Array":           reflect.ValueOf(reflect.Array),
		"ArrayOf":         reflect.ValueOf(reflect.ArrayOf),
		"Bool":            reflect.ValueOf(reflect.Bool),
		"BothDir":         reflect.ValueOf(reflect.BothDir),
		"Chan":            reflect.ValueOf(reflect.Chan),
		"ChanOf":          reflect.ValueOf(reflect.ChanOf),
		"Complex128":      reflect.ValueOf(reflect.Complex128),
		"Complex64":       reflect.ValueOf(reflect.Complex64),
		"Copy":            reflect.ValueOf(reflect.Copy),
		"DeepEqual":       reflect.ValueOf(reflect.DeepEqual),
		"Float32":         reflect.ValueOf(reflect.Float32),
		"Float64":         reflect.ValueOf(reflect.Float64),
		"Func":            reflect.ValueOf(reflect.Func),
		"FuncOf":          reflect.ValueOf(reflect.FuncOf),
		"Indirect":        reflect.ValueOf(reflect.Indirect),
		"Int":             reflect.ValueOf(reflect.Int),
		"Int16":           reflect.ValueOf(reflect.Int16),
		"Int32":           reflect.ValueOf(reflect.Int32),
		"Int64":           reflect.ValueOf(reflect.Int64),
		"Int8":            reflect.ValueOf(reflect.Int8),
		"Interface":       reflect.ValueOf(reflect.Interface),
		"Invalid":         reflect.ValueOf(reflect.Invalid),
		"MakeChan":        reflect.ValueOf(reflect.MakeChan),
		"MakeFunc":        reflect.ValueOf(reflect.MakeFunc),
		"MakeMap":         reflect.ValueOf(reflect.MakeMap),
		"MakeMapWithSize": reflect.ValueOf(reflect.MakeMapWithSize),
		"MakeSlice":       reflect.ValueOf(reflect.MakeSlice),
		"Map":             reflect.ValueOf(reflect.Map),
		"MapOf":           reflect.ValueOf(reflect.MapOf),
		"New":             reflect.ValueOf(reflect.New),
		"NewAt":           reflect.ValueOf(reflect.NewAt),
		"Pointer":         reflect.ValueOf(reflect.Pointer),
		"PointerTo":       reflect.ValueOf(reflect.PointerTo),
		"Ptr":             reflect.ValueOf(reflect.Ptr),
		"PtrTo":           reflect.ValueOf(reflect.PtrTo),
		"RecvDir":         reflect.ValueOf(reflect.RecvDir),
		"Select":          reflect.ValueOf(reflect.Select),
		"SelectDefault":   reflect.ValueOf(reflect.SelectDefault),
		"SelectRecv":      reflect.ValueOf(reflect.SelectRecv),
		"SelectSend":      reflect.ValueOf(reflect.SelectSend),
		"SendDir":         reflect.ValueOf(reflect.SendDir),
		"Slice":           reflect.ValueOf(reflect.Slice),
		"SliceOf":         reflect.ValueOf(reflect.SliceOf),
		"String":          reflect.ValueOf(reflect.String),
		"Struct":          reflect.ValueOf(reflect.Struct),
		"StructOf":        reflect.ValueOf(reflect.StructOf),
		"Swapper":         reflect.ValueOf(reflect.Swapper),
		"TypeOf":          reflect.ValueOf(reflect.TypeOf),
		"Uint":            reflect.ValueOf(reflect.Uint),
		"Uint16":          reflect.ValueOf(reflect.Uint16),
		"Uint32":          reflect.ValueOf(reflect.Uint32),
		"Uint64":          reflect.ValueOf(reflect.Uint64),
		"Uint8":           reflect.ValueOf(reflect.Uint8),
		"Uintptr":         reflect.ValueOf(reflect.Uintptr),
		"UnsafePointer":   reflect.ValueOf(reflect.UnsafePointer),
		"ValueOf":         reflect.ValueOf(reflect.ValueOf),
		"VisibleFields":   reflect.ValueOf(reflect.VisibleFields),
		"Zero":            reflect.ValueOf(reflect.Zero),

		"ChanDir":      reflect.ValueOf((*reflect.ChanDir)(nil)),
		"Kind":         reflect.ValueOf((*reflect.Kind)(nil)),
		"MapIter":      reflect.ValueOf((*reflect.MapIter)(nil)),
		"Method":       reflect.ValueOf((*reflect.Method)(nil)),
		"SelectCase":   reflect.ValueOf((*reflect.SelectCase)(nil)),
		"SelectDir":    reflect.ValueOf((*reflect.SelectDir)(nil)),
		"SliceHeader":  reflect.ValueOf((*reflect.SliceHeader)(nil)),
		"StringHeader": reflect.ValueOf((*reflect.StringHeader)(nil)),
		"StructField":  reflect.ValueOf((*reflect.StructField)(nil)),
		"StructTag":    reflect.ValueOf((*reflect.StructTag)(nil)),
		"Type":         reflect.ValueOf((*reflect.Type)(nil)),
		"Value":        reflect.ValueOf((*reflect.Value)(nil)),
		"ValueError":   reflect.ValueOf((*reflect.ValueError)(nil)),

		"_Type": reflect.ValueOf((*_reflect_Type)(nil)),
	}
}

type _reflect_Type struct {
	IValue           interface{}
	WAlign           func() int
	WAssignableTo    func(u reflect.Type) bool
	WBits            func() int
	WChanDir         func() reflect.ChanDir
	WComparable      func() bool
	WConvertibleTo   func(u reflect.Type) bool
	WElem            func() reflect.Type
	WField           func(i int) reflect.StructField
	WFieldAlign      func() int
	WFieldByIndex    func(index []int) reflect.StructField
	WFieldByName     func(name string) (reflect.StructField, bool)
	WFieldByNameFunc func(match func(string) bool) (reflect.StructField, bool)
	WImplements      func(u reflect.Type) bool
	WIn              func(i int) reflect.Type
	WIsVariadic      func() bool
	WKey             func() reflect.Type
	WKind            func() reflect.Kind
	WLen             func() int
	WMethod          func(a0 int) reflect.Method
	WMethodByName    func(a0 string) (reflect.Method, bool)
	WName            func() string
	WNumField        func() int
	WNumIn           func() int
	WNumMethod       func() int
	WNumOut          func() int
	WOut             func(i int) reflect.Type
	WPkgPath         func() string
	WSize            func() uintptr
	WString          func() string
}

func (W _reflect_Type) Align() int                       { _ = "STUB: not implemented"; return 0 }
func (W _reflect_Type) AssignableTo(u reflect.Type) bool { _ = "STUB: not implemented"; return false }
func (W _reflect_Type) Bits() int                        { _ = "STUB: not implemented"; return 0 }
func (W _reflect_Type) ChanDir() reflect.ChanDir {
	_ = "STUB: not implemented"
	return *new(reflect.ChanDir)
}
func (W _reflect_Type) Comparable() bool                  { _ = "STUB: not implemented"; return false }
func (W _reflect_Type) ConvertibleTo(u reflect.Type) bool { _ = "STUB: not implemented"; return false }
func (W _reflect_Type) Elem() reflect.Type                { _ = "STUB: not implemented"; return *new(reflect.Type) }
func (W _reflect_Type) Field(i int) reflect.StructField {
	_ = "STUB: not implemented"
	return *new(reflect.StructField)
}
func (W _reflect_Type) FieldAlign() int { _ = "STUB: not implemented"; return 0 }
func (W _reflect_Type) FieldByIndex(index []int) reflect.StructField {
	_ = "STUB: not implemented"
	return *new(reflect.StructField)
}
func (W _reflect_Type) FieldByName(name string) (reflect.StructField, bool) {
	_ = "STUB: not implemented"
	return *new(reflect.StructField), false
}

func (W _reflect_Type) FieldByNameFunc(match func(string) bool) (reflect.StructField, bool) {
	_ = "STUB: not implemented"
	return *new(reflect.StructField), false
}

func (W _reflect_Type) Implements(u reflect.Type) bool { _ = "STUB: not implemented"; return false }
func (W _reflect_Type) In(i int) reflect.Type          { _ = "STUB: not implemented"; return *new(reflect.Type) }
func (W _reflect_Type) IsVariadic() bool               { _ = "STUB: not implemented"; return false }
func (W _reflect_Type) Key() reflect.Type              { _ = "STUB: not implemented"; return *new(reflect.Type) }
func (W _reflect_Type) Kind() reflect.Kind             { _ = "STUB: not implemented"; return *new(reflect.Kind) }
func (W _reflect_Type) Len() int                       { _ = "STUB: not implemented"; return 0 }
func (W _reflect_Type) Method(a0 int) reflect.Method {
	_ = "STUB: not implemented"
	return *new(reflect.Method)
}
func (W _reflect_Type) MethodByName(a0 string) (reflect.Method, bool) {
	_ = "STUB: not implemented"
	return *new(reflect.Method), false
}
func (W _reflect_Type) Name() string   { _ = "STUB: not implemented"; return "" }
func (W _reflect_Type) NumField() int  { _ = "STUB: not implemented"; return 0 }
func (W _reflect_Type) NumIn() int     { _ = "STUB: not implemented"; return 0 }
func (W _reflect_Type) NumMethod() int { _ = "STUB: not implemented"; return 0 }
func (W _reflect_Type) NumOut() int    { _ = "STUB: not implemented"; return 0 }
func (W _reflect_Type) Out(i int) reflect.Type {
	_ = "STUB: not implemented"
	return *new(reflect.Type)
}
func (W _reflect_Type) PkgPath() string { _ = "STUB: not implemented"; return "" }
func (W _reflect_Type) Size() uintptr   { _ = "STUB: not implemented"; return 0 }
func (W _reflect_Type) String() string  { _ = "STUB: not implemented"; return "" }
