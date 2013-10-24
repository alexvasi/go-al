package al

// #include "goal.h"
import "C"
import ()

type Format C.ALenum

const (
	Mono8    Format = C.AL_FORMAT_MONO8
	Mono16          = C.AL_FORMAT_MONO16
	Stereo8         = C.AL_FORMAT_STEREO8
	Stereo16        = C.AL_FORMAT_STEREO16
	// Some implementations seem to have 32 and 24-bit options, but adding those won't
	// let this compile on either of my installs. Add them if you need them I guess
)

type Data struct {
	Format    Format
	Data      []byte
	Frequency int
}
