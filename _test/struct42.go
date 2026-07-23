package main

type T struct {
	t func(*T)
	y *xxx
}

func f(t *T) { _ = "STUB: not implemented"; return }

var x = &T{t: f}

type xxx struct{}

func main() {
	println("ok")
}
