package main

import "image/color"

type SSF struct {
	CPS2Game
}

func (game *SSF) GetName() string {
	return game.name
}

func makeSSF() SSF {
	var game SSF
	game.gfxROMSize = 0xc00000
	game.gfx_banks = []RomSrc{
		{"ssf.13m", "bf2a6d98a656d1cb5734da7836686242d3211137", 2, 0x000000, 0x200000, 0x000000, 8},
		{"ssf.15m", "4b302dbb66e8a5c2ad92798699391e981bada427", 2, 0x000000, 0x200000, 0x000002, 8},
		{"ssf.17m", "b21b1c749a8241440879bf8e7cb33968ccef97e5", 2, 0x000000, 0x200000, 0x000004, 8},
		{"ssf.19m", "4d320fc96d1ef0b9928a8ce801734245a4c097a5", 2, 0x000000, 0x200000, 0x000006, 8},
		{"ssf.14m", "0f4d26af338dab5dce5b7b34d32ad0c573434ace", 2, 0x000000, 0x100000, 0x800000, 8},
		{"ssf.16m", "f4456833fb396e6501f4174c0fe5fd63ea40a188", 2, 0x000000, 0x100000, 0x800002, 8},
		{"ssf.18m", "4b060501e56b9d61294748da5387cdae5280ec4d", 2, 0x000000, 0x100000, 0x800004, 8},
		{"ssf.20m", "32b11ba7a12004aff810d719bff7508204c7b7c0", 2, 0x000000, 0x100000, 0x800006, 8},
	}

	game.paletteAddr = 0

	game.name = "ssf"
	return game
}

func (game *SSF) Load() bool {
	if !game.CPS2Game.Load() {
		return false
	}

	var cammy Palette
	cammy.colors[0] = color.RGBA{0x11, 0x11, 0x11, 0xff}
	cammy.colors[1] = color.RGBA{0xFF, 0xDD, 0xCC, 0xff}
	cammy.colors[2] = color.RGBA{0xff, 0xcc, 0xaa, 0xff}
	cammy.colors[3] = color.RGBA{0xee, 0xaa, 0x66, 0xff}
	cammy.colors[4] = color.RGBA{0xdd, 0x88, 0x44, 0xff}
	cammy.colors[5] = color.RGBA{0xcc, 0x66, 0x33, 0xff}
	cammy.colors[6] = color.RGBA{0x99, 0x33, 0x11, 0xff}
	cammy.colors[7] = color.RGBA{0x66, 0x44, 0x44, 0xff}
	cammy.colors[8] = color.RGBA{0x77, 0x66, 0x66, 0xff}
	cammy.colors[9] = color.RGBA{0x88, 0x88, 0x88, 0xff}
	cammy.colors[10] = color.RGBA{0x99, 0xbb, 0xbb, 0xff}
	cammy.colors[11] = color.RGBA{0xaa, 0xdd, 0xdd, 0xff}
	cammy.colors[12] = color.RGBA{0xee, 0x88, 0x00, 0xff}
	cammy.colors[13] = color.RGBA{0xff, 0xff, 0x00, 0xff}
	cammy.colors[14] = color.RGBA{0xff, 0xff, 0xff, 0xff}
	cammy.colors[15] = color.RGBA{0xff, 0xff, 0xff, 0xff}

	game.w(144, &cammy)

	return true
}
