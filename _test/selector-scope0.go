package main

import (
	"fmt"
	"time"
)

func test(time string, t time.Time) string { _ = "STUB: not implemented"; return "" }

var zero = time.Time{}

func test2(time string) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func main() {
	str := test("test", time.Now())
	fmt.Println(str)

	str2 := test2("test2")
	fmt.Println(str2)
}
