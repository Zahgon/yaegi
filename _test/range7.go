package main

import (
	"fmt"
)

func someChan() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func main() {
	for _ = range someChan() {
		fmt.Println("success")
		return
	}
}
