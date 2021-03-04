package input

import "github.com/Nerzal/tinydom"

type InputType string

const (
	Password      InputType = "password"
	ButtonInput   InputType = "button"
	CheckboxInput InputType = "checkbox"
	ColorInput    InputType = "color"
	DateInput     InputType = "date"
	LocalInput    InputType = "datetime-local"
	EmailInput    InputType = "email"
	FileInput     InputType = "file"
	HiddenInput   InputType = "hidden"
	ImageInput    InputType = "image"
	MonthInput    InputType = "month"
	NumberInput   InputType = "number"
	PasswordInput InputType = "password"
	RadioInput    InputType = "radio"
	RangeInput    InputType = "range"
	ResetInput    InputType = "reset"
	SearchInput   InputType = "search"
	SubmitInput   InputType = "submit"
	TelInput      InputType = "tel"
	TextInput     InputType = "text"
	TimeInput     InputType = "time"
	UrlInput      InputType = "url"
	WeekInput     InputType = "week"
)

type Input struct {
	*tinydom.Element
}

func New(inputType InputType) *Input {
	input := tinydom.GetDocument().CreateElement("input")
	input.SetAttribute("type", string(inputType))
	return &Input{input}
}

func NewTextInput() *Input {
	return New(TextInput)
}

func (i *Input) Value() string {
	return i.Get("value").String()
}

func (i *Input) SetValue(v string) {
	i.Set("value", v)
}

func (i *Input) Name() string {
	return i.Get("name").String()
}

func (i *Input) SetName(n string) {
	i.Set("name", n)
}

func (i *Input) For() string {
	return i.Get("for").String()
}

func (i *Input) SetFor(value string) {
	i.Set("for", value)
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
	i.Set("onClick", function)
}
