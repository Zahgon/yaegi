package main

import "fmt"

func (t MyT) Test() string { _ = "STUB: not implemented"; return "" }

type MyT int

func main() {
	t := MyT(1)

	fmt.Println(t.Test())
}
