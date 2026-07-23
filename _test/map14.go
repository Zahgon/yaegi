package main

import "fmt"

var m = map[string]float64{"foo": 1.0}

func f(s string) bool { _ = "STUB: not implemented"; return false }

func main() {
	fmt.Println(f("foo"), f("bar"))
}
