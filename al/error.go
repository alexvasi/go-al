package al

// #include "goal.h"
import "C"
import ()

type Error struct {
	problem C.ALenum
}

func (err Error) Error() string {
	switch err.problem {
	case C.AL_INVALID_NAME:
		return "AL Error: Invalid Name"
	case C.AL_INVALID_ENUM:
		return "AL Error: Invalid Enum"
	case C.AL_INVALID_VALUE:
		return "AL Error: Invalid Value"
	case C.AL_INVALID_OPERATION:
		return "AL Error: Invalid Operation"
	case C.AL_OUT_OF_MEMORY:
		return "AL Error: Out of Memory"
	default:
		return "Unknown error"
	}
}

func (err Error) ProblemCode() int {
	return int(C.int(err.problem))
}

func GetError() error {
	if errVal := C.alGetError(); errVal == C.AL_NO_ERROR {
		return nil
	} else {
		return Error{errVal}
	}
}
