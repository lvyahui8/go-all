package main

import "testing"

type Api interface {
	SetName(name string)
}

type Member struct {
	name string
}

func (m *Member) SetName(name string) {
	m.name = name
}

func handle(api Api) {
	api.SetName("hello")
}

func TestHandle(t *testing.T) {
	m := Member{}
	handle(&m)
	t.Log(m.name)
}
