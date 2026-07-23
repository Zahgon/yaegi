package interp

import (
	"context"
	"go/ast"
	"go/token"
	"reflect"
)

type Program struct {
	pkgName string
	root    *node
	init    []*node
}

func (p *Program) PackageName() string { _ = "STUB: not implemented"; return "" }

func (interp *Interpreter) FileSet() *token.FileSet { _ = "STUB: not implemented"; return nil }

func (interp *Interpreter) Compile(src string) (*Program, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (interp *Interpreter) CompilePath(path string) (*Program, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (interp *Interpreter) compileSrc(src, name string, inc bool) (*Program, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (interp *Interpreter) CompileAST(n ast.Node) (*Program, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (interp *Interpreter) Execute(p *Program) (res reflect.Value, err error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}

func (interp *Interpreter) ExecuteWithContext(ctx context.Context, p *Program) (res reflect.Value, err error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}
