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
const PALETTE_ADDR = 0x8ACBA

// 0x8D07A - 0x8ACBA
func PaletteSliceFrom(rom []byte, paletteId int) *Palette {
	BASE := PALETTE_ADDR + paletteId * PALETTE_SIZE
	paletteSlice := rom[BASE : BASE + PALETTE_SIZE]
	return PaletteFrom(paletteSlice)
}

const CODE_ROMS_SIZE = 1 << 17 // 128 KiB
const CODE_ROMS_PER_BANK = 2
const CODE_BANKS = 4

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
				{"sf2e_30g.11e", "22558eb15e035b09b80935a32b8425d91cd79669"},
				{"sf2e_37g.11f", "bf1ccfe7cc1133f0f65556430311108722add1f2"},
			},
			{
				{"sf2e_31g.12e", "86a3954335310865b14ce8b4e0e4499feb14fc12"},
				{"sf2e_38g.12f", "6565946591a18eaf46f04c1aa449ee0ae9ac2901"},
			},
			{
				{"sf2e_28g.9e", "bbcef63f35e5bff3f373968ba1278dd6bd86b593"},
				{"sf2e_35g.9f", "507bda3e4519de237aca919cf72e543403ec9724"},
			},
			{
				{"sf2_29b.10e", "75f0827f4f7e9f292add46467f8d4fe19b2514c9"},
				{"sf2_36b.10f", "b807cc495bff3f95d03b061fc629c95f965cb6d8"},
			},
		}

		code_rom = make([]byte, CODE_BANKS*CODE_ROMS_PER_BANK*CODE_ROMS_SIZE)
		for i := 0; i < len(code_banks); i++ {
			offset := i * CODE_ROMS_SIZE * CODE_ROMS_PER_BANK
			desinterleave_code_bank(code_banks[i], code_rom[offset:offset+CODE_ROMS_SIZE*CODE_ROMS_PER_BANK])
		}
	}
	return PaletteSliceFrom(code_rom, paletteId)
}