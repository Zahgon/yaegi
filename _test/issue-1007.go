package main

type TypeA struct {
	B TypeB
}

type TypeB struct {
	C1 *TypeC
	C2 *TypeC
}

type TypeC struct {
	Val string
	D   *TypeD
	D2  *TypeD
}

type TypeD struct {
	Name string
}

func build() *TypeA { _ = "STUB: not implemented"; return nil }

func Bar(s string) string { _ = "STUB: not implemented"; return "" }

func main() {
	println(Bar("test"))
}
