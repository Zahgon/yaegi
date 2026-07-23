package main

import "fmt"

func SumInts(m map[string]int64) int64 { _ = "STUB: not implemented"; return 0 }

func SumFloats(m map[string]float64) float64 { _ = "STUB: not implemented"; return 0 }

func main() {

	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		SumInts(ints),
		SumFloats(floats))
}
