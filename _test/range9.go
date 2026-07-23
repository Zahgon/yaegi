package main

func main() {
	var c chan<- struct{} = make(chan struct{})

	for _ = range c {
	}
}
