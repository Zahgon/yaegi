package main

import (
	"fmt"
)

func md5Crypt(password, salt, magic []byte) []byte { _ = "STUB: not implemented"; return nil }

func main() {
	b := md5Crypt([]byte("1"), []byte("2"), []byte("3"))

	fmt.Println(b)
}
