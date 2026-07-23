package main

import "fmt"

type T [2]int

func F1(t *T) { _ = "STUB: not implemented"; return }

func main() {
	t := &T{}
	F1(t)
	fmt.Println(t)
}
