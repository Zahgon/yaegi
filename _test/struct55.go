package main

import (
	"log"
	"os"
)

type Logger struct {
	m []*log.Logger
}

func (l *Logger) Infof(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

func main() {
	l := &Logger{m: []*log.Logger{log.New(os.Stdout, "", log.Lmsgprefix)}}
	l.Infof("test %s", "test")
}
