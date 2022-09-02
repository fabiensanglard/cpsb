package main

import "unsafe"

type CPS2Game struct {
	Game
}

func (game *CPS2Game) unshuffleBytes(buf []uint64, len int) {
	if len == 2 {
		return
	}

	//if len%4 == 0 {
	//	panic(len)
	//}

	len = len / 2

	game.unshuffleBytes(buf, len)
	game.unshuffleBytes(buf[len:], len)

	for i := 0; i < len/2; i++ {
		t := buf[len/2+i]
		buf[len/2+i] = buf[len+i]
		buf[len+i] = t
	}
}

func (game *CPS2Game) unshuffle() bool {
	const banksize = 0x200000
	size := len(game.gfxROM)

	for i := 0; i < size; i += banksize {
		slice := game.gfxROM[i:]
		castedRom := *(*[]uint64)(unsafe.Pointer(&slice))
		game.unshuffleBytes(castedRom, banksize/8)
	}
	return true
}

func (game *CPS2Game) Load() bool {
	if !game.Game.Load() {
		return false
	}
	return game.unshuffle()
}
