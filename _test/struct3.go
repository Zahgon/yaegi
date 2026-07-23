package main

type T struct {
	f int
	g int
	h struct {
		k int
	}
}

func f(i int) int { _ = "STUB: not implemented"; return 0 }

func main() {
	a := T{}
	a.h.k = f(4)
	println(a.h.k)
}
