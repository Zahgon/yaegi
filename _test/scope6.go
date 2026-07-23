package main

import "fmt"

var a = [3]int{1, 2, 3}

func f() { _ = "STUB: not implemented"; return }

func main() {
	fmt.Println(a)
	a[1] = 5
	f()
}
