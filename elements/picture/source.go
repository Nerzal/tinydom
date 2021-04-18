package picture

import "github.com/Nerzal/tinydom"

// <source srcset="/media/cc0-images/surfer-240-200.jpg" media="(min-width: 800px)">

// Source is a html5 source element
// See https://developer.mozilla.org/en-US/docs/Web/HTML/Element/source for reference
type Source struct {
	*tinydom.Element
}

// New creates a new instance of Source
func NewSource() *Source {
	sourceElement := tinydom.GetDocument().CreateElement("source")

	return &Source{
		Element: sourceElement,
	}
}

// FromSourceElement creates a new Source based on the given element
func FromSourceElement(element *tinydom.Element) *Source {
	return &Source{
		Element: element,
	}
}

func (s *Source) SetType(mimetype string) *Source {
	s.SetAttribute("type", mimetype)
	return s
}

func (s *Source) Type() (bool, string) {
	return s.GetAttribute("type")
}

func (s *Source) SetSrcSet(value string) *Source {
	s.SetAttribute("srcset", value)
	return s
}

func (s *Source) SrcSet() (bool, string) {
	return s.GetAttribute("srcset")
}

func (s *Source) SetSizes(value ...string) *Source {
	s.SetMultiValueAttribute("sizes", value...)
	return s
}

func (s *Source) Sizes() (bool, string) {
	return s.GetAttribute("sizes")
}

func (s *Source) SetMedia(value string) *Source {
	s.SetAttribute("media", value)
	return s
}

func (s *Source) Media() (bool, string) {
	return s.GetAttribute("media")
}
