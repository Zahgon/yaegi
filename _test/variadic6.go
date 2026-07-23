package main

type A struct {
}

func (a A) f(vals ...bool) { _ = "STUB: not implemented"; return }

func main() {
	bools := []bool{true}
	a := A{}
	a.f(bools...)
}
