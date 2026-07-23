package interp

import (
	"io/fs"
)

type realFS struct{}

func (dir realFS) Open(name string) (fs.File, error) {
	_ = "STUB: not implemented"
	return *new(fs.File), nil
}
