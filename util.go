package main

import (
	"net/http"
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

// Get Content-Type of http.Response without charset
func getContentType(resp *http.Response) string {
	return strings.Split(resp.Header.Get("Content-Type"), ";")[0]
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
