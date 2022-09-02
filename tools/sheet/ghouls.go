package main

type Ghouls struct {
	Game
}

func (game Ghouls) GetName() string {
	return game.name
}

func makeGhouls() Ghouls {
	var game Ghouls
	game.gfxROMSize = 0x300000
	game.gfx_banks = []RomSrc{
		{"dm-05.3a", "c4945b603115f32b7346d72426571dc2d361159f", 2, 0x00000, 0x80000, 0x00000, 8},
		{"dm-07.3f", "212176947933fcfef991bc80ad5bd91718689ffe", 2, 0x00000, 0x80000, 0x00002, 8},
		{"dm-06.3c", "35bc9dec5ddbf064c30c951627581c16764456ac", 2, 0x00000, 0x80000, 0x00004, 8},
		{"dm-08.3g", "7d0c4736f16577afe9966447a18f039728f6fbdf", 2, 0x00000, 0x80000, 0x00006, 8},
	}

	//
	//ROM_REGION( 0x300000, "gfx", 0 )
	//ROM_LOAD64_BYTE( "09.4a",    0x200000, 0x10000, CRC(ae24bb19) SHA1(aa91c6ffe657b878e10e4e930457b530f7bb529b) )
	//ROM_LOAD64_BYTE( "18.7a",    0x200001, 0x10000, CRC(d34e271a) SHA1(55211fc2861dce32951f41624c9c99c09bf3b184) )
	//ROM_LOAD64_BYTE( "13.4e",    0x200002, 0x10000, CRC(3f70dd37) SHA1(9ecb4dec9d131e9fca15ded7d71986a9fcb62c19) )
	//ROM_LOAD64_BYTE( "22.7e",    0x200003, 0x10000, CRC(7e69e2e6) SHA1(4e0b4d2474beaa5c869c8f1a91893c79ac6e7f39) )
	//ROM_LOAD64_BYTE( "11.4c",    0x200004, 0x10000, CRC(37c9b6c6) SHA1(b2bb82537e335339846dbf9588cfacfdbdd75ee6) )
	//ROM_LOAD64_BYTE( "20.7c",    0x200005, 0x10000, CRC(2f1345b4) SHA1(14c450abcf9defa29c6f48e5ffd0b9d1e4a66a1d) )
	//ROM_LOAD64_BYTE( "15.4g",    0x200006, 0x10000, CRC(3c2a212a) SHA1(f8fa0e0e2d09ea37c54d1c2493752b4e97e3f534) )
	//ROM_LOAD64_BYTE( "24.7g",    0x200007, 0x10000, CRC(889aac05) SHA1(9301dcecee699e7f7672bb498122e1f4831ce536) )
	//ROM_LOAD64_BYTE( "10.4b",    0x280000, 0x10000, CRC(bcc0f28c) SHA1(02f587aa4ae71631f27b0e3aaf1829cdded1bdc2) )
	//ROM_LOAD64_BYTE( "19.7b",    0x280001, 0x10000, CRC(2a40166a) SHA1(dc4e75d7ed87ae5386d721a09113bba364740465) )
	//ROM_LOAD64_BYTE( "14.4f",    0x280002, 0x10000, CRC(20f85c03) SHA1(86385139a9b42270aade758bfe338525936f5671) )
	//ROM_LOAD64_BYTE( "23.7f",    0x280003, 0x10000, CRC(8426144b) SHA1(2dbf9625413b302fcdad5bef8733a9dfbfaead52) )
	//ROM_LOAD64_BYTE( "12.4d",    0x280004, 0x10000, CRC(da088d61) SHA1(67229eff2827a42af97a60ceb252e132e7f307bc) )
	//ROM_LOAD64_BYTE( "21.7d",    0x280005, 0x10000, CRC(17e11df0) SHA1(42fb15e9300b07fc5f4bc21744484869859b130c) )
	//ROM_LOAD64_BYTE( "16.4h",    0x280006, 0x10000, CRC(f187ba1c) SHA1(6d9441d04ecef2a9d9c7a2cc7781acd7904c2061) )
	//ROM_LOAD64_BYTE( "25.7h",    0x280007, 0x10000, CRC(29f79c78) SHA1(26000a58454a06c3016f99ebc3a79c52911a7070) )

	game.name = "ghouls"
	game.paletteAddr = 4

	game.codeROMSize = 0x100000
	game.code_banks = []RomSrc{
		{"dme_29.10h", "f21fcf88d2ebb7bc9e8885fde760a5d82f295c1a", 1, 0, 0x20000, 0x00000, 2},
		{"dme_30.10j", "3613699213db47bfeabedf87f12eb0fa7e5973b6", 1, 0, 0x20000, 0x00001, 2},
		{"dme_27.9h", "fa230bf5503487ec11d767485a18f0a55dcc13d2", 1, 0, 0x20000, 0x40000, 2},
		{"dme_28.9j", "a07786062358c89f3b4634b8822173261802290b", 1, 0, 0x20000, 0x40001, 2},
		{"dm-17.7j", "c51f1c38cdaed77ad715cedd845617a291ab2441", 2, 0, 0x80000, 0x80000, 0},
	}

	return game
}

func (game *Ghouls) Load() bool {
	if !game.Game.Load() {
		return false
	}
	return true
}
