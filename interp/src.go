package interp

import (
	"io/fs"
)

func (interp *Interpreter) importSrc(rPath, importPath string, skipTest bool) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (interp *Interpreter) rootFromSourceLocation() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (interp *Interpreter) pkgDir(goPath string, root, importPath string) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

const vendor = "vendor"

func previousRoot(filesystem fs.FS, rootPath, root string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func effectivePkg(root, p string) string { _ = "STUB: not implemented"; return "" }

func isPathRelative(s string) bool { _ = "STUB: not implemented"; return false }
