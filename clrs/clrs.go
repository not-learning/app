package clrs

import "image/color"

var (
	Green = color.RGBA{0, 150, 75, 255}
	White = color.White
)

func YCC(y, cb, cr uint8) color.Color {
	r, g, b := color.YCbCrToRGB(y, cb, cr)
	return color.RGBA{r, g, b, 255}
}
