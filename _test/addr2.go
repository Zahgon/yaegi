package main

import (
	"fmt"
)

type Email struct {
	Where string `xml:"where,attr"`
	Addr  string
}

func f(r interface{}) error { _ = "STUB: not implemented"; return nil }

func withPointerAsInterface(r interface{}) error { _ = "STUB: not implemented"; return nil }

func ff(s string, r interface{}) error { _ = "STUB: not implemented"; return nil }

func fff(s string, r interface{}) error { _ = "STUB: not implemented"; return nil }

func main() {
	data := `
		<Email where='work'>
			<Addr>bob@work.com</Addr>
		</Email>
	`
	v := Email{}
	err := f(&v)
	fmt.Println(err, v)

	vv := Email{}
	err = ff(data, &vv)
	fmt.Println(err, vv)

	vvv := Email{}
	err = ff(data, &vvv)
	fmt.Println(err, vvv)
}
