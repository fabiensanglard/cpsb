package main

import (
	"encoding/hex"
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"log"
	"os"
)

const TILE_WIDTH = 16
const TILE_HEIGHT = 16

const SHEET_WIDTH = 16
const SHEET_HEIGHT = 16

const TILES_SIZE = TILE_WIDTH * TILE_HEIGHT / 2

const PALETTE_SIZE = 32

type RomSrc struct {
	filename   string
	sha        string
	word_size  int
	src_offset int
	src_length int
	dst_offset int
	dst_stride int
}

type Game struct {
	palettes map[int][]*Palette

	gfx_banks  []RomSrc
	code_banks []RomSrc
	name       string

	gfxROMSize int
	gfxROM     []byte

	codeROMSize int
	codeROM     []byte

	paletteAddr int
	numPalettes int
}

func drawLine(line []byte, x int, y int, img *image.RGBA, palette *Palette) {
	if len(line) != TILE_WIDTH/2 {
		panic("Line buffer must be 8 bytes long")
	}
	cursor := 0
	for i := 0; i < 2; i++ {
		// Read four bytes
		bytes := make([]byte, 4)
		for j := 0; j < 4; j++ {
			bytes[j] = line[cursor]
			cursor += 1
		}

		// Get four index values
		var bits = []byte{128, 64, 32, 16, 8, 4, 2, 1}
		for j := 7; j >= 0; j-- {
			var b1, b2, b3, b4 byte
			if bytes[0]&bits[j] != 0 {
				b1 = 1
			} else {
				b1 = 0
			}
			if bytes[1]&bits[j] != 0 {
				b2 = 1
			} else {
				b2 = 0
			}
			if bytes[2]&bits[j] != 0 {
				b3 = 1
			} else {
				b3 = 0
			}
			if bytes[3]&bits[j] != 0 {
				b4 = 1
			} else {
				b4 = 0
			}
			var value = b4<<3 | b3<<2 | b2<<1 | b1
			// Write
			img.Set(x+j+i*8, y, palette.colors[value])
		}
	}
}

func (game *Game) drawTile(tile []byte, x int, y int, img *image.RGBA, sheetID int) {
	if len(tile) != TILES_SIZE {
		errorString := fmt.Sprintf("Tile buffer must be %d bytes long\n", TILES_SIZE)
		panic(errorString)
	}
	palette := game.GetPalette(sheetID, y/16*SHEET_WIDTH+x/16)
	for i := 0; i < TILE_HEIGHT; i++ {
		offset := i * TILE_WIDTH / 2
		drawLine(tile[offset:offset+TILE_WIDTH/2], x, y+i, img, palette)
	}
}

func (game *Game) dumpSheet(sheetID int, sheet []byte, folder string) {
	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{SHEET_WIDTH * TILE_WIDTH, SHEET_HEIGHT * TILE_HEIGHT}})
	for y := 0; y < SHEET_HEIGHT; y++ {
		for x := 0; x < SHEET_WIDTH; x++ {
			offset := (x + y*TILE_WIDTH) * TILES_SIZE
			game.drawTile(sheet[offset:offset+TILES_SIZE], x*TILE_WIDTH, y*TILE_HEIGHT, img, sheetID)
		}
	}
	filename := fmt.Sprintf("%s/%04d.png", folder, sheetID)
	f, _ := os.Create(filename)
	defer f.Close()
	png.Encode(f, img)
	png2svg(filename, fmt.Sprintf("%s/%04d.svg", folder, sheetID), sheetID)
}

func (game *Game) desinterleave(roms []RomSrc, dstROM []byte) bool {

	for _, rom := range roms {
		content, err := ioutil.ReadFile("./roms/" + rom.filename)
		if err != nil {
			fmt.Println("Unable to open '", rom.filename, "' for ", game.name)
			return false
		}
		hash := sha(content)
		hash_string := hex.EncodeToString(hash[:])
		if hash_string != rom.sha {
			fmt.Println(hash_string)
			fmt.Println(rom.filename)
			fmt.Println("File ", rom.filename, " bad sha. Got (", hash_string, ") but expected (", rom.sha, ")")
			return false
		}

		for j := 0; j < rom.src_length/rom.word_size; j++ {
			srcOffset := rom.src_offset + j*rom.word_size
			src := content[srcOffset : srcOffset+rom.word_size]

			dstOffset := rom.dst_offset + j*rom.dst_stride
			dst := dstROM[dstOffset : dstOffset+rom.word_size]

			for w := 0; w < rom.word_size; w++ {
				dst[w] = src[w]
			}
		}
	}
	return true
}

func (game *Game) Load() bool {
	fmt.Println("\nLoading GFX...", game.name)
	game.gfxROM = make([]byte, game.gfxROMSize)
	success := game.desinterleave(game.gfx_banks, game.gfxROM)
	if !success {
		return false
	}

	game.palettes = make(map[int][]*Palette)

	if game.code_banks != nil {
		fmt.Println("Loading Code...", game.name)
		game.codeROM = make([]byte, game.codeROMSize)
		game.desinterleave(game.code_banks, game.codeROM)
	}

	return true
}

func (game *Game) Extract() {
	game.ensureExtractFolder()
	// A sheet is 16 x 16 tile. Each tile is 16 x 16 texels.
	const SHEET_SIZE = SHEET_WIDTH * SHEET_HEIGHT * TILES_SIZE // 32 KiB
	// For each offset at the sprite location, dump a sheet
	numSheets := len(game.gfxROM) / SHEET_SIZE
	fmt.Println("Extracting ", numSheets, " sheets.")
	for i := 0; i < numSheets; i++ {
		offset := i * SHEET_SIZE
		sheet := game.gfxROM[offset : offset+SHEET_SIZE]
		game.dumpSheet(i, sheet, game.extractFolder())
	}
}

func (game *Game) ExtractPalette() {
	if game.codeROM == nil {
		return
	}
	filename := fmt.Sprintf("%s/palettes.html", game.extractFolder())
	f, _ := os.Create(filename)
	defer f.Close()
	numPalettes := game.numPalettes
	fmt.Println("Found ", numPalettes, " palettes")
	for i := 0; i < numPalettes; i++ {
		palette := game.RetrievePalette(i)
		f.WriteString(fmt.Sprintf("Palette %d<br/>\n", i))
		f.WriteString(palette.toHTML())
	}

	// Also save palette to disk
	filename = fmt.Sprintf("%s/code.bin", game.extractFolder())
	os.WriteFile(filename, game.codeROM, 0666)
}

func (game *Game) w(sheetID int, p *Palette) {
	sheet := make([]*Palette, 256)

	for i, _ := range sheet {
		sheet[i] = p
	}

	game.palettes[sheetID] = sheet
}

func (game *Game) s(sheetID int, tileID int, width int, height int, palette *Palette) {
	_, hasSheet := game.palettes[sheetID]
	if !hasSheet {
		game.w(sheetID, &greyPalette)
	}

	sheet := game.palettes[sheetID]

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			sheet[tileID+i*16+j] = palette
		}
	}
}

func (game *Game) u(sheetID int, tileID int, palette *Palette) {
	game.s(sheetID, tileID, 1, 1, palette)
}

func (game *Game) GetPalette(sheetID int, tileNumber int) *Palette {
	// If this sheet is unknown to us, just return default grey palette
	sheet, hasSheet := game.palettes[sheetID]
	if !hasSheet {
		return &greyPalette
	}

	// If the tileID in that sheet nil?
	if sheet[tileNumber] == nil {
		return &greyPalette
	}

	return sheet[tileNumber]
}

func (game *Game) RetrievePalette(paletteId int) *Palette {
	if game.codeROM == nil {
		return &greyPalette
	}

	base := game.paletteAddr + paletteId*PALETTE_SIZE
	paletteSlice := game.codeROM[base : base+PALETTE_SIZE]
	return PaletteFrom(paletteSlice)
}

func (game *Game) ensureExtractFolder() {
	var folder = game.extractFolder()
	err := os.MkdirAll(folder, 0777)
	if err != nil {
		log.Fatal(err)
	}
}

func (game *Game) extractFolder() string {
	return fmt.Sprintf("pics/%s", game.name)
}
