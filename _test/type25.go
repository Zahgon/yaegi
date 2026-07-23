package main

import (
	"errors"
	"sync/atomic"
)

type wrappedError struct {
	wrapped error
}

func (e wrappedError) Error() string { _ = "STUB: not implemented"; return "" }

func (e wrappedError) Unwrap() error { _ = "STUB: not implemented"; return nil }

var err atomic.Value

func getWrapped() *wrappedError { _ = "STUB: not implemented"; return nil }

func main() {
	err.Store(wrappedError{wrapped: errors.New("test")})

	e := getWrapped()
	if e != nil {
		println(e.Error())
		println(e.wrapped.Error())
	}
}
