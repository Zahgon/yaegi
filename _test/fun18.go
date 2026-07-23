package main

var m = map[string]int{"foo": 1, "bar": 2}

func f(s string) interface{} { _ = "STUB: not implemented"; return nil }

func main() {
	println(f("foo").(int))
}
