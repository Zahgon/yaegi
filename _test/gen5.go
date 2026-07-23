package main

import "fmt"

type Set[Elem comparable] struct {
	m map[Elem]struct{}
}

func Make[Elem comparable]() Set[Elem] { _ = "STUB: not implemented"; return nil }

func (s Set[Elem]) Add(v Elem) { _ = "STUB: not implemented"; return }

func main() {
	s := Make[int]()
	s.Add(1)
	fmt.Println(s)
}
