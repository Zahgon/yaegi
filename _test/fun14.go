package main

func f() (bool, int) { _ = "STUB: not implemented"; return false, 0 }

func g() (bool, int) { _ = "STUB: not implemented"; return false, 0 }

func main() {
	b, i := g()
	println(b, i)
}
