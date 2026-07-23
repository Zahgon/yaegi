package main

type adder func(int, int) int

func genAdd(k int) adder { _ = "STUB: not implemented"; return *new(adder) }

func main() {
	f := genAdd(5)
	g := genAdd(8)
	println(f(3, 4))
	println(g(3, 4))
}
