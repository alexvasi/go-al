package alc

// #include "goal.h"
import "C"
import ()

type Error struct {
	problem C.ALCenum
}

func (err Error) Error() string {
	switch err.problem {
	case C.ALC_INVALID_DEVICE:
		return "ALC Error: Invalid Device"
	case C.ALC_INVALID_CONTEXT:
		return "ALC Error: Invalid Context"
	case C.ALC_INVALID_ENUM:
		return "ALC Error: Invalid Enum"
	case C.ALC_INVALID_VALUE:
		return "ALC Error: Invalid Value"
	//case C.ALC_INVALID_OPERATION:
	//	return "ALC Error: Invalid Operation"
	case C.ALC_OUT_OF_MEMORY:
		return "ALC Error: Out of Memory"
	default:
		return "Unknown error"
	}
}

func (err Error) ProblemCode() int {
	return int(C.int(err.problem))
}

func GetError() error {
	// alcGetError wants an "ALCvoid" type, which go reports is of type *[0]byte.
	// I figure nil is good enough?
	if errVal := C.alcGetError(nil); errVal == C.ALC_NO_ERROR {
		return nil
	} else {
		return Error{errVal}
	}
}
