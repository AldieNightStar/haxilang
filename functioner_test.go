package haxilang

import "testing"

func TestFunctionerCall(t *testing.T) {
	e := NewEnvironment()
	InitEnvironment(e)
	e.SetCommand("x", func(e *Environment, args string) int64 { return 12 })

	f := NewFunctioner(e)

	f.Append("/set a /x")
	f.Append("/out /get a")

	e.SetCommand("fnc", f.GetFunction())

	n := e.RunCode("/fnc")
	if n != 12 {
		t.Fatal("Invalid value")
	}
}

func TestFunctionerFunctionCreation(t *testing.T) {
	e := NewEnvironment()
	InitEnvironment(e)

	code := `/func add
/+ /set a /argn 1
/+ /set b /argn 2
/+ /super /set c 3
/+ /out /math $a + $b
/endfunc

/set x /add 1 55`
	e.RunCode(code)

	if e.GetVariable("x") != 56 {
		t.Fatal("Invalid value of evaluation")
	}

	if e.GetVariable("a")+e.GetVariable("b") != 0 {
		t.Fatal("Values a/b are damaged")
	}

	if e.GetVariable("c") != 3 {
		t.Fatal("Variable 'c' did not set by 'super' command")
	}
}
