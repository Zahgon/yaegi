package main

type List struct {
	Next *List
	Num  int
}

func add(l *List, n int) *List { _ = "STUB: not implemented"; return nil }

func pr(l *List) { _ = "STUB: not implemented"; return }

func main() {
	a := add(nil, 0)
	pr(a)
	a = add(a, 1)
	pr(a)
	a = add(a, 2)
	pr(a)
}
