package tinydom

import (
	"syscall/js"
)

type Element struct {
	js.Value
}

func (e *Element) HasFocus() bool {
	return e.IsEqualNode(GetDocument().ActiveElement())
}

func (e *Element) AppendBefore(n *Element) {
	e.ParentNode().InsertBefore(n, e)
}

func (e *Element) AppendAfter(n *Element) {
	e.ParentNode().InsertBefore(n, e.NextSibling())
}

func (e *Element) AppendChild(child *Element) {
	e.Call("appendChild", child)
}

func (e *Element) AppendChildren(children ...*Element) {
	for _, child := range children {
		e.AppendChild(child)
	}
}

// AppendChildBr appends the child and adds an additional br
func (e *Element) AppendChildBr(child *Element) {
	e.Call("appendChild", child)
	e.Call("appendChild", GetDocument().CreateElement("br"))
}

func (e *Element) AppendChildrenBr(children ...*Element) {
	for _, child := range children {
		e.AppendChildBr(child)
	}
}

func (e *Element) Br() {
	br := GetDocument().CreateElement("br")
	e.AppendChild(br)
}

func (e *Element) RemoveAllChildNodes() {
	for e.HasChildNodes() {
		e.RemoveChild(e.LastChild())
	}
}

func (e *Element) SetId(id string) {
	e.Set("id", id)
}

func (e *Element) SetAttribute(key, value string) {
	e.Call("setAttribute", key, value)
}

func (e *Element) SetInnerHTML(value string) {
	e.Set("innerHTML", value)
}

func (e *Element) InnerHTML() string {
	return e.Get("innerHTML").String()
}

func (e *Element) OuterHTML() string {
	return e.Get("outerHTML").String()
}

func (e *Element) SetOuterHTML(html string) {
	e.Set("outerHTML", html)
}

func (e *Element) TagName() string {
	return e.Get("tagName").String()
}

// GetAttribute returns the searched attribute, returns false if the attribute wasn't found.
func (e *Element) GetAttribute(name string) (bool, string) {
	if !e.HasAttribute(name) {
		return false, ""
	}

	return true, e.Call("getAttribute", name).String()
}

func (e *Element) HasAttribute(name string) bool {
	return e.Call("hasAttribute", name).Bool()
}

func (e *Element) QuerySelector(selectors string) *Element {
	return &Element{e.Call("querySelector", selectors)}
}

func (e *Element) QuerySelectorAll(selectors string) []*Element {
	nodeList := e.Call("querySelectorAll", selectors)
	length := nodeList.Get("length").Int()

	nodes := make([]*Element, length)

	for i := 0; i < length; i++ {
		nodes[i] = &Element{nodeList.Call("item", i)}
	}

	return nodes
}

func (e *Element) GetElementsByTagName(tagName string) []*Element {
	nodeList := e.Call("getElementsByTagName", tagName)
	length := nodeList.Get("length").Int()

	nodes := make([]*Element, length)

	for i := 0; i < length; i++ {
		nodes[i] = &Element{nodeList.Call("item", i)}
	}

	return nodes
}

func (e *Element) ChildNodes() []*Element {
	nodeList := e.Get("childNodes")
	length := nodeList.Get("length").Int()
	var nodes []*Element
	for i := 0; i < length; i++ {
		nodes = append(nodes, &Element{nodeList.Call("item", i)})
	}
	return nodes
}

func (e *Element) FirstChild() *Element {
	return &Element{e.Get("firstChild")}
}

func (e *Element) LastChild() *Element {
	return &Element{e.Get("lastChild")}
}

func (e *Element) NextSibling() *Element {
	return &Element{e.Get("nextSibling")}
}

func (e *Element) NodeType() int {
	return e.Get("nodeType").Int()
}

func (e *Element) NodeValue() string {
	return e.Get("nodeValue").String()
}

func (e *Element) SetNodeValue(s string) {
	e.Set("nodeValue", s)
}

func (e *Element) ParentNode() *Element {
	return &Element{e.Get("parentNode")}
}

func (e *Element) TextContent() string {
	return e.Get("textContent").String()
}

func (e *Element) SetTextContent(s string) {
	e.Set("textContent", s)
}

func (e *Element) Contains(n *Element) bool {
	return e.Call("contains", n).Bool()
}

func (e *Element) HasChildNodes() bool {
	return e.Call("hasChildNodes").Bool()
}

func (e *Element) InsertBefore(newNode, referenceNode *Element) *Element {
	return &Element{e.Call("insertBefore", newNode, referenceNode)}
}

func (e *Element) IsEqualNode(n *Element) bool {
	return e.Call("isEqualNode", n).Bool()
}

func (e *Element) IsSameNode(n *Element) bool {
	return e.Call("isSameNode", n).Bool()
}

func (e *Element) LookupPrefix() string {
	return e.Call("lookupPrefix").String()
}

func (e *Element) Normalize() {
	e.Call("normalize")
}

func (e *Element) RemoveChild(c *Element) *Element {
	return &Element{e.Call("removeChild", c)}
}

func (e *Element) ReplaceChild(newChild, oldChild *Element) *Element {
	return &Element{e.Call("replaceChild", newChild, oldChild)}
}

func (e *Element) AddEventListener(t string, listener js.Func) {
	e.Call("addEventListener", t, listener)
}

func (e *Element) RemoveEventListener(t string, listener js.Func) {
	e.Call("removeEventListener", t, listener)
}

func (e *Element) Style() *CSS {
	return &CSS{e.Get("style")}
}

func (e *Element) Dataset() *Element {
	return &Element{e.Get("dataset")}
}

func (e *Element) Blur() {
	e.Call("blur")
}

func (e *Element) Focus() {
	e.Call("focus")
}
