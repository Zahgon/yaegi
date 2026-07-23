package main

import "fmt"

type Node struct {
	Name  string
	Child map[string]Node
}

func main() {
	a := Node{Name: "hello", Child: map[string]Node{}}
	a.Child["1"] = Node{Name: "world", Child: map[string]Node{}}
	fmt.Println(a)
	a.Child["1"].Child["1"] = Node{Name: "sunshine", Child: map[string]Node{}}
	fmt.Println(a)
}
