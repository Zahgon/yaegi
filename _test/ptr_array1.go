package main

type T [3]int

func F0(t *T) { _ = "STUB: not implemented"; return }

func main() {
	t := &T{1, 2, 3}
	F0(t)
}
