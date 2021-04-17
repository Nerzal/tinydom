package li

import "github.com/Nerzal/tinydom"

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
