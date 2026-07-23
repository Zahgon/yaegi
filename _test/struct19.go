package main

import "fmt"

type Config struct {
	Users        `json:"users,omitempty" mapstructure:","`
	UsersFile    string `json:"usersFile,omitempty"`
	Realm        string `json:"realm,omitempty"`
	RemoveHeader bool   `json:"removeHeader,omitempty"`
	HeaderField  string `json:"headerField,omitempty" export:"true"`
}

type Users []string

func CreateConfig() *Config { _ = "STUB: not implemented"; return nil }

func main() {
	c := CreateConfig()
	fmt.Println(c)
}
