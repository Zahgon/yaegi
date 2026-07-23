package main

type T1 struct {
	T2 *T2
}

func (t *T1) Get() string { _ = "STUB: not implemented"; return "" }

type T2 struct {
	Name string
}

func (t *T2) V() *T2 { _ = "STUB: not implemented"; return nil }

var defaultT2 = &T2{"no name"}

func main() {
	t := &T1{}
	println(t.Get())
}
