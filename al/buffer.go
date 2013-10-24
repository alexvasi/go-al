package al

// #include "goal.h"
import "C"
import (
	"unsafe"
)

type Buffer struct {
	buffer C.ALuint
}

func GenBuffers(numBuffers int) ([]Buffer, error) {
	if numBuffers == 0 {
		return nil, nil // AL guarantees that 0 will have no effect, so we'll save ourselves the C call
	}
	buf := make([]Buffer, numBuffers)
	C.alGenBuffers(C.ALsizei(C.int(numBuffers)), (*C.ALuint)(unsafe.Pointer(&buf[0]))) // Since a Buffer is a struct{C.ALuint} an array of Buffers is an array of ALuints
	if err := GetError(); err != nil {
		return nil, err
	}

	return buf, nil
}

// Not really any different from GenBuffers(1), except it returns a single buffer instead of a slice
func GenBuffer() (buf Buffer, err error) {
	C.alGenBuffers(C.ALsizei(C.int(1)), &buf.buffer)
	err = GetError()
	return
}

func DeleteBuffers(toDelete ...Buffer) error {
	if len(toDelete) == 0 {
		return nil
	}

	C.alDeleteBuffers(C.ALsizei(C.int(len(toDelete))), (*C.ALuint)(unsafe.Pointer(&toDelete[0])))
	return GetError()
}

func (buf Buffer) Delete() error {
	C.alDeleteBuffers(C.ALsizei(C.int(1)), &buf.buffer)
	return GetError()
}

func (buf Buffer) IsValid() (bool, error) {
	val := C.alIsBuffer(buf.buffer)
	err := GetError()
	if val == 1 {
		return true, err
	} else {
		return false, err
	}
}

func (buf Buffer) BufferData(data Data) error {
	C.alBufferData(buf.buffer, C.ALenum(data.Format), unsafe.Pointer(&data.Data[0]), C.ALsizei(C.int(len(data.Data))), C.ALsizei(C.int(data.Frequency)))
	return GetError()
}

func (buf Buffer) GetFrequency() (int, error) {
	var val C.ALint
	C.alGetBufferi(buf.buffer, C.AL_FREQUENCY, &val)
	return int(C.int(val)), GetError()
}

func (buf Buffer) GetBitDepth() (int, error) {
	var val C.ALint
	C.alGetBufferi(buf.buffer, C.AL_BITS, &val)
	return int(C.int(val)), GetError()
}

func (buf Buffer) GetChannels() (int, error) {
	var val C.ALint
	C.alGetBufferi(buf.buffer, C.AL_CHANNELS, &val)
	return int(C.int(val)), GetError()
}

func (buf Buffer) GetSize() (int, error) {
	var val C.ALint
	C.alGetBufferi(buf.buffer, C.AL_SIZE, &val)
	return int(C.int(val)), GetError()
}

func (buf Buffer) ID() int {
	return int(C.int(buf.buffer))
}
