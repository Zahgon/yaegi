package main

type Number interface {
	int | int64 | ~float64
}

func Sum[T Number](numbers []T) T { _ = "STUB: not implemented"; return *new(T) }

func main() {
	xs := []int{3, 5, 10}
	total := Sum(xs)
	println(total)
}
