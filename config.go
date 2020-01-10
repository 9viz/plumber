package main

var (
	App = map[string]string{
		"text/html":       "palemoon --new-window",
		"image/*":         "meh",
		"text/*":          "st -g 80x40+500+250 -e less",
		"application/pdf": "zathura",
		"audio/*":         "mpv",
		"video/*":         "mpv",
		"man":             "st -g 80x40+500+250 -e man",
		"search":          "palemoon --new-window --search",
	}

	fileCache = "/home/viz/usr/local/share/cache/plumb/filec"
	dirCache  = "/home/viz/usr/local/share/cache/plumb/dirc"
)
