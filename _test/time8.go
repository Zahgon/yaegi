package main

import (
	"time"
)

type durationValue time.Duration

func (d *durationValue) String() string { _ = "STUB: not implemented"; return "" }

func main() {
	var d durationValue
	println(d.String())
}
