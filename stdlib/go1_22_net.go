//go:build go1.22
// +build go1.22

package stdlib

import (
	"go/constant"
	"go/token"
	"net"
	"reflect"
	"time"
)

func init() {
	Symbols["net/net"] = map[string]reflect.Value{

		"CIDRMask":                   reflect.ValueOf(net.CIDRMask),
		"DefaultResolver":            reflect.ValueOf(&net.DefaultResolver).Elem(),
		"Dial":                       reflect.ValueOf(net.Dial),
		"DialIP":                     reflect.ValueOf(net.DialIP),
		"DialTCP":                    reflect.ValueOf(net.DialTCP),
		"DialTimeout":                reflect.ValueOf(net.DialTimeout),
		"DialUDP":                    reflect.ValueOf(net.DialUDP),
		"DialUnix":                   reflect.ValueOf(net.DialUnix),
		"ErrClosed":                  reflect.ValueOf(&net.ErrClosed).Elem(),
		"ErrWriteToConnected":        reflect.ValueOf(&net.ErrWriteToConnected).Elem(),
		"FileConn":                   reflect.ValueOf(net.FileConn),
		"FileListener":               reflect.ValueOf(net.FileListener),
		"FilePacketConn":             reflect.ValueOf(net.FilePacketConn),
		"FlagBroadcast":              reflect.ValueOf(net.FlagBroadcast),
		"FlagLoopback":               reflect.ValueOf(net.FlagLoopback),
		"FlagMulticast":              reflect.ValueOf(net.FlagMulticast),
		"FlagPointToPoint":           reflect.ValueOf(net.FlagPointToPoint),
		"FlagRunning":                reflect.ValueOf(net.FlagRunning),
		"FlagUp":                     reflect.ValueOf(net.FlagUp),
		"IPv4":                       reflect.ValueOf(net.IPv4),
		"IPv4Mask":                   reflect.ValueOf(net.IPv4Mask),
		"IPv4allrouter":              reflect.ValueOf(&net.IPv4allrouter).Elem(),
		"IPv4allsys":                 reflect.ValueOf(&net.IPv4allsys).Elem(),
		"IPv4bcast":                  reflect.ValueOf(&net.IPv4bcast).Elem(),
		"IPv4len":                    reflect.ValueOf(constant.MakeFromLiteral("4", token.INT, 0)),
		"IPv4zero":                   reflect.ValueOf(&net.IPv4zero).Elem(),
		"IPv6interfacelocalallnodes": reflect.ValueOf(&net.IPv6interfacelocalallnodes).Elem(),
		"IPv6len":                    reflect.ValueOf(constant.MakeFromLiteral("16", token.INT, 0)),
		"IPv6linklocalallnodes":      reflect.ValueOf(&net.IPv6linklocalallnodes).Elem(),
		"IPv6linklocalallrouters":    reflect.ValueOf(&net.IPv6linklocalallrouters).Elem(),
		"IPv6loopback":               reflect.ValueOf(&net.IPv6loopback).Elem(),
		"IPv6unspecified":            reflect.ValueOf(&net.IPv6unspecified).Elem(),
		"IPv6zero":                   reflect.ValueOf(&net.IPv6zero).Elem(),
		"InterfaceAddrs":             reflect.ValueOf(net.InterfaceAddrs),
		"InterfaceByIndex":           reflect.ValueOf(net.InterfaceByIndex),
		"InterfaceByName":            reflect.ValueOf(net.InterfaceByName),
		"Interfaces":                 reflect.ValueOf(net.Interfaces),
		"JoinHostPort":               reflect.ValueOf(net.JoinHostPort),
		"Listen":                     reflect.ValueOf(net.Listen),
		"ListenIP":                   reflect.ValueOf(net.ListenIP),
		"ListenMulticastUDP":         reflect.ValueOf(net.ListenMulticastUDP),
		"ListenPacket":               reflect.ValueOf(net.ListenPacket),
		"ListenTCP":                  reflect.ValueOf(net.ListenTCP),
		"ListenUDP":                  reflect.ValueOf(net.ListenUDP),
		"ListenUnix":                 reflect.ValueOf(net.ListenUnix),
		"ListenUnixgram":             reflect.ValueOf(net.ListenUnixgram),
		"LookupAddr":                 reflect.ValueOf(net.LookupAddr),
		"LookupCNAME":                reflect.ValueOf(net.LookupCNAME),
		"LookupHost":                 reflect.ValueOf(net.LookupHost),
		"LookupIP":                   reflect.ValueOf(net.LookupIP),
		"LookupMX":                   reflect.ValueOf(net.LookupMX),
		"LookupNS":                   reflect.ValueOf(net.LookupNS),
		"LookupPort":                 reflect.ValueOf(net.LookupPort),
		"LookupSRV":                  reflect.ValueOf(net.LookupSRV),
		"LookupTXT":                  reflect.ValueOf(net.LookupTXT),
		"ParseCIDR":                  reflect.ValueOf(net.ParseCIDR),
		"ParseIP":                    reflect.ValueOf(net.ParseIP),
		"ParseMAC":                   reflect.ValueOf(net.ParseMAC),
		"Pipe":                       reflect.ValueOf(net.Pipe),
		"ResolveIPAddr":              reflect.ValueOf(net.ResolveIPAddr),
		"ResolveTCPAddr":             reflect.ValueOf(net.ResolveTCPAddr),
		"ResolveUDPAddr":             reflect.ValueOf(net.ResolveUDPAddr),
		"ResolveUnixAddr":            reflect.ValueOf(net.ResolveUnixAddr),
		"SplitHostPort":              reflect.ValueOf(net.SplitHostPort),
		"TCPAddrFromAddrPort":        reflect.ValueOf(net.TCPAddrFromAddrPort),
		"UDPAddrFromAddrPort":        reflect.ValueOf(net.UDPAddrFromAddrPort),

		"Addr":                reflect.ValueOf((*net.Addr)(nil)),
		"AddrError":           reflect.ValueOf((*net.AddrError)(nil)),
		"Buffers":             reflect.ValueOf((*net.Buffers)(nil)),
		"Conn":                reflect.ValueOf((*net.Conn)(nil)),
		"DNSConfigError":      reflect.ValueOf((*net.DNSConfigError)(nil)),
		"DNSError":            reflect.ValueOf((*net.DNSError)(nil)),
		"Dialer":              reflect.ValueOf((*net.Dialer)(nil)),
		"Error":               reflect.ValueOf((*net.Error)(nil)),
		"Flags":               reflect.ValueOf((*net.Flags)(nil)),
		"HardwareAddr":        reflect.ValueOf((*net.HardwareAddr)(nil)),
		"IP":                  reflect.ValueOf((*net.IP)(nil)),
		"IPAddr":              reflect.ValueOf((*net.IPAddr)(nil)),
		"IPConn":              reflect.ValueOf((*net.IPConn)(nil)),
		"IPMask":              reflect.ValueOf((*net.IPMask)(nil)),
		"IPNet":               reflect.ValueOf((*net.IPNet)(nil)),
		"Interface":           reflect.ValueOf((*net.Interface)(nil)),
		"InvalidAddrError":    reflect.ValueOf((*net.InvalidAddrError)(nil)),
		"ListenConfig":        reflect.ValueOf((*net.ListenConfig)(nil)),
		"Listener":            reflect.ValueOf((*net.Listener)(nil)),
		"MX":                  reflect.ValueOf((*net.MX)(nil)),
		"NS":                  reflect.ValueOf((*net.NS)(nil)),
		"OpError":             reflect.ValueOf((*net.OpError)(nil)),
		"PacketConn":          reflect.ValueOf((*net.PacketConn)(nil)),
		"ParseError":          reflect.ValueOf((*net.ParseError)(nil)),
		"Resolver":            reflect.ValueOf((*net.Resolver)(nil)),
		"SRV":                 reflect.ValueOf((*net.SRV)(nil)),
		"TCPAddr":             reflect.ValueOf((*net.TCPAddr)(nil)),
		"TCPConn":             reflect.ValueOf((*net.TCPConn)(nil)),
		"TCPListener":         reflect.ValueOf((*net.TCPListener)(nil)),
		"UDPAddr":             reflect.ValueOf((*net.UDPAddr)(nil)),
		"UDPConn":             reflect.ValueOf((*net.UDPConn)(nil)),
		"UnixAddr":            reflect.ValueOf((*net.UnixAddr)(nil)),
		"UnixConn":            reflect.ValueOf((*net.UnixConn)(nil)),
		"UnixListener":        reflect.ValueOf((*net.UnixListener)(nil)),
		"UnknownNetworkError": reflect.ValueOf((*net.UnknownNetworkError)(nil)),

		"_Addr":       reflect.ValueOf((*_net_Addr)(nil)),
		"_Conn":       reflect.ValueOf((*_net_Conn)(nil)),
		"_Error":      reflect.ValueOf((*_net_Error)(nil)),
		"_Listener":   reflect.ValueOf((*_net_Listener)(nil)),
		"_PacketConn": reflect.ValueOf((*_net_PacketConn)(nil)),
	}
}

type _net_Addr struct {
	IValue   interface{}
	WNetwork func() string
	WString  func() string
}

func (W _net_Addr) Network() string { _ = "STUB: not implemented"; return "" }
func (W _net_Addr) String() string  { _ = "STUB: not implemented"; return "" }

type _net_Conn struct {
	IValue            interface{}
	WClose            func() error
	WLocalAddr        func() net.Addr
	WRead             func(b []byte) (n int, err error)
	WRemoteAddr       func() net.Addr
	WSetDeadline      func(t time.Time) error
	WSetReadDeadline  func(t time.Time) error
	WSetWriteDeadline func(t time.Time) error
	WWrite            func(b []byte) (n int, err error)
}

func (W _net_Conn) Close() error                       { _ = "STUB: not implemented"; return nil }
func (W _net_Conn) LocalAddr() net.Addr                { _ = "STUB: not implemented"; return *new(net.Addr) }
func (W _net_Conn) Read(b []byte) (n int, err error)   { _ = "STUB: not implemented"; return 0, nil }
func (W _net_Conn) RemoteAddr() net.Addr               { _ = "STUB: not implemented"; return *new(net.Addr) }
func (W _net_Conn) SetDeadline(t time.Time) error      { _ = "STUB: not implemented"; return nil }
func (W _net_Conn) SetReadDeadline(t time.Time) error  { _ = "STUB: not implemented"; return nil }
func (W _net_Conn) SetWriteDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }
func (W _net_Conn) Write(b []byte) (n int, err error)  { _ = "STUB: not implemented"; return 0, nil }

type _net_Error struct {
	IValue     interface{}
	WError     func() string
	WTemporary func() bool
	WTimeout   func() bool
}

func (W _net_Error) Error() string   { _ = "STUB: not implemented"; return "" }
func (W _net_Error) Temporary() bool { _ = "STUB: not implemented"; return false }
func (W _net_Error) Timeout() bool   { _ = "STUB: not implemented"; return false }

type _net_Listener struct {
	IValue  interface{}
	WAccept func() (net.Conn, error)
	WAddr   func() net.Addr
	WClose  func() error
}

func (W _net_Listener) Accept() (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}
func (W _net_Listener) Addr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }
func (W _net_Listener) Close() error   { _ = "STUB: not implemented"; return nil }

type _net_PacketConn struct {
	IValue            interface{}
	WClose            func() error
	WLocalAddr        func() net.Addr
	WReadFrom         func(p []byte) (n int, addr net.Addr, err error)
	WSetDeadline      func(t time.Time) error
	WSetReadDeadline  func(t time.Time) error
	WSetWriteDeadline func(t time.Time) error
	WWriteTo          func(p []byte, addr net.Addr) (n int, err error)
}

func (W _net_PacketConn) Close() error        { _ = "STUB: not implemented"; return nil }
func (W _net_PacketConn) LocalAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }
func (W _net_PacketConn) ReadFrom(p []byte) (n int, addr net.Addr, err error) {
	_ = "STUB: not implemented"
	return 0, *new(net.Addr), nil
}
func (W _net_PacketConn) SetDeadline(t time.Time) error      { _ = "STUB: not implemented"; return nil }
func (W _net_PacketConn) SetReadDeadline(t time.Time) error  { _ = "STUB: not implemented"; return nil }
func (W _net_PacketConn) SetWriteDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }
func (W _net_PacketConn) WriteTo(p []byte, addr net.Addr) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}
