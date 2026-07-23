package main

type F func(a *A)

type A struct {
	B string
	D
}

type D struct {
	*A
	E *A
	f F
}

func f1(a *A) { _ = "STUB: not implemented"; return }

func main() {
	a := &A{B: "b"}
	a.D = D{f: f1}
	a.f(a)
}
