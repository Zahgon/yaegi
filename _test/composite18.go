package main

import "fmt"

type fn func(string, string) bool

var funcs = []fn{
	cmpLessFn,
	cmpGreaterFn,
	nil,
}

func cmpLessFn(a string, b string) bool { _ = "STUB: not implemented"; return false }

func cmpGreaterFn(a string, b string) bool { _ = "STUB: not implemented"; return false }

func main() {
	for _, f := range funcs {
		if f == nil {
			continue
		}
		fmt.Println(f("a", "b"))
	}
}
