package main

type T [2]int

func F0(t *T) int { _ = "STUB: not implemented"; return 0 }

func main() {
	t := &T{1, 2}
	println(F0(t))
}
