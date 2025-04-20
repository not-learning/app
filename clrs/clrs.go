package clrs

import "image/color"

type Clr color.Color

var (
	Green Clr = color.RGBA{0, 150, 75, 255}
	Blue  Clr = color.RGBA{0, 75, 255, 255}
	White Clr = color.White
)

func YCC(y, cb, cr uint8) Clr {
	r, g, b := color.YCbCrToRGB(y, cb, cr)
	return color.RGBA{r, g, b, 255}
}
