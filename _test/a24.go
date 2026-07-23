package main

import "fmt"

func main() {
	for i := 0; i < 2; i++ {
		var buf [8]byte
		var x int
		fmt.Println(buf, x)
		for i := range buf {
			buf[i] = byte(i)
			x = i
		}
		fmt.Println(buf, x)
	}
}
