package main

import (
	"fmt"
)

type Hello struct{}

func (*Hello) Hi() string { _ = "STUB: not implemented"; return "" }

func main() {
	fmt.Println(&Hello{})
}
