package al

// #cgo windows LDFLAGS: -lOpenAL32
// #include "goal.h"
import "C"
import ()

type DistanceModel C.ALenum

const (
	NoDistanceModel         DistanceModel = C.AL_NONE
	InverseDistance                       = C.AL_INVERSE_DISTANCE
	InverseDistanceClampled               = C.AL_INVERSE_DISTANCE_CLAMPED
)

func GetDopplerFactor() (float32, error) {
	val := C.alGetFloat(C.AL_DOPPLER_FACTOR)
	return float32(C.float(val)), GetError()
}

func GetDopplerVelocity() (float32, error) {
	val := C.alGetFloat(C.AL_DOPPLER_VELOCITY)
	return float32(C.float(val)), GetError()
}

func GetDistanceModel() (DistanceModel, error) {
	model := C.alGetInteger(C.AL_DISTANCE_MODEL)
	return DistanceModel(model), GetError()
}

func GetVendor() (string, error) {
	val := C.alGetString(C.AL_VENDOR)
	return C.GoString((*C.char)(val)), GetError()
}

func GetVersion() (string, error) {
	val := C.alGetString(C.AL_VERSION)
	return C.GoString((*C.char)(val)), GetError()
}

func GetRenderer() (string, error) {
	val := C.alGetString(C.AL_RENDERER)
	return C.GoString((*C.char)(val)), GetError()
}

func GetExtensions() (string, error) {
	val := C.alGetString(C.AL_EXTENSIONS)
	return C.GoString((*C.char)(val)), GetError()
}

func SetDistanceModel(model DistanceModel) error {
	C.alDistanceModel(C.ALenum(model))
	return GetError()
}

// Default is 1.0
func SetDopplerFactor(factor float32) error {
	C.alDopplerFactor(C.ALfloat(C.float(factor)))
	return GetError()
}

// Default is 343.0
func SetDopplerVelocity(v float32) error {
	C.alDopplerVelocity(C.ALfloat(C.float(v)))
	return GetError()
}
