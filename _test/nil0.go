package main

import "fmt"

func f() (host, port string, err error) { _ = "STUB: not implemented"; return "", "", nil }

func main() {
	h, p, err := f()
	fmt.Println(h, p, err)
}
