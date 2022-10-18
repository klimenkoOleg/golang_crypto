package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height-1; py++ {
		y1 := float64(py)/height*(ymax-ymin) + ymin
		//y2 := float64(py+1)/height*(ymax-ymin) + ymin
		for px := 0; px < width-1; px++ {
			x1 := float64(px)/width*(xmax-xmin) + xmin
			//x2 := float64(px+1)/width*(xmax-xmin) + xmin
			z1 := complex(x1, y1)
			//z2 := complex(x1, y2)
			//z3 := complex(x2, y1)
			//z4 := complex(x2, y2)
			// Image point (px, py) represents complex value z.
			c1 := color.RGBAModel.Convert(zPower4Minusl(z1)).(color.RGBA)
			//c2 := color.RGBAModel.Convert(mandelbrot(z2)).(color.RGBA)
			//c3 := color.RGBAModel.Convert(mandelbrot(z3)).(color.RGBA)
			//c4 := color.RGBAModel.Convert(mandelbrot(z4)).(color.RGBA)

			//c := color.RGBA{(c1.R + c2.R + c3.R + c4.R) / 4,
			//	(c1.G + c2.G + c3.G + c4.G) / 4,
			//	(c1.B + c2.B + c3.B + c4.B) / 4,
			//	255}

			img.Set(px, py, c1)
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 35

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			//return color.Gray{255 - contrast*n}
			if n > 15 {
				return color.RGBA{255 - contrast*n, 0, 0, 255}
			} else if n > 10 {
				return color.RGBA{0, 255 - contrast*n, 0, 255}
			} else if n > 5 {
				return color.RGBA{0, 0, 255 - contrast*n, 255}
			} else {
				return color.RGBA{128, 129, 128, 255}
			}
		}
	}
	return color.Black
}

func zPower4Minusl(z complex128) color.Color {
	const iterations = 200
	const contrast = 35

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v*v*v + z
		if cmplx.Abs(v-1) < 1e-2 || cmplx.Abs(v) < 1e-2 || cmplx.Abs(v) < 1e-2 || cmplx.Abs(v) < 1e-2 {
			//return color.Gray{255 - contrast*n}
			if n > 15 {
				return color.RGBA{255 - contrast*n, 0, 0, 255}
			} /*else if n > 10 {
				return color.RGBA{0, 255 - contrast*n, 0, 255}
			} else if n > 5 {
				return color.RGBA{0, 0, 255 - contrast*n, 255}
			} else {
				return color.RGBA{128, 129, 128, 255}
			}*/
		}
	}
	return color.Black
}
