package main

import (
	"os/exec"
	"strings"
)

// ormap of strings.HasPrefix
func hasAnyPrefix(s string, prefix []string) bool {
	b := false
	for _, p := range prefix {
		b = b || strings.HasPrefix(s, p)
	}
	return b
}

// execute Cmd in the background through sh
func shExec(scmd string) {
	// using sh is a design decision
	cmd := exec.Command("sh", "-c", scmd)
	err := cmd.Start()
	if err != nil {
		panic(err)
	}
}
