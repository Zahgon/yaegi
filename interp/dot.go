package interp

import (
	"io"
)

func (n *node) astDot(out io.Writer, name string) { _ = "STUB: not implemented"; return }

func (n *node) cfgDot(out io.Writer) { _ = "STUB: not implemented"; return }

type nopCloser struct {
	io.Writer
}

func (nopCloser) Close() error { _ = "STUB: not implemented"; return nil }

func dotWriter(dotCmd string) io.WriteCloser {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser)
}

func defaultDotCmd(filePath, prefix string) string { _ = "STUB: not implemented"; return "" }
