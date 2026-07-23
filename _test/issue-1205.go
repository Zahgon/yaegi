package main

type Option interface {
	apply()
}

func f(opts ...Option) { _ = "STUB: not implemented"; return }

type T struct{}

func (t *T) apply() { _ = "STUB: not implemented"; return }

func main() {
	opt := []Option{&T{}}
	f(opt[0])
	f(opt...)
}
