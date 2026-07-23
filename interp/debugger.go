package interp

import (
	"context"
	"errors"
	"go/token"
	"reflect"
	"sync"
)

var (
	ErrNotLive = errors.New("not live")

	ErrRunning = errors.New("running")

	ErrNotRunning = errors.New("not running")
)

var rNodeType = reflect.TypeOf((*node)(nil)).Elem()

type Debugger struct {
	interp  *Interpreter
	events  func(*DebugEvent)
	context context.Context
	cancel  context.CancelFunc

	gWait *sync.WaitGroup
	gLock *sync.Mutex
	gID   int
	gLive map[int]*debugRoutine

	result reflect.Value
	err    error
}

type debugRoutine struct {
	id int

	mode    DebugEventReason
	running bool
	resume  chan struct{}

	fDepth int
	fStep  int
}

type nodeDebugData struct {
	program     *Program
	breakOnLine bool
	breakOnCall bool
}

type frameDebugData struct {
	g     *debugRoutine
	node  *node
	name  string
	kind  frameKind
	scope *scope
}

type frameKind int

const (
	frameRoot frameKind = iota + 1

	frameCall

	frameClosure
)

type DebugOptions struct {
	GoRoutineStartAt1 bool
}

type DebugEvent struct {
	debugger *Debugger
	reason   DebugEventReason
	frame    *frame
}

type DebugFrame struct {
	event  *DebugEvent
	frames []*frame
}

type DebugFrameScope struct {
	frame *frame
}

type DebugVariable struct {
	Name  string
	Value reflect.Value
}

type DebugGoRoutine struct {
	id int
}

type Breakpoint struct {
	Valid bool

	Position token.Position
}

type DebugEventReason int

const (
	debugRun DebugEventReason = iota

	DebugPause

	DebugBreak

	DebugEntry

	DebugStepInto

	DebugStepOver

	DebugStepOut

	DebugTerminate

	DebugEnterGoRoutine

	DebugExitGoRoutine
)

func (interp *Interpreter) Debug(ctx context.Context, prog *Program, events func(*DebugEvent), opts *DebugOptions) *Debugger {
	_ = "STUB: not implemented"
	return nil
}

func (dbg *Debugger) Wait() (reflect.Value, error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}

func (dbg *Debugger) enterGoRoutine() *debugRoutine { _ = "STUB: not implemented"; return nil }

func (dbg *Debugger) exitGoRoutine(g *debugRoutine) { _ = "STUB: not implemented"; return }

func (dbg *Debugger) getGoRoutine(id int) (*debugRoutine, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (dbg *Debugger) enterCall(nFunc, nCall *node, f *frame) { _ = "STUB: not implemented"; return }

func (dbg *Debugger) exitCall(nFunc, nCall *node, f *frame) { _ = "STUB: not implemented"; return }

func (dbg *Debugger) exec(n *node, f *frame) (stop bool) { _ = "STUB: not implemented"; return false }

func (dbg *Debugger) Continue(id int) error { _ = "STUB: not implemented"; return nil }

func (g *debugRoutine) setMode(reason DebugEventReason) { _ = "STUB: not implemented"; return }

func (dbg *Debugger) Step(id int, reason DebugEventReason) error {
	_ = "STUB: not implemented"
	return nil
}

func (dbg *Debugger) Interrupt(id int, reason DebugEventReason) bool {
	_ = "STUB: not implemented"
	return false
}

func (dbg *Debugger) Terminate() { _ = "STUB: not implemented"; return }

type BreakpointTarget func(*Debugger, func(*node))

func PathBreakpointTarget(path string) BreakpointTarget {
	_ = "STUB: not implemented"
	return *new(BreakpointTarget)
}

func ProgramBreakpointTarget(prog *Program) BreakpointTarget {
	_ = "STUB: not implemented"
	return *new(BreakpointTarget)
}

func AllBreakpointTarget() BreakpointTarget {
	_ = "STUB: not implemented"
	return *new(BreakpointTarget)
}

type breakpointSetup struct {
	roots []*node
	lines map[int]int
	funcs map[string]int
}

type BreakpointRequest func(*breakpointSetup, int)

func LineBreakpoint(line int) BreakpointRequest {
	_ = "STUB: not implemented"
	return *new(BreakpointRequest)
}

func FunctionBreakpoint(name string) BreakpointRequest {
	_ = "STUB: not implemented"
	return *new(BreakpointRequest)
}

func (dbg *Debugger) SetBreakpoints(target BreakpointTarget, requests ...BreakpointRequest) []Breakpoint {
	_ = "STUB: not implemented"
	return nil
}

func (dbg *Debugger) GoRoutines() []*DebugGoRoutine { _ = "STUB: not implemented"; return nil }

func (r *DebugGoRoutine) ID() int { _ = "STUB: not implemented"; return 0 }

func (r *DebugGoRoutine) Name() string { _ = "STUB: not implemented"; return "" }

func (evt *DebugEvent) GoRoutine() int { _ = "STUB: not implemented"; return 0 }

func (evt *DebugEvent) Reason() DebugEventReason {
	_ = "STUB: not implemented"
	return *new(DebugEventReason)
}

func (evt *DebugEvent) walkFrames(fn func([]*frame) bool) { _ = "STUB: not implemented"; return }

func (evt *DebugEvent) FrameDepth() int { _ = "STUB: not implemented"; return 0 }

func (evt *DebugEvent) Frames(start, end int) []*DebugFrame { _ = "STUB: not implemented"; return nil }

func (f *DebugFrame) Name() string { _ = "STUB: not implemented"; return "" }

func (f *DebugFrame) Position() token.Position {
	_ = "STUB: not implemented"
	return *new(token.Position)
}

func (f *DebugFrame) Program() *Program { _ = "STUB: not implemented"; return nil }

func (f *DebugFrame) Scopes() []*DebugFrameScope { _ = "STUB: not implemented"; return nil }

func (f *DebugFrameScope) IsClosure() bool { _ = "STUB: not implemented"; return false }

func (f *DebugFrameScope) Variables() []*DebugVariable { _ = "STUB: not implemented"; return nil }

func scanScope(sc *scope, index map[int]string) { _ = "STUB: not implemented"; return }
