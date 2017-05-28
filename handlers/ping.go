package handlers

import (
	"fmt"
	"os/exec"
	"strings"
)

func ping(args []string) string {
	if len(args) < 2 {
		return "please give a hostname"
	}
	host := args[1]
	fmt.Println(host)
	out, err := exec.Command("ping", host, "-c", "4", "-w", "10").Output()
	if len(out) == 0 && err != nil {
		return fmt.Sprintf("ping error: %v", err)
	}
	lines := strings.Split(string(out), "\n")
	return "ping results: " + host + " " + lines[len(lines)-3] + " " + lines[len(lines)-2]
}
