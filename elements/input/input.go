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
	return i.Get("autofocus").Bool()
}

func (i *Input) SetAutofocus(b bool) {
	i.Set("autofocus", b)
}

func (i *Input) Autocomplete() bool {
	return i.Get("autocomplete").Bool()
}

func (i *Input) SetAutocomplete(b bool) {
	i.Set("autocomplete", b)
}

func (i *Input) For() string {
	return i.Get("for").String()
}

func (i *Input) SetFor(value string) {
	i.Set("for", value)
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

func (i *Input) SetList(value string) {
	i.Set("list", value)
}

func (i *Input) Min() string {
	return i.Get("min").String()
}

func (i *Input) SetMin(min string) {
	i.Set("min", min)
}

func (i *Input) MaxLength() string {
	return i.Get("maxlength").String()
}

func (i *Input) SetMaxLength(length string) {
	i.Set("maxlength", length)
}

func (i *Input) Max() string {
	return i.Get("max").String()
}

func (i *Input) SetMax(max string) {
	i.Set("max", max)
}

func (i *Input) Checked() string {
	return i.Get("checked").String()
}

func (i *Input) SetChecked(checked string) {
	i.Set("checked", checked)
}

func (i *Input) Required() string {
	return i.Get("required").String()
}

func (i *Input) SetRequired(required string) {
	i.Set("required", required)
}

func (i *Input) Pattern() string {
	return i.Get("pattern").String()
}

func (i *Input) SetPattern(pattern string) {
	i.Set("pattern", pattern)
}

func (i *Input) Step() string {
	return i.Get("step").String()
}

func (i *Input) SetStep(step string) {
	i.Set("step", step)
}

func (i *Input) OnClick(function string) {
	i.Set("onclick", function)
}

func (i *Input) Value() string {
	return i.Get("value").String()
}

func (i *Input) SetValue(v string) {
	i.Set("value", v)
}

func (i *Input) Width() string {
	return i.Get("width").String()
}

func (i *Input) SetWidth(v string) {
	i.Set("width", v)
}

func (i *Input) Height() string {
	return i.Get("height").String()
}

func (i *Input) SetHeight(v string) {
	i.Set("height", v)
}
