package main

func main() {
	o := Coord{3, 4}
	println(o.dist())
}

func (c *Coord) dist() int { _ = "STUB: not implemented"; return 0 }

type Coord struct {
	x, y int
}
