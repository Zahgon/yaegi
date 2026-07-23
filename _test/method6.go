package main

type Sample struct {
	Name string
}

func (s Sample) foo(i int) { _ = "STUB: not implemented"; return }

var samples = []Sample{
	Sample{"hello"},
}

func main() {
	samples[0].foo(3)
}
