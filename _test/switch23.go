package main

func getType() string { _ = "STUB: not implemented"; return "" }

func main() {
	switch getType() {
	case "T1":
		println("T1")
	default:
		println("default")
	}
}
