package li

import "github.com/Nerzal/tinydom"

// Li is a li element
// See https://developer.mozilla.org/en-US/docs/Web/HTML/Element/li for reference
type Li struct {
	*tinydom.Element
}

func New() *Li {
	doc := tinydom.GetDocument()
	element := doc.CreateElement("li")

	return &Li{
		Element: element,
	}
}
