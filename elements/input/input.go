package input

import "github.com/Nerzal/tinydom"

type Input struct {
	*tinydom.Element
	iType InputType
}

func New(inputType InputType) *Input {
	input := tinydom.GetDocument().CreateElement("input")
	input.SetAttribute("type", string(inputType))
	return &Input{Element: input, iType: inputType}
}

func NewTextInput() *Input {
	return New(TextInput)
}

func FromElement(e *tinydom.Element) *Input {
	_, inputType := e.GetAttribute("type")

	return &Input{Element: e, iType: InputType(inputType)}
}

func (i *Input) Autofocus() bool {
	result, _ := i.GetAttribute("autofocus")
	return result
}

func (i *Input) SetAutofocus(b bool) *Input {
	i.SetAttribute("autofocus", b)
	return i
}

func (i *Input) Autocomplete() bool {
	result, _ := i.GetAttribute("autocomplete")
	return result
}

func (i *Input) SetAutocomplete(b bool) *Input {
	i.SetAttribute("autocomplete", b)
	return i
}

func (i *Input) For() string {
	return i.Get("for").String()
}

func (i *Input) SetFor(value string) *Input {
	i.SetAttribute("for", value)
	return i
}

func (i *Input) FormEnctype() string {
	return i.Get("formenctype").String()
}

func (i *Input) SetFormEnctype(value string) error {
	if i.iType != SubmitInput {
		return ErrInvalidAttribute
	}

	i.Set("formenctype", value)
	return nil
}

func (i *Input) FormTarget() string {
	return i.Get("formtarget").String()
}

func (i *Input) SetFormTarget(value tinydom.Target) error {
	if i.iType != SubmitInput {
		return ErrInvalidAttribute
	}

	i.Set("formtarget", value.String())
	return nil
}

func (i *Input) FormNoValidate() string {
	return i.Get("formnovalidate").String()
}

func (i *Input) SetFormNoValidate() error {
	if i.iType != SubmitInput {
		return ErrInvalidAttribute
	}

	i.Set("formnovalidate", "")
	return nil
}

func (i *Input) List() string {
	return i.Get("list").String()
}

func (i *Input) SetList(value string) *Input {
	i.Set("list", value)
	return i
}

func (i *Input) Min() string {
	return i.Get("min").String()
}

func (i *Input) SetMin(min string) *Input {
	i.Set("min", min)
	return i
}

func (i *Input) MaxLength() string {
	return i.Get("maxlength").String()
}

func (i *Input) SetMaxLength(length string) *Input {
	i.Set("maxlength", length)
	return i
}

func (i *Input) Max() string {
	return i.Get("max").String()
}

func (i *Input) SetMax(max string) *Input {
	i.Set("max", max)
	return i
}

func (i *Input) Checked() string {
	return i.Get("checked").String()
}

func (i *Input) SetChecked(checked string) *Input {
	i.Set("checked", checked)
	return i
}

func (i *Input) Required() string {
	return i.Get("required").String()
}

func (i *Input) SetRequired(required string) *Input {
	i.Set("required", required)
	return i
}

func (i *Input) Pattern() string {
	return i.Get("pattern").String()
}

func (i *Input) SetPattern(pattern string) *Input {
	i.Set("pattern", pattern)
	return i
}

func (i *Input) Step() string {
	return i.Get("step").String()
}

func (i *Input) SetStep(step string) *Input {
	i.Set("step", step)
	return i
}

func (i *Input) Value() string {
	return i.Get("value").String()
}

func (i *Input) SetValue(v string) *Input {
	i.Set("value", v)
	return i
}

func (i *Input) Width() string {
	return i.Get("width").String()
}

func (i *Input) SetWidth(v string) *Input {
	i.Set("width", v)
	return i
}

func (i *Input) Height() string {
	return i.Get("height").String()
}

func (i *Input) SetHeight(v string) *Input {
	i.Set("height", v)
	return i
}
