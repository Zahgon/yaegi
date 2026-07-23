package main

var (
	a = concat("hello", b)
	b = concat(" ", c, "!")
	c = d
	d = "world"
)

func concat(a ...string) string { _ = "STUB: not implemented"; return "" }

func main() {
	println(a)
}
