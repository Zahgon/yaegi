package main

type Coord struct {
	x, y int
}

func (c Coord) dist() int { _ = "STUB: not implemented"; return 0 }

type Point struct {
	Coord
	z int
}

func main() {
	o := Point{Coord{3, 4}, 5}
	f := o.dist
	println(f())
}
