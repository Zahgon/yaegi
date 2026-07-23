package main

type fn func(int)

func test(f fn, v int) { _ = "STUB: not implemented"; return }

func main() {
	a := 3
	test(func(i int) { println("f1", i, a) }, 21)
}
