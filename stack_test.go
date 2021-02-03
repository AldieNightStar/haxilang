package haxilang

import "testing"

func TestStack(t *testing.T) {
	e := NewEnvironment()
	InitEnvironment(e)

	e.RunCode("/push 32\n/push 12\n/set a /pop\n/set b /pop")

	a := e.GetVariable("a")
	b := e.GetVariable("b")
	if a != 12 || b != 32 {
		t.Fatal("Stack works not right")
	}
}
