package main

func main() {
	a := f2() && f1()
	println(a)
}

func f1() bool { _ = "STUB: not implemented"; return false }

func f2() bool { _ = "STUB: not implemented"; return false }
