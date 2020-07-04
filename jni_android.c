
#include <stdlib.h>
#include <stdio.h>
#include <jni.h>
#include "jni_android.h"
#include "_cgo_export.h"

jclass FindClass(JNIEnv *env, const char *name) {
	return (*env)->FindClass(env, name);
}

jclass GetClass(JNIEnv *env, jobject obj) {
	return (*env)->GetObjectClass(env, obj);
}

jmethodID GetMethod(JNIEnv *env, jclass clazz, const char *function, const char *signature) {
	return (*env)->GetMethodID(env, clazz, function, signature);
}

void CallVoidMethodA(JNIEnv *env, jobject obj, jmethodID method, jvalue *args) {
	(*env)->CallVoidMethodA(env, obj, method, args);
}

jobject CallObjectMethodA(JNIEnv *env, jobject obj, jmethodID method, jvalue *args) {
	return (*env)->CallObjectMethodA(env, obj, method, args);
}

jint GetEnv(JavaVM *vm, JNIEnv **env, jint version) {
	return (*vm)->GetEnv(vm, (void **)env, version);
}

jint AttachCurrentThread(JavaVM *vm, JNIEnv **p_env, void *thr_args) {
	return (*vm)->AttachCurrentThread(vm, p_env, thr_args);
}

jint DetachCurrentThread(JavaVM *vm) {
	return (*vm)->DetachCurrentThread(vm);
}

jsize GetArrayLength(JNIEnv *env, jarray arr) {
	return (*env)->GetArrayLength(env, arr);
}

jsize GetStringLength(JNIEnv *env, jstring str) {
	return (*env)->GetStringLength(env, str);
}

jsize GetStringUTFLength(JNIEnv *env, jstring str) {
	return (*env)->GetStringUTFLength(env, str);
}

const jchar *GetStringChars(JNIEnv *env, jstring str) {
	return (*env)->GetStringChars(env, str, NULL);
}

const jchar *GetStringUTFChars(JNIEnv *env, jstring str) {
	return (*env)->GetStringUTFChars(env, str, JNI_FALSE);
}

jbyte *GetByteArrayElements(JNIEnv *env, jbyteArray arr) {
	return (*env)->GetByteArrayElements(env, arr, NULL);
}

jobject NewGlobalRef(JNIEnv *env, jobject o) {
	return (*env)->NewGlobalRef(env, o);
}

jstring NewString(JNIEnv *env, const jchar *unicodeChars, int len) {
	return (*env)->NewString(env, unicodeChars, len);
}

jbyteArray NewByteArray(JNIEnv *env, jbyte* cdata, int datalen) {
	jbyteArray data = (*env)->NewByteArray(env, datalen);
	if (datalen > 0) {
		(*env)->SetByteArrayRegion(env, data, 0, datalen, cdata);
	}

	return data;
}

jthrowable ExceptionOccurred(JNIEnv *env) {
	return (*env)->ExceptionOccurred(env);
}

jboolean ExceptionCheck(JNIEnv *env) {
    return (*env)->ExceptionCheck(env);
}

void ExceptionClear(JNIEnv *env) {
	(*env)->ExceptionClear(env);
}


jint GetJavaVM(JNIEnv *env, JavaVM **vm) {
	return (*env)->GetJavaVM(env, vm);
}