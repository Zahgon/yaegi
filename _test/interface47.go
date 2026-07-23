package main

type Doer interface {
	Do() error
}

type T struct {
	Name string
}

func (t *T) Do() error { _ = "STUB: not implemented"; return nil }

func f() (Doer, error) { _ = "STUB: not implemented"; return *new(Doer), nil }

type Ev struct {
	doer func() (Doer, error)
}

func (e *Ev) do() { _ = "STUB: not implemented"; return }

func main() {
	e := &Ev{f}
	println(e != nil)
	e.do()
}
