package main

var db dbWrapper

type dbWrapper struct {
	DB *cmap
}

func (d *dbWrapper) get() *cmap { _ = "STUB: not implemented"; return nil }

type cmap struct {
	name string
}

func (c *cmap) f() { _ = "STUB: not implemented"; return }

func main() {
	db.DB = &cmap{name: "test"}
	db.get().f()
}
