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

func (i *Input) SetValue(s string) {
	i.Set("value", s)
}
