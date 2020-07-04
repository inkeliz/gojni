package jni

/*
#cgo LDFLAGS: -landroid

#include <jni.h>
#include <stdlib.h>
#include <jni_android.h>
*/
import "C"
import (
	"unicode/utf16"
	"unsafe"
)

// Value is an interface that converts the value to Java equivalent and convert the Java to Golang.
type Value interface {
	// Go converts the given obj to an Golang value, the function might stores the value inside the struct.
	//
	// Using:
	// val := new(Bytes)
	// Exec(env, obj, method, nil, val)
	//
	// It's possible to access the []byte by using:
	// val.Data
	//
	Go(env *C.JNIEnv, obj C.jobject)

	// Java converts the original Golang value to an Java equivalent then for *C.jvalue.
	Java(env *C.JNIEnv) *C.jvalue

	// Signature returns the type used in the signature (for getMethodID).
	Signature() string
}

type Bytes struct {
	Data []byte
}

// Go reads an C.jbyteArray and stores the result on `Data` field of the struct.
func (g *Bytes) Go(env *C.JNIEnv, obj C.jobject) {
	g.Data = C.GoBytes(unsafe.Pointer(C.GetByteArrayElements(env, C.jbyteArray(obj))), C.GetArrayLength(env, C.jarray(obj)))
}

// Java transforms the `Data` field to C.jvalue
func (g *Bytes) Java(env *C.JNIEnv) *C.jvalue {
	if g.Data == nil {
		return nil
	}

	v := g.JavaByteArray(env)
	return (*C.jvalue)(unsafe.Pointer(&v))
}

// JavaByteArray transforms the `Data` field to C.jbyteArray
func (g *Bytes) JavaByteArray(env *C.JNIEnv) C.jbyteArray {
	if g.Data == nil {
		return 0
	}

	return C.NewByteArray(env, (*C.jbyte)(unsafe.Pointer(&g.Data[0])), C.int(len(g.Data)))
}

// Signature returns "[B"
func (g *Bytes) Signature() string {
	return "[B"
}

type String struct {
	Data string
}

// Go reads an C.jstring and stores the result on `Data` field of the struct.
func (g *String) Go(env *C.JNIEnv, obj C.jobject) {
	g.Data = C.GoStringN((*C.char)(unsafe.Pointer(C.GetStringUTFChars(env, C.jstring(obj)))), C.GetStringUTFLength(env, C.jstring(obj)))
}

// Java transforms the `Data` field to C.jvalue
func (g *String) Java(env *C.JNIEnv) *C.jvalue {
	if g.Data == "" {
		return nil
	}

	s := g.JavaString(env)
	return (*C.jvalue)(unsafe.Pointer(&s))
}

// JavaByteArray transforms the `Data` field to C.jstring
func (g *String) JavaString(env *C.JNIEnv) C.jstring {
	if g.Data == "" {
		return 0
	}

	b := utf16.Encode([]rune(g.Data))
	return C.NewString(env, (*C.jchar)(unsafe.Pointer(&b[0])), C.int(len(b)))
}

// Signature returns "Ljava/lang/String;"
func (g *String) Signature() string {
	return "Ljava/lang/String;"
}

type JClass struct {
	Data C.jclass
}

func (g *JClass) Go(env *C.JNIEnv, obj C.jobject) {
	g.Data = C.jclass(obj)
}

func (g *JClass) Java(env *C.JNIEnv) *C.jvalue {
	return (*C.jvalue)(unsafe.Pointer(&g.Data))
}

func (g *JClass) Signature() string {
	return "Ljava/lang/Class;"
}