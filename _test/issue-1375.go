package main

type Option func(*Struct)

func WithOption(opt string) Option { _ = "STUB: not implemented"; return *new(Option) }

type Struct struct {
	opt string
}

func New(opts ...Option) *Struct { _ = "STUB: not implemented"; return nil }

func (s *Struct) ShowOption() { _ = "STUB: not implemented"; return }

func main() {
	opts := []Option{
		WithOption("test"),
	}
	s := New(opts...)
	s.ShowOption()
}
