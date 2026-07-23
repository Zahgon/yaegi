//go:build go1.21 && !go1.22
// +build go1.21,!go1.22

package stdlib

import (
	"net/smtp"
	"reflect"
)

func init() {
	Symbols["net/smtp/smtp"] = map[string]reflect.Value{

		"CRAMMD5Auth": reflect.ValueOf(smtp.CRAMMD5Auth),
		"Dial":        reflect.ValueOf(smtp.Dial),
		"NewClient":   reflect.ValueOf(smtp.NewClient),
		"PlainAuth":   reflect.ValueOf(smtp.PlainAuth),
		"SendMail":    reflect.ValueOf(smtp.SendMail),

		"Auth":       reflect.ValueOf((*smtp.Auth)(nil)),
		"Client":     reflect.ValueOf((*smtp.Client)(nil)),
		"ServerInfo": reflect.ValueOf((*smtp.ServerInfo)(nil)),

		"_Auth": reflect.ValueOf((*_net_smtp_Auth)(nil)),
	}
}

type _net_smtp_Auth struct {
	IValue interface{}
	WNext  func(fromServer []byte, more bool) (toServer []byte, err error)
	WStart func(server *smtp.ServerInfo) (proto string, toServer []byte, err error)
}

func (W _net_smtp_Auth) Next(fromServer []byte, more bool) (toServer []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (W _net_smtp_Auth) Start(server *smtp.ServerInfo) (proto string, toServer []byte, err error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}
