package main

type T struct {
	f func(*T)
}

func f1(t *T) { _ = "STUB: not implemented"; return }

func main() {
	t := &T{}
	f1(t)
	println(t.f != nil)
}
