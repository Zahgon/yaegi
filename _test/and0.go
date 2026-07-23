package main

func main() {
	a, b := 1, 2

	if f2() && f1() {
		println(a, b)
	}
}

func f1() bool { _ = "STUB: not implemented"; return false }

func f2() bool { _ = "STUB: not implemented"; return false }
