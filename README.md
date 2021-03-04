# TinyDom

TinyGo compatible DOM manipulation library. For use in WASM

This library is heavily based on [godom](https://github.com/siongui/godom). It was changed to be usable in [TinyGo](https://tinygo.org) projects.

## Example Usage

```go
package main

import (
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

 wait := make(chan struct{}, 0)
 <-wait
}

}
```

## Run the example

Simply use the makefile :)

> make example-app

Note: The example uses a wasm_exec.js from tinygo 1.17 :)

## Example Result

![grafik](https://user-images.githubusercontent.com/9110370/110029225-7931a480-7d34-11eb-9202-d3af100bdf98.png)
