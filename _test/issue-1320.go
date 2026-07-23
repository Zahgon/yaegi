package main

type Pooler interface {
	Get() string
}

type baseClient struct {
	connPool Pooler
}

type connPool struct {
	name string
}

func (c *connPool) Get() string { _ = "STUB: not implemented"; return "" }

func newBaseClient(i int, p Pooler) *baseClient { _ = "STUB: not implemented"; return nil }

func newConnPool() *connPool { _ = "STUB: not implemented"; return nil }

func main() {
	b := newBaseClient(0, newConnPool())
	println(b.connPool.(*connPool).name)
}
