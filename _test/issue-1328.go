package main

import (
	"crypto/sha1"
	"encoding/hex"
)

func main() {
	script := "hello"
	sumRaw := sha1.Sum([]byte(script))
	sum := hex.EncodeToString(sumRaw[:])
	println(sum)
}
