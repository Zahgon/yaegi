package main

func getArray() ([]int, error) { _ = "STUB: not implemented"; return nil, nil }

func getNum() (int, error) { _ = "STUB: not implemented"; return 0, nil }

func main() {
	if a, err := getNum(); err != nil {
		println("#1", a)
	} else if a, err := getArray(); err != nil {
		println("#2", a)
	}
	println("#3")
}
