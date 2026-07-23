package main

import "fmt"

func f1(ch chan string) { _ = "STUB: not implemented"; return }

func main() {
	ch := make(chan string, 1)
	f1(ch)

	for s := range ch {
		fmt.Println(s)
	}
}
