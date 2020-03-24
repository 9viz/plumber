package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

// Is path a valid file?
func IsFile(path string) bool {
	inf, err := os.Stat(path)
	return err == nil && !inf.IsDir()
}

// Is path a valid directory?
func IsDir(path string) bool {
	inf, err := os.Stat(path)
	return err == nil && inf.IsDir()
}

// Is name(section) a man page? Section is a string
// because manpages can have non numerical sections.
// Examples include perl's manpage.
func IsManPage(name string, section string) bool {
	// MANPATH isn't set by every distro (read: NixOS)
	// Investigate a way other than calling apropos and checking the exit code
	for _, m := range strings.Split(os.Getenv("MANPATH"), ":") {
		// MANPATH/manSECTION/name.SECTION
		if IsFile(m + "/man" + section + "/" + name + "." + section) {
			return true
		}
	}
	return false
}

// Is the string in the format of a manpage i.e.,
// is the string in the format name(section)?
func IsMan(str string) bool {
	m := strings.Split(str, "(")
	return len(m) == 2 && string(m[1][len(m[1])-1]) == ")"
}

// Parse string into name and section
// Returns an error when the string is not manpage-like
// Check IsMan to know manpage-like means
func ParseMan(str string) (string, string, error) {
	if !IsMan(str) {
		return "", "", fmt.Errorf("utils.go/ParseMan: %s is not manpage-like", str)
	}
	m := strings.Split(str, "(")
	return m[0], m[1][:len(m[1])-1], nil
}

// Is the string like a URL that can be viewed in a browser?
// It has to start with http://, https:// or ftp://
func IsUrl(str string) bool {
	b := false
	for _, p := range []string{"https://", "http://", "ftp://"} {
		b = b || strings.HasPrefix(str, p)
	}
	return b
}

// Get the mimetype of file residing in path
// Returns an error if the path is not a valid path or
// if the contents of the file cannot be read
func GetMimeType(path string) (string, error) {
	if !IsFile(path) {
		return "", fmt.Errorf("utils.go/GetMimeType: %s is not a file", path)
	}

	// Error is safe to ignore now
	f, _ := os.Open(path)

	buf := make([]byte, 50)
	_, err := f.Read(buf)
	if err != nil {
		return "", err
	}
	defer f.Close()

	return strings.Split(http.DetectContentType(buf), ";")[0], nil
}