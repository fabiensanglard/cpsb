package main

type Pang3 struct {
	Game
}

func (game *Pang3) GetName() string {
	return game.name
}

func makePang3() Pang3 {
	var game Pang3
	game.gfxROMSize = 0x400000 // 2 MiB each
	game.gfx_banks = []RomSrc{
		{"pa3-01m.2c", "fa491874068924c39bcc7de93dfda3b27f5d9613", 2, 0x000000, 0x100000, 0x00000, 8},
		{"pa3-01m.2c", "fa491874068924c39bcc7de93dfda3b27f5d9613", 2, 0x100000, 0x100000, 0x00004, 8},
		{"pa3-07m.2f", "cfe68e24632b53fb6cd6d03b2166d6b5ba28b778", 2, 0x000000, 0x100000, 0x00002, 8},
		{"pa3-07m.2f", "cfe68e24632b53fb6cd6d03b2166d6b5ba28b778", 2, 0x100000, 0x100000, 0x00006, 8},
	}
	game.name = "pang3"
	game.paletteAddr = 0

	return game
}
