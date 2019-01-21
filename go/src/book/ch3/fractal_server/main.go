package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math"
	"math/cmplx"
	"net/http"
	"strconv"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		xmin, ymin, xmax, ymax, zoomLevel := -2., -2., +2., +2., 1.
		for k, v := range r.Form {
			switch k {
			case "x1":
				xmin, _ = strconv.ParseFloat(v[0], 64)
			case "x2":
				xmax, _ = strconv.ParseFloat(v[0], 64)
			case "y1":
				ymin, _ = strconv.ParseFloat(v[0], 64)
			case "y2":
				ymax, _ = strconv.ParseFloat(v[0], 64)
			case "zoom":
				zoomLevel, _ = strconv.ParseFloat(v[0], 64)
			}
		}
		display(w, xmin, ymin, xmax, ymax, zoomLevel)
	}
	http.HandleFunc("/", handler)
	//!-http
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	return
}

func display(w io.Writer, xminBase, yminBase, xmaxBase, ymaxBase, zoomLevel float64) {
	const (
		width, height = 1024, 1024
	)

	xmin, ymin, xmax, ymax := zoom(xminBase, yminBase, xmaxBase, ymaxBase, zoomLevel)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, newton(z))
		}
	}
	png.Encode(w, img) // NOTE: ignoring errors
}

func zoom(x1, x2, y1, y2, z float64) (float64, float64, float64, float64) {
	return x1 / z, x2 / z, y1 / z, y2 / z
}

func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 5
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			val := 255 - contrast*i
			re := math.Round(real(z))
			im := math.Round(imag(z))
			if re == 1 {
				return color.RGBA{255, 0, 0, val}
			} else if re == -1 {
				return color.RGBA{0, 255, 0, val}
			} else if im == 1 {
				return color.RGBA{0, 0, 255, val}
			} else if im == -1 {
				return color.RGBA{255, 255, 255, val}
			}
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}
