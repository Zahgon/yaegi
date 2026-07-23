package syscall

import "reflect"

var Symbols = map[string]map[string]reflect.Value{}

func init() {
	Symbols["github.com/traefik/yaegi/stdlib/syscall/syscall"] = map[string]reflect.Value{
		"Symbols": reflect.ValueOf(Symbols),
	}
}

//go:generate ../../internal/cmd/extract/extract -exclude=^Exec,Exit,ForkExec,Kill,Ptrace,Reboot,Shutdown,StartProcess,Syscall syscall
