package haxilang

import (
	"strings"
)

// Functioner allows to run functions with haxilang code inside
type Functioner struct {
	env  *Environment
	code strings.Builder
}

// NewFunctioner - create new Functioner
func NewFunctioner(env *Environment) *Functioner {
	return &Functioner{
		env:  env,
		code: strings.Builder{},
	}
}

// Append to Functioner code
func (f *Functioner) Append(code string) {
	f.code.WriteString(code)
	if !strings.HasSuffix(code, "\n") {
		f.code.WriteString("\n")
	}
}

// GetFunction - returns function based function
func (f *Functioner) GetFunction() Command {
	e2 := f.env.CreateSubEnvironment()
	codeStr := f.code.String()

	var result int64
	e2.SetCommand("out", func(e *Environment, args string) int64 {
		result = e2.RunCode(args)
		return result
	})
	e2.SetCommand("super", func(e *Environment, args string) int64 {
		return f.env.RunCode(args)
	})
	e2.SetCommand("superSet", func(e *Environment, args string) int64 {
		arr := strings.SplitN(args, " ", 2)
		if len(arr) != 2 {
			return 0
		}
		f.env.SetVariable(arr[0], e2.RunCode(arr[1]))
		return 1
	})

	return func(e *Environment, args string) int64 {
		argnArray := strings.Split(args, " ")
		e2.SetCommand("argn", func(e *Environment, argnArgs string) int64 {
			num := e2.RunCode(argnArgs) - 1
			if num < 0 || int(num) > len(argnArray)-1 {
				return 0
			}
			return e.RunCode(argnArray[num])
		})
		e2.SetCommand("arg", func(e *Environment, _ string) int64 {
			return e.RunCode(args)
		})
		e2.SetCommand("argc", func(e *Environment, _ string) int64 {
			return int64(len(argnArray))
		})
		e2.RunCode(codeStr)
		return result
	}
}
