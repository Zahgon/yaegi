package main

type Map[K comparable, V any] struct {
	ж map[K]V
}

func (m Map[K, V]) Has(k K) bool { _ = "STUB: not implemented"; return false }

func main() {
	m := Map[string, float64]{}
	println(m.Has("test"))
}
