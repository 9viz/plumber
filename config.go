package main

var (
	App = map[string]string{
		"text/html":       "palemoon --new-window",
		"image/*":         "meh",
		"text/*":          "tab --parent-id emacsclient -c -a ''",
		"application/pdf": "zathura",
		"audio/*":         "mpv",
		"video/*":         "mpv",
		"man":             "st -g 80x40+500+250 -e man",
		"search":          "palemoon --new-window --search",
		"term":            "st -g 80x40+500+250",
	}

	fileCache = "/home/viz/usr/local/share/cache/plumb/filec"
	dirCache  = "/home/viz/usr/local/share/cache/plumb/dirc"
)
