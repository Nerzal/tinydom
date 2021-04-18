package img

import "github.com/Nerzal/tinydom"

// Img is the html5 element img
// See https://developer.mozilla.org/en-US/docs/Web/HTML/Element/img for reference
type Img struct {
	*tinydom.Element
}

func New(src, alt string) *Img {
	imgElement := tinydom.GetDocument().CreateElement("img").
		SetAttribute("src", src).
		SetAttribute("alt", alt)

	return &Img{
		Element: imgElement,
	}
}

func (i *Img) SetSrc(value string) *Img {
	i.SetAttribute("src", value)
	return i
}

func (i *Img) Src() (bool, string) {
	return i.GetAttribute("src")
}

func (i *Img) SetAlt(value string) *Img {
	i.SetAttribute("alt", value)
	return i
}

func (i *Img) Alt() (bool, string) {
	return i.GetAttribute("alt")
}
