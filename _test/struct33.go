package main

type T0 struct {
	name string
}

type lookupFunc func(s string) T0

type T1 struct {
	name string
	info lookupFunc
}

func (t T0) F1() bool { _ = "STUB: not implemented"; return false }

var t0 = T0{"t0"}

func look(s string) T0 { _ = "STUB: not implemented"; return *new(T0) }

var table = []*T1{{
	name: "bar",
	info: look,
},
}

func main() {
	info := table[0].info
	println(info("foo").F1())
}
