package main

import "time"

func forever() { _ = "STUB: not implemented"; return }

func main() {
	go forever()
	time.Sleep(1e4)
	println("bye")
}
