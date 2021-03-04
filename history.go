package tinydom

import (
	"syscall/js"
)

type History struct {
	js.Value
}

func (h *History) Length() int {
	return h.Get("length").Int()
}

func (h *History) Back() {
	h.Call("back")
}

func (h *History) Forward() {
	h.Call("forward")
}

func (h *History) Go(p int) {
	h.Call("go", p)
}

func (h *History) PushState(state interface{}, title, url string) {
	h.Call("pushState", state, title, url)
}

func (h *History) ReplaceState(state interface{}, title, url string) {
	h.Call("replaceState", state, title, url)
}
