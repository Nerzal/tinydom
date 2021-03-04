package tinydom

import "syscall/js"

type Navigator struct {
	js.Value
}

func (n *Navigator) Language() string {
	return n.Get("language").String()
}

func (n *Navigator) Languages() string {
	return n.Get("languages").String()
}
