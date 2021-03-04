package form

import "github.com/Nerzal/tinydom"

type FieldSet struct {
	*tinydom.Element
}

// NewFieldSet creates a new fieldset element.
// Optionally a legend can be passed. the string is being used as innerHTML
func NewFieldSet(legendArgs ...string) *FieldSet {
	doc := tinydom.GetDocument()
	result := doc.CreateElement("fieldset")

	for _, legend := range legendArgs {
		legendElem := doc.CreateElement("legend")
		legendElem.Set("value", legend)
		result.AppendChild(legendElem)
	}

	return &FieldSet{result}
}

// Check https://www.w3schools.com/html/html_form_elements.asp for valid tags
func (f *FieldSet) Append(elements ...*tinydom.Element) error {
	for i := range elements {
		element := elements[i]
		f.AppendChild(element)
	}

	return nil
}
