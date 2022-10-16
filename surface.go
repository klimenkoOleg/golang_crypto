package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.04       // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)]

var maxX, maxY float64

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, _ := corner(i+1, j)
			bx, by, _ := corner(i, j)
			cx, cy, _ := corner(i, j+1)
			dx, dy, _ := corner(i+1, j+1)
			/*var colorStr string
			if z > 0 {
				colorStr = fmt.Sprintf("#%x00", int((z)*16))
				//if z < 0.3 {
				//	colorStr = fmt.Sprintf("#%x00", int((1-z)*8))
				//
				//}
			} else {
				colorStr = fmt.Sprintf("#0%x0", int(-z*16))
			}*/
			//0.8 - 1
			//x - 16
			//17
			//fmt.Println(colorStr)
			//colorStr = "#f00"
			//fmt.Printf("<polygon style='fill:"+colorStr+"; stroke:#000; stroke-width: 1;' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
	//fmt.Println(maxX)
	//fmt.Println(maxY)
}

func corner(i, j int) (float64, float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	if x > maxX {
		maxX = x
	}
	if y > maxY {
		maxY = y
	}
	// Compute surface height z.
	z := f(x, y)
	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func f(x, y float64) float64 {

	//return x/15.*x/15. - y/15.*y/15.

	return math.Sin(x) + math.Sin(y)

	//r := math.Hypot(x, y) // distance from (0,0)
	//return math.Sin(r) / r
}
