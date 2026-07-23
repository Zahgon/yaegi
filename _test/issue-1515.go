package main

type I1 interface {
	I2
	Wrap() *S3
}

type I2 interface {
	F()
}

type S2 struct {
	I2
}

func newS2(i2 I2) I1 { _ = "STUB: not implemented"; return *new(I1) }

type S3 struct {
	base *S2
}

func (s *S2) Wrap() *S3 { _ = "STUB: not implemented"; return nil }

type T struct {
	name string
}

func (t *T) F() { _ = "STUB: not implemented"; return }

func main() {
	t := &T{"test"}
	s2 := newS2(t)
	s3 := s2.Wrap()
	s3.base.F()
}
