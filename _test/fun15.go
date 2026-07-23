package main

func f1(a int) interface{} { _ = "STUB: not implemented"; return nil }

func f2(a int64) interface{} { _ = "STUB: not implemented"; return nil }

func main() {
	c := f1(3)
	println(c.(int))
	b := f2(3)
	println(b.(int64))
}
