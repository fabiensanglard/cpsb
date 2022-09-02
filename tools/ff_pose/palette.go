package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"image/color"
	"io/ioutil"
)

type Palette = [16]color.RGBA

// Palette are stored in int16 with format Ignored, Red, Green, and Blue.
// IIIIRRRRGGGGBBBB
// To convert 0xX to 0xXX, we do (* 16 + 1) = * 17
func PaletteFrom(array []byte) *Palette {
	var palette Palette

	for i:= 0 ; i < 15 ; i++ {
		r := array[i * 2]
		gb := array[i * 2 + 1]
		palette[i] = color.RGBA{r & 0xF * 17, (gb >> 4) * 17, gb & 0xF * 17, 0xff}
	}
	palette[15] = color.RGBA{0x00, 0x00, 0x00, 0x00}
	return &palette
}

const PALETTE_SIZE = 32
const PALETTE_ADDR = 0xC0000

func PaletteSliceFrom(rom []byte, paletteId int) *Palette {
	BASE := PALETTE_ADDR + paletteId * PALETTE_SIZE
	paletteSlice := rom[BASE : BASE + PALETTE_SIZE]
	return PaletteFrom(paletteSlice)
}

const CODE_ROMS_SIZE = 1 << 17 // 128 KiB
const CODE_ROMS_PER_BANK = 2
const CODE_BANKS = 2

func sha(array []byte) []byte{
	h := sha1.New()
	h.Write(array)
	return h.Sum(nil)
}

func desinterleave_code_bank(roms [CODE_ROMS_PER_BANK][2]string, dst[] byte){
	var files[CODE_ROMS_PER_BANK][]byte
	for i, _ := range files {
		content, err := ioutil.ReadFile("./roms/" + roms[i][0])
		hash := sha(content)
		hash_string := hex.EncodeToString(hash[:])
		if hash_string != roms[i][1] {
			fmt.Println(hash_string)
			fmt.Println(roms[i][1])
			panic("Unexpected file")
		}

		if err != nil {
			panic(err)
		}
		files[i] = content
	}

	var cursor = 0
	for i:=0 ; i < CODE_ROMS_SIZE ; i++ {
		for _, f := range files {
			dst[cursor]   = f[i]
			cursor += 1
		}
	}


}

var code_rom []byte
func GetPalette(paletteId int) *Palette {
	if len(code_rom) == 0 {
		var code_banks = [CODE_BANKS][CODE_ROMS_PER_BANK][2]string{
			{
				{"ff_36.11f", "0756ae576a1f6d5b8b22f8630dca40ef38567ea6"},
				{"ff_42.11h", "5045a467f3e228c02b4a355b52f58263ffa90113"},
			},
			{
				{"ff_37.12f", "38f44434c8befd623953ae23d6e5ff4e201d6627"},
				{"ffe_43.12h", "de16873d1639ac1738be0937270b108a9914f263"},
			},
		}

		code_rom = make([]byte, CODE_BANKS*CODE_ROMS_PER_BANK*CODE_ROMS_SIZE)
		for i := 0; i < len(code_banks); i++ {
			offset := i * CODE_ROMS_SIZE * CODE_ROMS_PER_BANK
			desinterleave_code_bank(code_banks[i], code_rom[offset:offset+CODE_ROMS_SIZE*CODE_ROMS_PER_BANK])
		}

		// Ugly hack but I am too tired to make it the right way.
		// We need to add the last ROM which is 0x80000 long at offset 0x80000.
		// And also byte swap all WORDs.
		last_rom, err := ioutil.ReadFile("./roms/" + "ff-32m.8h")
		if err != nil {
			panic(err)
		}
		if len(last_rom) != 0x80000 {
			panic("ROM size is not 0x80000")
		}
		for i := 0; i < len(last_rom) -1; i+=2 {
			v1 := last_rom[i+0]
			v2 := last_rom[i+1]
			last_rom[i+0] = v2
			last_rom[i+1] = v1
		}

		new_code_rom := make([]byte, len(code_rom) + 0x80000)
		copy(new_code_rom, code_rom)
		copy(new_code_rom[0x80000:], last_rom)
		code_rom = new_code_rom
		ioutil.WriteFile( "./roms/code.bin", code_rom,0666)
	}

	return PaletteSliceFrom(code_rom, paletteId)
}