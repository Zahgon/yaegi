package main

import "fmt"

type CustomError string

func (s CustomError) Error() string { _ = "STUB: not implemented"; return "" }

func NewCustomError(errorText string) CustomError {
	_ = "STUB: not implemented"
	return *new(CustomError)
}

func fail() (err error) { _ = "STUB: not implemented"; return nil }

func main() {
	fmt.Println(fail())
	var myError error
	myError = NewCustomError("ok")
	fmt.Println(myError)
}
