//go:build go1.22
// +build go1.22

package stdlib

import (
	"reflect"
	"sort"
)

func init() {
	Symbols["sort/sort"] = map[string]reflect.Value{

		"Find":              reflect.ValueOf(sort.Find),
		"Float64s":          reflect.ValueOf(sort.Float64s),
		"Float64sAreSorted": reflect.ValueOf(sort.Float64sAreSorted),
		"Ints":              reflect.ValueOf(sort.Ints),
		"IntsAreSorted":     reflect.ValueOf(sort.IntsAreSorted),
		"IsSorted":          reflect.ValueOf(sort.IsSorted),
		"Reverse":           reflect.ValueOf(sort.Reverse),
		"Search":            reflect.ValueOf(sort.Search),
		"SearchFloat64s":    reflect.ValueOf(sort.SearchFloat64s),
		"SearchInts":        reflect.ValueOf(sort.SearchInts),
		"SearchStrings":     reflect.ValueOf(sort.SearchStrings),
		"Slice":             reflect.ValueOf(sort.Slice),
		"SliceIsSorted":     reflect.ValueOf(sort.SliceIsSorted),
		"SliceStable":       reflect.ValueOf(sort.SliceStable),
		"Sort":              reflect.ValueOf(sort.Sort),
		"Stable":            reflect.ValueOf(sort.Stable),
		"Strings":           reflect.ValueOf(sort.Strings),
		"StringsAreSorted":  reflect.ValueOf(sort.StringsAreSorted),

		"Float64Slice": reflect.ValueOf((*sort.Float64Slice)(nil)),
		"IntSlice":     reflect.ValueOf((*sort.IntSlice)(nil)),
		"Interface":    reflect.ValueOf((*sort.Interface)(nil)),
		"StringSlice":  reflect.ValueOf((*sort.StringSlice)(nil)),

		"_Interface": reflect.ValueOf((*_sort_Interface)(nil)),
	}
}

type _sort_Interface struct {
	IValue interface{}
	WLen   func() int
	WLess  func(i int, j int) bool
	WSwap  func(i int, j int)
}

func (W _sort_Interface) Len() int               { _ = "STUB: not implemented"; return 0 }
func (W _sort_Interface) Less(i int, j int) bool { _ = "STUB: not implemented"; return false }
func (W _sort_Interface) Swap(i int, j int)      { _ = "STUB: not implemented"; return }
