package main

import (
	"time"
)

type MyTime struct {
	time.Time
	index int
}

func (m MyTime) Foo() { _ = "STUB: not implemented"; return }

func (m *MyTime) Bar() { _ = "STUB: not implemented"; return }

func main() {
	t := MyTime{}
	t.Time = time.Date(2009, time.November, 10, 23, 4, 5, 0, time.UTC)
	t.Foo()
	t.Bar()
	(&t).Bar()
}
