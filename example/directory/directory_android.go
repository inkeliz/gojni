// +build android

package main

import (
	"github.com/inkeliz/gojni"
)

//go:generate javac -source 8 -target 8 -bootclasspath $ANDROID_HOME\platforms\android-29\android.jar -classpath ..\..\register_android.jar -d classes directory_android.java
//go:generate jar cf directory_android.jar -C classes .

func pathDirectory(c Environment) (path string, err error) {
	sess, err := jni.NewSessionRegistered("com/inkeliz/example/directory")
	if err != nil {
		return path, err
	}

	function := androidEnvironment(c)

 	// If function is "Download" it will call "com/inkeliz/example/directory . Download()":
	return jni.NewCall(sess, function).ReturnString(nil)
}

func androidEnvironment(c Environment) string {
	switch c {
	case Download:
		return "Download"
	case Pictures:
		return "Pictures"
	case Music:
		return "Music"
	case Video:
		return "Video"
	case Documents:
		return "Document"
	default:
		return "Document"
	}
}
