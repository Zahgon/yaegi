package main

import (
	"fmt"
)

type Config struct {
	A string
}

var conf *Config = &Config{}

func SetConfig() func(*Config) { _ = "STUB: not implemented"; return nil }

func main() {
	conf := &Config{
		A: "foo",
	}

	fmt.Println(conf.A)
}
