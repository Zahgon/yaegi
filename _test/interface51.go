package main

type Error interface {
	error
	Message() string
}

type T struct {
	Msg string
}

func (t *T) Error() string   { _ = "STUB: not implemented"; return "" }
func (t *T) Message() string { _ = "STUB: not implemented"; return "" }

func newError() Error { _ = "STUB: not implemented"; return *new(Error) }

func main() {
	e := newError()
	println(e.Error())
}
