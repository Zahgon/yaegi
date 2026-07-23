package main

type T int

func (t T) Error() string { _ = "STUB: not implemented"; return "" }

var invalidT T

func main() {
	var err error
	if err > invalidT {
		println("ok")
	}
}
