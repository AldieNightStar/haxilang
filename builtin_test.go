package haxilang

import "testing"

const mathErr = "Math inner functional error"

func TestSetAndMath(t *testing.T) {
	e := NewEnvironment()
	InitEnvironment(e)

	n := int64(0)
	n += e.RunCode("/set a 100\n/set b 200\n/set c /math $a + $b")

	if e.GetVariable("a") != 100 || e.GetVariable("b") != 200 {
		t.Fatal("Built in set is not working")
	}

	if e.GetVariable("c") != 300 {
		t.Fatal("Built in adding failed")
	}

	if n != 0 {
		t.Fatal("Return value should be 0")
	}
}

func TestMath(t *testing.T) {
	e := NewEnvironment()
	InitEnvironment(e)
	e.SetVariable("a", 124)
	e.SetVariable("b", 20)

	runAndValidateReturn(t, e, "/math 2 + 2", mathErr, 4)
	runAndValidateReturn(t, e, "/math 3 - 2", mathErr, 1)
	runAndValidateReturn(t, e, "/math 3 * 2", mathErr, 6)
	runAndValidateReturn(t, e, "/math 8 / 2", mathErr, 4)
	runAndValidateReturn(t, e, "/math 8 / 4", mathErr, 2)
	runAndValidateReturn(t, e, "/math 9 % 4", mathErr, 1)
	runAndValidateReturn(t, e, "/math $a + 10", mathErr, 134)
	runAndValidateReturn(t, e, "/math $a - $b", mathErr, 104)
	runAndValidateReturn(t, e, "/math $a * $b", mathErr, 2480)
	runAndValidateReturn(t, e, "/math $a / 2", mathErr, 62)
	runAndValidateReturn(t, e, "/math $a % 20", mathErr, 4)
	runAndValidateReturn(t, e, "/math $b - 5", mathErr, 15)
	runAndValidateReturn(t, e, "/math $b ^ 2", mathErr, 400)
}

func runAndValidateReturn(t *testing.T, e *Environment, code, messageIfFail string, expNumber int64) {
	n := e.RunCode(code)
	if n != expNumber {
		t.Fatal(messageIfFail + ". Code: " + code)
	}
}

func TestGet(t *testing.T) {
	e := NewEnvironment()
	InitEnvironment(e)

	e.SetVariable("x", 12)

	n := e.GetVariable("x")
	n += e.RunCode("/get x")

	if n != 24 {
		t.Fatal("Number is not valid")
	}
}
