package main

type S struct {
	q chan struct{}
}

func (s *S) Send() { _ = "STUB: not implemented"; return }

func main() {
	s := &S{q: make(chan struct{}, 1)}
	s.Send()
	println("bye")
}
