package main

type Enabler interface {
	Enabled() bool
}

type Logger struct {
	core Enabler
}

func (log *Logger) GetCore() Enabler { _ = "STUB: not implemented"; return *new(Enabler) }

type T struct{}

func (t *T) Enabled() bool { _ = "STUB: not implemented"; return false }

func main() {
	base := &Logger{&T{}}
	println(base.GetCore().Enabled())
}
