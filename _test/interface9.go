package main

import "fmt"

type Int int

func (I Int) String() string { _ = "STUB: not implemented"; return "" }

func main() {
	var i Int
	var st fmt.Stringer = i
	fmt.Println(st.String())
}
