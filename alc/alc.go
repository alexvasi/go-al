package alc

// #include "goal.h"
// #cgo windows LDFLAGS: -lOpenAL32
import "C"
import ()

type Context struct {
	context *C.ALCcontext
}

func CreateDefaultContext(dev Device) (Context, error) {
	return Context{C.alcCreateContext(dev.device, nil)}, GetError()

}

func CreateContext(dev Device, sync bool, frequency, refresh int) (Context, error) {
	var syncval int
	if sync {
		syncval = 1
	} else {
		syncval = 0
	}
	args := [...]C.ALCint{C.ALC_SYNC, (C.ALCint)(syncval), C.ALC_FREQUENCY, (C.ALCint)(frequency), C.ALC_REFRESH, (C.ALCint)(refresh), C.ALC_INVALID}
	return Context{C.alcCreateContext(dev.device, (*C.ALCint)(&args[0]))}, GetError()
}

func CurrentContext() Context {
	return Context{C.alcGetCurrentContext()}
}

func UnsetCurrentContext() error {
	C.alcMakeContextCurrent(nil)
	return GetError()
}

func UnsetProcessingContext() error {
	C.alcProcessContext(nil)
	return GetError()
}

func (con Context) MakeCurrent() error {
	C.alcMakeContextCurrent(con.context)
	return GetError()
}

func (con Context) Process() error {
	C.alcProcessContext(con.context)
	return GetError()
}

func (con Context) Suspend() error {
	C.alcSuspendContext(con.context)
	return GetError()
}

func (con Context) Destroy() error {
	C.alcDestroyContext(con.context)
	return GetError()
}

func (con Context) Equal(eq Context) bool {
	return con.context == eq.context
}
