package main

type adder func(int, int) int

func genAdd(k int) adder { _ = "STUB: not implemented"; return *new(adder) }

func main() {
	f := genAdd(5)
	println(f(3, 4))
}
