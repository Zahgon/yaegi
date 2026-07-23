package main

type fii interface {
	Hello()
}

type Boo struct {
	Name string
}

type Bar struct{}

func (b Bar) Hello() { _ = "STUB: not implemented"; return }

func (b Boo) Hello() { _ = "STUB: not implemented"; return }

func inCall(foo fii) { _ = "STUB: not implemented"; return }

func main() {
	boo := Boo{"foo"}
	inCall(boo)
	inCall(Bar{})
}
