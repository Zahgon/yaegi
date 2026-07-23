//go:build go1.21 && !go1.22
// +build go1.21,!go1.22

package stdlib

import (
	"crypto/cipher"
	"reflect"
)

func init() {
	Symbols["crypto/cipher/cipher"] = map[string]reflect.Value{

		"NewCBCDecrypter":     reflect.ValueOf(cipher.NewCBCDecrypter),
		"NewCBCEncrypter":     reflect.ValueOf(cipher.NewCBCEncrypter),
		"NewCFBDecrypter":     reflect.ValueOf(cipher.NewCFBDecrypter),
		"NewCFBEncrypter":     reflect.ValueOf(cipher.NewCFBEncrypter),
		"NewCTR":              reflect.ValueOf(cipher.NewCTR),
		"NewGCM":              reflect.ValueOf(cipher.NewGCM),
		"NewGCMWithNonceSize": reflect.ValueOf(cipher.NewGCMWithNonceSize),
		"NewGCMWithTagSize":   reflect.ValueOf(cipher.NewGCMWithTagSize),
		"NewOFB":              reflect.ValueOf(cipher.NewOFB),

		"AEAD":         reflect.ValueOf((*cipher.AEAD)(nil)),
		"Block":        reflect.ValueOf((*cipher.Block)(nil)),
		"BlockMode":    reflect.ValueOf((*cipher.BlockMode)(nil)),
		"Stream":       reflect.ValueOf((*cipher.Stream)(nil)),
		"StreamReader": reflect.ValueOf((*cipher.StreamReader)(nil)),
		"StreamWriter": reflect.ValueOf((*cipher.StreamWriter)(nil)),

		"_AEAD":      reflect.ValueOf((*_crypto_cipher_AEAD)(nil)),
		"_Block":     reflect.ValueOf((*_crypto_cipher_Block)(nil)),
		"_BlockMode": reflect.ValueOf((*_crypto_cipher_BlockMode)(nil)),
		"_Stream":    reflect.ValueOf((*_crypto_cipher_Stream)(nil)),
	}
}

type _crypto_cipher_AEAD struct {
	IValue     interface{}
	WNonceSize func() int
	WOpen      func(dst []byte, nonce []byte, ciphertext []byte, additionalData []byte) ([]byte, error)
	WOverhead  func() int
	WSeal      func(dst []byte, nonce []byte, plaintext []byte, additionalData []byte) []byte
}

func (W _crypto_cipher_AEAD) NonceSize() int { _ = "STUB: not implemented"; return 0 }
func (W _crypto_cipher_AEAD) Open(dst []byte, nonce []byte, ciphertext []byte, additionalData []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (W _crypto_cipher_AEAD) Overhead() int { _ = "STUB: not implemented"; return 0 }
func (W _crypto_cipher_AEAD) Seal(dst []byte, nonce []byte, plaintext []byte, additionalData []byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

type _crypto_cipher_Block struct {
	IValue     interface{}
	WBlockSize func() int
	WDecrypt   func(dst []byte, src []byte)
	WEncrypt   func(dst []byte, src []byte)
}

func (W _crypto_cipher_Block) BlockSize() int                 { _ = "STUB: not implemented"; return 0 }
func (W _crypto_cipher_Block) Decrypt(dst []byte, src []byte) { _ = "STUB: not implemented"; return }
func (W _crypto_cipher_Block) Encrypt(dst []byte, src []byte) { _ = "STUB: not implemented"; return }

type _crypto_cipher_BlockMode struct {
	IValue       interface{}
	WBlockSize   func() int
	WCryptBlocks func(dst []byte, src []byte)
}

func (W _crypto_cipher_BlockMode) BlockSize() int { _ = "STUB: not implemented"; return 0 }
func (W _crypto_cipher_BlockMode) CryptBlocks(dst []byte, src []byte) {
	_ = "STUB: not implemented"
	return
}

type _crypto_cipher_Stream struct {
	IValue        interface{}
	WXORKeyStream func(dst []byte, src []byte)
}

func (W _crypto_cipher_Stream) XORKeyStream(dst []byte, src []byte) {
	_ = "STUB: not implemented"
	return
}
