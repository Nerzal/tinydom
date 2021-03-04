package form

import "github.com/Nerzal/tinydom"

type Form struct {
	*tinydom.Element
}

func New() *Form {
	result := tinydom.GetDocument().CreateElement("form")
	return &Form{result}
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
