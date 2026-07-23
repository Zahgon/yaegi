package main

type Sample struct {
	Name string
	Foo  []string
}

func (s *Sample) foo(j int) { _ = "STUB: not implemented"; return }

var samples = []Sample{
	Sample{"hello", []string{"world"}},
}

func main() {
	samples[0].foo(3)
}
