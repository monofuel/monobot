package handlers

import (
	"fmt"
	"os"
)

var commandMap = make(map[string]func([]string) string)

func init() {
	commandMap["randomLink"] = randLink
	commandMap["threest"] = toDo

	commandMap["quit"] = quit
}

//SetHandler manually sets an external handler
func SetHandler(key string, f func([]string) string) {

	_, ok := commandMap[key]
	if ok {
		fmt.Println("overriding handler:", key)
	}
	commandMap[key] = f
}

//HandleCommand handle a specified command line and return the output
func HandleCommand(args []string) string {
	if len(args) == 0 {
		return "refer to help"
	}
	f, ok := commandMap[args[0]]
	if !ok {
		return "no such command"
	}
	return f(args)
}

func toDo(line []string) string {
	return "not implemented"
}

func quit(line []string) string {
	fmt.Println("ordered to quit")
	os.Exit(-1)
	return ""
}
