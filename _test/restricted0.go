package main

import (
	"fmt"
	"log"
)

func main() {
	defer func() {
		r := recover()
		fmt.Println("recover:", r)
	}()
	log.Fatal("log.Fatal does not exit")
	println("not printed")
}
