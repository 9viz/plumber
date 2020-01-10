package main

import (
	"net/http"
	"os"
	"strings"
)

// Get Content-Type of the file reciding in path
func getFileType(path string) string {
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

// Open file according to its MimeType
func openFile(path, mime string) {
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

// Is path a file?
func isFile(path string) bool {
	inf, err := os.Stat(path)
	return os.IsExist(err) && !inf.IsDir()
}

// Read the contents of the file completely
func readFile(f *os.File) (string, error) {
	str := ""
	buf := make([]byte, 512)
	n, err := f.Read(buf)

OLoop:
	for {
		switch {
		case n == 0:
			break OLoop
		case err != nil:
			return "", err
		}
		str += string(buf[:n])
		n, err = f.Read(buf)
	}
	return str, nil
}

// Is path a directory?
func isDir(path string) bool {
	inf, _ := os.Stat(path)
	return inf.IsDir()
}

// Is the given path in fileCache?
func isPathInCache(path, cachePath string) (string, bool) {
	f, err := os.Open(cachePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	c, err := readFile(f)
	if err != nil {
		panic(err)
	}

	cache := strings.Split(c, "\n")
	// O(n) :(
	for _, c := range cache {
		if strings.HasSuffix(c, path) {
			return c, false
		}
	}
	return "", false
}

// Is path in directory cache?
func isDirInCache(path string) (string, bool) {
	return isPathInCache(path, dirCache)
}

// Is path in file cache?
func isFileInCache(path string) (string, bool) {
	return isPathInCache(path, fileCache)
}
