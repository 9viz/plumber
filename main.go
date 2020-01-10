package main

import (
	"io"
	"net/http"
	"os"
	"strings"
)

// Get string from user. Could be argv[1:] or PLUMB env var
func getString() string {
	if len(os.Args) > 1 {
		return strings.Join(os.Args[1:], " ")
	}
	return os.Getenv("PLUMB")
}

// Handle the url - opens it in the desired application
func handleHttp(url string) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	ct := strings.Split(resp.Header.Get("Content-Type"), ";")[0]
	switch {
	case ct == "text/html":
		shExec(App[ct] + url)
	case hasAnyPrefix(ct, []string{"audio/", "video/"}):
		shExec("mpv " + url)
	default:
		defer resp.Body.Close()
		f, err := os.Create("/tmp/plumb")
		if err != nil {
			panic(err)
		}
		defer f.Close()

		_, err = io.Copy(f, resp.Body)
		if err != nil {
			panic(err)
		}
		openFile("/tmp/plumb", ct)
	}
}

// Try to open the file. Returns true if succeded
func handleFile(path string) bool {
	if isFile(path) {
		openFile(path, getFileType(path))
		return true
	}
	p, b := isFileInCache(path)
	if b {
		openFile(p, getFileType(p))
		return true
	}
	return false
}

// Try to open directory in terminal. Returns true if succeded
func handleDir(path string) bool {
	if isDir(path) {
		err := os.Chdir(path)
		if err != nil {
			panic(err)
		}
		shExec(App["term"])
		return true
	}
	p, b := isDirInCache(path)
	if b {
		err := os.Chdir(p)
		if err != nil {
			panic(err)
		}
		shExec(App["term"])
		return true
	}
	return false
}

// Other stuff that could be done with the string
func other(s string) {
	switch {
	case isMan(s):
		n, s := parseMan(s)
		shExec(App["man"] + " " + s + " " + n)
	case handleFile(s):
		os.Exit(0)
	case handleDir(s):
		os.Exit(0)
	default:
		shExec(App["search"] + " " + s)
	}
}

func main() {
	str := getString()
	switch {
	case str == "":
		os.Exit(0) // quit quietly
	case strings.HasPrefix(str, "https://www.youtube.com/watch?v="):
		shExec("ytdl -o - " + str + " | mpv -")
	case hasAnyPrefix(str, []string{"http://", "https://"}):
		handleHttp(str)
	default:
		other(str)
	}
}
