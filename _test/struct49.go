package main

type S struct {
	ts map[string][]*T
}

type T struct {
	s *S
}

func (c *S) getT(addr string) (t *T, ok bool) { _ = "STUB: not implemented"; return nil, false }

func main() {
	s := &S{
		ts: map[string][]*T{},
	}
	s.ts["test"] = append(s.ts["test"], &T{s: s})

	t, ok := s.getT("test")
	println(t != nil, ok)
}
