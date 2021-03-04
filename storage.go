package tinydom

import "syscall/js"

type Storage struct {
	js.Value
}

var LocalStorage = &Storage{GetWindow().Get("localStorage")}
var SessionStorage = &Storage{GetWindow().Get("sessionStorage")}

func (s *Storage) Length() int {
	return s.Get("length").Int()
}

func (s *Storage) Key(index int) string {
	return s.Call("key", index).String()
}

func (s *Storage) GetItem(key string) string {
	return s.Call("getItem", key).String()
}

func (s *Storage) SetItem(key, value string) {
	s.Call("setItem", key, value)
}

func (s *Storage) RemoveItem(key string) {
	s.Call("removeItem", key)
}

func (s *Storage) Clear() {
	s.Call("clear")
}

func (s *Storage) KeyExists(key string) bool {
	return !s.Call("getItem", key).IsNull()
}
