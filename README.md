# TinyDom

TinyGo compatible DOM manipulation library. For use in WASM

This library is heavily based on [godom](https://github.com/siongui/godom). It was changed to be usable in [TinyGo](https://tinygo.org) projects.

## Example Usage:

```go
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
```

## Run the example: 

Simply use the makefile :) 

> make example-app