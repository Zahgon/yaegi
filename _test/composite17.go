package main

import (
	"html/template"
)

var str = `{{ stringOr .Data "test" }}`

func main() {
	_, err := template.New("test").
		Funcs(template.FuncMap{
			"stringOr": stringOr,
		}).
		Parse(str)
	if err != nil {
		println(err.Error())
		return
	}
	println("success")
}

func stringOr(v, def string) string { _ = "STUB: not implemented"; return "" }
