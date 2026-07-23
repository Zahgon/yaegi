package main

type I2 interface {
	I2() string
}

type I interface {
	I2
}

type S struct{}

func (*S) I2() string { _ = "STUB: not implemented"; return "" }

func main() {
	var i I
	_, ok := i.(*S)
	println(ok)
}
