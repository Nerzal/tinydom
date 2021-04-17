package nav

import (
	"github.com/Nerzal/tinydom"
	"github.com/Nerzal/tinydom/elements/li"
)

type Nav struct {
	*tinydom.Element
	ulElement *tinydom.Element
}

func New() *Nav {
	doc := tinydom.GetDocument()
	element := doc.CreateElement("nav")
	ulElement := doc.CreateElement("ul")

	element.AppendChild(ulElement)

	return &Nav{
		Element:   element,
		ulElement: ulElement,
	}
}

func (n *Nav) AppendListItem(item *li.Li) *Nav {
	n.ulElement.AppendChild(item.Element)
	return n
}
