package main

func main() {
	type foo struct {
		yolo string
	}

	type foo struct{}
	var bar foo
	println(bar)
}
