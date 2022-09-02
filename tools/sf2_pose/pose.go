package main

import (
	"fmt"
	"image/color"
	"os"
)
type Coordinate struct {
	x, y int32
}

type Rectangle struct {
	min Coordinate
	max Coordinate
}

type Pose struct {
	filename string
	rom []byte
	svg string
	rec Rectangle
	palette *Palette
}

func (a *Pose) Init(filename string, rom []byte, palette *Palette) {
	//var err error
	//a.file, err = os.Create("pics/" + filename)
	//check(err)
	a.rom = rom
	a.filename = "pics/" + filename
	a.palette = palette
}

func Max(x int32, y int32) int32 {
	if x > y {
		return x
	}
	return y
}

func Min(x int32, y int32) int32 {
	if x < y {
		return x
	}
	return y
}

func (a *Pose) GetColor(sheedID int32, tileID int32, x int32, y int32) color.RGBA {
	// A sheet is 16 x 16 tile. Each tile is 16 x 16 texels.
	const SHEET_SIZE = SHEET_WIDTH * SHEET_HEIGHT  * TILES_SIZE // 32 KiB

	offset := sheedID * SHEET_SIZE
	sheet := a.rom[offset:offset + SHEET_SIZE]

	offset = tileID * TILES_SIZE
	tile := sheet[offset:offset + TILES_SIZE]

	const LINE_SIZE = 16 / 2
	offset = y * LINE_SIZE
	line := tile[offset : offset + LINE_SIZE]

	offset = (x / 8) * 4
	bucket := line[offset : offset + 4]

	// Get four index values
	var bits = []byte{128, 64, 32, 16, 8 , 4 , 2 , 1}
	//mask := byte(1 << (x % 8))
	mask := bits[x % 8]

	var b1, b2, b3, b4 byte
	if bucket[0] & mask != 0 {b1 = 1} else {b1 = 0}
	if bucket[1] & mask != 0 {b2 = 1} else {b2 = 0}
	if bucket[2] & mask != 0 {b3 = 1} else {b3 = 0}
	if bucket[3] & mask != 0 {b4 = 1} else {b4 = 0}
	var value = b4 << 3 | b3 << 2 | b2 << 1 | b1

	return a.palette[value]
}

func (a *Pose) AddBlock(sheedID int32, tileID int32, width int32, height int32, x int32, y int32) {
   for w:= int32(0) ; w < width ; w++ {
	   for h := int32(0) ; h < height; h++ {
		   a.AddTile(sheedID, tileID + w + h * 0x10, x + w * 16, y + h * 16)
	   }
   }
}

func (a *Pose) AddTile(sheedID int32, tileID int32, tileX int32, tileY int32) {
	a.svg += `<g>`
	for x :=  int32(0) ; x < 16 ; x++ {
		for y:= int32(0) ; y < 16 ; y++ {
			color := a.GetColor(sheedID, tileID, x, y)
			txt_color := fmt.Sprintf("%02x%02x%02x", color.R, color.G, color.B)
			var opacity int32
			if color.A == 255 {
			  opacity = 1
		    } else {
				opacity = 0
			}
			//style := fmt.Sprintf(`fill:#%s;fill-opacity:1;stroke:#000000;stroke-width:1;stroke-opacity:1`, txt_color, opacity)
			style := fmt.Sprintf(`fill:#%s;fill-opacity:%d;`, txt_color, opacity)
			rect := fmt.Sprintf(`<rect x="%d" y="%d" style="%s" width="1" height="1" />
            `, tileX + x, tileY + y, style)
			a.svg += rect
		}
	}


	for x :=  int32(0) ; x < 16 ; x++ {
		style := "stroke-width:0.05;" //stroke-dasharray:1;
		line := fmt.Sprintf(`<line x1="%d" y1="%d" x2="%d" y2="%d" style="%s" stroke="black" />
        `,tileX + x, tileY, tileX + x, tileY + 16,style)
		a.svg += line
	}

	for y :=  int32(0) ; y < 16 ; y++ {
		style := "stroke-width:0.05;" //stroke-dasharray:1;
		line := fmt.Sprintf(`<line x1="%d" y1="%d" x2="%d" y2="%d" style="%s" stroke="black" />
        `,tileX, tileY + y, tileX + 16, tileY + y,style)
		a.svg += line
	}


	style := `fill-opacity:0;stroke:#000000;stroke-width:0.3;stroke-opacity:1`
	rect := fmt.Sprintf(`<rect x="%d" y="%d" style="%s" width="16" height="16" />
    `, tileX, tileY, style)
	a.svg += rect
	a.svg += `</g>`

	// stroke-dasharray:none; | 1
	a.rec.min.x = Min(a.rec.min.x, tileX)
	a.rec.min.y = Min(a.rec.min.y, tileY)

	a.rec.max.x = Max(a.rec.max.x, tileX + 16)
	a.rec.max.y = Max(a.rec.max.y, tileY + 16)
}

func (a *Pose) Finalize() {
	width := a.rec.max.x - a.rec.min.x
	height := a.rec.max.y - a.rec.min.y
	prolog := fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8" standalone="no"?>
    <svg width="%d.5" height="%d.5" 
     xmlns="http://www.w3.org/2000/svg"
     xmlns:svg="http://www.w3.org/2000/svg">
    `, width, height)
	a.svg = prolog + a.svg + `</svg>`
	os.WriteFile(a.filename, []byte(a.svg), 666)
}


