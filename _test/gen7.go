package main

func MapKeys[K comparable, V any](m map[K]V) []K { _ = "STUB: not implemented"; return nil }

func main() {
	var m = map[int]string{1: "2", 2: "4", 4: "8"}

	println(len(MapKeys))
}
