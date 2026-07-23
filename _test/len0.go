package main

func f(a []int) interface{} { _ = "STUB: not implemented"; return nil }

func main() {
	a := []int{1, 2}
	println(f(a).(int))
}
