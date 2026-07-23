//go:build go1.21 && !go1.22
// +build go1.21,!go1.22

package stdlib

import (
	"encoding/xml"
	"go/constant"
	"go/token"
	"reflect"
)

func init() {
	Symbols["encoding/xml/xml"] = map[string]reflect.Value{

		"CopyToken":       reflect.ValueOf(xml.CopyToken),
		"Escape":          reflect.ValueOf(xml.Escape),
		"EscapeText":      reflect.ValueOf(xml.EscapeText),
		"HTMLAutoClose":   reflect.ValueOf(&xml.HTMLAutoClose).Elem(),
		"HTMLEntity":      reflect.ValueOf(&xml.HTMLEntity).Elem(),
		"Header":          reflect.ValueOf(constant.MakeFromLiteral("\"<?xml version=\\\"1.0\\\" encoding=\\\"UTF-8\\\"?>\\n\"", token.STRING, 0)),
		"Marshal":         reflect.ValueOf(xml.Marshal),
		"MarshalIndent":   reflect.ValueOf(xml.MarshalIndent),
		"NewDecoder":      reflect.ValueOf(xml.NewDecoder),
		"NewEncoder":      reflect.ValueOf(xml.NewEncoder),
		"NewTokenDecoder": reflect.ValueOf(xml.NewTokenDecoder),
		"Unmarshal":       reflect.ValueOf(xml.Unmarshal),

		"Attr":                 reflect.ValueOf((*xml.Attr)(nil)),
		"CharData":             reflect.ValueOf((*xml.CharData)(nil)),
		"Comment":              reflect.ValueOf((*xml.Comment)(nil)),
		"Decoder":              reflect.ValueOf((*xml.Decoder)(nil)),
		"Directive":            reflect.ValueOf((*xml.Directive)(nil)),
		"Encoder":              reflect.ValueOf((*xml.Encoder)(nil)),
		"EndElement":           reflect.ValueOf((*xml.EndElement)(nil)),
		"Marshaler":            reflect.ValueOf((*xml.Marshaler)(nil)),
		"MarshalerAttr":        reflect.ValueOf((*xml.MarshalerAttr)(nil)),
		"Name":                 reflect.ValueOf((*xml.Name)(nil)),
		"ProcInst":             reflect.ValueOf((*xml.ProcInst)(nil)),
		"StartElement":         reflect.ValueOf((*xml.StartElement)(nil)),
		"SyntaxError":          reflect.ValueOf((*xml.SyntaxError)(nil)),
		"TagPathError":         reflect.ValueOf((*xml.TagPathError)(nil)),
		"Token":                reflect.ValueOf((*xml.Token)(nil)),
		"TokenReader":          reflect.ValueOf((*xml.TokenReader)(nil)),
		"UnmarshalError":       reflect.ValueOf((*xml.UnmarshalError)(nil)),
		"Unmarshaler":          reflect.ValueOf((*xml.Unmarshaler)(nil)),
		"UnmarshalerAttr":      reflect.ValueOf((*xml.UnmarshalerAttr)(nil)),
		"UnsupportedTypeError": reflect.ValueOf((*xml.UnsupportedTypeError)(nil)),

		"_Marshaler":       reflect.ValueOf((*_encoding_xml_Marshaler)(nil)),
		"_MarshalerAttr":   reflect.ValueOf((*_encoding_xml_MarshalerAttr)(nil)),
		"_Token":           reflect.ValueOf((*_encoding_xml_Token)(nil)),
		"_TokenReader":     reflect.ValueOf((*_encoding_xml_TokenReader)(nil)),
		"_Unmarshaler":     reflect.ValueOf((*_encoding_xml_Unmarshaler)(nil)),
		"_UnmarshalerAttr": reflect.ValueOf((*_encoding_xml_UnmarshalerAttr)(nil)),
	}
}

type _encoding_xml_Marshaler struct {
	IValue      interface{}
	WMarshalXML func(e *xml.Encoder, start xml.StartElement) error
}

func (W _encoding_xml_Marshaler) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	_ = "STUB: not implemented"
	return nil
}

type _encoding_xml_MarshalerAttr struct {
	IValue          interface{}
	WMarshalXMLAttr func(name xml.Name) (xml.Attr, error)
}

func (W _encoding_xml_MarshalerAttr) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	_ = "STUB: not implemented"
	return *new(xml.Attr), nil
}

type _encoding_xml_Token struct {
	IValue interface{}
}

type _encoding_xml_TokenReader struct {
	IValue interface{}
	WToken func() (xml.Token, error)
}

func (W _encoding_xml_TokenReader) Token() (xml.Token, error) {
	_ = "STUB: not implemented"
	return *new(xml.Token), nil
}

type _encoding_xml_Unmarshaler struct {
	IValue        interface{}
	WUnmarshalXML func(d *xml.Decoder, start xml.StartElement) error
}

func (W _encoding_xml_Unmarshaler) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	_ = "STUB: not implemented"
	return nil
}

type _encoding_xml_UnmarshalerAttr struct {
	IValue            interface{}
	WUnmarshalXMLAttr func(attr xml.Attr) error
}

func (W _encoding_xml_UnmarshalerAttr) UnmarshalXMLAttr(attr xml.Attr) error {
	_ = "STUB: not implemented"
	return nil
}
