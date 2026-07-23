package main

type Root struct {
	Name string
}

func (r *Root) Hello() { _ = "STUB: not implemented"; return }

type One = Root

func main() {
	one := &One{Name: "one"}
	displayOne(one)
	displayRoot(one)

	root := &Root{Name: "root"}
	displayOne(root)
	displayRoot(root)
}

func displayOne(val *One) { _ = "STUB: not implemented"; return }

func displayRoot(val *Root) { _ = "STUB: not implemented"; return }
