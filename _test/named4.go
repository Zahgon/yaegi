package main

import (
	"net/http"
)

type A http.Header

func (a A) Test1() { _ = "STUB: not implemented"; return }

type B A

func (b B) Test2() { _ = "STUB: not implemented"; return }

func (b B) Test3() { _ = "STUB: not implemented"; return }

func main() {
	b := B{}

	b.Test2()
	b["test"] = []string{"a", "b"}
	b.Test3()
}
