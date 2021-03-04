package main

import (
	"github.com/Nerzal/tinydom"
	"github.com/Nerzal/tinydom/elements"
)

func main() {
	document := tinydom.GetDocument()

	body := document.GetElementById("body-component")

	h1 := document.CreateElement("h1")
	h1.SetInnerHTML("Welcome to tinydom - Hello TinyWorld <3")
	body.AppendChild(h1)

	h2 := document.CreateElement("h1")
	h2.SetInnerHTML("Yes! I do compile with TinyGo!")
	body.AppendChild(h2)

	input := elements.NewTextInput()
	body.AppendChild(input.Element)

	wait := make(chan struct{}, 0)
	<-wait
}
