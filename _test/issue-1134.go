package main

type I interface {
	Hello()
}

type T struct {
	Name  string
	Child []*T
}

func (t *T) Hello() { _ = "STUB: not implemented"; return }

func main() {
	var i I = new(T)
	i.Hello()
}
