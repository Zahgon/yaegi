package main

func f(i *int) { _ = "STUB: not implemented"; return }

func main() {
	var a int = 2
	f(&a)
	println(a)
}
