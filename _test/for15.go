package main

func f() int { _ = "STUB: not implemented"; return 0 }

func main() {
	for i := f(); ; {
		println("in loop")
		if i > 0 {
			break
		}
	}
}
