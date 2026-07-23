package main

func test() error { _ = "STUB: not implemented"; return nil }

func main() {
	if err := test(); nil == err {
		println("err is nil")
	}
}
