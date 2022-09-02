package main

type CarrierAirWing struct {
	Game
}

func (game CarrierAirWing) GetName() string {
	return game.name
}

func makeCAW() CarrierAirWing {
	var game CarrierAirWing
	game.gfxROMSize = 0x400000
	game.gfx_banks = []RomSrc{
		{"ca-5m.7a", "d355ea64ff29d228dcbfeee72bcf11882bf1cd9d", 2, 0, 0x80000, 0x00000, 8},
		{"ca-7m.9a", "bdb6820b81fbce77d7eacb01777af7c380490402", 2, 0, 0x80000, 0x00002, 8},
		{"ca-1m.3a", "5f62cd551b6a230edefd81fa60c10c84186ca804", 2, 0, 0x80000, 0x00004, 8},
		{"ca-3m.5a", "c31f0e78f49d94ea9dea20eb0cbd98a6c613bcbf", 2, 0, 0x80000, 0x00006, 8},
	}

	game.name = "carrierairwing"
	game.paletteAddr = 0

	return game
}

func (game *CarrierAirWing) Load() bool {
	if !game.Game.Load() {
		return false
	}
	return true
}
