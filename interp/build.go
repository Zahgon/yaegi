package interp

import (
	"go/ast"
	"go/build"
)

func (interp *Interpreter) buildOk(ctx *build.Context, name, src string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func buildLineOk(ctx *build.Context, line string) (ok bool) {
	_ = "STUB: not implemented"
	return false
}

func buildOptionOk(ctx *build.Context, tag string) bool { _ = "STUB: not implemented"; return false }

func buildTagOk(ctx *build.Context, s string) (r bool) { _ = "STUB: not implemented"; return false }

func setYaegiTags(ctx *build.Context, comments []*ast.CommentGroup) {
	_ = "STUB: not implemented"
	return
}

func contains(tags []string, tag string) bool { _ = "STUB: not implemented"; return false }

func goMinorVersion(ctx *build.Context) int { _ = "STUB: not implemented"; return 0 }

func skipFile(ctx *build.Context, p string, skipTest bool) bool {
	_ = "STUB: not implemented"
	return false
}

var knownOs = map[string]bool{
	"aix":       true,
	"android":   true,
	"darwin":    true,
	"dragonfly": true,
	"freebsd":   true,
	"illumos":   true,
	"ios":       true,
	"js":        true,
	"linux":     true,
	"netbsd":    true,
	"openbsd":   true,
	"plan9":     true,
	"solaris":   true,
	"wasip1":    true,
	"windows":   true,
}

var knownArch = map[string]bool{
	"386":      true,
	"amd64":    true,
	"arm":      true,
	"arm64":    true,
	"loong64":  true,
	"mips":     true,
	"mips64":   true,
	"mips64le": true,
	"mipsle":   true,
	"ppc64":    true,
	"ppc64le":  true,
	"s390x":    true,
	"wasm":     true,
}
