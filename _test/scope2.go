package main

var a int = 1

func f() { _ = "STUB: not implemented"; return }

func main() {
	println(a)
	a := 2
	println(a)
	f()
}
