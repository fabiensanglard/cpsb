package main

import (
	"image/color"
)
import "crypto/sha1"

func sha(array []byte) []byte {
	h := sha1.New()
	h.Write(array)
	return h.Sum(nil)
}

type SF2 struct {
	Game
}

func (game *SF2) GetName() string {
	return game.name
}

func makeSF2() SF2 {
	var game SF2

	game.gfxROMSize = 0x600000
	game.gfx_banks = []RomSrc{
		{"sf2-5m.4a", "b9194fb337b30502c1c9501cd6c64ae4035544d4", 2, 0, 0x80000, 0x0000000, 8},
		{"sf2-7m.6a", "3759b851ac0904ec79cbb67a2264d384b6f2f9f9", 2, 0, 0x80000, 0x0000002, 8},
		{"sf2-1m.3a", "520840d727161cf09ca784919fa37bc9b54cc3ce", 2, 0, 0x80000, 0x0000004, 8},
		{"sf2-3m.5a", "2360cff890551f76775739e2d6563858bff80e41", 2, 0, 0x80000, 0x0000006, 8},

		{"sf2-6m.4c", "357c2275af9133fd0bd6fbb1fa9ad5e0b490b3a2", 2, 0, 0x80000, 0x200000, 8},
		{"sf2-8m.6c", "baa92b91cf616bc9e2a8a66adc777ffbf962a51b", 2, 0, 0x80000, 0x200002, 8},
		{"sf2-2m.3c", "2eea16673e60ba7a10bd4d8f6c217bb2441a5b0e", 2, 0, 0x80000, 0x200004, 8},
		{"sf2-4m.5c", "f787aab98668d4c2c54fc4ba677c0cb808e4f31e", 2, 0, 0x80000, 0x200006, 8},

		{"sf2-13m.4d", "5669b845f624b10e7be56bfc89b76592258ce48b", 2, 0, 0x80000, 0x400000, 8},
		{"sf2-15m.6d", "9af9df0826988872662753e9717c48d46f2974b0", 2, 0, 0x80000, 0x400002, 8},
		{"sf2-9m.3d", "a6a7f4725e52678cbd8d557285c01cdccb2c2602", 2, 0, 0x80000, 0x400004, 8},
		{"sf2-11m.5d", "f9a92d614e8877d648449de2612fc8b43c85e4c2", 2, 0, 0x80000, 0x400006, 8},
	}

	game.codeROMSize = 0x400000
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
    game.numPalettes = 300

	game.name = "sf2"
	return game
}

func (game *SF2) Load() bool {
	if !game.Game.Load() {
		return false
	}

	font := game.RetrievePalette(0)

	ryu := game.RetrievePalette(1)
	hon := game.RetrievePalette(2)
	bla := game.RetrievePalette(3)
	gui := game.RetrievePalette(4)
	ken := game.RetrievePalette(5)
	chu := game.RetrievePalette(6)
	zan := game.RetrievePalette(7)
	dal := game.RetrievePalette(8)

	// Fireball
	fb := game.RetrievePalette(0xE)

	dic := game.RetrievePalette(0x90)
	box := game.RetrievePalette(0xB0)
	sag := game.RetrievePalette(0xA0)
	dac := game.RetrievePalette(0xC0)

	flam := game.RetrievePalette(15)
	elec := game.RetrievePalette(161)

	warrier := game.RetrievePalette(0x11E)
	logo := game.RetrievePalette(0x11F)
	health_bar := game.RetrievePalette(12)
	credit := game.RetrievePalette(273)
	shimo := game.RetrievePalette(274)
	shimo.colors[10] = color.RGBA{0xFF, 0xFF, 0xFF, 0xff}
	shimo.colors[12] = color.RGBA{0xFF, 0xFF, 0xFF, 0xff}

	game.w(0, ryu)
	game.s(0, 0xED, 3, 2, ken)

	game.w(1, ken)
	game.s(1, 0x0, 4, 5, ryu)
	game.s(1, 0x48, 8, 11, ryu)
	game.s(1, 0xA1, 3, 6, ryu)
	game.s(1, 0xD0, 1, 1, ryu)
	game.s(1, 0x84, 4, 6, ryu)

	game.w(2, ryu)

	for i := 4; i <= 9; i++ {
		game.w(i, hon)
	}
	game.s(4, 0x01, 7, 5, flam)

	for i := 10; i <= 13; i++ {
		game.w(i, bla)
	}

	for i := 14; i <= 16; i++ {
		game.w(i, hon)
	}

	game.s(14, 0x67, 3, 1, fb)
	game.s(14, 0x5B, 3, 1, fb)
	game.s(14, 0x4F, 1, 1, fb)

	game.w(17, ryu)

	for i := 18; i <= 22; i++ {
		game.w(i, bla)
	}

	for i := 23; i <= 33; i++ {
		game.w(i, zan)
	}

	game.s(33, 0xD0, 16, 3, flam)
	game.s(33, 0x8A, 6, 5, flam)

	for i := 34; i <= 39; i++ {
		game.w(i, hon)
	}

	for i := 40; i <= 53; i++ {
		game.w(i, dal)
	}

	game.s(51, 0x60, 3, 6, chu)
	game.s(51, 0x63, 1, 2, chu)
	game.s(51, 0x3C, 1, 1, bla)
	game.s(51, 0x4C, 3, 1, bla)
	game.s(51, 0x4D, 2, 3, bla)

	for i := 54; i <= 59; i++ {
		game.w(i, dac)
	}

	for i := 60; i <= 70; i++ {
		game.w(i, ryu)
	}

	game.s(69, 0x0D, 1, 1, credit)
	game.s(69, 0x40, 1, 1, credit)
	game.s(69, 0x43, 1, 1, credit)
	game.s(69, 0x4E, 1, 1, credit)
	game.s(69, 0xC2, 1, 1, credit)
	game.s(69, 0x54, 1, 1, credit)
	game.s(69, 0x58, 1, 1, credit)
	game.s(69, 0x5F, 1, 1, credit)
	game.s(69, 0xA1, 1, 1, credit)
	game.s(69, 0xA6, 1, 2, shimo)
	game.s(69, 0xB6, 3, 1, shimo)
	game.s(69, 0xBE, 2, 1, credit)

	game.s(69, 0x98, 1, 1, ken)
	game.s(69, 0xAA, 1, 1, ken)
	game.s(69, 0xB9, 2, 1, ken)

	game.s(69, 0xC2, 2, 2, credit)
	game.s(69, 0xC4, 1, 1, credit)

	game.s(69, 0xC5, 2, 1, chu)
	game.s(69, 0xD4, 2, 2, chu)
	game.s(69, 0xE0, 1, 1, chu)
	game.s(69, 0xE3, 1, 1, chu)
	game.s(69, 0xE6, 2, 1, chu)
	game.s(69, 0xF8, 1, 1, chu)
	game.s(69, 0xCC, 4, 4, chu)

	game.w(71, ken)
	game.s(71, 0xC0, 2, 2, shimo)
	game.s(71, 0x38, 2, 1, shimo)
	game.s(71, 0x8D, 3, 4, flam)
	game.s(71, 0xCF, 1, 4, flam)
	game.s(71, 0x2D, 3, 1, chu)

	for i := 72; i <= 77; i++ {
		game.w(i, box)
	}

	for i := 78; i <= 83; i++ {
		game.w(i, sag)
	}
	game.s(78, 0x00, 4, 1, box)
	game.s(78, 0x10, 3, 4, box)

	game.w(84, chu)

	for i := 85; i <= 89; i++ {
		game.w(i, bla)
	}

	for i := 90; i <= 98; i++ {
		game.w(i, dic)
	}

	game.s(99, 0x24, 4, 8, elec)

	for i := 100; i <= 101; i++ {
		game.w(i, gui)
	}

	for i := 102; i <= 111; i++ {
		game.w(i, chu)
	}

	for i := 112; i <= 113; i++ {
		game.w(i, bla)
	}

	for i := 114; i <= 122; i++ {
		game.w(i, gui)
	}

	game.s(123, 0xC8, 8, 2, warrier)

	game.w(128, font)
	game.w(129, font)

	game.s(129, 0xDE, 2, 2, health_bar)
	game.s(129, 0xF0, 16, 1, health_bar)

	game.s(130, 0x40, 16, 7, logo)
	game.s(130, 0xB1, 7, 1, logo)
	game.s(130, 0xC0, 9, 3, logo)
	game.s(130, 0xF0, 8, 1, logo)

	for i := 136; i <= 143; i++ {
		game.w(i, zan)
	}

	return true
}

//
//const CODE_ROMS_SIZE = 1 << 17 // 128 KiB
//const CODE_ROMS_PER_BANK = 2
//const CODE_BANKS = 4
//
//func desinterleave_code_bank(roms [][2]string, dst []byte) {
//	var files [CODE_ROMS_PER_BANK][]byte
//	for i, _ := range files {
//		content, err := ioutil.ReadFile("./roms/" + roms[i][0])
//		hash := sha(content)
//		hash_string := hex.EncodeToString(hash[:])
//		if hash_string != roms[i][1] {
//			fmt.Println(hash_string)
//			fmt.Println(roms[i][1])
//			panic("Unexpected file")
//		}
//
//		if err != nil {
//			panic(err)
//		}
//		files[i] = content
//	}
//
//	var cursor = 0
//	for i := 0; i < CODE_ROMS_SIZE; i++ {
//		for _, f := range files {
//			dst[cursor] = f[i]
//			cursor += 1
//		}
//	}
//
//}
