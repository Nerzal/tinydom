package input

import "github.com/Nerzal/tinydom"

type DataList struct {
	*tinydom.Element
}

func NewDataList(id string, options []string) *DataList {
	doc := tinydom.GetDocument()

	result := doc.CreateElement("datalist")
	result.SetId(id)

	for i := range options {
		option := options[i]

		optElem := doc.CreateElement("option")
		optElem.Set("value", option)
		result.AppendChild(optElem)
	}

	return &DataList{result}
}
