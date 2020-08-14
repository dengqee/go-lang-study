package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var red = color.RGBA{0xff, 0x00, 0x00, 0xff}
var green = color.RGBA{0x00, 0xff, 0x00, 0xff}
var bule = color.RGBA{0x00, 0x00, 0xff, 0xff}
var palette = []color.Color{color.White, red, green, bule}

const (
	whiteIndex = 0 //first color in palette
)

func main() {
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time
	// Thanks to Randal1 McPherson for pointing out the omission.
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			if int(t)%3 == 1 {
				img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
					1)
			} else if int(t)%3 == 2 {
				img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
					2)

			} else {
				img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
					3)

			}
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
