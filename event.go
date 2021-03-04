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

func (e *Event) Code() string {
	return e.Get("code").String()
}

func (e *Event) Key() string {
	return e.Get("key").String()
}

func (e *Event) KeyCode() int {
	return e.Get("keyCode").Int()
}
