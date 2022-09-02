package main

import (
	"fmt"
)
import "io/ioutil"

const ROM_CHIP_SIZE = 1 << 19 // 512 KiB
const TOTAL_ROM_SIZE = ROM_CHIP_SIZE * 12

const ROMS_PER_BANK = 4
const BANKS = 3
const BANK_SIZE = ROMS_PER_BANK * ROM_CHIP_SIZE

const TILE_WIDTH = 16
const TILE_HEIGHT = 16

const SHEET_WIDTH = 16
const SHEET_HEIGHT = 16

const TILES_SIZE = TILE_WIDTH * TILE_HEIGHT / 2

var gfx_banks = [BANKS][ROMS_PER_BANK][2]string{
	{
		{"sf2_06.bin", "b9194fb337b30502c1c9501cd6c64ae4035544d4"},
		{"sf2_08.bin", "3759b851ac0904ec79cbb67a2264d384b6f2f9f9"},
		{"sf2_05.bin", "520840d727161cf09ca784919fa37bc9b54cc3ce"},
		{"sf2_07.bin", "2360cff890551f76775739e2d6563858bff80e41"},
	},

	{
		{"sf2_15.bin", "357c2275af9133fd0bd6fbb1fa9ad5e0b490b3a2"},
		{"sf2_17.bin", "baa92b91cf616bc9e2a8a66adc777ffbf962a51b"},
		{"sf2_14.bin", "2eea16673e60ba7a10bd4d8f6c217bb2441a5b0e"},
		{"sf2_16.bin", "f787aab98668d4c2c54fc4ba677c0cb808e4f31e"},
	},

	{
		{"sf2_25.bin", "5669b845f624b10e7be56bfc89b76592258ce48b"},
		{"sf2_27.bin", "9af9df0826988872662753e9717c48d46f2974b0"},
		{"sf2_24.bin", "a6a7f4725e52678cbd8d557285c01cdccb2c2602"},
		{"sf2_26.bin", "f9a92d614e8877d648449de2612fc8b43c85e4c2"},
	},
}

func desinterleave(roms [ROMS_PER_BANK][2]string, dst[] byte){
	var files[ROMS_PER_BANK][]byte
	for i, _ := range files {
		content, err := ioutil.ReadFile("./roms/" + roms[i][0])
		if err != nil {
			panic(err)
		}
		files[i] = content
	}

	var cursor = 0
	for i:=0 ; i < ROM_CHIP_SIZE ; i+=2 {
		for _, f := range files {
			dst[cursor]   = f[i]
			dst[cursor+1] = f[i+1]
			cursor += 2
		}
	}
}

func main() {

	fmt.Println("Pose starting...")

	// Interleave ROMS into a single address space
	// Each ROM is 512 KiB. SF2 has 12 chips for a total of 6MiB
	var rom = make([]byte, TOTAL_ROM_SIZE)
	for i := 0 ; i < BANKS; i++ {
		offset := i * BANK_SIZE
		desinterleave(gfx_banks[i], rom[offset:offset + BANK_SIZE])
	}

	palette := GetPalette(1)

	var ryu Pose
	ryu.Init("ryu.svg", rom, palette)
	ryu.AddBlock(0x00, 0x64, 4, 6, 0, 0)
	ryu.Finalize()

	var kenWin Pose
	kenWin.Init("kenWin.svg", rom, GetPalette(5))
	kenWin.AddBlock(0x01, 0x84, 4, 6, 0, 32)
	kenWin.AddBlock(0x01, 0x75, 3, 1, 16, 16)
	kenWin.AddBlock(0x01, 0x66, 2, 1, 32, 0)
	kenWin.Finalize()

	var kenPatch Pose
	kenPatch.Init("kenPatch.svg", rom, GetPalette(5))
	kenPatch.AddBlock(0x01, 0x70, 2, 2, 0, 0)
	kenPatch.Finalize()

	var ryuWin Pose
	ryuWin.Init("ryuWin.svg", rom, palette)
	ryuWin.AddBlock(0x01, 0x84, 4, 6, 0, 32)
	ryuWin.AddBlock(0x01, 0x75, 3, 1, 16, 16)
	ryuWin.AddBlock(0x01, 0x66, 2, 1, 32, 0)
	ryuWin.Finalize()

	sagat_palette := GetPalette(0xA0)
	var sagat Pose
	sagat.Init("sagatTigerP.svg", rom, sagat_palette)
	sagat.AddBlock(0x53, 0x20, 2, 1, 16, 0)
	sagat.AddBlock(0x53, 0x30, 3, 5, 16, 16)
	sagat.AddBlock(0x53, 0x33, 1, 2, 0, 16)
	sagat.AddBlock(0x53, 0x34, 1, 2, 0, 64)
	sagat.AddBlock(0x53, 0x81, 2, 1, 32, 96)
	sagat.AddBlock(0x53, 0x84, 2, 1, 64 + 16, 96)

	sagat.AddBlock(0x53, 0x53, 1, 2, 64, 64 -16)
	sagat.AddBlock(0x53, 0x73, 3, 1, 64, 64 + 16)
	sagat.Finalize()

	var honda Pose
	honda.Init("honda.svg", rom, GetPalette(2))
	honda.AddBlock(0x06, 0x08, 8, 9, 0, 0)
	honda.Finalize()

	var chunLi Pose
	chunLi.Init("chunLi.svg", rom, GetPalette(6))
	chunLi.AddBlock(0x66, 0xA6, 6, 6, 0, 0)
	chunLi.Finalize()

	var honda2 Pose
	honda2.Init("honda2.svg", rom, GetPalette(2))
	honda2.AddBlock(0x0e, 0x52, 4, 5, 0, 0)
	honda2.AddBlock(0x0e, 0x61, 1, 5, -16, 16)
	honda2.AddBlock(0x0e, 0x60, 1, 2, -32, 16)

	honda2.AddBlock(0x0e, 0x66, 1, 5, 64, 16)
	// Far right leg
	honda2.AddBlock(0x0e, 0x77, 1, 4, 80, 32)

	// Hands
	honda2.AddBlock(0x0e, 0x56, 2, 1, -64, 32)
	honda2.AddBlock(0x0e, 0x50, 2, 1, -64, 16)
	honda2.AddBlock(0x0e, 0xa2, 1, 1, 0, 80)
	honda2.Finalize()

	var zanghief Pose
	zanghief.Init("zanghief.svg", rom, GetPalette(7))
	zanghief.AddBlock(0x89, 0x05, 6, 8, 0, 0)
	zanghief.Finalize()

	var guileCalve Pose
	guileCalve.Init("guileCalve.svg", rom, GetPalette(4))
	guileCalve.AddBlock(0x77, 0x52, 1, 7, 0, 0)
	guileCalve.AddBlock(0x77, 0xb1, 1, 1, -16, 16*6)
	guileCalve.AddBlock(0x77, 0x53, 1, 6, 16, 0)
	guileCalve.AddBlock(0x77, 0x64, 2, 6, 32, 16)
	guileCalve.AddBlock(0x77, 0x95, 2, 3, 3*16, 4 * 16)
	guileCalve.AddBlock(0x77, 0xA7, 1, 2, 5 * 16, 5 * 16)
	guileCalve.Finalize()
}



