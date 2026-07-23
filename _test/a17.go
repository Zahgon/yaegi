package main

import "fmt"

func main() {
	a := make([]int, 2, 7)
	fmt.Println(a, len(a), cap(a))
}
