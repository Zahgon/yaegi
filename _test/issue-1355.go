package main

import "github.com/traefik/yaegi/_test/p2"

func f(i interface{}) { _ = "STUB: not implemented"; return }

func main() {
	var v *p2.T
	var i interface{}

	i = v
	_, ok := i.(p2.I)
	println("ok:", ok)
	f(v)
}
