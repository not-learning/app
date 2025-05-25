package clrs

import "image/color"

type Clr color.Color

var (
	Green Clr = color.RGBA{0, 150, 75, 255}
	Blue  Clr = color.RGBA{0, 125, 255, 255}
	White Clr = color.White
	Black Clr = color.Black
)

func RGB(r, g, b uint8) Clr {
	return Clr(color.RGBA{r, g, b, 255})
}

func YCC(y, cb, cr uint8) Clr {
	r, g, b := color.YCbCrToRGB(y, cb, cr)
	return color.RGBA{r, g, b, 255}
}
