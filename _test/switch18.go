package main

type fii interface {
	Hello()
}

type Bir struct{}

func (b Bir) Yo() { _ = "STUB: not implemented"; return }

func (b Bir) Hello() { _ = "STUB: not implemented"; return }

type Boo struct {
	Name string
}

func (b Boo) Hello() { _ = "STUB: not implemented"; return }

type Bar struct{}

func (b Bar) Hello() { _ = "STUB: not implemented"; return }

func inCall(foo fii) { _ = "STUB: not implemented"; return }

func main() {
	boo := Bir{}
	inCall(boo)
	inCall(Bar{})
}
