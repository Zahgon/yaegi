package main

import "fmt"

func main() {
	var buf [8]byte
	for i := 0; i < 2; i++ {
		for i := range buf {
			buf[i] += byte(i)
		}
		fmt.Println(buf)
	}
}
