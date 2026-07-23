package main

func main() {
	var foo int
	foo = 2

	type foo struct{}
	var bar foo
	println(bar)
}
