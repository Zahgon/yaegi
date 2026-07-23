package main

type foo func(b int)

func boo(b int) { _ = "STUB: not implemented"; return }

func main() {
	var f foo

	f = boo
	f(4)
}
