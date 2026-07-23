package main

import (
	"fmt"
)

func MapOf[K comparable, V any](m map[K]V) Map[K, V] { _ = "STUB: not implemented"; return nil }

type Map[K comparable, V any] struct {
	ж map[K]V
}

func (v MapView) Int() Map[string, int] { _ = "STUB: not implemented"; return nil }

type VMap struct {
	Int map[string]int
}

type MapView struct {
	ж *VMap
}

func main() {
	mv := MapView{&VMap{}}
	fmt.Println(mv.ж)
}
