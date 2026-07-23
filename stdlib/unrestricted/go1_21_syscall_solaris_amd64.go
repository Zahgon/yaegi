//go:build go1.21 && !go1.22
// +build go1.21,!go1.22

package unrestricted

import (
	"reflect"
	"syscall"
)

func init() {
	Symbols["syscall/syscall"] = map[string]reflect.Value{

		"Exec":         reflect.ValueOf(syscall.Exec),
		"Exit":         reflect.ValueOf(syscall.Exit),
		"ForkExec":     reflect.ValueOf(syscall.ForkExec),
		"Kill":         reflect.ValueOf(syscall.Kill),
		"RawSyscall":   reflect.ValueOf(syscall.RawSyscall),
		"Shutdown":     reflect.ValueOf(syscall.Shutdown),
		"StartProcess": reflect.ValueOf(syscall.StartProcess),
		"Syscall":      reflect.ValueOf(syscall.Syscall),
	}
}
