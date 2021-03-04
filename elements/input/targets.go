package input

type Target string

const (
	Blank  Target = "_blank"
	Self   Target = "_self"
	Parent Target = "_parent"
	Top    Target = "_top"
)

func (t Target) String() string {
	return string(t)
}
