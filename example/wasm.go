package main

import (
	"github.com/Nerzal/tinydom"
	"github.com/Nerzal/tinydom/elements/form"
	"github.com/Nerzal/tinydom/elements/input"
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

	br := document.CreateElement("br")
	body.AppendChild(br)
	body.AppendChild(br)

	textInput := input.NewTextInput()
	body.AppendChild(textInput.Element)

	body.AppendChild(br)
	body.AppendChild(br)

	body.AppendChild(br)
	body.AppendChild(br)
	body.AppendChild(br)
	body.AppendChild(br)

	myForm := form.New()
	label := document.CreateElement("label")
	label.SetInnerHTML("Name")
	myForm.AppendChildBr(label)
	myForm.AppendChildBr(textInput.Element)

	passwordLabel := document.CreateElement("label")
	passwordInput := input.New(input.PasswordInput)
	myForm.AppendChildBr(passwordLabel)
	myForm.AppendChildBr(passwordInput.Element)

	body.AppendChild(myForm.Element)

	wait := make(chan struct{}, 0)
	<-wait
}
