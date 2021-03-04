package tinydom

import "syscall/js"

type XSLTProcessor struct {
	js.Value
}

func NewXSLTProcessor() *XSLTProcessor {
	return &XSLTProcessor{GetWindow().Get("XSLTProcessor").New()}
}

func (x *XSLTProcessor) ImportStylesheet(node *Element) {
	x.Call("importStylesheet", node)
}

func (x *XSLTProcessor) TransformToFragment(node, document *Element) *Element {
	return &Element{x.Call("transformToFragment", node, document)}
}

func (x *XSLTProcessor) TransformToDocument(node *Element) *Element {
	return &Element{x.Call("transformToDocument", node)}
}
