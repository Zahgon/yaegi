package main

func send(c chan<- bool) { _ = "STUB: not implemented"; return }

func main() {
	channel := make(chan bool)
	go send(channel)
	if <-channel {
		println("ok")
	} else {
		println("nok")
	}
}
