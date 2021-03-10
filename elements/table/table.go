package table

import "github.com/Nerzal/tinydom"

type Table struct {
	*tinydom.Element
}

func New() *Table {
	element := tinydom.GetDocument().CreateElement("table")
	return &Table{element}
}

func (t *Table) SetHeader(header ...string) *Table {
	doc := tinydom.GetDocument()
	head := doc.CreateElement("thead")

	tr := doc.CreateElement("tr")
	head.AppendChild(tr)

	for _, thead := range header {
		headElement := doc.CreateElement("th").SetInnerHTML(thead)
		tr.AppendChild(headElement)
	}

	t.AppendChild(head)
	return t
}

func (t *Table) SetBody(element *tinydom.Element) *Table {
	t.AppendChild(element)
	return t
}

func (t *Table) Body() *tinydom.Element {
	return t.FindChildNode("tbody")
}
