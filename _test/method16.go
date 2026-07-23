package main

type Cheese struct {
	property string
}

func (t *Cheese) Hello(param string) { _ = "STUB: not implemented"; return }

func main() {
	(*Cheese).Hello(&Cheese{property: "value"}, "param")
}
