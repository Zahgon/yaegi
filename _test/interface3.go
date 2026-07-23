package main

type fii interface {
	Hello()
}

type Boo struct {
	Name string
}

func (b Boo) Hello() { _ = "STUB: not implemented"; return }

func inCall(foo fii) { _ = "STUB: not implemented"; return }

func main() {
	boo := Boo{"foo"}
	inCall(boo)
}
