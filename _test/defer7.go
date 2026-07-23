package main

import "fmt"

func f1(in, out []string) { _ = "STUB: not implemented"; return }

func main() {
	in := []string{"foo", "bar"}
	out := make([]string, 2)
	f1(in, out)

	fmt.Println(out)
}
