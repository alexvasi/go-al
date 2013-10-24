package al

// #include "goal.h"
import "C"
import (
	"unsafe"
)

func GetListenerPosition() ([3]float32, error) {
	var val [3]float32
	C.alGetListenerfv(C.AL_POSITION, (*C.ALfloat)((*C.float)(&val[0])))
	return val, GetError()
}

func GetListenerVelocity() ([3]float32, error) {
	var val [3]float32
	C.alGetListenerfv(C.AL_VELOCITY, (*C.ALfloat)((*C.float)(&val[0])))
	return val, GetError()
}

func GetListenerOrientation() ([3]float32, error) {
	var val [3]float32
	C.alGetListenerfv(C.AL_ORIENTATION, (*C.ALfloat)((*C.float)(&val[0])))
	return val, GetError()
}

func GetListenerPosition3f() (float32, float32, float32, error) {
	var x, y, z float32
	C.alGetListener3f(C.AL_POSITION, (*C.ALfloat)((*C.float)(&x)), (*C.ALfloat)((*C.float)(&y)), (*C.ALfloat)((*C.float)(&z)))
	return x, y, z, GetError()
}

func GetListenerVelocity3f() (float32, float32, float32, error) {
	var x, y, z float32
	C.alGetListener3f(C.AL_VELOCITY, (*C.ALfloat)((*C.float)(&x)), (*C.ALfloat)((*C.float)(&y)), (*C.ALfloat)((*C.float)(&z)))
	return x, y, z, GetError()
}

func GetListenerGain() (float32, error) {
	var val float32
	C.alGetListenerf(C.AL_GAIN, (*C.ALfloat)((*C.float)(&val)))
	return val, GetError()
}

func SetListenerPosition(val [3]float32) error {
	C.alListenerfv(C.AL_POSITION, (*C.ALfloat)((*C.float)(&val[0])))
	return GetError()
}

func SetListenerVelocity(val [3]float32) error {
	C.alListenerfv(C.AL_VELOCITY, (*C.ALfloat)((*C.float)(&val[0])))
	return GetError()
}

func SetListenerOrientation(at, up [3]float32) error {
	tmp := [2][3]float32{at, up}
	C.alListenerfv(C.AL_ORIENTATION, (*C.ALfloat)(unsafe.Pointer(&tmp[0][0])))
	return GetError()
}

func SetListenerPosition3f(x, y, z float32) error {
	C.alListener3f(C.AL_POSITION, (C.ALfloat)((C.float)(x)), (C.ALfloat)((C.float)(y)), (C.ALfloat)((C.float)(z)))
	return GetError()
}

func SetListenerVelocity3f(x, y, z float32) error {
	C.alListener3f(C.AL_VELOCITY, (C.ALfloat)((C.float)(x)), (C.ALfloat)((C.float)(y)), (C.ALfloat)((C.float)(z)))
	return GetError()
}

func SetListenerGain(val float32) error {
	C.alListenerf(C.AL_GAIN, C.ALfloat(C.float(val)))
	return GetError()
}
