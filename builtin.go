package haxilang

import (
	"fmt"
	"math"
	"strings"
)

type _stdCommand func(e *Environment, args string) int64
type _stdMathOp func(a, b int64) int64

var mathOps = map[string]_stdMathOp{
	"+": func(a, b int64) int64 { return a + b },
	"-": func(a, b int64) int64 { return a - b },
	"*": func(a, b int64) int64 { return a * b },
	"/": func(a, b int64) int64 { return a / b },
	"%": func(a, b int64) int64 { return a % b },
	"^": func(a, b int64) int64 { return int64(math.Pow(float64(a), float64(b))) },
}

func _stdPrint(e *Environment, args string) int64 {
	fmt.Printf("%d\n", e.RunCode(args))
	return 0
}

func _stdSet(e *Environment, args string) int64 {
	arr := strings.SplitN(args, " ", 2)
	if (len(arr)) != 2 {
		return 0
	}
	e.SetVariable(arr[0], e.RunCode(arr[1]))
	return 1
}

func _stdGet(e *Environment, args string) int64 {
	return e.GetVariable(args)
}

func _stdMath(e *Environment, args string) int64 {
	arr := strings.Split(args, " ")
	if len(arr) != 3 {
		if len(arr) == 1 {
			return e.RunCode(arr[0])
		}
		return 0
	}
	mathOp, ok := mathOps[arr[1]]
	if !ok {
		return 0
	}
	return mathOp(e.RunCode(arr[0]), e.RunCode(arr[2]))
}

func _stdPush(e *Environment, args string) int64 {
	e.stack.Push(e.RunCode(args))
	return 0
}

func _stdPop(e *Environment, args string) int64 {
	el := e.stack.Pop()
	n, ok := el.(int64)
	if !ok {
		return 0
	}
	return n
}

// InitEnvironment - Initialize standard functions for the language
func InitEnvironment(e *Environment) {
	e.SetCommand("print", _stdPrint)
	e.SetCommand("set", _stdSet)
	e.SetCommand("get", _stdGet)
	e.SetCommand("math", _stdMath)
	e.SetCommand("push", _stdPush)
	e.SetCommand("pop", _stdPop)

	// ================
	// Special Commands
	// ================

	// Function creation
	// ===================
	var functioner *Functioner
	var functionerName string
	e.SetCommand("func", func(e *Environment, args string) int64 {
		if args == "" {
			return 0
		}
		functionerName = args
		functioner = NewFunctioner(e)
		return 1
	})
	e.SetCommand("+", func(e *Environment, args string) int64 {
		if args == "" {
			return 0
		}
		functioner.Append(args)
		return 1
	})
	e.SetCommand("endfunc", func(e *Environment, args string) int64 {
		if functioner == nil || functionerName == "" {
			return 0
		}
		f := functioner.GetFunction()
		e.SetCommand(functionerName, f)

		functioner = nil
		functionerName = ""
		return 1
	})
}
