package stdlib

import (
	"errors"
	"io"
	"log"
	"os"
)

var errRestricted = errors.New("restricted")

func osExit(code int) { _ = "STUB: not implemented"; return }

func osFindProcess(pid int) (*os.Process, error) { _ = "STUB: not implemented"; return nil, nil }

func logFatal(v ...interface{})            { _ = "STUB: not implemented"; return }
func logFatalf(f string, v ...interface{}) { _ = "STUB: not implemented"; return }
func logFatalln(v ...interface{})          { _ = "STUB: not implemented"; return }

type logLogger struct {
	l *log.Logger
}

func logNew(out io.Writer, prefix string, flag int) *logLogger {
	_ = "STUB: not implemented"
	return nil
}

func (l *logLogger) Fatal(v ...interface{})            { _ = "STUB: not implemented"; return }
func (l *logLogger) Fatalf(f string, v ...interface{}) { _ = "STUB: not implemented"; return }
func (l *logLogger) Fatalln(v ...interface{})          { _ = "STUB: not implemented"; return }

func (l *logLogger) Flags() int                        { _ = "STUB: not implemented"; return 0 }
func (l *logLogger) Output(d int, s string) error      { _ = "STUB: not implemented"; return nil }
func (l *logLogger) Panic(v ...interface{})            { _ = "STUB: not implemented"; return }
func (l *logLogger) Panicf(f string, v ...interface{}) { _ = "STUB: not implemented"; return }
func (l *logLogger) Panicln(v ...interface{})          { _ = "STUB: not implemented"; return }
func (l *logLogger) Prefix() string                    { _ = "STUB: not implemented"; return "" }
func (l *logLogger) Print(v ...interface{})            { _ = "STUB: not implemented"; return }
func (l *logLogger) Printf(f string, v ...interface{}) { _ = "STUB: not implemented"; return }
func (l *logLogger) Println(v ...interface{})          { _ = "STUB: not implemented"; return }
func (l *logLogger) SetFlags(flag int)                 { _ = "STUB: not implemented"; return }
func (l *logLogger) SetOutput(w io.Writer)             { _ = "STUB: not implemented"; return }
func (l *logLogger) Writer() io.Writer                 { _ = "STUB: not implemented"; return *new(io.Writer) }
