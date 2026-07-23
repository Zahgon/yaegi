package main

func main() {
	if a := f(); a > 0 {
		println(a)
	}
}

func f() int { _ = "STUB: not implemented"; return 0 }
