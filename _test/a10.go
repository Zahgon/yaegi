package main

type Sample struct {
	Name string
}

var samples = []Sample{}

func f(i int) { _ = "STUB: not implemented"; return }

func main() {
	samples = append(samples, Sample{Name: "test"})
	f(0)
}
