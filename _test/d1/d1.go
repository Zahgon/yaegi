package d1

type T struct {
	Name string
}

func (t *T) F() { _ = "STUB: not implemented"; return }

func NewT(s string) *T { _ = "STUB: not implemented"; return nil }
