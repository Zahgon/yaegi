package main

type Channel chan string

type T struct {
	Channel
}

func send(c Channel) { _ = "STUB: not implemented"; return }

func main() {
	t := &T{}
	t.Channel = make(Channel)
	go send(t.Channel)
	msg := <-t.Channel
	println(msg)
}
