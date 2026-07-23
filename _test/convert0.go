package main

type T struct {
	v int
}

type comparator func(T, T) bool

func sort(items []T, comp comparator) { _ = "STUB: not implemented"; return }

func compT(t0, t1 T) bool { _ = "STUB: not implemented"; return false }

func main() {
	a := []T{}
	sort(a, comparator(compT))
}
