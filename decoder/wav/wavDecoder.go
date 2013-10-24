package wav

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/Jragonmiris/go-al/al"
	"os"
)

type WavHeader struct {
	_, _, _       [4]byte
	_, _          [4]byte
	AudioFormat   int16
	NumChannels   int16
	SampleRate    int32
	ByteRate      int32
	BlockAlign    int16
	BitsPerSample int16
	_, _          [4]byte // 44 bytes
}

type WavData struct {
	*WavHeader
	Data []byte
}

func LoadWavFile(filename string) (*WavData, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	info, err := os.Stat(filename)
	if err != nil {
		return nil, err
	}
	size := info.Size() - 44

	header := &WavHeader{}
	err = binary.Read(file, binary.LittleEndian, header)
	if err != nil {
		return nil, err
	}

	if header.AudioFormat != 1 {
		return nil, errors.New(fmt.Sprintf("Unsupported audio format, file is probably compressed %d", header.AudioFormat))
	}

	data := &WavData{header, make([]byte, size)}
	_, err = file.Read(data.Data)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func ToALData(rawDat *WavData) (dat al.Data, err error) {
	dat.Data = rawDat.Data
	dat.Frequency = int(rawDat.SampleRate)
	switch {
	case rawDat.NumChannels == 1 && rawDat.BitsPerSample == 8:
		dat.Format = al.Mono8
	case rawDat.NumChannels == 1 && rawDat.BitsPerSample == 16:
		dat.Format = al.Mono16
	case rawDat.NumChannels == 2 && rawDat.BitsPerSample == 8:
		dat.Format = al.Stereo8
	case rawDat.NumChannels == 2 && rawDat.BitsPerSample == 16:
		dat.Format = al.Stereo16
	default:
		return dat, errors.New(fmt.Sprintf("Incorrect number of channels or BitsPerSample. Numchannels: %d (expected 1 or 2), BitsPerSample: %d (expected 8 or 16)", rawDat.NumChannels, rawDat.BitsPerSample))
	}

	return
}

func ToEmptyALData(rawDat *WavHeader) (dat al.Data, err error) {
	dat.Data = make([]byte, rawDat.SampleRate)
	dat.Frequency = int(rawDat.SampleRate)
	switch {
	case rawDat.NumChannels == 1 && rawDat.BitsPerSample == 8:
		dat.Format = al.Mono8
	case rawDat.NumChannels == 1 && rawDat.BitsPerSample == 16:
		dat.Format = al.Mono16
	case rawDat.NumChannels == 2 && rawDat.BitsPerSample == 8:
		dat.Format = al.Stereo8
	case rawDat.NumChannels == 2 && rawDat.BitsPerSample == 16:
		dat.Format = al.Stereo16
	default:
		return dat, errors.New(fmt.Sprintf("Incorrect number of channels or BitsPerSample. Numchannels: %d (expected 1 or 2), BitsPerSample: %d (expected 8 or 16)", rawDat.NumChannels, rawDat.BitsPerSample))
	}

	return
}

func StreamFromFile(filename string) (requestChan chan<- []byte, fillChan <-chan []byte, header *WavHeader, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, nil, err
	}

	info, err := os.Stat(filename)
	if err != nil {
		return nil, nil, nil, err
	}
	size := info.Size() - 44

	header = &WavHeader{}
	err = binary.Read(file, binary.LittleEndian, header)
	if err != nil {
		return nil, nil, nil, err
	}

	if header.AudioFormat != 1 {
		return nil, nil, nil, errors.New(fmt.Sprintf("Unsupported audio format, file is probably compressed %d", header.AudioFormat))
	}

	myRequestChan := make(chan []byte)
	myFillChan := make(chan []byte)

	requestChan = myRequestChan
	fillChan = myFillChan

	frameSize := header.SampleRate
	go func() {
		defer file.Close()
		for i := int64(0); i <= size; i += int64(frameSize) {
			dat, ok := <-myRequestChan
			if !ok {
				close(myFillChan)
				return
			}
			n, err := file.Read(dat)
			if err != nil {
				if n > 0 {
					myFillChan <- dat
				}
				close(myFillChan)
				return
			}
			myFillChan <- dat
		}
		close(myFillChan)
	}()

	return requestChan, fillChan, header, nil
}

func StreamFromBytes(audio []byte, frameSize int) (datChan chan []byte) {
	buf := bytes.NewReader(audio)

	datChan = make(chan []byte)

	go func() {
		for i := 0; i <= len(audio); i += frameSize {
			dat, ok := <-datChan
			if !ok {
				return
			}
			n, err := buf.Read(dat)
			if err != nil {
				if n > 0 {
					datChan <- dat
				}
				close(datChan)
				return
			}
			datChan <- dat
		}

		close(datChan)
	}()

	return datChan
}
