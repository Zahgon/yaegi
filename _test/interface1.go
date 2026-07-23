package main

import "fmt"

type fii interface {
	Hello()
}

type Boo struct {
	Name string
}

func (b *Boo) Hello() { _ = "STUB: not implemented"; return }

func inCall(foo fii) { _ = "STUB: not implemented"; return }

func main() {
	fmt.Println("in")
	boo := &Boo{"foo"}
	inCall(boo)
}
