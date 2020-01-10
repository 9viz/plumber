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

func main() {
	str := getString()

	switch {
	case str == "":
		os.Exit(0) // quit quietly
	case strings.HasPrefix(str, "https://www.youtube.com/watch?v="):
		shExec("ytdl -o - " + str + " | mpv -")
	case hasAnyPrefix(str, []string{"http://", "https://"}):
		handleHttp(str)
	// TODO: check if str is in file,dir cache or a man page
	// case isfile, isdir:
	default:
		shExec(App["search"] + " " + str)
	}
}
