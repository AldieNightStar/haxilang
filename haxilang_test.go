package haxilang

import "testing"

func TestCommandRegisterRight(t *testing.T) {
	argsPassed := ""

	c := func(e *Environment, args string) int64 {
		argsPassed = args
		return 32
	}
	e := NewEnvironment()
	e.SetCommand("c", c)

	n := e.RunCode("/c abc")

	if argsPassed != "abc" {
		t.Fatal("Arguments is not right")
	}
	if n != 32 {
		t.Fatal("Return value of the function isn't right")
	}

}

func TestCommandWithoutArguments(t *testing.T) {
	c := func(e *Environment, args string) int64 {
		if args != "" {
			t.Fatal("Arguments is not empty")
		}
		return 16
	}
	e := NewEnvironment()
	e.SetCommand("cmd", c)

	n := e.RunCode("/cmd")

	if n != 16 {
		t.Fatal("Return value is not right")
	}
}

func TestMultipleCommandsShouldNotReturnNumebr(t *testing.T) {
	c := func(e *Environment, a string) int64 { return 3 }
	e := NewEnvironment()
	e.SetCommand("a", c)
	e.SetCommand("b", c)

	n := e.RunCode("/a\n/b")
	if n != 0 {
		t.Fatal("Return value should be 0")
	}
}

func TestGetVariableNumber(t *testing.T) {
	e := NewEnvironment()
	e.SetVariable("a", 44)

	n := e.RunCode("$a")
	if n != 44 {
		t.Fatal("variable getting bad")
	}
}

func TestSimpleNumberNoVariable(t *testing.T) {
	e := NewEnvironment()

	n := e.RunCode("36")
	if n != 36 {
		t.Fatal("Number without variable is not valid")
	}
}

func TestMultipleFunctionsRun(t *testing.T) {
	e := NewEnvironment()

	a := 0
	f := func(e *Environment, args string) int64 {
		a++
		return 0
	}
	f2 := func(e *Environment, args string) int64 {
		a *= 100
		return 0
	}

	e.SetCommand("addOne", f)
	e.SetCommand("mul100", f2)

	n := e.RunCode("/addOne\n/addOne\n/addOne\n/mul100")
	if n != 0 {
		t.Fatal("Return value should be 0")
	}
	if a != 300 {
		t.Fatal("Result is not valid")
	}
}

func TestSubEnvironmentVariables(t *testing.T) {
	e := NewEnvironment()
	e.SetVariable("a", 100)
	e.SetVariable("b", 200)

	e2 := e.CreateSubEnvironment()
	e2.SetVariable("a", 22)

	n := e2.GetVariable("a")
	if n != 22 {
		t.Fatal("Value is wrong")
	}

	n = e.GetVariable("a")
	if n != 100 {
		t.Fatal("Value is damaged in super Environment")
	}

	n = e2.GetVariable("b")
	if n != 200 {
		t.Fatal("Value is not inherited")
	}
}

func TestSubEnvironmentCommands(t *testing.T) {
	e := NewEnvironment()
	e.SetCommand("x", func(e *Environment, args string) int64 {
		return 3
	})

	e2 := e.CreateSubEnvironment()
	n := e2.RunCode("/x")

	if n != 3 {
		t.Fatal("Value is wrong")
	}
}

func TestGetCommand(t *testing.T) {
	e := NewEnvironment()
	e.SetCommand("ax", func(e *Environment, args string) int64 {
		return 3
	})

	e2 := e.CreateSubEnvironment()

	n := e.GetCommand("ax")(nil, "")
	n += e2.GetCommand("ax")(nil, "")
	n += e2.CreateSubEnvironment().GetCommand("ax")(nil, "")

	if n != 9 {
		t.Fatal("Value from e2 is wrong")
	}
}

func TestCallCommand(t *testing.T) {
	e := NewEnvironment()

	e.SetCommand("ret", func(e *Environment, args string) int64 {
		return 5
	})

	e2 := e.CreateSubEnvironment()

	n := e.RunCode("/ret")
	n += e2.RunCode("/ret")
	n += e2.CreateSubEnvironment().RunCode("/ret")

	if n != 15 {
		t.Fatal("Call is not valid at least at one of Environments")
	}

}

func TestEmptyLines(t *testing.T) {
	e := NewEnvironment()
	abc := 0
	e.SetCommand("abc", func(e *Environment, args string) int64 { abc += 50; return 100 })

	n := e.RunCode("")
	n += e.RunCode("\n\n/abc\n\n/abc\n/abc\n\n\n\n\n\n\n\n/abc")
	if n != 0 {
		t.Fatal("Invalid value")
	}
	if abc != 200 {
		t.Fatal("Not valid function call count")
	}
}
