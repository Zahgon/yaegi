package main

func f(a []int) interface{} { _ = "STUB: not implemented"; return nil }

func g(a []int) int { _ = "STUB: not implemented"; return 0 }

func main() {
	a := []int{1, 2}
	println(g(a))
	println(f(a).(int))
}
