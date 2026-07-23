package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	d := sha1.New()
	d.Write([]byte("password"))
	a := d.Sum(nil)
	fmt.Println(a)
}
