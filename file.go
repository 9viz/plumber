package main

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
