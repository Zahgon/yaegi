package main

type sample struct {
	count int
}

func run(inf interface{}, name string) { _ = "STUB: not implemented"; return }

func main() {
	a := sample{2}
	println(a.count)
	run(a, "truc")
}
