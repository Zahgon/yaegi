package main

type Ti func(*T) X

type T1 struct {
	t Ti
}

type T struct {
	t Ti
	y *xxx
}

func f(t *T) X { _ = "STUB: not implemented"; return *new(X) }

type X struct{ Name string }

type xxx struct{}

var x = &T1{t: f}

func main() {
	println("ok")
}
