package main

import (
	"fmt"
	"net"
)

type ipNetValue net.IPNet

func (ipnet *ipNetValue) Set(value string) error { _ = "STUB: not implemented"; return nil }

func main() {
	v := ipNetValue{}
	fmt.Println(v)
}
