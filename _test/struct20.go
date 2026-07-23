package main

type SecretProvider func(user, realm string) string

type BasicAuth struct {
	Realm   string
	Secrets SecretProvider
}

func (a *BasicAuth) CheckAuth() string { _ = "STUB: not implemented"; return "" }

func secretBasic(user, realm string) string { _ = "STUB: not implemented"; return "" }

func main() {
	b := &BasicAuth{"test", secretBasic}
	s := b.CheckAuth()
	println(s)
}
