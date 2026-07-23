//go:build go1.22
// +build go1.22

package stdlib

import (
	"crypto/ecdh"
	"io"
	"reflect"
)

func init() {
	Symbols["crypto/ecdh/ecdh"] = map[string]reflect.Value{

		"P256":   reflect.ValueOf(ecdh.P256),
		"P384":   reflect.ValueOf(ecdh.P384),
		"P521":   reflect.ValueOf(ecdh.P521),
		"X25519": reflect.ValueOf(ecdh.X25519),

		"Curve":      reflect.ValueOf((*ecdh.Curve)(nil)),
		"PrivateKey": reflect.ValueOf((*ecdh.PrivateKey)(nil)),
		"PublicKey":  reflect.ValueOf((*ecdh.PublicKey)(nil)),

		"_Curve": reflect.ValueOf((*_crypto_ecdh_Curve)(nil)),
	}
}

type _crypto_ecdh_Curve struct {
	IValue         interface{}
	WGenerateKey   func(rand io.Reader) (*ecdh.PrivateKey, error)
	WNewPrivateKey func(key []byte) (*ecdh.PrivateKey, error)
	WNewPublicKey  func(key []byte) (*ecdh.PublicKey, error)
}

func (W _crypto_ecdh_Curve) GenerateKey(rand io.Reader) (*ecdh.PrivateKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (W _crypto_ecdh_Curve) NewPrivateKey(key []byte) (*ecdh.PrivateKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (W _crypto_ecdh_Curve) NewPublicKey(key []byte) (*ecdh.PublicKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
