package main

import "fmt"

func main() {
	a, b := 1, 1
	for i := 0; i < 10; i++ {
		fmt.Println(a)
		a, b = b, a+b
	}
}
