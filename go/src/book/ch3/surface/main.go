// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
)

const (
	cells   = 100         // number of grid cells
	xyrange = 30.0        // axis ranges (-xyrange..+xyrange)
	angle   = math.Pi / 6 // angle of x, y axes (=30°)
	zfactor = 1
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		width, height := 600., 320. // canvas size in pixels
		color := "red"
		for k, v := range r.Form {
			switch k {
			case "width":
				width, _ = strconv.ParseFloat(v[0], 64)
			case "height":
				height, _ = strconv.ParseFloat(v[0], 64)
			case "color":
				color = v[0]
			}
		}
		surface(w, width, height, color)
	}
	http.HandleFunc("/", handler)
	//!-http
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	return

}

func surface(w io.Writer, width, height float64, color string) {
	xyscale := width / 2 / xyrange // pixels per x or y unit
	zscale := height * 0.4         // pixels per z unit
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, width, height, xyscale, zscale)
			bx, by := corner(i, j, width, height, xyscale, zscale)
			cx, cy := corner(i, j+1, width, height, xyscale, zscale)
			dx, dy := corner(i+1, j+1, width, height, xyscale, zscale)
			points := [8]float64{ax, ay, bx, by, cx, cy, dx, dy}
			valid := true
			for _, val := range points {
				if math.IsNaN(val) {
					valid = false
				}
			}
			if !valid {
				continue
			}

			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='%s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}
	fmt.Fprintln(w, "</svg>")
}

func corner(i, j int, width, height, xyscale, zscale float64) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Pow(math.Sin(x), 2.)/r + math.Pow(math.Sin(y), 2.)/r
}

//!-
