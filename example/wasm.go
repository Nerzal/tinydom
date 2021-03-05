package main

import (
	"github.com/Nerzal/tinydom"
	"github.com/Nerzal/tinydom/elements/form"
	"github.com/Nerzal/tinydom/elements/href"
	"github.com/Nerzal/tinydom/elements/input"
	"github.com/Nerzal/tinydom/elements/media"
)

func main() {
	document := tinydom.GetDocument()

	body := document.GetElementById("body-component")

	h1 := document.CreateElement("h1")
	h1.SetInnerHTML("Welcome to tinydom - Hello TinyWorld <3")
	body.Br()

	h2 := document.CreateElement("h1")
	h2.SetInnerHTML("Yes! I do compile with TinyGo!")
	body.Br()

	body.Br()

	body.Br()

	myForm := form.New()
	label := document.CreateElement("label")
	label.SetInnerHTML("Name:")
	textInput := input.NewTextInput()

	passwordLabel := document.CreateElement("label")
	passwordLabel.SetInnerHTML("Password:")
	passwordInput := input.New(input.PasswordInput)

	submitInput := input.New(input.SubmitInput)

	err := myForm.Append(label, textInput.Element, passwordLabel, passwordInput.Element, submitInput.Element)
	if err != nil {
		println(err.Error())
	}

	body.AppendChild(myForm.Element)

	body.Br()
	body.Br()

	video := media.NewVideoParams(640, 360, true, false, true, &media.VideoSource{
		Source: "video.mp4",
		Type:   media.MP4,
	})

	body.AppendChild(video.Element)
	body.Br()
	body.Br()

	link := href.New("https://www.bigbuckbunny.org/", "Big Buck Bunny")
	body.AppendChild(link.Element)

	wait := make(chan struct{}, 0)
	<-wait
}
