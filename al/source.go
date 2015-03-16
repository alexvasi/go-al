package al

// #include "goal.h"
import "C"
import (
	"sort"
	"unsafe"
)

type SourceState C.ALenum

const (
	Playing SourceState = C.AL_PLAYING
	Stopped             = C.AL_STOPPED
	Paused              = C.AL_PAUSED
	Initial             = C.AL_INITIAL
)

type Source struct {
	source C.ALuint
}

func GenSources(numSources int) ([]Source, error) {
	if numSources == 0 {
		return nil, nil // AL guarantees that 0 will have no effect, so we'll save ourselves the C call
	}
	buf := make([]Source, numSources)
	C.alGenSources(C.ALsizei(C.int(numSources)), (*C.ALuint)(unsafe.Pointer(&buf[0])))
	if err := GetError(); err != nil {
		return nil, err
	}

	return buf, nil
}

// Not really any different from GenSources(1), except it returns a single buffer instead of a slice
func GenSource() (source Source, err error) {
	C.alGenSources(C.ALsizei(C.int(1)), &source.source)
	err = GetError()
	return
}

func DeleteSources(toDelete ...Source) error {
	if len(toDelete) == 0 {
		return nil
	}

	C.alDeleteSources(C.ALsizei(C.int(len(toDelete))), (*C.ALuint)(unsafe.Pointer(&toDelete[0])))
	return GetError()
}

type sourceSorter []Source

func (ss sourceSorter) Len() int {
	return len(ss)
}

func (ss sourceSorter) Less(i, j int) bool {
	return -ss[i].source < -ss[j].source
}

func (ss sourceSorter) Swap(i, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}

func SafelyDeleteSources(toDelete ...Source) error {
	sort.Sort(sourceSorter(toDelete))
	return DeleteSources(toDelete...)
}

func PlaySources(sources ...Source) error {
	if len(sources) == 0 {
		return nil
	}
	C.alSourcePlayv(C.ALsizei(C.int(len(sources))), (*C.ALuint)(unsafe.Pointer(&sources[0])))
	return GetError()
}

func StopSources(sources ...Source) error {
	if len(sources) == 0 {
		return nil
	}
	C.alSourceStopv(C.ALsizei(C.int(len(sources))), (*C.ALuint)(unsafe.Pointer(&sources[0])))
	return GetError()
}

func PauseSources(sources ...Source) error {
	if len(sources) == 0 {
		return nil
	}
	C.alSourcePausev(C.ALsizei(C.int(len(sources))), (*C.ALuint)(unsafe.Pointer(&sources[0])))
	return GetError()
}

func RewindSources(sources ...Source) error {
	if len(sources) == 0 {
		return nil
	}
	C.alSourceRewindv(C.ALsizei(C.int(len(sources))), (*C.ALuint)(unsafe.Pointer(&sources[0])))
	return GetError()
}

func (source Source) Delete() error {
	C.alDeleteSources(C.ALsizei(C.int(1)), &source.source)
	return GetError()
}

func (source Source) IsValid() (bool, error) {
	val := C.alIsBuffer(source.source)
	err := GetError()
	if val == 1 {
		return true, err
	} else {
		return false, err
	}
}

func (source Source) Play() error {
	C.alSourcePlay(source.source)
	return GetError()
}

func (source Source) Pause() error {
	C.alSourcePause(source.source)
	return GetError()
}

func (source Source) Stop() error {
	C.alSourceStop(source.source)
	return GetError()
}

func (source Source) Rewind() error {
	C.alSourceRewind(source.source)
	return GetError()
}

func (source Source) QueueBuffers(buffers ...Buffer) error {
	if len(buffers) == 0 {
		return nil
	}
	C.alSourceQueueBuffers(source.source, C.ALsizei(C.int(len(buffers))), (*C.ALuint)(unsafe.Pointer(&buffers[0])))
	return GetError()
}

func (source Source) UnqueueBuffers(buffers ...Buffer) error {
	if len(buffers) == 0 {
		return nil
	}
	C.alSourceUnqueueBuffers(source.source, C.ALsizei(C.int(len(buffers))), (*C.ALuint)(unsafe.Pointer(&buffers[0])))
	return GetError()
}

func (source Source) SetPitch(val float32) error {
	C.alSourcef(source.source, C.AL_PITCH, C.ALfloat(C.float(val)))
	return GetError()
}

func (source Source) SetGain(val float32) error {
	C.alSourcef(source.source, C.AL_GAIN, C.ALfloat(C.float(val)))
	return GetError()
}

func (source Source) SetMaxDist(val float32) error {
	C.alSourcef(source.source, C.AL_MAX_DISTANCE, C.ALfloat(C.float(val)))
	return GetError()
}

func (source Source) SetRolloff(val float32) error {
	C.alSourcef(source.source, C.AL_ROLLOFF_FACTOR, C.ALfloat(C.float(val)))
	return GetError()
}

func (source Source) SetReferenceDist(val float32) error {
	C.alSourcef(source.source, C.AL_REFERENCE_DISTANCE, C.ALfloat(C.float(val)))
	return GetError()
}

func (source Source) SetMinGain(val float32) error {
	C.alSourcef(source.source, C.AL_MIN_GAIN, C.ALfloat(C.float(val)))
	return GetError()
}

func (source Source) SetMaxGain(val float32) error {
	C.alSourcef(source.source, C.AL_MAX_GAIN, C.ALfloat(C.float(val)))
	return GetError()
}

func (source Source) SetConeOuterGain(val float32) error {
	C.alSourcef(source.source, C.AL_CONE_OUTER_GAIN, C.ALfloat(C.float(val)))
	return GetError()
}

func (source Source) SetPosition(val [3]float32) error {
	C.alSourcefv(source.source, C.AL_POSITION, (*C.ALfloat)(unsafe.Pointer(&val[0])))
	return GetError()
}

func (source Source) SetVelocity(val [3]float32) error {
	C.alSourcefv(source.source, C.AL_VELOCITY, (*C.ALfloat)(unsafe.Pointer(&val[0])))
	return GetError()
}

func (source Source) SetDirection(val [3]float32) error {
	C.alSourcefv(source.source, C.AL_DIRECTION, (*C.ALfloat)(unsafe.Pointer(&val[0])))
	return GetError()
}

func (source Source) SetPosition3f(x, y, z float32) error {
	C.alSource3f(source.source, C.AL_POSITION, C.ALfloat(C.float(x)), C.ALfloat(C.float(y)), C.ALfloat(C.float(z)))
	return GetError()
}

func (source Source) SetVelocity3f(x, y, z float32) error {
	C.alSource3f(source.source, C.AL_VELOCITY, C.ALfloat(C.float(x)), C.ALfloat(C.float(y)), C.ALfloat(C.float(z)))
	return GetError()
}

func (source Source) SetDirection3f(x, y, z float32) error {
	C.alSource3f(source.source, C.AL_DIRECTION, C.ALfloat(C.float(x)), C.ALfloat(C.float(y)), C.ALfloat(C.float(z)))
	return GetError()
}

func (source Source) SetSourceRelative(isRelative bool) error {
	var i int
	if isRelative {
		i = 1
	} else {
		i = 0
	}
	C.alSourcei(source.source, C.AL_SOURCE_RELATIVE, C.ALint(C.int(i)))
	return GetError()
}

func (source Source) SetConeInnerAngle(i int) error {
	C.alSourcei(source.source, C.AL_CONE_INNER_ANGLE, C.ALint(C.int(i)))
	return GetError()
}

func (source Source) SetConeOuterAngle(i int) error {
	C.alSourcei(source.source, C.AL_CONE_OUTER_ANGLE, C.ALint(C.int(i)))
	return GetError()
}

func (source Source) SetLooping(isLooping bool) error {
	var i int
	if isLooping {
		i = 1
	} else {
		i = 0
	}
	C.alSourcei(source.source, C.AL_LOOPING, C.ALint(C.int(i)))
	return GetError()
}

func (source Source) SetBuffer(buf Buffer) error {
	C.alSourcei(source.source, C.AL_BUFFER, C.ALint(buf.buffer))
	return GetError()
}

func (source Source) ClearBuffer(buf Buffer) error {
	C.alSourcei(source.source, C.AL_BUFFER, C.ALint(C.int(0)))
	return GetError()
}

func (source Source) SetState(state SourceState) error {
	C.alSourcei(source.source, C.AL_SOURCE_STATE, C.ALint(state))
	return GetError()
}

/* Getters */

func (source Source) GetPitch() (float32, error) {
	var val float32
	C.alGetSourcef(source.source, C.AL_PITCH, (*C.ALfloat)((*C.float)(&val)))
	return val, GetError()
}

func (source Source) GetGain() (float32, error) {
	var val float32
	C.alGetSourcef(source.source, C.AL_GAIN, (*C.ALfloat)((*C.float)(&val)))
	return val, GetError()
}

func (source Source) GetMinGain() (float32, error) {
	var val float32
	C.alGetSourcef(source.source, C.AL_MIN_GAIN, (*C.ALfloat)((*C.float)(&val)))
	return val, GetError()
}

func (source Source) GetMaxGain() (float32, error) {
	var val float32
	C.alGetSourcef(source.source, C.AL_MAX_GAIN, (*C.ALfloat)((*C.float)(&val)))
	return val, GetError()
}

func (source Source) GetMaxDistance() (float32, error) {
	var val float32
	C.alGetSourcef(source.source, C.AL_MAX_DISTANCE, (*C.ALfloat)((*C.float)(&val)))
	return val, GetError()
}

func (source Source) GetRolloff() (float32, error) {
	var val float32
	C.alGetSourcef(source.source, C.AL_ROLLOFF_FACTOR, (*C.ALfloat)((*C.float)(&val)))
	return val, GetError()
}

func (source Source) GetConeOuterGain() (float32, error) {
	var val float32
	C.alGetSourcef(source.source, C.AL_CONE_OUTER_GAIN, (*C.ALfloat)((*C.float)(&val)))
	return val, GetError()
}

func (source Source) GetConeOuterAngle() (float32, error) {
	var val float32
	C.alGetSourcef(source.source, C.AL_CONE_OUTER_ANGLE, (*C.ALfloat)((*C.float)(&val)))
	return val, GetError()
}

func (source Source) GetConeInnerAngle() (float32, error) {
	var val float32
	C.alGetSourcef(source.source, C.AL_CONE_INNER_ANGLE, (*C.ALfloat)((*C.float)(&val)))
	return val, GetError()
}

func (source Source) GetReferenceDistance() (float32, error) {
	var val float32
	C.alGetSourcef(source.source, C.AL_REFERENCE_DISTANCE, (*C.ALfloat)((*C.float)(&val)))
	return val, GetError()
}

func (source Source) GetPosition() ([3]float32, error) {
	var val [3]float32
	C.alGetSourcefv(source.source, C.AL_POSITION, (*C.ALfloat)((*C.float)(&val[0])))
	return val, GetError()
}

func (source Source) GetVelocity() ([3]float32, error) {
	var val [3]float32
	C.alGetSourcefv(source.source, C.AL_VELOCITY, (*C.ALfloat)((*C.float)(&val[0])))
	return val, GetError()
}

func (source Source) GetDirection() ([3]float32, error) {
	var val [3]float32
	C.alGetSourcefv(source.source, C.AL_DIRECTION, (*C.ALfloat)((*C.float)(&val[0])))
	return val, GetError()
}

func (source Source) IsSourceRelative() (bool, error) {
	var val C.ALint
	C.alGetSourcei(source.source, C.AL_SOURCE_RELATIVE, &val)
	if val == C.AL_TRUE {
		return true, GetError()
	} else {
		return false, GetError()
	}
}

func (source Source) GetBuffer() (Buffer, error) {
	var buf Buffer
	C.alGetSourcei(source.source, C.AL_BUFFER, (*C.ALint)(unsafe.Pointer(&buf.buffer)))
	return buf, GetError()
}

func (source Source) GetState() (SourceState, error) {
	var val C.ALint
	C.alGetSourcei(source.source, C.AL_SOURCE_STATE, &val)
	return SourceState(val), GetError()
}

func (source Source) GetSourceState() (SourceState, error) {
	var val C.ALint
	C.alGetSourcei(source.source, C.AL_BUFFER, &val)
	return SourceState(val), GetError()
}

func (source Source) BuffersQueued() (int, error) {
	var val C.ALint
	C.alGetSourcei(source.source, C.AL_BUFFERS_QUEUED, &val)
	return int(C.int(val)), GetError()
}

func (source Source) BuffersProcessed() (int, error) {
	var val C.ALint
	C.alGetSourcei(source.source, C.AL_BUFFERS_PROCESSED, &val)
	return int(C.int(val)), GetError()
}
