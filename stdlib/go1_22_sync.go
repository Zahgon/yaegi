//go:build go1.22
// +build go1.22

package stdlib

import (
	"reflect"
	"sync"
)

func init() {
	Symbols["sync/sync"] = map[string]reflect.Value{

		"NewCond":  reflect.ValueOf(sync.NewCond),
		"OnceFunc": reflect.ValueOf(sync.OnceFunc),

		"Cond":      reflect.ValueOf((*sync.Cond)(nil)),
		"Locker":    reflect.ValueOf((*sync.Locker)(nil)),
		"Map":       reflect.ValueOf((*sync.Map)(nil)),
		"Mutex":     reflect.ValueOf((*sync.Mutex)(nil)),
		"Once":      reflect.ValueOf((*sync.Once)(nil)),
		"Pool":      reflect.ValueOf((*sync.Pool)(nil)),
		"RWMutex":   reflect.ValueOf((*sync.RWMutex)(nil)),
		"WaitGroup": reflect.ValueOf((*sync.WaitGroup)(nil)),

		"_Locker": reflect.ValueOf((*_sync_Locker)(nil)),
	}
}

type _sync_Locker struct {
	IValue  interface{}
	WLock   func()
	WUnlock func()
}

func (W _sync_Locker) Lock()   { _ = "STUB: not implemented"; return }
func (W _sync_Locker) Unlock() { _ = "STUB: not implemented"; return }
