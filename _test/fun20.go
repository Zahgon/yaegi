package main

import "fmt"

var myerr error = fmt.Errorf("bar")

func ferr() error { _ = "STUB: not implemented"; return nil }

func foo() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func main() {
	a, b := foo()
	fmt.Println(a, b)
}
