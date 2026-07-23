package main

func adder() func(int) int { _ = "STUB: not implemented"; return nil }

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		println(pos(i), neg(-2*i))
	}
}
