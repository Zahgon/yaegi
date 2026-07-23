package main

type Sample struct {
	Name string
}

func (s *Sample) foo(i int) { _ = "STUB: not implemented"; return }

func main() {
	sample := Sample{"hello"}
	s := &sample
	s.foo(3)
}
