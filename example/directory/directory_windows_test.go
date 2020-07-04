package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestPathDirectory(t *testing.T) {
	s, err := PathDirectory(Video)
	if err != nil {
		t.Fatal(err)
	}

	if d := filepath.Join(os.Getenv(`USERPROFILE`), "Videos"); s != d {
		t.Log("Needs manual inspection")
		t.Log("Current folder is:", s)
		t.Log("Default folder is:", d)
		t.Log("It's not an error if you have change the default folder")
	}
}