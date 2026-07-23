package main

import "fmt"

var (
	t *S
	_ I = t
	_ J = t
)

type S struct {
	Name string
}

func (s *S) F() int { _ = "STUB: not implemented"; return 0 }
func (s *S) G() int { _ = "STUB: not implemented"; return 0 }
func (s *S) Ri() I  { _ = "STUB: not implemented"; return *new(I) }
func (s *S) Rj() J  { _ = "STUB: not implemented"; return *new(J) }

type J interface {
	I
	G() int
	Rj() J
}

type I interface {
	F() int
	Ri() I
}

func main() {
	var j J
	fmt.Println(j)
}
