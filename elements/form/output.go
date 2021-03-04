package form

import "github.com/Nerzal/tinydom"

type Output struct {
	*tinydom.Element
}

func NewOutput(name string, forValue ...string) *Output {
	output := tinydom.GetDocument().CreateElement("output")
	output.Set("name", name)

	if forValue != nil {
		forAttribute := ""

		for i, value := range forValue {
			forAttribute += value
			if i != len(forValue) {
				forAttribute += " "
			}
		}
	}

	return &Output{output}
}
