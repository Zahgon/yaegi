package main

import "fmt"

func f1(m map[string]string) { _ = "STUB: not implemented"; return }

func main() {
	m := map[string]string{
		"foo": "bar",
		"baz": "bat",
	}
	f1(m)

	fmt.Println(m)
}
