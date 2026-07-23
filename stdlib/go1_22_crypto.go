//go:build go1.22
// +build go1.22

package stdlib

import (
	"crypto"
	"io"
	"reflect"
)

func init() {
	Symbols["crypto/crypto"] = map[string]reflect.Value{

		"BLAKE2b_256":  reflect.ValueOf(crypto.BLAKE2b_256),
		"BLAKE2b_384":  reflect.ValueOf(crypto.BLAKE2b_384),
		"BLAKE2b_512":  reflect.ValueOf(crypto.BLAKE2b_512),
		"BLAKE2s_256":  reflect.ValueOf(crypto.BLAKE2s_256),
		"MD4":          reflect.ValueOf(crypto.MD4),
		"MD5":          reflect.ValueOf(crypto.MD5),
		"MD5SHA1":      reflect.ValueOf(crypto.MD5SHA1),
		"RIPEMD160":    reflect.ValueOf(crypto.RIPEMD160),
		"RegisterHash": reflect.ValueOf(crypto.RegisterHash),
		"SHA1":         reflect.ValueOf(crypto.SHA1),
		"SHA224":       reflect.ValueOf(crypto.SHA224),
		"SHA256":       reflect.ValueOf(crypto.SHA256),
		"SHA384":       reflect.ValueOf(crypto.SHA384),
		"SHA3_224":     reflect.ValueOf(crypto.SHA3_224),
		"SHA3_256":     reflect.ValueOf(crypto.SHA3_256),
		"SHA3_384":     reflect.ValueOf(crypto.SHA3_384),
		"SHA3_512":     reflect.ValueOf(crypto.SHA3_512),
		"SHA512":       reflect.ValueOf(crypto.SHA512),
		"SHA512_224":   reflect.ValueOf(crypto.SHA512_224),
		"SHA512_256":   reflect.ValueOf(crypto.SHA512_256),

		"Decrypter":     reflect.ValueOf((*crypto.Decrypter)(nil)),
		"DecrypterOpts": reflect.ValueOf((*crypto.DecrypterOpts)(nil)),
		"Hash":          reflect.ValueOf((*crypto.Hash)(nil)),
		"PrivateKey":    reflect.ValueOf((*crypto.PrivateKey)(nil)),
		"PublicKey":     reflect.ValueOf((*crypto.PublicKey)(nil)),
		"Signer":        reflect.ValueOf((*crypto.Signer)(nil)),
		"SignerOpts":    reflect.ValueOf((*crypto.SignerOpts)(nil)),

		"_Decrypter":     reflect.ValueOf((*_crypto_Decrypter)(nil)),
		"_DecrypterOpts": reflect.ValueOf((*_crypto_DecrypterOpts)(nil)),
		"_PrivateKey":    reflect.ValueOf((*_crypto_PrivateKey)(nil)),
		"_PublicKey":     reflect.ValueOf((*_crypto_PublicKey)(nil)),
		"_Signer":        reflect.ValueOf((*_crypto_Signer)(nil)),
		"_SignerOpts":    reflect.ValueOf((*_crypto_SignerOpts)(nil)),
	}
}

type _crypto_Decrypter struct {
	IValue   interface{}
	WDecrypt func(rand io.Reader, msg []byte, opts crypto.DecrypterOpts) (plaintext []byte, err error)
	WPublic  func() crypto.PublicKey
}

func (W _crypto_Decrypter) Decrypt(rand io.Reader, msg []byte, opts crypto.DecrypterOpts) (plaintext []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (W _crypto_Decrypter) Public() crypto.PublicKey {
	_ = "STUB: not implemented"
	return *new(crypto.PublicKey)
}

type _crypto_DecrypterOpts struct {
	IValue interface{}
}

type _crypto_PrivateKey struct {
	IValue interface{}
}

type _crypto_PublicKey struct {
	IValue interface{}
}

type _crypto_Signer struct {
	IValue  interface{}
	WPublic func() crypto.PublicKey
	WSign   func(rand io.Reader, digest []byte, opts crypto.SignerOpts) (signature []byte, err error)
}

func (W _crypto_Signer) Public() crypto.PublicKey {
	_ = "STUB: not implemented"
	return *new(crypto.PublicKey)
}
func (W _crypto_Signer) Sign(rand io.Reader, digest []byte, opts crypto.SignerOpts) (signature []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type _crypto_SignerOpts struct {
	IValue    interface{}
	WHashFunc func() crypto.Hash
}

func (W _crypto_SignerOpts) HashFunc() crypto.Hash {
	_ = "STUB: not implemented"
	return *new(crypto.Hash)
}
