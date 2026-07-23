package main

type T1 struct {
	Name string
}

func (t *T1) genAdd(k int) func(int) int { _ = "STUB: not implemented"; return nil }

var t = &T1{"test"}

func main() {
	f := t.genAdd(4)
	println(f(5))
}
