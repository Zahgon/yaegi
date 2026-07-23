package main

type T struct {
	f func(*T)
}

func f1(t *T) { _ = "STUB: not implemented"; return }

func f2(t *T) { _ = "STUB: not implemented"; return }

func main() {
	println("ok")
}
