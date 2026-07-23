package main

type Myint int

func (i Myint) Double() { _ = "STUB: not implemented"; return }

type Boo interface {
	Double()
}

func f(boo Boo) { _ = "STUB: not implemented"; return }

func main() {
	var i Myint = 3
	f(i)
}
