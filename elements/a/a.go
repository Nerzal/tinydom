package a

import "github.com/Nerzal/tinydom"

// A is an a element
// See https://developer.mozilla.org/en-US/docs/Web/HTML/Element/a for reference
type A struct {
	*tinydom.Element
}

// New creates a new instance of A
func New(href, description string) *A {
	aElement := tinydom.GetDocument().
		CreateElement("a").
		SetAttribute("href", href).
		SetInnerHTML(description)

	return &A{
		Element: aElement,
	}
}

func (a *A) SetTarget(value string) *A {
	a.Element.SetAttribute("target", value)
	return a
}

func (a *A) Target() (bool, string) {
	return a.GetAttribute("target")
}

func (a *A) SetRel(values ...Rel) *A {
	var value Rel

	valueCount := len(values)

	for i, rel := range values {
		value += rel

		if i < valueCount {
			value += " "
		}
	}

	a.Element.SetAttribute("rel", value)
	return a
}

func (a *A) Rel() (bool, string) {
	return a.GetAttribute("rel")
}
