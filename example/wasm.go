package main

import (
	"syscall/js"

	"github.com/Nerzal/tinydom"
	"github.com/Nerzal/tinydom/elements/form"
	"github.com/Nerzal/tinydom/elements/href"
	"github.com/Nerzal/tinydom/elements/input"
	"github.com/Nerzal/tinydom/elements/label"
	"github.com/Nerzal/tinydom/elements/media"
)

var video *media.Video

func main() {
	document := tinydom.GetDocument()

	body := document.GetElementById("body-component")

	h1 := document.CreateElement("h1")
	h1.SetInnerHTML("Welcome to tinydom - Hello TinyWorld <3")
	body.AppendChildBr(h1)
	body.Br()

	h2 := document.CreateElement("h1")
	h2.SetInnerHTML("Yes! I do compile with TinyGo!")
	body.AppendChildBr(h2)
	body.Br()

	body.Br()

	body.Br()

	myForm := form.New()

	nameLabel := label.New()
	nameLabel.SetInnerHTML("Name:")
	textInput := input.NewTextInput()

	passwordLabel := label.New()
	passwordLabel.SetInnerHTML("Password:")
	passwordInput := input.New(input.PasswordInput)

	submitInput := input.New(input.SubmitInput)

	err := myForm.Append(nameLabel.Element, textInput.Element, passwordLabel.Element, passwordInput.Element, submitInput.Element)
	if err != nil {
		println(err.Error())
	}

	body.AppendChild(myForm.Element)

	body.Br()
	body.Br()
	largeButton := input.New(input.ButtonInput)
	largeButton.SetValue("Large")
	largeButton.AddEventListener("click", js.FuncOf(large))

	smallButton := input.New(input.ButtonInput)
	smallButton.SetValue("Small")
	smallButton.AddEventListener("click", js.FuncOf(small))

	body.AppendChildren(smallButton.Element, largeButton.Element)
	body.Br()

	video = media.NewVideoParams(640, 360, true, false, true, &media.VideoSource{
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

func large(this js.Value, args []js.Value) interface{} {
	println("large")
	video.SetWidth(1280)
	video.SetHeight(720)
	return nil
}

func small(this js.Value, args []js.Value) interface{} {
	println("small")
	video.SetWidth(640)
	video.SetHeight(360)
	return nil
}
