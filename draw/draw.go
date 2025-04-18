package draw

import (
	l "github.com/not-learning/lmnts"
	v "github.com/not-learning/app/vec"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// TODO: have lists prepared for drawing
// TODO: deal with myClr

type Img struct { // todo a better name
	vertices []ebiten.Vertex
	indices  []uint16
	tr       *ebiten.Image
	x0, y0 float32
	//w, h float32
}

func Init(x1, y1, x2, y2 float32) *Img {
	img := &Img{}
	img.tr = ebiten.NewImage(2, 2)
	img.tr.Fill(color.RGBA{1, 1, 1, 1})
	w, h := x2-x1, y2-y1
	img.x0 = x1 + w/2
	img.y0 = y1 + h/2
	/*img.w = x2
	img.h = y2//*/
	return img
}

func (img *Img) Arc(scr *ebiten.Image, x, y, r, φ1, φ2 float32, clr color.Color) {
	var path vector.Path
	path.Arc(x+img.x0, -y+img.y0, r, -φ1, -φ2, vector.CounterClockwise)

	op := &vector.StrokeOptions{}
	op.Width = 1.5
	img.vertices, img.indices = path.AppendVerticesAndIndicesForStroke(img.vertices[:0], img.indices[:0], op)

	red, green, blue, _ := clr.RGBA()
	for i := range img.vertices {
		img.vertices[i].SrcX = 1
		img.vertices[i].SrcY = 1
		img.vertices[i].ColorR = float32(red)
		img.vertices[i].ColorG = float32(green)
		img.vertices[i].ColorB = float32(blue)
		img.vertices[i].ColorA = 1
	}

	trop := &ebiten.DrawTrianglesOptions{}
	trop.AntiAlias = true
	trop.FillRule = ebiten.FillRuleNonZero

	scr.DrawTriangles(img.vertices, img.indices, img.tr, trop)
}

func (img *Img) Poly(scr *ebiten.Image, crds []*v.VecF32, clr color.Color) {
	var path vector.Path
	for i, v := range crds {
		if i == 0 {
			path.MoveTo(v.X, v.Y)
			continue
		}
		path.LineTo(v.X, v.Y)
	}

	op := &vector.StrokeOptions{}
	op.Width = 1.5
	img.vertices, img.indices = path.AppendVerticesAndIndicesForStroke(img.vertices[:0], img.indices[:0], op)

	red, green, blue, _ := clr.RGBA()
	for i := range img.vertices {
		/*img.vertices[i].DstX += v.X
		img.vertices[i].DstY += v.Y//*/
		img.vertices[i].SrcX = 1
		img.vertices[i].SrcY = 1
		img.vertices[i].ColorR = float32(red)
		img.vertices[i].ColorG = float32(green)
		img.vertices[i].ColorB = float32(blue)
		img.vertices[i].ColorA = 1
	}

	trop := &ebiten.DrawTrianglesOptions{}
	trop.AntiAlias = true
	trop.FillRule = ebiten.FillRuleNonZero

	scr.DrawTriangles(img.vertices, img.indices, img.tr, trop)
}

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
}//*/
