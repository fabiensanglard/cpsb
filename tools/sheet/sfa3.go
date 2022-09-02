package main

type SF3A3 struct {
	CPS2Game
}

func (game *SF3A3) GetName() string {
	return game.name
}

func makeSFA3() SF3A3 {
	var game SF3A3
	game.gfxROMSize = 0x2000000
	game.gfx_banks = []RomSrc{
		{"sz3.13m", "c69e0ee22537312909dacc86d2e4be319d54e426", 2, 0x000000, 0x400000, 0x0000000, 8},
		{"sz3.15m", "f4ac4bfe830dc7df9fe4f680e4e0c053e7cbd8fe", 2, 0x000000, 0x400000, 0x0000002, 8},
		{"sz3.17m", "37f331fbb1284db446faecade6f484f58c0e1b2a", 2, 0x000000, 0x400000, 0x0000004, 8},
		{"sz3.19m", "f14136564648f006c1b74afda78349f260524b5f", 2, 0x000000, 0x400000, 0x0000006, 8},

		{"sz3.14m", "9e0ce43380b776c7a03872bafd4856f6fa60bda7", 2, 0x000000, 0x400000, 0x1000000, 8},
		{"sz3.16m", "7918204dc457f7a146d8fb8cf7242dfed3109fd8", 2, 0x000000, 0x400000, 0x1000002, 8},
		{"sz3.18m", "c18c56822b90a71ca5fbdf3440eb2671011f3d8f", 2, 0x000000, 0x400000, 0x1000004, 8},
		{"sz3.20m", "af60a5116c1ca9050366a35ea29128921867f3cc", 2, 0x000000, 0x400000, 0x1000006, 8},
	}

	game.paletteAddr = 0

	game.name = "sfa3"
	return game
}
