package tinydom

import "syscall/js"

// Document wraps the JavaScript document element, which is usually fetched by js.Global().Get("document")
type Document struct {
	js.Value
}

var doc = js.Global().Get("document")

func GetDocument() *Document {
	return &Document{doc}
}

func (e *Document) ActiveElement() *Element {
	return &Element{e.Get("activeElement")}
}

func (e *Document) DocumentElement() *Element {
	return &Element{e.Get("documentElement")}
}

func (d *Document) CreateElement(tag string) *Element {
	return &Element{d.Call("createElement", tag)}
}

func (d *Document) CreateTextNode(textContent string) *Element {
	return &Element{d.Call("createTextNode", textContent)}
}

func (d *Document) GetElementById(id string) *Element {
	return &Element{d.Call("getElementById", id)}
}

func (d *Document) Write(markup string) {
	d.Call("write", markup)
}
