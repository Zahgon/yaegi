package main

func f1(a int) interface{} { _ = "STUB: not implemented"; return nil }

func f2(a int) interface{} { _ = "STUB: not implemented"; return nil }

func main() {
	c := f2(3)
	println(c.(int))
}
