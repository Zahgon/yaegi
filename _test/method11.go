package main

func BlockSize() string { _ = "STUB: not implemented"; return "" }

type Cipher struct{}

func (c *Cipher) BlockSize() string { _ = "STUB: not implemented"; return "" }

func main() {
	println(BlockSize())
	s := Cipher{}
	println(s.BlockSize())
}
