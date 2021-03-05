package href

import "github.com/Nerzal/tinydom"

type Href struct {
	*tinydom.Element
}

func New(link, innerHTML string) *Href {
	a := tinydom.GetDocument().CreateElement("a")
	a.Set("href", link)
	a.Set("target", "_blank")
	a.SetInnerHTML(innerHTML)
	return &Href{a}
}
