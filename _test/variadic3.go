package main

import "fmt"

func f(a ...int) int { _ = "STUB: not implemented"; return 0 }

func main() {
	fmt.Println(f(1, 2, 3, 4))
}
