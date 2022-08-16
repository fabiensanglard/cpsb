package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"math/rand"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		println("Usage: desync src dst")
		return
	}

	input, err := os.Open(os.Args[1])
	if err != nil {
		println("Unable to open", os.Args[1])
		log.Fatal(err)
	}
	defer input.Close()

	src, err := png.Decode(input)
	if err != nil {
		log.Fatal(err)
	}

	upLeft := image.Point{0, 0}
	lowRight := image.Point{src.Bounds().Max.X, src.Bounds().Max.Y}
	dst := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	yOffsetSrc := 50
	xOffsetBase := 20
	for y := 0; y < src.Bounds().Max.Y-yOffsetSrc; y++ {
		xOffsetDst := xOffsetBase
		xOffsetDst += rand.Intn(10) // 0-10
		xOffsetDst += int(math.Cos(float64(y)/10.0) * 5)
		for x := 0; x < src.Bounds().Max.X; x++ {
			if x < xOffsetDst {
				dst.Set(x, y, color.RGBA{0, 0, 0, 0xff})
			} else {
				color := src.At(x-xOffsetDst, yOffsetSrc+y)
				dst.Set(x, y, color)
			}
		}
	}

	for y := src.Bounds().Max.Y - yOffsetSrc; y < src.Bounds().Max.Y; y++ {
		for x := 0; x < src.Bounds().Max.X; x++ {
			dst.Set(x, y, color.RGBA{0, 0, 0, 0xff})
		}
	}

	f, _ := os.Create(os.Args[2])
	png.Encode(f, dst)
}
