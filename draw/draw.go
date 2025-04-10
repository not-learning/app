package draw

import (
	l "github.com/not-learning/lmnts"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// TODO: have lists prepared for drawing
// TODO: deal with myClr

func Draw(scr *ebiten.Image, lls ...*l.Lmnt) {
	i := 0
	for _, k := range lls {
		i++
		n := float64(i) / 22.0
		j := 2 * math.Pi
		var a, b uint8 = uint8(255 * math.Sin(j*n)), uint8(255 * math.Cos(j*n))
		x1, y1, x2, y2 := k.Rect()
		vector.StrokeRect(scr, x1, y1, x2-x1, y2-y1, 2, myClr(a, b), true)
	}
}

func TestDraw(scr *ebiten.Image) func(*l.Lmnt) {
	i := 0
	return func(el *l.Lmnt) {
		i++
		n := float64(i) / 22.0
		j := 2 * math.Pi
		x1, y1, x2, y2 := el.Rect()
		a, b := uint8(255*math.Sin(j*n)), uint8(255*math.Cos(j*n))
		vector.StrokeRect(scr, x1, y1, x2-x1, y2-y1, 1.5, myClr(a, b), true) // todo thickness to 1?
		vector.StrokeLine(scr, x1, y1, x2, y2, 1.5, myClr(a, b), true)
	}
}

func myClr(x, y uint8) color.Color {
	r, g, b := color.YCbCrToRGB(255, x, y)
	return color.RGBA{r, g, b, 255}
}
