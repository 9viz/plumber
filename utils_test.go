package main

import (
	"fmt"
	"testing"
)

func TestIsDir(t *testing.T) {
	testData := []struct {
		dir      string
		expected bool
	}{
		{"/root", true},
		{"/home/viz", true},
		{"/home/viz/.profile", false},
		{"/home/viz/lib", true},
	}

	for _, d := range testData {
		t.Run(fmt.Sprintf("%s", d.dir), func(t *testing.T) {
			res := IsDir(d.dir)
			if res != d.expected {
				t.Errorf("Expected %t, got %t", d.expected, res)
			}
		})
	}
}

func TestIsManPage(t *testing.T) {
	testData := []struct {
		name, section string
		expected      bool
	}{
		{"st", "1", true},
		{"dmenu", "1", true},
		{"go", "1", false},
		{"st", "2", false},
	}

	for _, d := range testData {
		t.Run(fmt.Sprintf("%s.%s", d.name, d.section), func(t *testing.T) {
			res := IsManPage(d.name, d.section)
			if res != d.expected {
				t.Errorf("Expected %t, got %t", d.expected, res)
			}
		})
	}
}

func TestIsMan(t *testing.T) {
	testData := []struct {
		str      string
		expected bool
	}{
		{"st(1)", true},
		{"dmenu(asd", false},
		{"go", false},
		{"st(1p)", true},
	}

	for _, d := range testData {
		t.Run(fmt.Sprintf("%s", d.str), func(t *testing.T) {
			res := IsMan(d.str)
			if res != d.expected {
				t.Errorf("Expected %t, got %t", d.expected, res)
			}
		})
	}
}

func TestParseMan(t *testing.T) {
	testData := []struct {
		str      string
		expected bool
	}{
		{"st(1)", true},
		{"dmenu(asd", false},
		{"go", false},
		{"st(1p)", true},
	}

	for _, d := range testData {
		t.Run(fmt.Sprintf("%s", d.str), func(t *testing.T) {
			_, _, err := ParseMan(d.str)
			res := err == nil
			if res != d.expected {
				t.Errorf("Expected %t, got %t (error %s)", d.expected, res, err)
			}
		})
	}
}

func TestIsUrl(t *testing.T) {
	testData := []struct {
		url      string
		expected bool
	}{
		{"https://youtube.com/", true},
		{"asdjaskdj", false},
		{"/home/viz/.config", false},
		{"ftp://youtube.com", true},
		{"http://youtube.com", true},
	}

	for _, d := range testData {
		t.Run(fmt.Sprintf("%s", d.url), func(t *testing.T) {
			res := IsUrl(d.url)
			if res != d.expected {
				t.Errorf("Expected %t, got %t", d.expected, res)
			}
		})
	}
}

func TestGetMimeType(t *testing.T) {
	testData := []struct {
		path, expected string
	}{
		{"/home/viz/src/go/plumber/utils.go", "text/plain"},
		{"/home/viz/med/img/art/ta2017_the_letter.jpg", "image/jpeg"},
	}

	for _, d := range testData {
		t.Run(fmt.Sprintf("%s", d.path), func(t *testing.T) {
			res, err := GetMimeType(d.path)
			if res != d.expected || err != nil {
				t.Errorf("Expected %s, got %s (error %s)", d.expected, res, err)
			}
		})
	}
}