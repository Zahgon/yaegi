package main

import "fmt"

func f() (string, error) { _ = "STUB: not implemented"; return "", nil }

func main() {
	_, err := f()
	if err != nil {
		fmt.Println(err.Error())
	}
}
