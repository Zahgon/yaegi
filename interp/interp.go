package interp

import (
	"context"
	"go/build"
	"go/scanner"
	"go/token"
	"io"
	"io/fs"
	"reflect"
	"sync"
)

type node struct {
	debug      *nodeDebugData
	child      []*node
	anc        *node
	param      []*itype
	start      *node
	tnext      *node
	fnext      *node
	interp     *Interpreter
	index      int64
	findex     int
	level      int
	nleft      int
	nright     int
	kind       nkind
	pos        token.Pos
	sym        *symbol
	typ        *itype
	recv       *receiver
	types      []reflect.Type
	scope      *scope
	action     action
	exec       bltn
	gen        bltnGenerator
	val        interface{}
	rval       reflect.Value
	ident      string
	redeclared bool
	meta       interface{}
}

func (n *node) shouldBreak() bool { _ = "STUB: not implemented"; return false }

func (n *node) setProgram(p *Program) { _ = "STUB: not implemented"; return }

func (n *node) setBreakOnCall(v bool) { _ = "STUB: not implemented"; return }

func (n *node) setBreakOnLine(v bool) { _ = "STUB: not implemented"; return }

type receiver struct {
	node  *node
	val   reflect.Value
	index []int
}

type frame struct {
	id uint64

	debug *frameDebugData

	root *frame
	anc  *frame
	data []reflect.Value

	mutex     sync.RWMutex
	deferred  [][]reflect.Value
	recovered interface{}
	done      reflect.SelectCase
}

func newFrame(anc *frame, length int, id uint64) *frame { _ = "STUB: not implemented"; return nil }

func (f *frame) runid() uint64      { _ = "STUB: not implemented"; return 0 }
func (f *frame) setrunid(id uint64) { _ = "STUB: not implemented"; return }
func (f *frame) clone() *frame      { _ = "STUB: not implemented"; return nil }

type Exports map[string]map[string]reflect.Value

type imports map[string]map[string]*symbol

type opt struct {
	dotCmd       string
	context      build.Context
	stdin        io.Reader
	stdout       io.Writer
	stderr       io.Writer
	args         []string
	env          map[string]string
	filesystem   fs.FS
	astDot       bool
	cfgDot       bool
	noRun        bool
	fastChan     bool
	specialStdio bool
	unrestricted bool
}

type Interpreter struct {
	id uint64

	nindex int64

	name string

	opt
	cancelChan bool
	fset       *token.FileSet
	binPkg     Exports
	rdir       map[string]bool
	mapTypes   map[reflect.Value][]reflect.Type

	mutex    sync.RWMutex
	frame    *frame
	universe *scope
	scopes   map[string]*scope
	srcPkg   imports
	pkgNames map[string]string
	done     chan struct{}
	roots    []*node
	generic  map[string]*node

	hooks *hooks

	debugger *Debugger
}

const (
	mainID     = "main"
	selfPrefix = "github.com/traefik/yaegi"
	selfPath   = selfPrefix + "/interp/interp"

	DefaultSourceName = "_.go"

	Test = false

	NoTest = true
)

var Self *Interpreter

var Symbols = Exports{
	selfPath: map[string]reflect.Value{
		"New": reflect.ValueOf(New),

		"Interpreter": reflect.ValueOf((*Interpreter)(nil)),
		"Options":     reflect.ValueOf((*Options)(nil)),
		"Panic":       reflect.ValueOf((*Panic)(nil)),
	},
}

func init() { Symbols[selfPath]["Symbols"] = reflect.ValueOf(Symbols) }

type _error struct {
	IValue interface{}
	WError func() string
}

func (w _error) Error() string { _ = "STUB: not implemented"; return "" }

type Panic struct {
	Value interface{}

	Callers []uintptr

	Stack []byte
}

func (e Panic) Error() string { _ = "STUB: not implemented"; return "" }

func (n *node) Walk(in func(n *node) bool, out func(n *node)) { _ = "STUB: not implemented"; return }

type Options struct {
	GoPath string

	BuildTags []string

	Stdin          io.Reader
	Stdout, Stderr io.Writer

	Args []string

	Env []string

	SourcecodeFilesystem fs.FS

	Unrestricted bool
}

func New(options Options) *Interpreter { _ = "STUB: not implemented"; return nil }

const (
	bltnAlignof  = "unsafe.Alignof"
	bltnAppend   = "append"
	bltnCap      = "cap"
	bltnClose    = "close"
	bltnComplex  = "complex"
	bltnImag     = "imag"
	bltnCopy     = "copy"
	bltnDelete   = "delete"
	bltnLen      = "len"
	bltnMake     = "make"
	bltnNew      = "new"
	bltnOffsetof = "unsafe.Offsetof"
	bltnPanic    = "panic"
	bltnPrint    = "print"
	bltnPrintln  = "println"
	bltnReal     = "real"
	bltnRecover  = "recover"
	bltnSizeof   = "unsafe.Sizeof"
)

func initUniverse() *scope { _ = "STUB: not implemented"; return nil }

func (interp *Interpreter) resizeFrame() { _ = "STUB: not implemented"; return }

func (interp *Interpreter) Eval(src string) (res reflect.Value, err error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}

func (interp *Interpreter) EvalPath(path string) (res reflect.Value, err error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}

func (interp *Interpreter) EvalPathWithContext(ctx context.Context, path string) (res reflect.Value, err error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}

func (interp *Interpreter) EvalTest(path string) error { _ = "STUB: not implemented"; return nil }

func isFile(filesystem fs.FS, path string) bool { _ = "STUB: not implemented"; return false }

func (interp *Interpreter) eval(src, name string, inc bool) (res reflect.Value, err error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}

func (interp *Interpreter) EvalWithContext(ctx context.Context, src string) (reflect.Value, error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}

func (interp *Interpreter) stop() { _ = "STUB: not implemented"; return }

func (interp *Interpreter) runid() uint64 { _ = "STUB: not implemented"; return 0 }

func ignoreScannerError(e *scanner.Error, s string) bool { _ = "STUB: not implemented"; return false }

func (interp *Interpreter) ImportUsed() { _ = "STUB: not implemented"; return }

func key2name(name string) string { _ = "STUB: not implemented"; return "" }

func fixKey(k string) string { _ = "STUB: not implemented"; return "" }

func (interp *Interpreter) REPL() (reflect.Value, error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}

func doPrompt(out io.Writer) func(v reflect.Value) { _ = "STUB: not implemented"; return nil }

func getPrompt(in io.Reader, out io.Writer) func(reflect.Value) {
	_ = "STUB: not implemented"
	return nil
}
