package main

type Transformer interface {
	Reset()
}

type Encoder struct {
	Transformer
}

type nop struct{}

func (nop) Reset() { _ = "STUB: not implemented"; return }

func f(e Transformer) { _ = "STUB: not implemented"; return }

func main() {
	e := Encoder{Transformer: nop{}}
	f(e)
}
