package main

type Channel chan string

func send(c Channel) { _ = "STUB: not implemented"; return }

func main() {
	channel := make(Channel)
	go send(channel)
	msg := <-channel
	println(msg)
}
