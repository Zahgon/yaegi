package p6

import (
	"net/netip"
)

type Slice[T any] struct {
	x []T
}

type IPPrefixSlice struct {
	x Slice[netip.Prefix]
}

func (v Slice[T]) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (v IPPrefixSlice) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
