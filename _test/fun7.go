package main

import (
	goflag "flag"
)

func Foo(goflag *goflag.Flag) { _ = "STUB: not implemented"; return }

func main() {
	g := &goflag.Flag{}
	Foo(g)
}
