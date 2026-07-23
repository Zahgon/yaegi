package named3

import (
	"net/http"
)

type T struct {
	A string
}

func (t *T) Print() { _ = "STUB: not implemented"; return }

type A http.Header

func (a A) ForeachKey() error { _ = "STUB: not implemented"; return nil }

func (a A) Set(k string, v []string) { _ = "STUB: not implemented"; return }
