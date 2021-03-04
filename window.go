package tinydom

import "syscall/js"

type Window struct {
	js.Value
}

var win = js.Global().Get("window")

func GetWindow() *Window {
	return &Window{win}
}

func (w *Window) Location() *Location {
	return &Location{w.Get("location")}
}

func (w *Window) Navigator() *Navigator {
	return &Navigator{w.Get("navigator")}
}

func (w *Window) History() *History {
	return &History{w.Get("history")}
}

func (w *Window) Alert(message string) {
	w.Call("alert", message)
}

func (w *Window) PushState(state interface{}, title, URL string) {
	w.Get("history").Call("pushState", state, title, URL)
}

func (w *Window) ReplaceState(state interface{}, title, URL string) {
	w.Get("history").Call("replaceState", state, title, URL)
}

func (w *Window) PageXOffset() float64 {
	return w.Get("pageXOffset").Float()
}

func (w *Window) PageYOffset() float64 {
	return w.Get("pageYOffset").Float()
}

func (w *Window) ScrollX() float64 {
	return w.Get("scrollX").Float()
}

func (w *Window) ScrollY() float64 {
	return w.Get("scrollY").Float()
}
