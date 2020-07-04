package jni

/*
#import <jni.h>
#import "jni_android.h"
*/
import "C"

type Call struct {
	session *Session
	method  string
}

// NewCall defines the method/function which want to be called
// it doesn't call the method. It's just an wrapper for Exec,
// making more friendly to retrieve values.
//
// You need to use `ReturnVoid`, `ReturnString` (...) to call
// the java method.
func NewCall(sess *Session, method string) (resp *Call) {
	return &Call{session: sess, method:  method}
}

// ReturnVoid executes the void-method with args as argument of the method.
func (c *Call) ReturnVoid(args Value) (err error) {
	return c.session.Exec(c.method, nil, args)
}

// ReturnString executes the method with args as argument of the method, returning string as response
func (c *Call) ReturnString(args Value) (resp string, err error) {
	value := new(String)

	err = c.session.Exec(c.method, value, args)
	return value.Data, err
}

// ReturnBytes executes the method with args as argument of the method, returning byte-slice as response
func (c *Call) ReturnBytes(args Value) (resp []byte, err error) {
	value := new(Bytes)

	err = c.session.Exec(c.method, value, args)
	return value.Data, err
}
