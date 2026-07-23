package main

type Coord struct{ x, y int }

func f(c Coord) int { _ = "STUB: not implemented"; return 0 }

func main() {
	c := Coord{3, 4}
	println(f(c))
}
