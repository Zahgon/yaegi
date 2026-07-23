package main

func main() {
	for i := 1; i <= 2; i++ {
		var x, y int
		println(x, y)
		x, y = i, 2*i
		println(x, y)
	}
}
