package main

type T struct {
	Name string
}

func f(t interface{}) { _ = "STUB: not implemented"; return }

func main() {
	f(&T{"truc"})
}
