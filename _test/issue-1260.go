package main

import (
	"fmt"
	"io"
	"os"
)

type WriteSyncer interface {
	io.Writer
	Sync() error
}

type Sink interface {
	WriteSyncer
	io.Closer
}

func newFileSink(path string) (Sink, error) { _ = "STUB: not implemented"; return *new(Sink), nil }

type Sink1 struct{ name string }

func (s Sink1) Write(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }
func (s Sink1) Sync() error                 { _ = "STUB: not implemented"; return nil }
func (s Sink1) Close() error                { _ = "STUB: not implemented"; return nil }
func newS1(name string) Sink                { _ = "STUB: not implemented"; return *new(Sink) }
func newS1p(name string) Sink               { _ = "STUB: not implemented"; return *new(Sink) }

type Sink2 struct{ name string }

func (s *Sink2) Write(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }
func (s *Sink2) Sync() error                 { _ = "STUB: not implemented"; return nil }
func (s *Sink2) Close() error                { _ = "STUB: not implemented"; return nil }
func newS2(name string) Sink                 { _ = "STUB: not implemented"; return *new(Sink) }

func main() {
	tmpfile, err := os.CreateTemp("", "xxx")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpfile.Name())
	closers := []io.Closer{}
	sink, err := newFileSink(tmpfile.Name())
	if err != nil {
		panic(err)
	}
	closers = append(closers, sink)

	s1p := newS1p("ptr")
	s1 := newS1("struct")
	s2 := newS2("ptr2")
	closers = append(closers, s1p, s1, s2)
	for _, closer := range closers {
		fmt.Println(closer.Close())
	}
}
