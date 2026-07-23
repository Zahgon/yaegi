package main

type T struct {
	f int
	g int
}

func f(i int) int { _ = "STUB: not implemented"; return 0 }

func main() {
	a := T{7, f(4)}
	println(a.f, a.g)
}
