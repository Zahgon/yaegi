package main

func main() {
	var foo struct {
		yolo string
	}

	type foo struct{}
	var bar foo
	println(bar)
}
