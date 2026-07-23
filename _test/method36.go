package main

type I interface{ Hello() }

type T struct{ Name string }

func (t *T) Hello() { _ = "STUB: not implemented"; return }

type FT func(i I)

type ST struct{ Handler FT }

func newF() FT { _ = "STUB: not implemented"; return *new(FT) }

func main() {
	st := &ST{}
	st.Handler = newF()
	st.Handler(&T{"test"})
}
