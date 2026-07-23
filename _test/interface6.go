package main

type Myint int

func (i Myint) Double() { _ = "STUB: not implemented"; return }

type Boo interface {
	Double()
}

func f(boo Boo) { _ = "STUB: not implemented"; return }

func g(i int) Boo { _ = "STUB: not implemented"; return *new(Boo) }

func main() {
	f(g(4))
}
