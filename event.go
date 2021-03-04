package tinydom

import "syscall/js"

type Event struct {
	js.Value
}

func (e *Event) Target() *Element {
	return &Element{e.Get("target")}
}

func (e *Event) PreventDefault() {
	e.Call("preventDefault")
}

func (e *Event) StopImmediatePropagation() {
	e.Call("stopImmediatePropagation")
}

func (e *Event) StopPropagation() {
	e.Call("stopPropagation")
}

func (o *Event) AddEventListener(t string, listener func(Event), args ...interface{}) {
	if len(args) == 1 {
		o.Call("addEventListener", t, listener, args[0])
	} else {
		o.Call("addEventListener", t, listener)
	}
}

func (o *Event) RemoveEventListener(t string, listener func(Event), args ...interface{}) {
	if len(args) == 1 {
		o.Call("removeEventListener", t, listener, args[0])
	} else {
		o.Call("removeEventListener", t, listener)
	}
}
