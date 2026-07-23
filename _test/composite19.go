package main

import "fmt"

type fn func(string, string) bool

var funcs = map[string]fn{
	"less":    cmpLessFn,
	"greater": cmpGreaterFn,
	"none":    nil,
}

func cmpLessFn(a string, b string) bool { _ = "STUB: not implemented"; return false }

func cmpGreaterFn(a string, b string) bool { _ = "STUB: not implemented"; return false }

func main() {
	for _, n := range []string{"less", "greater", "none"} {
		f := funcs[n]
		if f == nil {
			continue
		}
		fmt.Println(f("a", "b"))
	}
}
