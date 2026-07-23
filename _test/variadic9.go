package main

import "fmt"

func Sprintf(format string, a ...interface{}) string { _ = "STUB: not implemented"; return "" }

func main() {
	fmt.Println(Sprintf("Hello %s", "World!"))
}
