package main

func genFunc() (f func()) { _ = "STUB: not implemented"; return nil }

func main() {
	println(genFunc() == nil)
}
