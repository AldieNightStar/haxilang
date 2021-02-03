package haxilang

import "strings"

// Arguments - returned by Parse(...)
type Arguments struct {
	Args    string
	Command string
}

// Parse - method to parse the line and return arguments
func Parse(line string) *Arguments {
	isCommand := false
	isSpacePresent := false
	args := &Arguments{}
	if line == "" {
		return nil
	}
	if strings.HasPrefix(line, "/") {
		isCommand = true
	}
	if strings.Contains(line, " ") {
		isSpacePresent = true
	}
	if isCommand {
		line = line[1:]
	} else {
		args.Args = line
		return args
	}
	if isSpacePresent {
		arr := strings.SplitN(line, " ", 2)
		args.Command = arr[0]
		if len(arr) > 1 {
			args.Args = arr[1]
		}
	} else {
		args.Command = line
	}
	return args
}
