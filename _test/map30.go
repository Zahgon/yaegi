package main

import "strings"

func f(s string) string { _ = "STUB: not implemented"; return "" }

func g(s string) string { _ = "STUB: not implemented"; return "" }

var methods = map[string]func(string) string{
	"f": f,
	"h": strings.ToLower,
}

func main() {
	methods["i"] = strings.ToUpper
	methods["g"] = g
	println(methods["f"]("test"))
	println(methods["g"]("test"))
	println(methods["i"]("test"))
	println(methods["h"]("TEST"))
}
