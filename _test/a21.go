package main

import "fmt"

func main() {
	a := []byte("hello")
	fmt.Println(a)
	a = append(a, '=')
	fmt.Println(a)
}
