//go:build go1.22
// +build go1.22

package stdlib

import (
	"io/fs"
	"reflect"
	"time"
)

func init() {
	Symbols["io/fs/fs"] = map[string]reflect.Value{

		"ErrClosed":          reflect.ValueOf(&fs.ErrClosed).Elem(),
		"ErrExist":           reflect.ValueOf(&fs.ErrExist).Elem(),
		"ErrInvalid":         reflect.ValueOf(&fs.ErrInvalid).Elem(),
		"ErrNotExist":        reflect.ValueOf(&fs.ErrNotExist).Elem(),
		"ErrPermission":      reflect.ValueOf(&fs.ErrPermission).Elem(),
		"FileInfoToDirEntry": reflect.ValueOf(fs.FileInfoToDirEntry),
		"FormatDirEntry":     reflect.ValueOf(fs.FormatDirEntry),
		"FormatFileInfo":     reflect.ValueOf(fs.FormatFileInfo),
		"Glob":               reflect.ValueOf(fs.Glob),
		"ModeAppend":         reflect.ValueOf(fs.ModeAppend),
		"ModeCharDevice":     reflect.ValueOf(fs.ModeCharDevice),
		"ModeDevice":         reflect.ValueOf(fs.ModeDevice),
		"ModeDir":            reflect.ValueOf(fs.ModeDir),
		"ModeExclusive":      reflect.ValueOf(fs.ModeExclusive),
		"ModeIrregular":      reflect.ValueOf(fs.ModeIrregular),
		"ModeNamedPipe":      reflect.ValueOf(fs.ModeNamedPipe),
		"ModePerm":           reflect.ValueOf(fs.ModePerm),
		"ModeSetgid":         reflect.ValueOf(fs.ModeSetgid),
		"ModeSetuid":         reflect.ValueOf(fs.ModeSetuid),
		"ModeSocket":         reflect.ValueOf(fs.ModeSocket),
		"ModeSticky":         reflect.ValueOf(fs.ModeSticky),
		"ModeSymlink":        reflect.ValueOf(fs.ModeSymlink),
		"ModeTemporary":      reflect.ValueOf(fs.ModeTemporary),
		"ModeType":           reflect.ValueOf(fs.ModeType),
		"ReadDir":            reflect.ValueOf(fs.ReadDir),
		"ReadFile":           reflect.ValueOf(fs.ReadFile),
		"SkipAll":            reflect.ValueOf(&fs.SkipAll).Elem(),
		"SkipDir":            reflect.ValueOf(&fs.SkipDir).Elem(),
		"Stat":               reflect.ValueOf(fs.Stat),
		"Sub":                reflect.ValueOf(fs.Sub),
		"ValidPath":          reflect.ValueOf(fs.ValidPath),
		"WalkDir":            reflect.ValueOf(fs.WalkDir),

		"DirEntry":    reflect.ValueOf((*fs.DirEntry)(nil)),
		"FS":          reflect.ValueOf((*fs.FS)(nil)),
		"File":        reflect.ValueOf((*fs.File)(nil)),
		"FileInfo":    reflect.ValueOf((*fs.FileInfo)(nil)),
		"FileMode":    reflect.ValueOf((*fs.FileMode)(nil)),
		"GlobFS":      reflect.ValueOf((*fs.GlobFS)(nil)),
		"PathError":   reflect.ValueOf((*fs.PathError)(nil)),
		"ReadDirFS":   reflect.ValueOf((*fs.ReadDirFS)(nil)),
		"ReadDirFile": reflect.ValueOf((*fs.ReadDirFile)(nil)),
		"ReadFileFS":  reflect.ValueOf((*fs.ReadFileFS)(nil)),
		"StatFS":      reflect.ValueOf((*fs.StatFS)(nil)),
		"SubFS":       reflect.ValueOf((*fs.SubFS)(nil)),
		"WalkDirFunc": reflect.ValueOf((*fs.WalkDirFunc)(nil)),

		"_DirEntry":    reflect.ValueOf((*_io_fs_DirEntry)(nil)),
		"_FS":          reflect.ValueOf((*_io_fs_FS)(nil)),
		"_File":        reflect.ValueOf((*_io_fs_File)(nil)),
		"_FileInfo":    reflect.ValueOf((*_io_fs_FileInfo)(nil)),
		"_GlobFS":      reflect.ValueOf((*_io_fs_GlobFS)(nil)),
		"_ReadDirFS":   reflect.ValueOf((*_io_fs_ReadDirFS)(nil)),
		"_ReadDirFile": reflect.ValueOf((*_io_fs_ReadDirFile)(nil)),
		"_ReadFileFS":  reflect.ValueOf((*_io_fs_ReadFileFS)(nil)),
		"_StatFS":      reflect.ValueOf((*_io_fs_StatFS)(nil)),
		"_SubFS":       reflect.ValueOf((*_io_fs_SubFS)(nil)),
	}
}

type _io_fs_DirEntry struct {
	IValue interface{}
	WInfo  func() (fs.FileInfo, error)
	WIsDir func() bool
	WName  func() string
	WType  func() fs.FileMode
}

func (W _io_fs_DirEntry) Info() (fs.FileInfo, error) {
	_ = "STUB: not implemented"
	return *new(fs.FileInfo), nil
}
func (W _io_fs_DirEntry) IsDir() bool       { _ = "STUB: not implemented"; return false }
func (W _io_fs_DirEntry) Name() string      { _ = "STUB: not implemented"; return "" }
func (W _io_fs_DirEntry) Type() fs.FileMode { _ = "STUB: not implemented"; return *new(fs.FileMode) }

type _io_fs_FS struct {
	IValue interface{}
	WOpen  func(name string) (fs.File, error)
}

func (W _io_fs_FS) Open(name string) (fs.File, error) {
	_ = "STUB: not implemented"
	return *new(fs.File), nil
}

type _io_fs_File struct {
	IValue interface{}
	WClose func() error
	WRead  func(a0 []byte) (int, error)
	WStat  func() (fs.FileInfo, error)
}

func (W _io_fs_File) Close() error                { _ = "STUB: not implemented"; return nil }
func (W _io_fs_File) Read(a0 []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }
func (W _io_fs_File) Stat() (fs.FileInfo, error) {
	_ = "STUB: not implemented"
	return *new(fs.FileInfo), nil
}

type _io_fs_FileInfo struct {
	IValue   interface{}
	WIsDir   func() bool
	WModTime func() time.Time
	WMode    func() fs.FileMode
	WName    func() string
	WSize    func() int64
	WSys     func() any
}

func (W _io_fs_FileInfo) IsDir() bool        { _ = "STUB: not implemented"; return false }
func (W _io_fs_FileInfo) ModTime() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }
func (W _io_fs_FileInfo) Mode() fs.FileMode  { _ = "STUB: not implemented"; return *new(fs.FileMode) }
func (W _io_fs_FileInfo) Name() string       { _ = "STUB: not implemented"; return "" }
func (W _io_fs_FileInfo) Size() int64        { _ = "STUB: not implemented"; return 0 }
func (W _io_fs_FileInfo) Sys() any           { _ = "STUB: not implemented"; return *new(any) }

type _io_fs_GlobFS struct {
	IValue interface{}
	WGlob  func(pattern string) ([]string, error)
	WOpen  func(name string) (fs.File, error)
}

func (W _io_fs_GlobFS) Glob(pattern string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
func (W _io_fs_GlobFS) Open(name string) (fs.File, error) {
	_ = "STUB: not implemented"
	return *new(fs.File), nil
}

type _io_fs_ReadDirFS struct {
	IValue   interface{}
	WOpen    func(name string) (fs.File, error)
	WReadDir func(name string) ([]fs.DirEntry, error)
}

func (W _io_fs_ReadDirFS) Open(name string) (fs.File, error) {
	_ = "STUB: not implemented"
	return *new(fs.File), nil
}
func (W _io_fs_ReadDirFS) ReadDir(name string) ([]fs.DirEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type _io_fs_ReadDirFile struct {
	IValue   interface{}
	WClose   func() error
	WRead    func(a0 []byte) (int, error)
	WReadDir func(n int) ([]fs.DirEntry, error)
	WStat    func() (fs.FileInfo, error)
}

func (W _io_fs_ReadDirFile) Close() error                { _ = "STUB: not implemented"; return nil }
func (W _io_fs_ReadDirFile) Read(a0 []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }
func (W _io_fs_ReadDirFile) ReadDir(n int) ([]fs.DirEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
func (W _io_fs_ReadDirFile) Stat() (fs.FileInfo, error) {
	_ = "STUB: not implemented"
	return *new(fs.FileInfo), nil
}

type _io_fs_ReadFileFS struct {
	IValue    interface{}
	WOpen     func(name string) (fs.File, error)
	WReadFile func(name string) ([]byte, error)
}

func (W _io_fs_ReadFileFS) Open(name string) (fs.File, error) {
	_ = "STUB: not implemented"
	return *new(fs.File), nil
}
func (W _io_fs_ReadFileFS) ReadFile(name string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type _io_fs_StatFS struct {
	IValue interface{}
	WOpen  func(name string) (fs.File, error)
	WStat  func(name string) (fs.FileInfo, error)
}

func (W _io_fs_StatFS) Open(name string) (fs.File, error) {
	_ = "STUB: not implemented"
	return *new(fs.File), nil
}
func (W _io_fs_StatFS) Stat(name string) (fs.FileInfo, error) {
	_ = "STUB: not implemented"
	return *new(fs.FileInfo), nil
}

type _io_fs_SubFS struct {
	IValue interface{}
	WOpen  func(name string) (fs.File, error)
	WSub   func(dir string) (fs.FS, error)
}

func (W _io_fs_SubFS) Open(name string) (fs.File, error) {
	_ = "STUB: not implemented"
	return *new(fs.File), nil
}
func (W _io_fs_SubFS) Sub(dir string) (fs.FS, error) {
	_ = "STUB: not implemented"
	return *new(fs.FS), nil
}
