package main

type Float interface {
	~float32 | ~float64
}

func add[T Float](a, b T) float64 { _ = "STUB: not implemented"; return 0 }

func main() {
	println(add(1, 2))
}
