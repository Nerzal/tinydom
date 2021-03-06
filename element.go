package tinydom

import (
	"errors"
	"strings"
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

func (e *Element) SetId(id string) *Element {
	e.Set("id", id)
	return e
}

func (e *Element) SetAttribute(key, value interface{}) *Element {
	e.Call("setAttribute", key, value)
	return e
}

func (e *Element) SetClass(values ...string) *Element {
	return e.SetMultiValueAttribute("class", values...)
}

// ErrClassAlreadyExisting is being thrown when trying to append the same class multiple times
var ErrClassAlreadyExisting = errors.New("tried to append class multiple times")

func (e *Element) AppendClass(values ...string) error {
	existing, currentClasses := e.Class()

	if !existing {
		e.SetClass(values...)
		return nil
	}

	for _, newClass := range values {
		for _, existingclass := range currentClasses {
			if newClass == existingclass {
				return ErrClassAlreadyExisting
			}
		}
	}

	newClass := append(currentClasses, values...)
	e.SetClass(newClass...)

	return nil
}

func (e *Element) Class() (bool, []string) {
	exists, attributeValues := e.GetAttribute("class")
	if !exists {
		return false, nil
	}

	splittedValues := strings.Split(attributeValues, " ")

	result := make([]string, len(attributeValues))
	for i, value := range splittedValues {
		result[i] = value
	}

	return true, result
}

func (e *Element) SetMultiValueAttribute(attributeName string, values ...string) *Element {
	var value string

	valueCount := len(values)

	for i, rel := range values {
		value += rel

		if i < valueCount {
			value += " "
		}
	}

	e.SetAttribute(attributeName, value)
	return e
}

func (e *Element) SetInnerHTML(value string) *Element {
	e.Set("innerHTML", value)
	return e
}

func (e *Element) InnerHTML() string {
	return e.Get("innerHTML").String()
}

func (e *Element) OuterHTML() string {
	return e.Get("outerHTML").String()
}

func (e *Element) SetOuterHTML(html string) *Element {
	e.Set("outerHTML", html)
	return e
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

func (e *Element) FindChildNode(tag string) *Element {
	children := e.ChildNodes()
	for _, child := range children {
		if child.TagName() == tag {
			return child
		}
	}

	return nil
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

func (e *Element) SetNodeValue(s string) *Element {
	e.Set("nodeValue", s)
	return e
}
func (e *Element) Name() string {
	return e.Get("name").String()
}

func (e *Element) SetName(n string) *Element {
	e.Set("name", n)
	return e
}

func (e *Element) ParentNode() *Element {
	return &Element{e.Get("parentNode")}
}

func (e *Element) TextContent() string {
	return e.Get("textContent").String()
}

func (e *Element) SetTextContent(s string) *Element {
	e.Set("textContent", s)
	return e
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

func (e *Element) AddEventListener(t string, listener js.Func) *Element {
	e.Call("addEventListener", t, listener)
	return e
}

func (e *Element) RemoveEventListener(t string, listener js.Func) *Element {
	e.Call("removeEventListener", t, listener)
	return e
}

func (e *Element) Style() *CSS {
	return &CSS{e.Get("style")}
}

func (e *Element) Dataset() *Element {
	return &Element{e.Get("dataset")}
}

func (e *Element) Blur() *Element {
	e.Call("blur")
	return e
}

func (e *Element) Focus() *Element {
	e.Call("focus")
	return e
}
