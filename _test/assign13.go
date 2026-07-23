package main

import "fmt"

func getStr() string { _ = "STUB: not implemented"; return "" }

func main() {
	m := make(map[string]string, 0)
	m["a"] = fmt.Sprintf("%v", 0.1)
	m["b"] = string(fmt.Sprintf("%v", 0.1))
	m["c"] = getStr()

	fmt.Println(m)
}
