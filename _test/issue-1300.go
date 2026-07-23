package main

const buflen = 512

type T struct {
	buf []byte
}

func f(t *T) { _ = "STUB: not implemented"; return }

func main() {
	s := T{}
	println(cap(s.buf))
	f(&s)
	println(cap(s.buf))
}
