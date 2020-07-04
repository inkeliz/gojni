#include <jni.h>

jclass FindClass(JNIEnv *env, const char *name);
jclass GetClass(JNIEnv *env, jobject obj);


jmethodID GetMethod(JNIEnv *env, jclass clazz, const char *name, const char *sig);


void CallVoidMethodA(JNIEnv *env, jobject obj, jmethodID method, jvalue *args);
jobject CallObjectMethodA(JNIEnv *env, jobject obj, jmethodID method, jvalue *args);


jint GetEnv(JavaVM *vm, JNIEnv **env, jint version);


jint AttachCurrentThread(JavaVM *vm, JNIEnv **p_env, void *thr_args);
jint DetachCurrentThread(JavaVM *vm);


jstring NewString(JNIEnv *env, const jchar *unicodeChars, int len);
const jchar *GetStringChars(JNIEnv *env, jstring str);
const jchar *GetStringUTFChars(JNIEnv *env, jstring string);
jsize GetStringLength(JNIEnv *env, jstring str);
jsize GetStringUTFLength(JNIEnv *env, jstring str);

jbyteArray NewByteArray(JNIEnv *env, jbyte* cdata, int datalen);
jbyte *GetByteArrayElements(JNIEnv *env, jbyteArray arr);
jsize GetArrayLength(JNIEnv *env, jarray arr);

jobject NewGlobalRef(JNIEnv *env, jobject o);

jthrowable ExceptionOccurred(JNIEnv *env);
jboolean ExceptionCheck(JNIEnv *env);
void ExceptionClear(JNIEnv *env);

jint GetJavaVM(JNIEnv *env, JavaVM **vm);
