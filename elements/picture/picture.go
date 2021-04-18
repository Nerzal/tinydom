package picture

import (
	"github.com/Nerzal/tinydom"
	"github.com/Nerzal/tinydom/elements/img"
)

// <picture>
//     <source srcset="/media/cc0-images/surfer-240-200.jpg"
//             media="(min-width: 800px)">
//       <source srcset="/media/cc0-images/surfer-240-200.png"
//             media="(min-width: 800px)">
//       <source srcset="/media/cc0-images/surfer-240-200.webgl"
//             media="(min-width: 800px)">
//     <img src="/media/cc0-images/painted-hand-298-332.jpg" alt="" />
// </picture>

// Picture is the html5 picture element
// See https://developer.mozilla.org/en-US/docs/Web/HTML/Element/picture for reference
type Picture struct {
	*tinydom.Element
	img     *img.Img
	sources []*Source
}

// New creates a new instance of Picture
// image is the path to the image
// sources are the different image type sources
func New(image *img.Img, sources ...*Source) *Picture {
	doc := tinydom.GetDocument()
	pictureElement := doc.CreateElement("picture")

	result := &Picture{
		Element: pictureElement,
		sources: sources,
		img:     image,
	}

	result.AppendSources(sources...)

	result.AppendChild(image.Element)

	return result
}

func (p *Picture) AppendSources(sources ...*Source) *Picture {
	if sources == nil {
		return p
	}

	for _, source := range sources {
		p.AppendChild(source.Element)
	}

	return p
}

func (p *Picture) Sources() []*Source {
	elements := p.GetElementsByTagName("source")

	result := make([]*Source, len(elements))

	for i, element := range elements {
		result[i] = FromSourceElement(element)
	}

	return result
}
