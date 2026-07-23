package main

type T struct {
	t *T1
	y *xxx
}

type T1 struct {
	T
}

func f(t *T) { _ = "STUB: not implemented"; return }

var x = &T1{}

type xxx struct{}

func main() {
	println("ok")
}
