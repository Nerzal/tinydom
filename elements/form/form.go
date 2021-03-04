package form

import "github.com/Nerzal/tinydom"

type Form struct {
	*tinydom.Element
}

func New() *Form {
	result := tinydom.GetDocument().CreateElement("form")
	return &Form{result}
}

// Check https://www.w3schools.com/html/html_form_elements.asp for valid tags
func (f *Form) Append(elements ...*tinydom.Element) error {
	for i := range elements {
		element := elements[i]

		switch element.TagName() {
		case "input", "label", "select", "textarea", "div", "button", "fieldset", "legend", "datalist", "output", "option", "optgroup":
			f.AppendChild(element)
			return nil
		default:
			return ErrInvalidTagAppended
		}
	}

	return nil
}

func (f *Form) SetMethod(method Method) {
	f.SetAttribute("method", method.String())
}

func (f *Form) SetAction(action string) {
	f.SetAttribute("action", action)
}

func (f *Form) SetNoValidate() {
	f.Set("novalidate", nil)
}
