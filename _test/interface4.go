package main

type fii interface {
	Hello()
}

type Boo struct {
	Name string
}

type Bir struct {
	Boo
}

func (b Boo) Hello() { _ = "STUB: not implemented"; return }

func inCall(foo fii) { _ = "STUB: not implemented"; return }

func main() {
	bir := Bir{Boo{"foo"}}
	inCall(bir)
}
