package main

type fn func(int)

func test(f fn, v int) { _ = "STUB: not implemented"; return }

func main() {
	f1 := func(i int) { println("f1", i) }
	test(f1, 21)
}
