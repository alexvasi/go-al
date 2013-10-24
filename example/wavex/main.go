package main

import (
	"fmt"
	"github.com/Jragonmiris/go-al/al"
	"github.com/Jragonmiris/go-al/alc"
	"github.com/Jragonmiris/go-al/decoder/wav"
	"time"
)

func main() {

	device, err := alc.OpenDefaultDevice()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(device.DeviceSpecifier())

	context, err := alc.CreateDefaultContext(device)
	if err != nil {
		panic(err)
	}

	if err = context.MakeCurrent(); err != nil {
		panic(err)
	}
	if err = context.Process(); err != nil {
		panic(err)
	}

	fmt.Println(al.GetVersion())
	fmt.Println(al.GetVendor())
	fmt.Println(al.GetExtensions())

	source, err := al.GenSource()
	rawDat, err := wav.LoadWavFile("gameover.wav")

	if err != nil {
		panic(err)
	}

	dat, err := wav.ToALData(rawDat)
	if err != nil {
		panic(err)
	}

	buf, _ := al.GenBuffer()
	err = buf.BufferData(dat)
	if err != nil {
		panic(err)
	}

	err = source.SetBuffer(buf)
	if err != nil {
		panic(err)
	}
	source.SetPitch(1.0)
	source.SetGain(1.0)
	source.SetLooping(false)

	err = source.Play()

	time.Sleep(1500 * time.Millisecond)

	source.Stop()
	source.Delete()
	buf.Delete()
	context.Destroy()
	device.Close()
}
