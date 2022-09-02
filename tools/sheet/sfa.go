package main

type SFA struct {
	CPS2Game
}

func (game *SFA) GetName() string {
	return game.name
}

func makeSFA() SFA {
	var game SFA
	game.gfxROMSize = 0x1000000
	game.gfx_banks = []RomSrc{
		{"sfz.14m", "5eb28c8de57acfeaefebdd01509c7d9ba5244705", 2, 0x000000, 0x200000, 0x0000000, 8},
		{"sfz.16m", "07588f1ba6addc04fef3274c971174aaf3e632ab", 2, 0x000000, 0x200000, 0x0000002, 8},
		{"sfz.18m", "ce25dad542308691dbe9606b26279bbd59ea4b81", 2, 0x000000, 0x200000, 0x0000004, 8},
		{"sfz.20m", "f054e95df650a891ef56da8bfb31cb2c945a9aed", 2, 0x000000, 0x200000, 0x0000006, 8},
	}

	game.paletteAddr = 0

	game.name = "sfa"
	return game
}
