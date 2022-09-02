package main

import (
	"fmt"
	"image/color"
)
import "io/ioutil"

const ROM_CHIP_SIZE = 1 << 19 // 512 KiB
const TOTAL_ROM_SIZE = ROM_CHIP_SIZE * 12

const ROMS_PER_BANK = 4
const BANKS = 1
const BANK_SIZE = ROMS_PER_BANK * ROM_CHIP_SIZE

const TILE_WIDTH = 16
const TILE_HEIGHT = 16

const SHEET_WIDTH = 16
const SHEET_HEIGHT = 16

const TILES_SIZE = TILE_WIDTH * TILE_HEIGHT / 2

var gfx_banks = [BANKS][ROMS_PER_BANK][2]string{
	{
		{"ff-5m.7a", "7868f5801347340867720255f8380548ad1a65a7"},
		{"ff-7m.9a", "f7b00a3ca8cb85264ab293089f9f540a8292b49c"},
		{"ff-1m.3a", "5ce16af72858a57aefbf6efed820c2c51935882a"},
		{"ff-3m.5a", "df5f3d3aa96a7a33ff22f2a31382942c4c4f1111"},
	},
}

func desinterleave(roms [ROMS_PER_BANK][2]string, dst[] byte){
	var files[ROMS_PER_BANK][]byte
	for i, _ := range files {
		path := "./roms/" + roms[i][0]
		content, err := ioutil.ReadFile(path)
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



	damned := GetPalette(31)
	var damned1 Pose
	damned1.Init("damned1.svg", rom, damned)
	damned1.AddBlock(0x1b, 0x28, 6, 3, 0, 0)
	damned1.Finalize()

	var damned2 Pose
	damned2.Init("damned2.svg", rom, damned)
	damned2.AddBlock(0x1b, 0x9d, 3, 4, 0, 0)
	damned2.AddBlock(0x1b, 0x1e, 2, 4, 48, 0)
	damned2.Finalize()


	oriber := GetPalette(50) // G. Oriber = 50, Wong Who = 51, Bill Bull = 52
	wong := GetPalette(51) // G. Oriber = 50, Wong Who = 51, Bill Bull = 52
	bill := GetPalette(52) // G. Oriber = 50, Wong Who = 51, Bill Bull = 52


	var oriberBody Pose
	oriberBody.Init("oriber.svg", rom, oriber)
	oriberBody.AddBlock(0x16, 0x00, 5, 7, 0, 0)
	oriberBody.Finalize()

	var wongBody Pose
	wongBody.Init("wong.svg", rom, wong)
	wongBody.AddBlock(0x16, 0x00, 5, 7, 0, 0)
	wongBody.Finalize()

	var billBody Pose
	billBody.Init("bill.svg", rom, bill)
	billBody.AddBlock(0x16, 0x00, 5, 7, 0, 0)
	billBody.Finalize()



	var oriberHead Pose
	oriberHead.Init("oriberHead.svg", rom, oriber)
	oriberHead.AddBlock(0x16, 0x01, 2, 2, 0, 0)
	oriberHead.Finalize()

	var wongHead Pose
	wongHead.Init("wongHead.svg", rom, wong)
	wongHead.AddBlock(0x17, 0x66, 2, 2, 0, 0)
	wongHead.Finalize()

	var billHead Pose
	billHead.Init("billHead.svg", rom, bill)
	billHead.AddBlock(0x17, 0x68, 2, 2, 0, 0)
	billHead.Finalize()

	var greyPalette = [16]color.RGBA{
		color.RGBA{0x00, 0x00, 0x00, 0xff},
		color.RGBA{0x22, 0x22, 0x22, 0xff},
		color.RGBA{0x33, 0x33, 0x33, 0xff},
		color.RGBA{0x44, 0x44, 0x44, 0xff},
		color.RGBA{0x55, 0x55, 0x55, 0xff},
		color.RGBA{0x66, 0x66, 0x66, 0xff},
		color.RGBA{0x77, 0x77, 0x77, 0xff},
		color.RGBA{0x88, 0x88, 0x88, 0xff},
		color.RGBA{0x99, 0x99, 0x99, 0xff},
		color.RGBA{0xaa, 0xaa, 0xaa, 0xff},
		color.RGBA{0xbb, 0xbb, 0xbb, 0xff},
		color.RGBA{0xcc, 0xcc, 0xcc, 0xff},
		color.RGBA{0xdd, 0xdd, 0xdd, 0xff},
		color.RGBA{0xee, 0xee, 0xee, 0xff},
		color.RGBA{0xff, 0xff, 0xff, 0xff},
		color.RGBA{0x00, 0x00, 0x00, 0x00},
	}

	var bodyBase Pose
	bodyBase.Init("bodyBase.svg", rom, &greyPalette)
	bodyBase.AddBlock(0x16, 0x00, 5, 7, 0, 0)
	bodyBase.Finalize()

	haggard_palette := GetPalette(2)
	var haggard Pose
	haggard.Init("haggard.svg", rom, haggard_palette)
	haggard.AddBlock(0xa, 0x08, 5, 7, 0, 0)
	haggard.Finalize()

	cody_palette := GetPalette(1)
	var cody Pose
	cody.Init("cody.svg", rom, cody_palette)
	cody.AddBlock(0x7, 0x60, 3, 7, 0, 0)
	cody.Finalize()

	barel_palette := GetPalette(62)
	var barel Pose
	barel.Init("barel.svg", rom, barel_palette)
	barel.AddBlock(23, 0xb8, 8, 5, 0, 0)
	barel.Finalize()

	axel_palette := GetPalette(23)
	var axel_flying Pose
	axel_flying.Init("axel_flying.svg", rom, axel_palette)
	axel_flying.AddBlock(16, 0xd7, 7, 3, 0, 0)
	axel_flying.Finalize()
	var axel_feets Pose
	axel_feets.Init("axel_feets.svg", rom, axel_palette)
	axel_feets.AddBlock(16, 0x01, 2, 1, 0, 0)
	axel_feets.Finalize()

	var axel_full1 Pose
	axel_full1.Init("axel_full1.svg", rom, axel_palette)
	axel_full1.AddBlock(15, 0x00, 16, 16, 0, 0)
	axel_full1.Finalize()

	var axel_full2 Pose
	axel_full2.Init("axel_full2.svg", rom, axel_palette)
	axel_full2.AddBlock(16, 0x00, 16, 16, 0, 0)
	axel_full2.Finalize()

	slash_palette := GetPalette(24)
	var slash_flying Pose
	slash_flying.Init("slash_flying.svg", rom, slash_palette)
	slash_flying.AddBlock(16, 0xd7, 7, 3, 0, 0)
	slash_flying.Finalize()

	var slash_feets Pose
	slash_feets.Init("slash_feets.svg", rom, slash_palette)
	slash_feets.AddBlock(16, 0x01, 2, 1, 0, 0)
	slash_feets.Finalize()

	var slash_koed Pose
	slash_koed.Init("slash_koed.svg", rom, slash_palette)
	slash_koed.AddBlock(15, 0x60, 8, 2, 0, 0)
	slash_koed.Finalize()

	var slash_head Pose
	slash_head.Init("slash_head.svg", rom, slash_palette)
	slash_head.AddBlock(15, 0x48, 2, 2, 0, 0)
	slash_head.Finalize()

	jake_palette := GetPalette(58)
	var jake Pose
	jake.Init("jake.svg", rom, jake_palette)
	jake.AddBlock(12, 0x0, 16, 16, 0, 0)
	jake.Finalize()
	var jake2 Pose
	jake2.Init("jake2.svg", rom, jake_palette)
	jake2.AddBlock(24, 0x0, 16, 16, 0, 0)
	jake2.Finalize()

	simon_palette := GetPalette(59)
	var simon Pose
	simon.Init("simon.svg", rom, simon_palette)
	simon.AddBlock(12, 0x0, 16, 16, 0, 0)
	simon.Finalize()
	var simon2 Pose
	simon2.Init("simon2.svg", rom, simon_palette)
	simon2.AddBlock(24, 0x0, 16, 16, 0, 0)
	simon2.Finalize()

	dug_palette := GetPalette(48)
	var dug Pose
	dug.Init("dug.svg", rom, dug_palette)
	dug.AddBlock(12, 0x0, 16, 16, 0, 0)
	dug.Finalize()
	var dug2 Pose
	dug2.Init("dug2.svg", rom, dug_palette)
	dug2.AddBlock(13, 0x0, 16, 16, 0, 0)
	dug2.Finalize()
	var dug3 Pose
	dug3.Init("dug3.svg", rom, dug_palette)
	dug3.AddBlock(24, 0x0, 16, 16, 0, 0)
	dug3.Finalize()

	bred_palette := GetPalette(46)
	var bred Pose
	bred.Init("bred.svg", rom, bred_palette)
	bred.AddBlock(12, 0x0, 16, 16, 0, 0)
	bred.Finalize()
	var bred2 Pose
	bred2.Init("bred2.svg", rom, bred_palette)
	bred2.AddBlock(24, 0x0, 16, 16, 0, 0)
	bred2.Finalize()


	dust_palette := GetPalette(25)
	var dust Pose
	dust.Init("dust.svg", rom, dust_palette)
	dust.AddBlock(19, 0x0, 6, 2, 0, 0)
	dust.Finalize()
}




