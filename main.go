package main

import (
	"io"
	"net/http"
	"os"
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

// Get string from user. Could be argv[1:] or PLUMB env var
func getString() string {
	if len(os.Args) > 1 {
		return strings.Join(os.Args[1:], " ")
	}
	return os.Getenv("PLUMB")
}

// Get Content-Type of the file reciding in path
func getFileContentType(path string) string {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	buf := make([]byte, 50)
	_, err = f.Read(buf)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	return strings.Split(http.DetectContentType(buf), ";")[0]
}

// execute Cmd in the background through sh
func shExec(scmd string) {
	// TODO: DONT USE SH
	cmd := exec.Command("sh", "-c", scmd)
	err := cmd.Start()
	if err != nil {
		panic(err)
	}
}

// Open file according to its MimeType
func openFile(path string) {
	mime := getFileContentType(path)

OLoop:
	for k, v := range App {
		switch {
		case strings.HasSuffix(k, "*"):
			if strings.HasPrefix(mime, k[:len(k)-2]) {
				shExec(v + " " + path)
				break OLoop
			}
		default:
			if mime == k {
				shExec(v + " " + path)
				break OLoop
			}
		}
	}
}

// Handle the url - opens it in the desired application
func handleHttp(url string) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	ct := getContentType(resp)
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
		openFile("/tmp/plumb")
	}
}

func main() {
	str := getString()

	switch {
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
