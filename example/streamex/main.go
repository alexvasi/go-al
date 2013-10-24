package main

import (
	"fmt"
	"github.com/Jragonmiris/go-al/al"
	"github.com/Jragonmiris/go-al/alc"
	"github.com/Jragonmiris/go-al/decoder/wav"
	"os"
)

func main() {

	// Initial setup, open device and initialize context
	device, err := alc.OpenDefaultDevice()
	if err != nil {
		fmt.Println(err)
	}

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

	// Print version information and such, mostly for fun
	deviceName, _ := device.DeviceSpecifier()
	version, _ := al.GetVersion()
	vendor, _ := al.GetVendor()
	extensions, _ := al.GetExtensions()
	fmt.Println("Using default device: ", deviceName)
	fmt.Println("OpenAL version: ", version)
	fmt.Println("Device vendor: ", vendor)
	fmt.Println("Extensions available: ", extensions)

	// Make our default audio source, it makes more sense when doing 3D positional audio
	// since it defines where the sound is coming from, but for such a simple endeavor
	// it just sits there
	source, err := al.GenSource()
	if err != nil {
		panic(err)
	}

	// Set the Pitch and Gain (volume), because you won't hear anything if you don't
	source.SetPitch(1.0)
	source.SetGain(1.0)
	source.SetLooping(false)

	// Get a few buffers to gradually fill the audio stream, like how Youtube videos and such buffer
	// 3 is just an arbitrary number
	bufs, err := al.GenBuffers(3)
	if err != nil {
		panic(err)
	}

	defer Cleanup(source, bufs, context, device)

	// Find the name of whatever .wav file we want
	if len(os.Args) < 2 {
		panic("Need to know what wav file to play! List it as an argument")
	}

	// Begin streaming that file (in another goroutine)
	requestChan, fillChan, header, err := wav.StreamFromFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	// Convert the wav metadata into a struct easy to pass into OpenAL
	bufData, err := wav.ToEmptyALData(header)
	if err != nil {
		panic(err)
	}

	var ok bool = true

	// Initial fill
	// We can either use a func and return for loop control,
	// or use a labelled break. I prefer the anonymous func and return, but
	// it's really just a matter of taste
	func() {
		for i := 0; i < 3; i++ {
			select {
			case requestChan <- bufData.Data:
			case _, ok = <-fillChan:
				if !ok {
					return
				}
			}

			bufData.Data, ok = <-fillChan
			if !ok {
				return
			}

			bufs[i].BufferData(bufData)
		}
	}()

	// Only do this if we still have data left
	if ok {

		// Setup
		err = source.QueueBuffers(bufs...)
		if err != nil {
			panic(err)
		}

		err = source.Play()
		if err != nil {
			panic(err)
		}

		//Keep filling the stream
		func() {
			for i := 0; ; i = (i + 1) % 3 {

				// Ugly spin loop, but there's really no other way to poll if AL is done with some data,
				// if we don't check for this, we'll get an error when we try to unqueue a buffer
				for j, _ := source.BuffersProcessed(); j < 1; j, _ = source.BuffersProcessed() {
				}

				err = source.UnqueueBuffers(bufs[i])
				if err != nil {
					panic(err)
				}

				select {
				case requestChan <- bufData.Data:
				case _, ok = <-fillChan:
					if !ok {
						return
					}
				}

				bufData.Data, ok = <-fillChan
				if !ok {
					return
				}

				err = bufs[i].BufferData(bufData)
				if err != nil {
					panic(err)
				}

				err = source.QueueBuffers(bufs[i])
				if err != nil {
					panic(err)
				}

				if state, _ := source.GetSourceState(); state != al.Playing {
					source.Play()
				}
			}
		}()

	} else {
		// This doesn't work well for short files
		return
	}

	close(requestChan)

	// Keep playing until it's done
	for state, _ := source.GetSourceState(); state == al.Playing; state, _ = source.GetSourceState() {
		source.Play()
	}

}

func Cleanup(source al.Source, bufs []al.Buffer, context alc.Context, device alc.Device) {
	// Cleanup
	// Should be done in this order to ensure proper closing of the audio device
	source.Stop()
	source.Delete()
	al.DeleteBuffers(bufs...)
	context.Destroy()
	device.Close()
}
