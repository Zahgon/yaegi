package main

type foobar struct {
	callback func(string) func()
}

func cb(text string) func() { _ = "STUB: not implemented"; return nil }

func main() {

	cb("Hi from inline callback!")()

	asVarTest1 := cb("Hi from asVarTest1 callback!")
	asVarTest1()

	asVarTest2 := cb
	asVarTest2("Hi from asVarTest2 callback!")()

	asStructField := &foobar{callback: cb}
	asStructField.callback("Hi from struct field callback!")()
}
