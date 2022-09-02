package main

type ForgottenUE struct {
	Game
}

func (game ForgottenUE) GetName() string {
	return game.name
}

func makeForgottenUE() ForgottenUE {
	var game ForgottenUE
	game.gfxROMSize = 0x400000
	game.gfx_banks = []RomSrc{
		{"lw-01.9d", "178ffd6da7bf845e30abf1bfc38a469cd319a73f", 2, 0, 0x80000, 0x00000, 8},
		{"lw-08.9f", "d57cee1fc508db2677e84882fb814e4d9ad20543", 2, 0, 0x80000, 0x00002, 8},
		{"lw-05.9e", "11147afc475904848458425661473586dd6f60cc", 2, 0, 0x80000, 0x00004, 8},
		{"lw-12.9g", "d63a1331fda2365f090fa31950098f321a720ea8", 2, 0, 0x80000, 0x00006, 8},
		{"lw-02.12d", "d3e6c971de0477ec4e178adc82508208dd8b397f", 2, 0, 0x80000, 0x200000, 8},
		{"lw-09.12f", "95e61af338945e690f2a82746feba3871ea224eb", 2, 0, 0x80000, 0x200002, 8},
		{"lw-06.12e", "6fd8f4a3ab070733b52365ab1945bf86acb2bf62", 2, 0, 0x80000, 0x200004, 8},
		{"lw-13.12g", "00f2c0050fd106276ea5398511c5861ebfbc0d10", 2, 0, 0x80000, 0x200006, 8},
	}

	game.name = "forgottnue"
	game.paletteAddr = 0

	return game
}

func (game *ForgottenUE) Load() bool {
	if !game.Game.Load() {
		return false
	}
	return true
}
