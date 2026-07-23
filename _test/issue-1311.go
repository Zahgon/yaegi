package main

type T struct {
	v interface{}
}

func f() (ret int64, err error) { _ = "STUB: not implemented"; return 0, nil }

func main() {
	t := &T{}
	t.v, _ = f()
	println(t.v.(int64))
}
