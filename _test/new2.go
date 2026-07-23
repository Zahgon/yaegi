package main

func f() interface{} { _ = "STUB: not implemented"; return nil }

func main() {
	a := f()
	println(*(a.(*int)))
}
