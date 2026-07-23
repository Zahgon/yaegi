package main

type Ti func(*T)

type T1 struct {
	t Ti
}

type T struct {
	t Ti
	y *xxx
}

func f(t *T) { _ = "STUB: not implemented"; return }

type xxx struct{}

var x = &T1{t: f}

func main() {
	x.t = f
	println("ok")
}
