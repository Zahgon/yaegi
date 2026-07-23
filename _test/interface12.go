package main

type I1 interface {
	Truc()
}

type T1 struct{}

func (T1) Truc() { _ = "STUB: not implemented"; return }

var x I1 = T1{}

func main() {
	x.Truc()
}
