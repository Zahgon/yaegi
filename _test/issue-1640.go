package main

func ShortVariableDeclarations() (i int, err error) { _ = "STUB: not implemented"; return 0, nil }

func main() {
	_, er := ShortVariableDeclarations()
	if er != nil {
		println("ShortVariableDeclarations ok")
	} else {
		println("ShortVariableDeclarations not ok")
	}
}
