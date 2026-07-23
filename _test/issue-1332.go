package main

func run(fn func(name string)) { _ = "STUB: not implemented"; return }

type T2 struct {
	name string
}

func (t *T2) f(s string) { _ = "STUB: not implemented"; return }

func main() {
	t2 := &T2{"foo"}
	run(t2.f)
}
