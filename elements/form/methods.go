package form

type Method string

const (
	GET  Method = "get"
	POST Method = "post"
)

func (m Method) String() string {
	return string(m)
}
