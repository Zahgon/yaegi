package main

type Root struct {
	Name string
}

type One struct {
	Root
}

type Hi interface {
	Hello() string
}

type Hey interface {
	Hello() string
}

func (r *Root) Hello() string { _ = "STUB: not implemented"; return "" }

func main() {

	var one Hey = &One{Root{Name: "test2"}}
	println(one.(Hi).Hello())
}
