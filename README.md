# TinyDom

TinyGo compatible DOM manipulation library. For use in WASM

This library is heavily based on [godom](https://github.com/siongui/godom). It was changed to be usable in [TinyGo](https://tinygo.org) projects.

## Example Usage:

```go
package main

import (
	"github.com/Nerzal/tinydom"
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

	dateInput := input.New(input.DateInput)
	body.AppendChild(dateInput.Element)

	wait := make(chan struct{}, 0)
	<-wait
}
```

## Run the example: 

Simply use the makefile :) 

> make example-app

## Example Result:

![grafik](https://user-images.githubusercontent.com/9110370/110025878-3f5e9f00-7d30-11eb-9f9e-8cf29c495eff.png)

