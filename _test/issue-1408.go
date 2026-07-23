package main

type (
	Number  = int32
	Number2 = Number
)

func f(n Number2) { _ = "STUB: not implemented"; return }

func main() {
	var n Number = 5
	f(n)
}
