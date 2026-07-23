package main

import "fmt"

func f() (int, int) { _ = "STUB: not implemented"; return 0, 0 }

func main() {
	fmt.Println(f())
}
