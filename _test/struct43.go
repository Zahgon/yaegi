package main

type T struct {
	t func(*T)
	y *xxx
}

func f(t *T) { _ = "STUB: not implemented"; return }

type xxx struct{}

func main() {
	x := &T{}
	x.t = f
	println("ok")
}
