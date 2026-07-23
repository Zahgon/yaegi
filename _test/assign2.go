package main

func main() {
	r := uint32(2000000000)
	r = hello(r)
	println(r)
}

func hello(r uint32) uint32 { _ = "STUB: not implemented"; return 0 }
