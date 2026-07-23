package main

import (
	"encoding/json"
	"fmt"
)

type MyJsonMarshaler struct{ n int }

func (m MyJsonMarshaler) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func main() {
	ch := make(chan json.Marshaler, 1)
	ch <- MyJsonMarshaler{2}
	m, err := json.Marshal(<-ch)
	fmt.Println(string(m), err)
}
