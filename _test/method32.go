package main

func main() {
	var a = []func(string){bar}
	b := a[0]
	b("bar")
}

func bar(a string) { _ = "STUB: not implemented"; return }
