package elements

import tinydom "github.com/Nerzal/TinyDom"

type Input struct {
	*tinydom.Element
}

func NewTextInput() *Input {
	tinydom.GetDocument().CreateElement("input")

	return &Input{}
}

func (i *Input) Value() string {
	return i.Get("value").String()
}

func (i *Input) SetValue(s string) {
	i.Set("value", s)
}
