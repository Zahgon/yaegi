package main

import (
	"fmt"
	"go/ast"
)

func NewBadExpr() ast.Expr { _ = "STUB: not implemented"; return *new(ast.Expr) }

func main() {
	fmt.Printf("%T\n", NewBadExpr().(*ast.BadExpr))
}
