//go:build go1.22
// +build go1.22

package stdlib

import (
	"go/build/constraint"
	"reflect"
)

func init() {
	Symbols["go/build/constraint/constraint"] = map[string]reflect.Value{

		"GoVersion":      reflect.ValueOf(constraint.GoVersion),
		"IsGoBuild":      reflect.ValueOf(constraint.IsGoBuild),
		"IsPlusBuild":    reflect.ValueOf(constraint.IsPlusBuild),
		"Parse":          reflect.ValueOf(constraint.Parse),
		"PlusBuildLines": reflect.ValueOf(constraint.PlusBuildLines),

		"AndExpr":     reflect.ValueOf((*constraint.AndExpr)(nil)),
		"Expr":        reflect.ValueOf((*constraint.Expr)(nil)),
		"NotExpr":     reflect.ValueOf((*constraint.NotExpr)(nil)),
		"OrExpr":      reflect.ValueOf((*constraint.OrExpr)(nil)),
		"SyntaxError": reflect.ValueOf((*constraint.SyntaxError)(nil)),
		"TagExpr":     reflect.ValueOf((*constraint.TagExpr)(nil)),

		"_Expr": reflect.ValueOf((*_go_build_constraint_Expr)(nil)),
	}
}

type _go_build_constraint_Expr struct {
	IValue  interface{}
	WEval   func(ok func(tag string) bool) bool
	WString func() string
}

func (W _go_build_constraint_Expr) Eval(ok func(tag string) bool) bool {
	_ = "STUB: not implemented"
	return false
}
func (W _go_build_constraint_Expr) String() string { _ = "STUB: not implemented"; return "" }
