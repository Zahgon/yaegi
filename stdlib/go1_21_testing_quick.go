//go:build go1.21 && !go1.22
// +build go1.21,!go1.22

package stdlib

import (
	"math/rand"
	"reflect"
	"testing/quick"
)

func init() {
	Symbols["testing/quick/quick"] = map[string]reflect.Value{

		"Check":      reflect.ValueOf(quick.Check),
		"CheckEqual": reflect.ValueOf(quick.CheckEqual),
		"Value":      reflect.ValueOf(quick.Value),

		"CheckEqualError": reflect.ValueOf((*quick.CheckEqualError)(nil)),
		"CheckError":      reflect.ValueOf((*quick.CheckError)(nil)),
		"Config":          reflect.ValueOf((*quick.Config)(nil)),
		"Generator":       reflect.ValueOf((*quick.Generator)(nil)),
		"SetupError":      reflect.ValueOf((*quick.SetupError)(nil)),

		"_Generator": reflect.ValueOf((*_testing_quick_Generator)(nil)),
	}
}

type _testing_quick_Generator struct {
	IValue    interface{}
	WGenerate func(rand *rand.Rand, size int) reflect.Value
}

func (W _testing_quick_Generator) Generate(rand *rand.Rand, size int) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}
