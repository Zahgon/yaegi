package main

import "fmt"

type Error interface {
	error
	Code() string
}

type MyError Error

type T struct {
	Name string
}

func (t *T) Error() string { _ = "STUB: not implemented"; return "" }
func (t *T) Code() string  { _ = "STUB: not implemented"; return "" }

func newT(s string) MyError { _ = "STUB: not implemented"; return *new(MyError) }

func main() {
	t := newT("foo")
	fmt.Println(t.Code())
}
