package handlers

import (
	"fmt"
	"os"
	"strings"
)

var commandMap = make(map[string]func([]string) string)

func init() {
	commandMap["randomLink"] = randLink
	commandMap["threest"] = toDo
	commandMap["help"] = help
	commandMap["ping"] = ping

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

func help(line []string) string {
	//list commands
	var output []string
	for key, _ := range commandMap {
		output = append(output, key)
	}
	return strings.Join(output, ",")
}

func toDo(line []string) string {
	return "not implemented"
}

func quit(line []string) string {
	fmt.Println("ordered to quit")
	os.Exit(-1)
	return ""
}
