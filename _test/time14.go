package main

import (
	"fmt"
	"time"
)

var t time.Time

func f() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func main() {
	fmt.Println(f())
}
