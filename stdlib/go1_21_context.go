//go:build go1.21 && !go1.22
// +build go1.21,!go1.22

package stdlib

import (
	"context"
	"reflect"
	"time"
)

func init() {
	Symbols["context/context"] = map[string]reflect.Value{

		"AfterFunc":         reflect.ValueOf(context.AfterFunc),
		"Background":        reflect.ValueOf(context.Background),
		"Canceled":          reflect.ValueOf(&context.Canceled).Elem(),
		"Cause":             reflect.ValueOf(context.Cause),
		"DeadlineExceeded":  reflect.ValueOf(&context.DeadlineExceeded).Elem(),
		"TODO":              reflect.ValueOf(context.TODO),
		"WithCancel":        reflect.ValueOf(context.WithCancel),
		"WithCancelCause":   reflect.ValueOf(context.WithCancelCause),
		"WithDeadline":      reflect.ValueOf(context.WithDeadline),
		"WithDeadlineCause": reflect.ValueOf(context.WithDeadlineCause),
		"WithTimeout":       reflect.ValueOf(context.WithTimeout),
		"WithTimeoutCause":  reflect.ValueOf(context.WithTimeoutCause),
		"WithValue":         reflect.ValueOf(context.WithValue),
		"WithoutCancel":     reflect.ValueOf(context.WithoutCancel),

		"CancelCauseFunc": reflect.ValueOf((*context.CancelCauseFunc)(nil)),
		"CancelFunc":      reflect.ValueOf((*context.CancelFunc)(nil)),
		"Context":         reflect.ValueOf((*context.Context)(nil)),

		"_Context": reflect.ValueOf((*_context_Context)(nil)),
	}
}

type _context_Context struct {
	IValue    interface{}
	WDeadline func() (deadline time.Time, ok bool)
	WDone     func() <-chan struct{}
	WErr      func() error
	WValue    func(key any) any
}

func (W _context_Context) Deadline() (deadline time.Time, ok bool) {
	_ = "STUB: not implemented"
	return *new(time.Time), false
}
func (W _context_Context) Done() <-chan struct{} { _ = "STUB: not implemented"; return nil }
func (W _context_Context) Err() error            { _ = "STUB: not implemented"; return nil }
func (W _context_Context) Value(key any) any     { _ = "STUB: not implemented"; return *new(any) }
