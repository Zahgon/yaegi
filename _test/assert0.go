package main

import (
	"fmt"
	"time"
)

type MyWriter interface {
	Write(p []byte) (i int, err error)
}

type DummyWriter interface {
	Write(p []byte) (i int, err error)
}

type TestStruct struct{}

func (t TestStruct) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func usesWriter(w MyWriter) { _ = "STUB: not implemented"; return }

type MyStringer interface {
	String() string
}

type DummyStringer interface {
	String() string
}

func usesStringer(s MyStringer) { _ = "STUB: not implemented"; return }

func main() {

	var t DummyWriter
	t = TestStruct{}
	var tw MyWriter
	var ok bool
	tw, ok = t.(MyWriter)
	if !ok {
		fmt.Println("TestStruct does not implement MyWriter")
	} else {
		fmt.Println("TestStruct implements MyWriter")
		usesWriter(tw)
	}
	n, _ := t.(MyWriter).Write([]byte("hello world"))
	fmt.Println(n)

	if _, ok := t.(MyWriter); !ok {
		fmt.Println("TestStruct does not implement MyWriter")
		return
	} else {
		fmt.Println("TestStruct implements MyWriter")
	}

	var tt DummyStringer
	tt = time.Nanosecond
	var myD MyStringer
	myD, ok = tt.(MyStringer)
	if !ok {
		fmt.Println("time.Nanosecond does not implement MyStringer")
	} else {
		fmt.Println("time.Nanosecond implements MyStringer")
		usesStringer(myD)
	}
	fmt.Println(tt.(MyStringer).String())

	if _, ok := tt.(MyStringer); !ok {
		fmt.Println("time.Nanosecond does not implement MyStringer")
	} else {
		fmt.Println("time.Nanosecond implements MyStringer")
	}

}
