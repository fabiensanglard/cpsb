package main

type SF2HF struct {
	Game
}

func (game *SF2HF) GetName() string {
	return game.name
}

func makeSF2HF() SF2HF {
	var game SF2HF
	game.gfxROMSize = 0x600000
	game.gfx_banks = []RomSrc{
		{"s92-1m.3a", "f370f25c96ad2b94f8c53d6b7139100285a25bef", 2, 0x000000, 0x80000, 0x00000, 8},
		{"s92-3m.5a", "2fb42a242f60ba7e74009b5a90eb26e035ba1e82", 2, 0x000000, 0x80000, 0x00002, 8},
		{"s92-2m.4a", "4c7d944fef200fdfcaf57758b901b5511188ed2e", 2, 0x000000, 0x80000, 0x00004, 8},
		{"s92-4m.6a", "27d3796429338d82a8de246a0ea06dd487a87768", 2, 0x000000, 0x80000, 0x00006, 8},

		{"s92-5m.7a", "4b696d66c51611e43522bed752654314e76d33b6", 2, 0x000000, 0x80000, 0x200000, 8},
		{"s92-7m.9a", "ebdf1f5e2638eed3a65dda82b1ed9151a355f4c9", 2, 0x000000, 0x80000, 0x200002, 8},
		{"s92-6m.8a", "4a4961bb68c3a1ce15f9d393d9c03ecb2466cc29", 2, 0x000000, 0x80000, 0x200004, 8},
		{"s92-8m.10a", "520390420da3a0271ba90b0a933e65143265e5cf", 2, 0x000000, 0x80000, 0x200006, 8},

		{"s92-10m.3c", "2868c31121b1c7564e9767b9a19cdbf655c7ed1d", 2, 0x000000, 0x80000, 0x400000, 8},
		{"s92-12m.5c", "648a59706b93c84b4206a968ecbdc3e834c476f6", 2, 0x000000, 0x80000, 0x400002, 8},
		{"s92-11m.4c", "ed6143f8737013b6ef1684e37c05e037e7a80dae", 2, 0x000000, 0x80000, 0x400004, 8},
		{"s92-13m.6c", "0083c0ffaf6fe7659ff0cf822be4346cd6e61329", 2, 0x000000, 0x80000, 0x400006, 8},
	}

	// Same as sf2 ROM
	game.codeROMSize = 0x100000
	game.code_banks = []RomSrc{
		{"sf2e_30g.11e", "22558eb15e035b09b80935a32b8425d91cd79669", 1, 0, 0x20000, 0x00000, 2},
		{"sf2e_37g.11f", "bf1ccfe7cc1133f0f65556430311108722add1f2", 1, 0, 0x20000, 0x00001, 2},

		{"sf2e_31g.12e", "86a3954335310865b14ce8b4e0e4499feb14fc12", 1, 0, 0x20000, 0x40000, 2},
		{"sf2e_38g.12f", "6565946591a18eaf46f04c1aa449ee0ae9ac2901", 1, 0, 0x20000, 0x40001, 2},

		{"sf2e_28g.9e", "bbcef63f35e5bff3f373968ba1278dd6bd86b593", 1, 0, 0x20000, 0x80000, 2},
		{"sf2e_35g.9f", "507bda3e4519de237aca919cf72e543403ec9724", 1, 0, 0x20000, 0x80001, 2},

		{"sf2_29b.10e", "75f0827f4f7e9f292add46467f8d4fe19b2514c9", 1, 0, 0x20000, 0xc0000, 2},
		{"sf2_36b.10f", "b807cc495bff3f95d03b061fc629c95f965cb6d8", 1, 0, 0x20000, 0xc0000, 2},
	}
	game.paletteAddr = 0x8ACBA

	game.name = "sf2hf"
	return game
}

func (game *SF2HF) Load() bool {
	if !game.Game.Load() {
		return false
	}
	game.w(0x00, game.RetrievePalette(1))
	game.s(123, 0xC8, 8, 2, game.RetrievePalette(0x11E))
	return true
}
