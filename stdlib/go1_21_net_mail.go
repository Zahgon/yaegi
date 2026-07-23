//go:build go1.21 && !go1.22
// +build go1.21,!go1.22

package stdlib

import (
	"net/mail"
	"reflect"
)

func init() {
	Symbols["net/mail/mail"] = map[string]reflect.Value{

		"ErrHeaderNotPresent": reflect.ValueOf(&mail.ErrHeaderNotPresent).Elem(),
		"ParseAddress":        reflect.ValueOf(mail.ParseAddress),
		"ParseAddressList":    reflect.ValueOf(mail.ParseAddressList),
		"ParseDate":           reflect.ValueOf(mail.ParseDate),
		"ReadMessage":         reflect.ValueOf(mail.ReadMessage),

		"Address":       reflect.ValueOf((*mail.Address)(nil)),
		"AddressParser": reflect.ValueOf((*mail.AddressParser)(nil)),
		"Header":        reflect.ValueOf((*mail.Header)(nil)),
		"Message":       reflect.ValueOf((*mail.Message)(nil)),
	}
}
