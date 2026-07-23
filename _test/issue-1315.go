package main

type Intf interface {
	M()
}

type T struct {
	s string
}

func (t *T) M() { _ = "STUB: not implemented"; return }

func f(i interface{}) { _ = "STUB: not implemented"; return }

func main() {
	var i Intf
	var k interface{} = 1
	i = &T{"hello"}
	f(i)
	f(k)
}
