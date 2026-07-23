package main

import "fmt"

type WidgetEvent struct {
	Nothing string
}

type WidgetControl interface {
	HandleEvent(e *WidgetEvent)
}

type Button struct{}

func (b *Button) HandleEvent(e *WidgetEvent) { _ = "STUB: not implemented"; return }

type WindowEvent struct {
	Something int
}

type Window struct {
	Widget WidgetControl
}

func (w *Window) HandleEvent(e *WindowEvent) { _ = "STUB: not implemented"; return }

func main() {
	window := &Window{
		Widget: &Button{},
	}
	windowevent := &WindowEvent{}

	window.HandleEvent(windowevent)
	fmt.Println("OK!")
}
