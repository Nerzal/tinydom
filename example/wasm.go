package main

import tinydom "github.com/Nerzal/TinyDom"

func main() {
	document := tinydom.GetDocument()

	body := document.GetElementById("body-component")

	wait := make(chan struct{}, 0)
	<-wait
}
