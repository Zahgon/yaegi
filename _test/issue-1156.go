package main

type myInterface interface {
	myFunc() string
}

type V struct{}

func (v *V) myFunc() string { _ = "STUB: not implemented"; return "" }

type U struct {
	v myInterface
}

func (u *U) myFunc() string { _ = "STUB: not implemented"; return "" }

func main() {
	x := V{}
	y := myInterface(&x)
	y = &U{y}
	println(y.myFunc())
}
