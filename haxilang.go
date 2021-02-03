package haxilang

import (
	"strconv"
	"strings"
)

// Environment - need to have memory for the language
type Environment struct {
	parent *Environment
	mem    map[string]int64
	cmds   map[string]Command
	stack  *Stack
}

// Command for the lang
type Command func(e *Environment, args string) int64

// NewEnvironment - create new environment
func NewEnvironment() *Environment {
	return &Environment{
		nil,
		make(map[string]int64),
		make(map[string]Command),
		NewStack(),
	}
}

// RunCode of code
func (e *Environment) RunCode(line string) int64 {
	if strings.Contains(line, "\n") {
		line = strings.ReplaceAll(line, "\r", "")
		line = strings.Trim(line, " \t")
		lines := strings.Split(line, "\n")
		for i := 0; i < len(lines); i++ {
			e.RunCode(lines[i])
		}
		return 0
	}
	a := Parse(line)
	if a == nil {
		return 0
	}
	if a.Command == "" {
		if strings.HasPrefix(a.Args, "$") {
			return e.GetVariable(a.Args[1:])
		}
		i, err := strconv.ParseInt(a.Args, 10, 64)
		if err != nil {
			return 0
		}
		return i
	}
	cmd := e.GetCommand(a.Command)
	if cmd == nil {
		return 0
	}
	return cmd(e, a.Args)
}

// SetCommand used to add commands to environment
func (e *Environment) SetCommand(name string, cmd Command) {
	e.cmds[name] = cmd
}

// GetCommand used to get commands from environment
func (e *Environment) GetCommand(name string) Command {
	cmd, ok := e.cmds[name]
	if !ok {
		if e.parent != nil {
			cmd := e.parent.GetCommand(name)
			return cmd
		}
		return nil
	}
	return cmd
}

// SetVariable for Environment
func (e *Environment) SetVariable(name string, value int64) {
	e.mem[name] = value
}

// GetVariable from Environment
func (e *Environment) GetVariable(name string) int64 {
	if name == "" {
		return 0
	}
	n, ok := e.mem[name]
	if !ok {
		if e.parent != nil {
			return e.parent.GetVariable(name)
		}
		return 0
	}
	return n
}

// GetStack - Allows to get Environment stack
func (e *Environment) GetStack() *Stack {
	return e.stack
}

// CreateSubEnvironment - Creates Environment which inherits existing one
func (e *Environment) CreateSubEnvironment() *Environment {
	env := NewEnvironment()
	env.parent = e
	return env
}

// GetParent - returns the Parent Environment of current
func (e *Environment) GetParent() *Environment {
	return e.parent
}
