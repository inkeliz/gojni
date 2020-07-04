package jni

/*
#cgo LDFLAGS: -landroid

#include <jni.h>
#include <stdlib.h>
#include <jni_android.h>
*/
import "C"
import (
	"errors"
	"fmt"
	"runtime"
	"unsafe"
)

var (
	ErrClassNotRegistered = errors.New("unknown class")
	ErrImpossibleCall     = errors.New("impossible to call function")
	ErrMethodNotExist     = errors.New("impossible to get mid")
	ErrClassNotExist      = errors.New("impossible to get class")
	ErrNoVM               = errors.New("no jvm found")
)

type Env struct {
	Env    *C.JNIEnv
	detach bool
	jvm    *C.JavaVM
}

// NewEnv creates an new C.JNIEnv inside given JVM
// It's not really supposed to be used outside that package
//
// You need to use runtime.LockOSThread() to lock the thread.
// You need to close that C.JNIEnv by calling `Env.Close()`.
func NewEnv(jvm *C.JavaVM) (*Env, error) {
	if jvm == nil {
		return nil, ErrNoVM
	}

	e := &Env{jvm: jvm}

	if res := C.GetEnv(e.jvm, &e.Env, C.JNI_VERSION_1_6); res != C.JNI_OK {
		if res != C.JNI_EDETACHED || C.AttachCurrentThread(e.jvm, &e.Env, nil) != C.JNI_OK {
			return nil, ErrNoVM
		}
		e.detach = true
	}

	return e, nil
}

// Error returns an error that happens on the last call, if any
func (e *Env) Error() error {
	return nil
}

// Close closes detaches the thread
func (e *Env) Close() {
	if !e.detach || e.jvm == nil {
		return
	}

	C.DetachCurrentThread(e.jvm)
}

type Session struct {
	vm    *C.JavaVM
	class C.jobject
	env   *Env
}

// NewSession creates an Session struct using *C.JavaVM and C.jobject
func NewSession(jvm *C.JavaVM, class C.jobject) (sess *Session) {
	return &Session{vm: jvm, class: class}
}

// NewSessionEnv creates an Session using the already available Env (C.JNIEnv)
// and the class.
//
// That seems useful inside native calls (from Java to Go).
func NewSessionEnv(e Env, class C.jobject) (sess *Session) {
	return &Session{env: &e, class: class}
}

// NewSessionRegistered creates an Session using the JavaJVM and the name
// of the library, which was previously registered using `Register`
//
// For consistency, the name must be `/com/something` rather than `com.something`.
//
// It only works if the Java calls `Register(this)` of `github.com.inkeliz.gojni.register_android`
func NewSessionRegistered(name string) (sess *Session, err error) {
	class, ok := classes[name]
	if !ok {
		return nil, ErrClassNotRegistered
	}

	return &Session{vm: class.JVM, class: class.Class}, nil
}

// Exec is an wrapper for Exec, which accepts the name of the method (name of the function) and the Output and Input
// If the method is void uses nil for `output`.
// If there's no argument, uses nil for `input`.
//
// Multiples arguments are not currently supported.
func (s *Session) Exec(method string, output Value, input Value) (err error) {
	if s.env == nil {
		runtime.LockOSThread()
		defer runtime.UnlockOSThread()

		s.env, err = NewEnv(s.vm)
		if err != nil {
			return err
		}
		defer s.env.Close()
	}

	return Exec(s.env, s.class, method, output, input)
}

// Exec call the given method on the given class, using the provided output and input arguments
// It's calls `getMethodID` inside the function, you don't need to do that on your side.
//
// You can use `nil` if there's no output or input.
func Exec(env *Env, class C.jobject, method string, output Value, input Value) (err error) {
	mid, err := GetMid(env, class, method, signature(output, input))
	if err != nil {
		return err
	}

	return CallMethod(env, class, mid, output, input)
}

func CallMethod(env *Env, class C.jobject, mid C.jmethodID, output Value, input Value) (err error) {
	if output == nil {
		if input == nil {
			C.CallVoidMethodA(env.Env, class, mid, nil)
		} else {
			C.CallVoidMethodA(env.Env, class, mid, input.Java(env.Env))
		}
	} else {
		if input == nil {
			output.Go(env.Env, C.CallObjectMethodA(env.Env, class, mid, nil))
		} else {
			output.Go(env.Env, C.CallObjectMethodA(env.Env, class, mid, input.Java(env.Env)))
		}
	}

	if env.Error() != nil {
		return ErrImpossibleCall
	}

	return nil
}

// GetClass calls GetObjectClass
func GetClass(env *Env, obj C.jobject) (class C.jclass, err error) {
	class = C.GetClass(env.Env, obj)
	if env.Error() != nil {
		return class, ErrClassNotExist
	}

	return class, nil
}

// GetMethod calls GetMethodID
func GetMethod(env *Env, class C.jclass, method, signature string) (mid C.jmethodID, err error) {
	met, sig := C.CString(method), C.CString(signature)
	defer C.free(unsafe.Pointer(met))
	defer C.free(unsafe.Pointer(sig))

	mid = C.GetMethod(env.Env, class, met, sig)
	if env.Error() != nil {
		return mid, ErrMethodNotExist
	}

	return mid, nil
}

// GetMid is an easier calls GetClass and GetMethod.
func GetMid(env *Env, obj C.jobject, method, signature string) (mid C.jmethodID, err error) {
	clazz, err := GetClass(env, obj)
	if err != nil {
		return mid, err
	}

	return GetMethod(env, clazz, method, signature)
}

func GetJavaVM(env *Env) *C.JavaVM {
	var v *C.JavaVM

	C.GetJavaVM(env.Env, &v)

	return v
}

func signature(output Value, input ...Value) string {
	i, o := "", "V"

	for _, in := range input {
		if in != nil {
			i += in.Signature()
		}
	}

	if output != nil {
		o = output.Signature()
	}

	return fmt.Sprintf(`(%s)%s`, i, o)
}
