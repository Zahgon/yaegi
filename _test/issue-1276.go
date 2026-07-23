package main

import (
	"flag"
)

type customFlag struct{}

func (cf customFlag) String() string { _ = "STUB: not implemented"; return "" }

func (cf customFlag) Set(string) error { _ = "STUB: not implemented"; return nil }

func main() {
	flag.Var(customFlag{}, "cf", "custom flag")
	flag.Parse()
	println("Hello, playground")
}
