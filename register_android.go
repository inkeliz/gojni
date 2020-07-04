package jni

//go:generate javac -source 8 -target 8 -bootclasspath $ANDROID_HOME\platforms\android-29\android.jar -d classes register_android.java
//go:generate jar cf register_android.jar -C classes .

/*
#cgo LDFLAGS: -landroid

#include <jni.h>
#include <stdlib.h>
#include <jni_android.h>
*/
import "C"
import (
	"strings"
)

var classes = make(map[string]cachedEntry)

type cachedEntry struct {
	Class C.jobject
	JVM   *C.JavaVM
}

//export Java_github_com_inkeliz_gojni_register_1android_Register
func Java_github_com_inkeliz_gojni_register_1android_Register(env *C.JNIEnv, class C.jclass, p C.jobject) {
	cn := new(JClass)
	if err := Exec(&Env{Env: env}, p, `getClass`, cn, nil); err != nil {
		panic(err)
	}

	sn := new(String)
	if err := Exec(&Env{Env: env}, C.jobject(cn.Data), `getName`, sn, nil); err != nil {
		panic(err)
	}

	sn.Data = strings.Replace(sn.Data, ".", "/", -1)

	classes[sn.Data] = cachedEntry{
		Class: C.jobject(C.NewGlobalRef(env, p)),
		JVM:   GetJavaVM(&Env{Env: env}),
	}
}
