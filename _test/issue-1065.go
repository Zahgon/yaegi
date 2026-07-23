package main

import "fmt"

type AST struct {
	Num      int
	Children []AST
}

func newAST(num int, root AST, children ...AST) AST { _ = "STUB: not implemented"; return *new(AST) }

func main() {
	ast := newAST(1, AST{}, AST{})
	fmt.Println(ast)
}
