package main

type T1 struct{}

func (t1 T1) f() { _ = "STUB: not implemented"; return }

func (t1 T1) g() { _ = "STUB: not implemented"; return }

type T2 struct {
	T1
}

func (t2 T2) f() { _ = "STUB: not implemented"; return }

type I interface {
	f()
}

func printType(i I) { _ = "STUB: not implemented"; return }

func main() {
	println("T1")
	printType(T1{})
	println("T2")
	printType(T2{})
}
