package main

import "time"

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {

		time.Sleep(1e9)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2e9)
		c2 <- "two"
	}()

	msg1 := <-c1
	println(msg1)

	msg2 := <-c2
	println(msg2)
}
