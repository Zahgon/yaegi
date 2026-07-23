package main

const BlockSize = 8

type Cipher struct{}

func (c *Cipher) BlockSize() int { _ = "STUB: not implemented"; return 0 }

func main() {
	println(BlockSize)
	s := Cipher{}
	println(s.BlockSize())
}
