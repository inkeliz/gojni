package main

type Environment uint8

const (
	Pictures Environment = iota
	Download
	Music
	Video
	Documents
)

func PathDirectory(c Environment) (path string, err error) {
	return pathDirectory(c)
}