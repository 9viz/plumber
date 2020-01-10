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

// Is s a man page?
func isMan(s string) bool {
	p := strings.Split(s, "(")
	if len(p) == 1 {
		return false
	}

	cmd := exec.Command("apropos", "-f", p[0])
	err := cmd.Run()
	if err != nil {
		panic(err)
	}

	return cmd.ProcessState.ExitCode() == 0
}

// "Parse" man string and return name and section
func parseMan(s string) (string, string) {
	p := strings.Split(s, "(")
	e := p[1][:len(p[1])-1] // Everything except )
	return p[0], e
}
