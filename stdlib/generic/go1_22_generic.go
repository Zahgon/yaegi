//go:build go1.22
// +build go1.22

package generic

import _ "embed"

//go:embed go1_22_cmp_cmp.go.txt
var cmpSource string

//go:embed go1_22_maps_maps.go.txt
var mapsSource string

//go:embed go1_22_slices_slices.go.txt
var slicesSource string

var Sources = [...]string{
	cmpSource,
	mapsSource,
	slicesSource,
}
