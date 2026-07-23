package main

func b() string { _ = "STUB: not implemented"; return "" }

func main() {
	var x int
	x = "a" + b()
}
