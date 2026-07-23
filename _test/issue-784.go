package main

import "fmt"

type Filter interface {
	Bounds(srcBounds string) (dstBounds string)
}

type GIFT struct {
	Filters []Filter
}

func New(filters ...Filter) *GIFT { _ = "STUB: not implemented"; return nil }

func (g *GIFT) Bounds(srcBounds string) (dstBounds string) { _ = "STUB: not implemented"; return "" }

func main() {
	var filters []Filter
	bounds := "foo"
	g := New(filters...)
	fmt.Println(g.Bounds(bounds))
}
