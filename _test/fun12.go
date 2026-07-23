package main

func use(interface{}) { _ = "STUB: not implemented"; return }

func main() {
	z := map[string]interface{}{"a": 5}
	use(z)
	println("bye")
}
