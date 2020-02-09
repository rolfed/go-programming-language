package main

import (
	"log"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"net/http"
	"fmt"
	"strconv"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // First color in palette
	blackIndex = 1 // Next color in palette
) 

func main() {
	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	cycles, err := strconv.Atoi(r.FormValue("cycles"))

	if err !=  {
		fmt.Print("Query param missing \n")
		cycles = 3
	}

  lissajous(w, cycles)
}

func lissajous(out io.Writer, cycles int) {
	const (
		res 	 = 0.001 	// Angular resolution
		size	 = 100 		// Image canvas colvers [-size..+size]
		nframes = 64 // Number of animation frames
		delay = 0 	// Delay between frames in 10ms units
	) 

	freq := rand.Float64() * 3.0 // Relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // Phase difference

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)

			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
