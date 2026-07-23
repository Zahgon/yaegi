package main

type Foo struct {
	val int
}

func f(p *Foo) { _ = "STUB: not implemented"; return }

func main() {
	var a = Foo{3}
	f(&a)
	println(a.val)
}
