package haxilang

import "testing"

func TestGeneral(t *testing.T) {
	e := NewEnvironment()
	InitEnvironment(e)

	code := `
/set a 300
/set b 400
/set c 300

/func multipleAdd
/+ /superSet assert_argc_res /argc
/+ /set a /argn 1
/+ /set b /argn 2
/+ /set c /argn 3
/+ /set r /math $a + $b
/+ /set r /math $r + $c
/+ /out $r
/endfunc

/set result /multipleAdd $a $b $c 0
`

	e.RunCode(code)

	if e.GetCommand("multipleAdd") == nil {
		t.Fatal("Function did not created successfuly")
	}

	if e.GetVariable("assert_argc_res") != 4 {
		t.Fatal("Assert argc fails")
	}
	if e.GetVariable("result") != 1000 {
		t.Fatal("Final result is invalid")
	}
	if e.GetVariable("r") != 0 {
		t.Fatal("Variable 'r' is damaged")
	}
}
