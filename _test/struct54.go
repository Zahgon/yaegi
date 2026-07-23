package main

type S struct {
	t *T
}

func newS() *S { _ = "STUB: not implemented"; return nil }

type T struct {
	u map[string]*U
}

type U struct {
	a int
}

func main() {
	s := newS()
	_ = s

	println("ok")
}
