package main

type T struct{}

func (t *T) Error() string { _ = "STUB: not implemented"; return "" }
func (*T) Foo()            { _ = "STUB: not implemented"; return }

var invalidT = &T{}

func main() {
	var err error
	if err != invalidT {
		println("ok")
	}
}
