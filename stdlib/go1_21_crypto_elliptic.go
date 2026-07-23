//go:build go1.21 && !go1.22
// +build go1.21,!go1.22

package stdlib

import (
	"crypto/elliptic"
	"math/big"
	"reflect"
)

func init() {
	Symbols["crypto/elliptic/elliptic"] = map[string]reflect.Value{

		"GenerateKey":         reflect.ValueOf(elliptic.GenerateKey),
		"Marshal":             reflect.ValueOf(elliptic.Marshal),
		"MarshalCompressed":   reflect.ValueOf(elliptic.MarshalCompressed),
		"P224":                reflect.ValueOf(elliptic.P224),
		"P256":                reflect.ValueOf(elliptic.P256),
		"P384":                reflect.ValueOf(elliptic.P384),
		"P521":                reflect.ValueOf(elliptic.P521),
		"Unmarshal":           reflect.ValueOf(elliptic.Unmarshal),
		"UnmarshalCompressed": reflect.ValueOf(elliptic.UnmarshalCompressed),

		"Curve":       reflect.ValueOf((*elliptic.Curve)(nil)),
		"CurveParams": reflect.ValueOf((*elliptic.CurveParams)(nil)),

		"_Curve": reflect.ValueOf((*_crypto_elliptic_Curve)(nil)),
	}
}

type _crypto_elliptic_Curve struct {
	IValue          interface{}
	WAdd            func(x1 *big.Int, y1 *big.Int, x2 *big.Int, y2 *big.Int) (x *big.Int, y *big.Int)
	WDouble         func(x1 *big.Int, y1 *big.Int) (x *big.Int, y *big.Int)
	WIsOnCurve      func(x *big.Int, y *big.Int) bool
	WParams         func() *elliptic.CurveParams
	WScalarBaseMult func(k []byte) (x *big.Int, y *big.Int)
	WScalarMult     func(x1 *big.Int, y1 *big.Int, k []byte) (x *big.Int, y *big.Int)
}

func (W _crypto_elliptic_Curve) Add(x1 *big.Int, y1 *big.Int, x2 *big.Int, y2 *big.Int) (x *big.Int, y *big.Int) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (W _crypto_elliptic_Curve) Double(x1 *big.Int, y1 *big.Int) (x *big.Int, y *big.Int) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (W _crypto_elliptic_Curve) IsOnCurve(x *big.Int, y *big.Int) bool {
	_ = "STUB: not implemented"
	return false
}
func (W _crypto_elliptic_Curve) Params() *elliptic.CurveParams {
	_ = "STUB: not implemented"
	return nil
}
func (W _crypto_elliptic_Curve) ScalarBaseMult(k []byte) (x *big.Int, y *big.Int) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (W _crypto_elliptic_Curve) ScalarMult(x1 *big.Int, y1 *big.Int, k []byte) (x *big.Int, y *big.Int) {
	_ = "STUB: not implemented"
	return nil, nil
}
