package handlers

import "strings"

var commandMap = make(map[string]func(string) string)

func init() {
	commandMap["randomLink"] = randLink
}

//HandleCommand handle a specified command line and return the output
func HandleCommand(line string) string {
	args := strings.Fields(line)
	if len(args) == 0 {
		return "refer to help"
	}
	f, ok := commandMap[args[0]]
	if !ok {
		return "no such command"
	}
	return f(line)
}
