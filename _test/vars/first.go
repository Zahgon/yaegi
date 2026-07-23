package vars

var (
	A = concat("hello", B)
	C = D
)

func concat(a ...string) string { _ = "STUB: not implemented"; return "" }
