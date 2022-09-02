package main

type FFight struct {
	Game
}

func (game FFight) GetName() string {
	return game.name
}

func makeFFight() FFight {
	var game FFight
	game.gfxROMSize = 0x200000
	game.gfx_banks = []RomSrc{
		{"ff-5m.7a", "7868f5801347340867720255f8380548ad1a65a7", 2, 0x00000, 0x80000, 0x00000, 8},
		{"ff-7m.9a", "f7b00a3ca8cb85264ab293089f9f540a8292b49c", 2, 0x00000, 0x80000, 0x00002, 8},
		{"ff-1m.3a", "5ce16af72858a57aefbf6efed820c2c51935882a", 2, 0x00000, 0x80000, 0x00004, 8},
		{"ff-3m.5a", "df5f3d3aa96a7a33ff22f2a31382942c4c4f1111", 2, 0x00000, 0x80000, 0x00006, 8},
	}
	game.name = "ffight"

	game.codeROMSize = 0x400000
	game.code_banks = []RomSrc{
		{"ff_36.11f", "0756ae576a1f6d5b8b22f8630dca40ef38567ea6", 1, 0, 0x20000, 0x00000, 2},
		{"ff_42.11h", "5045a467f3e228c02b4a355b52f58263ffa90113", 1, 0, 0x20000, 0x00001, 2},

		{"ff_37.12f", "38f44434c8befd623953ae23d6e5ff4e201d6627", 1, 0, 0x20000, 0x40000, 2},
		{"ffe_43.12h", "de16873d1639ac1738be0937270b108a9914f263", 1, 0, 0x20000, 0x40001, 2},

		{"ff-32m.8h", "d3362dadded31ccb7eaf71ef282d698d18edd722", 1, 0, 0x80000, 0x80000, 1},
	}
	game.paletteAddr = 0xC0000
	game.numPalettes = 300
	return game
}

func (game *FFight) Load() bool {
	if !game.Game.Load() {
		return false
	}

	// Swap the latest part 16-bit WORDS.
	// This is a weird ROM but i am too lazy to make my deinterlave() smarter.
	// So I just hack it in here.
	for i := 0; i < 0x80000 -1; i+=2 {
		v1 := game.codeROM[0x80000+i+0]
		v2 := game.codeROM[0x80000+i+1]
		game.codeROM[0x80000+i+0] = v2
		game.codeROM[0x80000+i+1] = v1
	}

	guy := game.RetrievePalette(0)
	cody := game.RetrievePalette(1)
	hagard := game.RetrievePalette(2)
	dude := game.RetrievePalette(50) // G. Oriber = 50, Wong Who = 51, Bill Bull = 52
	damned := game.RetrievePalette(31)
    barrel := game.RetrievePalette(62)

	axl := game.RetrievePalette(24) // axl 23, slash = 24
	bred := game.RetrievePalette(46) // jake = 58, simons = 59, dug = 48, bred = 46

	dust := game.RetrievePalette(25)

	for i := 0; i < 4; i++ {
		game.w(i, guy)
	}


	for i := 4; i < 8; i++ {
		game.w(i, cody)
	}


	for i := 9; i < 12; i++ {
		game.w(i, hagard)
	}

	for i := 15; i < 17; i++ {
		game.w(i, axl)
	}

	for i := 19; i < 20; i++ {
		game.w(i, dust)
	}

	for i := 24; i < 25; i++ {
		game.w(i, bred)
	}

	for i := 22; i < 24; i++ {
		game.w(i, dude)
	}

	for i := 23; i < 24; i++ {
		game.w(i, barrel)
	}

	for i := 26; i < 28; i++ {
		game.w(i, damned)
	}



	return true
}
