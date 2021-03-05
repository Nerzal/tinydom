package label

import "github.com/Nerzal/tinydom"

type Label struct {
	*tinydom.Element
}

func New() *Label {
	return &Label{tinydom.GetDocument().CreateElement("label")}
}

func (l *Label) SetFor(value string) *Label {
	l.SetAttribute("for", value)
	return l
}

func (l *Label) For() string {
	return l.Get("for").String()
}
