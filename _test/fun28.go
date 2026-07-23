package main

type T struct {
	name string
}

func finalize(t *T) { _ = "STUB: not implemented"; return }

func newT() *T { _ = "STUB: not implemented"; return nil }

func main() {
	t := newT()
	println(t != nil)
}
