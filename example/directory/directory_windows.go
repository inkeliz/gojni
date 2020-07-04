//+build windows

package main

import (
	"golang.org/x/sys/windows/registry"
)

func pathDirectory(c Environment) (path string, err error) {
	k, err := registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\Microsoft\Windows\CurrentVersion\Explorer\Shell Folders`, registry.QUERY_VALUE)
	if err != nil {
		return path, err
	}

	defer k.Close()

	path, _, err = k.GetStringValue(winEnvironment(c))
	if err != nil {
		return path, err
	}

	return path, nil
}

func winEnvironment(c Environment) string {
	switch c {
	case Download:
		return "{374DE290-123F-4565-9164-39C4925E467B}"
	case Pictures:
		return "My Pictures"
	case Music:
		return "My Music"
	case Video:
		return "My Video"
	case Documents:
		return "Personal"
	default:
		return "Personal"
	}
}
