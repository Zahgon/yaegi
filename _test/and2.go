package main

func f() bool { _ = "STUB: not implemented"; return false }

func main() {
	var (
		cl = 0
		ct = "some text"
		ce = ""
	)
	if ce == "" && (cl == 0 || cl > 1000) && (ct == "" || f()) {
		println("ok")
	}
}
