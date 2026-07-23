package main

func f1(a int) int { _ = "STUB: not implemented"; return 0 }

func f2(a int) interface{} { _ = "STUB: not implemented"; return nil }

func main() {
	c := f2(3)
	println(c.(int))
}
