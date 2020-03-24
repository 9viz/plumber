package main

import (
	"os"
	"strings"
)

// Is path a valid file?
func IsFile(path string) bool {
	inf, err := os.Stat(path)
	return err == nil && !inf.IsDir()
}

// Is name(section) a man page? Section is a string
// because manpages can have non numerical sections.
// Examples include perl's manpage.
func IsManPage(name string, section string) bool {
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
