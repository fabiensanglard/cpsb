package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

const ADD_STORAGE_SIZE = 8
const NUM_PHRASES = 127

type Address struct {
	start uint32
	end   uint32
}

func (a Address) String() string {
	return fmt.Sprintf("0x%06X 0x%06X", a.start, a.end)
}

func (a Address) len() uint32 {
	return a.end - a.start
}

func writeAddress(addr Address, dst [ADD_STORAGE_SIZE]byte) {
	dst[0] = byte(addr.start & 0xFF0000 >> 16)
	dst[1] = byte(addr.start & 0x00FF00 >> 8)
	dst[2] = byte(addr.start & 0x0000FF)
	dst[3] = byte(addr.end & 0xFF0000 >> 16)
	dst[4] = byte(addr.end & 0x00FF00 >> 8)
	dst[5] = byte(addr.end & 0x0000FF)
	dst[6] = 0
	dst[7] = 0
}

func readAddress(src []byte) *Address {
	add := Address{}
	add.start = uint32(src[0]) << 16
	add.start |= uint32(src[1]) << 8
	add.start |= uint32(src[2]) << 0
	add.end = uint32(src[3]) << 16
	add.end |= uint32(src[4]) << 8
	add.end |= uint32(src[5]) << 0
	return &add
}

func dumpMsm6295ROM(rom []byte) {
	totalSize := 0
	for i := 1; i <= NUM_PHRASES; i++ {
		addrSlice := rom[i*ADD_STORAGE_SIZE : i*ADD_STORAGE_SIZE+ADD_STORAGE_SIZE]
		addr := readAddress(addrSlice)
		if addr.len() == 0 {
			continue
		}
		//fmt.Printf("%v\n", addr)
		fmt.Println("#{addr}\n")
		phrase := rom[addr.start:addr.end]
		println("Found phrase ", i, "at [", addr.start, "-", addr.end, "] size=", len(phrase), "duration =", len(phrase)*2000/7575)
		totalSize += len(phrase)
		filename := fmt.Sprintf("samples/s%d.oki", i)
		ioutil.WriteFile(filename, phrase, 0777)
	}
	println("Total size=", totalSize)
	println("Total duration=", totalSize*2000/7575)
}

// Convert WAV to RAW PCM
// ffmpeg -i sample.wav -f s16be -ar 7575 -acodec pcm_s16le sample.pcm
// https://github.com/fabiensanglard/adpcm
// ./adpcm oe sample.wav sample.oki

// Convert RAW PCM to WAV
// https://stackoverflow.com/questions/11986279/can-ffmpeg-convert-audio-from-raw-pcm-to-wav
// find -exec adpcm od {} {}.oki \;
// find -name "*.raw" -exec ffmpeg -f s16le -ar 7575 -ac 1 -i {} {}.wav \;

func main() {
	fmt.Println("msm6295 is starting...")
	os.MkdirAll("samples", 0777)

	args := []string{"sf2_18.11c", "sf2_19.12c"} //os.Args[1:]
	files := make([][]byte, len(args))

	totalSize := 0
	for i := 0; i < len(files); i++ {
		dat, err := ioutil.ReadFile(args[i])
		if err != nil {
			panic(err)
		}
		files[i] = dat
		totalSize += len(dat)
	}

	romSize := totalSize / len(files)
	// Interleave
	rom := make([]byte, totalSize)
	fmt.Printf("Rom size=0x%X\n", len(rom))

	for i, file := range files {
		copy(rom[i*romSize:], file)
	}

	dumpMsm6295ROM(rom)

	//rom := make([]byte, 0x40000)
	//
	//copy(rom[0x33:], dat)
	//
	//sample_add := makeAddress(0x33, 0x33+17096)
	//for i := 1; i < 128; i++ {
	//	add_loc := rom[i*8 : i*8+8]
	//	add_loc[0] = sample_add.start[0]
	//	add_loc[1] = sample_add.start[1]
	//	add_loc[2] = sample_add.start[2]
	//	add_loc[3] = sample_add.end[0]
	//	add_loc[4] = sample_add.end[1]
	//	add_loc[5] = sample_add.end[2]
	//	add_loc[6] = 0
	//	add_loc[7] = 0
	//}
}
