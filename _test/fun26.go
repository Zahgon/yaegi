package main

type F func() (int, error)

func f1() (int, error) { _ = "STUB: not implemented"; return 0, nil }

func f2(a string, f F) { _ = "STUB: not implemented"; return }

func main() {
	f2("hello", F(f1))
}
