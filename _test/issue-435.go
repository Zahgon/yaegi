package main

type Foo int

func (f Foo) String() string { _ = "STUB: not implemented"; return "" }

func print1(arg interface{}) { _ = "STUB: not implemented"; return }

func main() {
	var arg Foo = 3
	var f = print1
	f(arg)
}
