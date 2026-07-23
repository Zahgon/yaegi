package main

type Coord struct {
	x, y int
}

func (c Coord) dist() int { _ = "STUB: not implemented"; return 0 }

type Point struct {
	Coord
	z int
}

type Tpoint struct {
	t int
	Point
}

func main() {
	o := Tpoint{0, Point{Coord{3, 4}, 5}}
	println(o.dist())
}
