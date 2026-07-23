//go:build go1.21 && !go1.22
// +build go1.21,!go1.22

package stdlib

import (
	"container/heap"
	"reflect"
)

func init() {
	Symbols["container/heap/heap"] = map[string]reflect.Value{

		"Fix":    reflect.ValueOf(heap.Fix),
		"Init":   reflect.ValueOf(heap.Init),
		"Pop":    reflect.ValueOf(heap.Pop),
		"Push":   reflect.ValueOf(heap.Push),
		"Remove": reflect.ValueOf(heap.Remove),

		"Interface": reflect.ValueOf((*heap.Interface)(nil)),

		"_Interface": reflect.ValueOf((*_container_heap_Interface)(nil)),
	}
}

type _container_heap_Interface struct {
	IValue interface{}
	WLen   func() int
	WLess  func(i int, j int) bool
	WPop   func() any
	WPush  func(x any)
	WSwap  func(i int, j int)
}

func (W _container_heap_Interface) Len() int               { _ = "STUB: not implemented"; return 0 }
func (W _container_heap_Interface) Less(i int, j int) bool { _ = "STUB: not implemented"; return false }
func (W _container_heap_Interface) Pop() any               { _ = "STUB: not implemented"; return *new(any) }
func (W _container_heap_Interface) Push(x any)             { _ = "STUB: not implemented"; return }
func (W _container_heap_Interface) Swap(i int, j int)      { _ = "STUB: not implemented"; return }
