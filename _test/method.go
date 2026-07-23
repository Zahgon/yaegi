package main

type Coord struct {
	x, y int
}

func (c Coord) dist() int { _ = "STUB: not implemented"; return 0 }

func main() {
	o := Coord{3, 4}
	println(o.dist())
}
