package main

func send(c chan<- string) { _ = "STUB: not implemented"; return }

func main() {
	channel := make(chan string)
	go send(channel)
	msg := <-channel
	println(msg)
}
