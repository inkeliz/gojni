//+build darwin

package main

import (
	"github.com/adrg/xdg"
)

func pathDirectory(c Environment) (path string, err error) {
	switch c {
	case Download:
		return xdg.UserDirs.Download, nil
	case Pictures:
		return xdg.UserDirs.Pictures, nil
	case Music:
		return xdg.UserDirs.Music, nil
	case Video:
		return xdg.UserDirs.Videos, nil
	case Documents:
		return xdg.UserDirs.Documents, nil
	default:
		return xdg.UserDirs.Documents, nil
	}
}
