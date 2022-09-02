package main

import "image/color"
import "fmt"

type Palette struct {
	colors [16]color.RGBA
}

//type Palette = [16]color.RGBA

// Palette are stored in int16 with format Ignored, Red, Green, and Blue.
// IIIIRRRRGGGGBBBB
// To convert 0xX to 0xXX, we do (* 16 + 1) = * 17
func PaletteFrom(array []byte) *Palette {
	var palette Palette

	for i := 0; i < 15; i++ {
		r := array[i*2]
		gb := array[i*2+1]
		palette.colors[i] = color.RGBA{r & 0xF * 17, (gb >> 4) * 17, gb & 0xF * 17, 0xff}
	}
	palette.colors[15] = color.RGBA{0x00, 0x00, 0x00, 0x00}
	return &palette
}

var greyPalette = Palette{[16]color.RGBA{
	color.RGBA{0x00, 0x00, 0x00, 0xff},
	color.RGBA{0x22, 0x22, 0x22, 0xff},
	color.RGBA{0x33, 0x33, 0x33, 0xff},
	color.RGBA{0x44, 0x44, 0x44, 0xff},
	color.RGBA{0x55, 0x55, 0x55, 0xff},
	color.RGBA{0x66, 0x66, 0x66, 0xff},
	color.RGBA{0x77, 0x77, 0x77, 0xff},
	color.RGBA{0x88, 0x88, 0x88, 0xff},
	color.RGBA{0x99, 0x99, 0x99, 0xff},
	color.RGBA{0xaa, 0xaa, 0xaa, 0xff},
	color.RGBA{0xbb, 0xbb, 0xbb, 0xff},
	color.RGBA{0xcc, 0xcc, 0xcc, 0xff},
	color.RGBA{0xdd, 0xdd, 0xdd, 0xff},
	color.RGBA{0xee, 0xee, 0xee, 0xff},
	color.RGBA{0xff, 0xff, 0xff, 0xff},
	color.RGBA{0x00, 0x00, 0x00, 0x00},
}}

func (palette *Palette) toHTML() string {
	header := `<table width="50%" border="1" cellpadding="0" cellspacing="0" style="border-collapse:collapse;"><tbody><tr>
  `

	for i := 0; i < 8; i++ {
		header += `<td class="t" style="background-color: #`
		header += fmt.Sprintf("%02x%02x%02xFF", palette.colors[i].R, palette.colors[i].G, palette.colors[i].B)
		header += `">&nbsp;</td>
    `
	}
	header += `</tr>
  <tr>
  `
	for i := 8; i < 15; i++ {
		header += `<td class="t" style="background-color: #`
		header += fmt.Sprintf("%02x%02x%02xFF", palette.colors[i].R, palette.colors[i].G, palette.colors[i].B)
		header += `">&nbsp;</td>
    `
	}
	// Last color is transparent
	header += `<td class="t" style="background-color: #000000">&nbsp;</td>`

	header += `</tr></tbody></table>
  `
	return header
}
