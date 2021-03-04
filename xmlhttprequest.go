package tinydom

import "syscall/js"

type XMLHttpRequest struct {
	js.Value
}

func NewXMLHttpRequest() *XMLHttpRequest {
	return &XMLHttpRequest{GetWindow().Get("XMLHttpRequest").New()}
}

func (x *XMLHttpRequest) ResponseText() string {
	return x.Get("responseText").String()
}

func (x *XMLHttpRequest) ResponseURL() string {
	return x.Get("responseURL").String()
}

func (x *XMLHttpRequest) ResponseXML() *Element {
	return &Element{x.Get("responseXML")}
}

func (x *XMLHttpRequest) StatusText() string {
	return x.Get("statusText").String()
}

func (x *XMLHttpRequest) WithCredentials() bool {
	return x.Get("withCredentials").Bool()
}

func (x *XMLHttpRequest) SetWithCredentials(c bool) {
	x.Set("withCredentials", c)
}

func (x *XMLHttpRequest) Abort() {
	x.Call("abort")
}

func (x *XMLHttpRequest) Open(method, url string, args ...interface{}) {
	if len(args) == 1 {
		// args[0] should be bool async
		x.Call("open", method, url, args[0])
	} else {
		x.Call("open", method, url)
	}
}

func (x *XMLHttpRequest) OverrideMimeType(mimeType string) {
	x.Call("overrideMimeType", mimeType)
}

func (x *XMLHttpRequest) Send(args ...interface{}) {
	if len(args) == 1 {
		x.Call("send", args[0])
	} else {
		x.Call("send")
	}
}

func (x *XMLHttpRequest) SetRequestHeader(header, value string) {
	x.Call("setRequestHeader", header, value)
}
