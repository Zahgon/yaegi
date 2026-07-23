package main

import "fmt"

func check() (result bool, err error) { _ = "STUB: not implemented"; return false, nil }

func main() {
	result, error := check()
	fmt.Println(result, error)
}
