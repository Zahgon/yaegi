package main

type Kind int

const (
	None Kind = 0
	Left Kind = 1 << iota
	Right
	Both Kind = Left | Right
)

func main() {
	println(None, Left, Right, Both)
}
