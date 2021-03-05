package input

type InputType string

const (
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
