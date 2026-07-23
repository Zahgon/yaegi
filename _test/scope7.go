package main

import "fmt"

var a = []int{1, 2, 3}

func f() { _ = "STUB: not implemented"; return }

func main() {
	fmt.Println(a)
	a = []int{6, 7}
	f()
}
