package main

import "fmt"

func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	_ = "STUB: not implemented"
	return *new(V)
}

func main() {

	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats[string, int64](ints),
		SumIntsOrFloats[string, float64](floats))
}
