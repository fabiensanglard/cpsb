package main

type Strider struct {
	Game
}

func (game Strider) GetName() string {
	return game.name
}

func makeStrider() Strider {
	var game Strider
	game.gfxROMSize = 0x400000
	game.gfx_banks = []RomSrc{
		{"st-2.8a", "5e619fd5f3f1181e32a8fd9dbb4661d74ff8a484", 2, 0, 0x80000, 0x00000, 8},
		{"st-11.10a", "593cec513de40ff802084d54313bb25a4561e25d", 2, 0, 0x80000, 0x00002, 8},
		{"st-5.4a", "6cbfa30b2852fd117d117beefba434ce41d24c2f", 2, 0, 0x80000, 0x00004, 8},
		{"st-9.6a", "cf71c62348ca6b404279e87a6686cb3a842eb381", 2, 0, 0x80000, 0x00006, 8},
		{"st-1.7a", "e6f65af7cc3295be9efaaded352e7ae6320b4133", 2, 0, 0x80000, 0x200000, 8},
		{"st-10.9a", "bb0926dc484dae4f64c5e5a6bce20afdc7aeba55", 2, 0, 0x80000, 0x200002, 8},
		{"st-4.3a", "5c5a079baa694927c33d0e0c23e5ff09d6c9d985", 2, 0, 0x80000, 0x200004, 8},
		{"st-8.5a", "759b8b1fc7a5c4b00d74a27c2dd11667db44b09e", 2, 0, 0x80000, 0x200006, 8},
	}

	game.name = "strider"
	game.paletteAddr = 0

	return game
}

func (game *Strider) Load() bool {
	if !game.Game.Load() {
		return false
	}
	return true
}
