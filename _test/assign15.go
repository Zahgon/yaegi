package main

func main() {
	var c chan<- struct{} = make(chan struct{})
	var d <-chan struct{} = c

	_ = d
}
