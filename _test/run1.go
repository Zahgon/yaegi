package main

func f() (int, int) { _ = "STUB: not implemented"; return 0, 0 }

func g(i, j int) int { _ = "STUB: not implemented"; return 0 }

func main() {
	println(g(f()))
}
