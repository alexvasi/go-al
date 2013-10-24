package alc

// #include "goal.h"
import "C"
import (
	"unsafe"
)

type Device struct {
	device *C.ALCdevice
}

func OpenDevice(deviceName string) (Device, error) {
	return Device{C.alcOpenDevice((*C.ALCchar)(C.CString(deviceName)))}, GetError()
}

func DefaultDeviceName() (string, error) {
	return C.GoString((*C.char)(C.alcGetString(nil, C.ALC_DEFAULT_DEVICE_SPECIFIER))), GetError()
}

func OpenDefaultDevice() (Device, error) {
	name, err := DefaultDeviceName()

	if err != nil {
		return Device{}, err
	}

	return OpenDevice(name)

}

func (dev Device) Close() error {
	C.alcCloseDevice(dev.device)
	return GetError()
}

func (dev Device) DeviceSpecifier() (string, error) {
	return C.GoString((*C.char)(C.alcGetString(dev.device, C.ALC_DEVICE_SPECIFIER))), GetError()
}

func (dev Device) Extensions() (string, error) {
	return C.GoString((*C.char)(C.alcGetString(dev.device, C.ALC_EXTENSIONS))), GetError()
}

func (dev Device) MajorVersion() (int, error) {
	var dest *C.int
	C.alcGetIntegerv(dev.device, C.ALC_MAJOR_VERSION, C.ALCsizei(1), (*C.ALCint)(dest))
	return int(*dest), GetError()
}

func (dev Device) MinorVersion() (int, error) {
	var dest *C.int
	C.alcGetIntegerv(dev.device, C.ALC_MINOR_VERSION, C.ALCsizei(1), (*C.ALCint)(dest))
	return int(*dest), GetError()
}

func (dev Device) AttributesSize() (int, error) {
	var dest *C.int
	C.alcGetIntegerv(dev.device, C.ALC_ATTRIBUTES_SIZE, C.ALCsizei(1), (*C.ALCint)(dest))
	return int(*dest), GetError()
}

// Temporary, I have no idea what format this comes out in.
// After I do I'll make something more sane
func (dev Device) Attributes(size int) ([]int, error) {
	dest := make([]C.int, size)
	C.alcGetIntegerv(dev.device, C.ALC_ALL_ATTRIBUTES, C.ALCsizei(size), (*C.ALCint)(&dest[0]))

	toReturn := make([]int, size)
	for i, dat := range dest {
		toReturn[i] = int(dat)
	}

	return toReturn, GetError()
}

func (dev Device) AllAttributes() ([]int, error) {
	size, err := dev.AttributesSize()
	if err != nil {
		return nil, err
	}
	return dev.Attributes(size)
}

func (dev Device) IsExtensionPresent(extName string) (bool, error) {
	val := C.int(C.alcIsExtensionPresent(dev.device, (*C.ALCchar)(C.CString(extName))))
	if val == 1 {
		return true, GetError()
	} else {
		return false, GetError()
	}
}

func (dev Device) procAddress(funcName string) (unsafe.Pointer, error) {
	return unsafe.Pointer(C.alcGetProcAddress(dev.device, (*C.ALCchar)(C.CString(funcName)))), GetError()
}

func (dev Device) enumValue(name string) (int, error) {
	return int(C.int(C.alcGetEnumValue(dev.device, (*C.ALCchar)(C.CString(name))))), GetError()
}
