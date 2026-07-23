package main

func isSeparator(c byte) bool { _ = "STUB: not implemented"; return false }

func main() {
	s := "max-age=20"
	for _, c := range []byte(s) {
		println(string(c), isSeparator(c))
	}
}
