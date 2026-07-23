//go:build go1.21 && !go1.22
// +build go1.21,!go1.22

package stdlib

import (
	"fmt"
	"reflect"
)

func init() {
	Symbols["fmt/fmt"] = map[string]reflect.Value{

		"Append":       reflect.ValueOf(fmt.Append),
		"Appendf":      reflect.ValueOf(fmt.Appendf),
		"Appendln":     reflect.ValueOf(fmt.Appendln),
		"Errorf":       reflect.ValueOf(fmt.Errorf),
		"FormatString": reflect.ValueOf(fmt.FormatString),
		"Fprint":       reflect.ValueOf(fmt.Fprint),
		"Fprintf":      reflect.ValueOf(fmt.Fprintf),
		"Fprintln":     reflect.ValueOf(fmt.Fprintln),
		"Fscan":        reflect.ValueOf(fmt.Fscan),
		"Fscanf":       reflect.ValueOf(fmt.Fscanf),
		"Fscanln":      reflect.ValueOf(fmt.Fscanln),
		"Print":        reflect.ValueOf(fmt.Print),
		"Printf":       reflect.ValueOf(fmt.Printf),
		"Println":      reflect.ValueOf(fmt.Println),
		"Scan":         reflect.ValueOf(fmt.Scan),
		"Scanf":        reflect.ValueOf(fmt.Scanf),
		"Scanln":       reflect.ValueOf(fmt.Scanln),
		"Sprint":       reflect.ValueOf(fmt.Sprint),
		"Sprintf":      reflect.ValueOf(fmt.Sprintf),
		"Sprintln":     reflect.ValueOf(fmt.Sprintln),
		"Sscan":        reflect.ValueOf(fmt.Sscan),
		"Sscanf":       reflect.ValueOf(fmt.Sscanf),
		"Sscanln":      reflect.ValueOf(fmt.Sscanln),

		"Formatter":  reflect.ValueOf((*fmt.Formatter)(nil)),
		"GoStringer": reflect.ValueOf((*fmt.GoStringer)(nil)),
		"ScanState":  reflect.ValueOf((*fmt.ScanState)(nil)),
		"Scanner":    reflect.ValueOf((*fmt.Scanner)(nil)),
		"State":      reflect.ValueOf((*fmt.State)(nil)),
		"Stringer":   reflect.ValueOf((*fmt.Stringer)(nil)),

		"_Formatter":  reflect.ValueOf((*_fmt_Formatter)(nil)),
		"_GoStringer": reflect.ValueOf((*_fmt_GoStringer)(nil)),
		"_ScanState":  reflect.ValueOf((*_fmt_ScanState)(nil)),
		"_Scanner":    reflect.ValueOf((*_fmt_Scanner)(nil)),
		"_State":      reflect.ValueOf((*_fmt_State)(nil)),
		"_Stringer":   reflect.ValueOf((*_fmt_Stringer)(nil)),
	}
}

type _fmt_Formatter struct {
	IValue  interface{}
	WFormat func(f fmt.State, verb rune)
}

func (W _fmt_Formatter) Format(f fmt.State, verb rune) { _ = "STUB: not implemented"; return }

type _fmt_GoStringer struct {
	IValue    interface{}
	WGoString func() string
}

func (W _fmt_GoStringer) GoString() string { _ = "STUB: not implemented"; return "" }

type _fmt_ScanState struct {
	IValue      interface{}
	WRead       func(buf []byte) (n int, err error)
	WReadRune   func() (r rune, size int, err error)
	WSkipSpace  func()
	WToken      func(skipSpace bool, f func(rune) bool) (token []byte, err error)
	WUnreadRune func() error
	WWidth      func() (wid int, ok bool)
}

func (W _fmt_ScanState) Read(buf []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}
func (W _fmt_ScanState) ReadRune() (r rune, size int, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}
func (W _fmt_ScanState) SkipSpace() { _ = "STUB: not implemented"; return }
func (W _fmt_ScanState) Token(skipSpace bool, f func(rune) bool) (token []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (W _fmt_ScanState) UnreadRune() error         { _ = "STUB: not implemented"; return nil }
func (W _fmt_ScanState) Width() (wid int, ok bool) { _ = "STUB: not implemented"; return 0, false }

type _fmt_Scanner struct {
	IValue interface{}
	WScan  func(state fmt.ScanState, verb rune) error
}

func (W _fmt_Scanner) Scan(state fmt.ScanState, verb rune) error {
	_ = "STUB: not implemented"
	return nil
}

type _fmt_State struct {
	IValue     interface{}
	WFlag      func(c int) bool
	WPrecision func() (prec int, ok bool)
	WWidth     func() (wid int, ok bool)
	WWrite     func(b []byte) (n int, err error)
}

func (W _fmt_State) Flag(c int) bool                   { _ = "STUB: not implemented"; return false }
func (W _fmt_State) Precision() (prec int, ok bool)    { _ = "STUB: not implemented"; return 0, false }
func (W _fmt_State) Width() (wid int, ok bool)         { _ = "STUB: not implemented"; return 0, false }
func (W _fmt_State) Write(b []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

type _fmt_Stringer struct {
	IValue  interface{}
	WString func() string
}

func (W _fmt_Stringer) String() string { _ = "STUB: not implemented"; return "" }
