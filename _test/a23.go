package main

import "fmt"

func main() {
	var key, salt [32]byte
	for i := range key {
		key[i] = byte(i)
		salt[i] = byte(i + 32)
	}
	fmt.Println(key)
	fmt.Println(salt)
}
