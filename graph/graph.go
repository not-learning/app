package graph

import (
	"image/color"
	"math"

	"github.com/not-learning/app/clrs"
	"github.com/not-learning/app/fonts"
	"github.com/not-learning/app/vec"

	"github.com/not-learning/lmnts"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// TODO: have lists prepared for drawing
// TODO: deal with myClr

type Graph struct {
	*fonts.Font

	vertices []ebiten.Vertex
	indices  []uint16
	tr       *ebiten.Image

	x0, y0   float32
}

func Init() *Graph {
	g := &Graph{}
	g.Font = fonts.InitFont()
	g.tr = ebiten.NewImage(2, 2)
	g.tr.Fill(color.RGBA{1, 1, 1, 1})
	return g
}

func (g *Graph) SetOrigin(x0, y0 float32) { g.x0, g.y0 = x0, y0 }

func (g *Graph) CirEmp(scr *ebiten.Image, x, y, r float32, clr clrs.Clr) {
	vector.StrokeCircle(scr, x+g.x0, -y+g.y0, r, 1, clr, true)
}

func (g *Graph) CirFull(scr *ebiten.Image, x, y, r float32, clr clrs.Clr) {
	vector.DrawFilledCircle(scr, x+g.x0, -y+g.y0, r, clr, true)
}

func (g *Graph) Arc(scr *ebiten.Image, x, y, r, φ1, φ2 float32, clr clrs.Clr) {
	var path vector.Path
	path.Arc(x+g.x0, -y+g.y0, r, -φ1, -φ2, vector.CounterClockwise)

	op := &vector.StrokeOptions{}
	op.Width = 1.5
	g.vertices, g.indices = path.AppendVerticesAndIndicesForStroke(g.vertices[:0], g.indices[:0], op)

	red, green, blue, _ := clr.RGBA()
	for i := range g.vertices {
		g.vertices[i].SrcX = 1
		g.vertices[i].SrcY = 1
		g.vertices[i].ColorR = float32(red / 256)   //+ float32(i)
		g.vertices[i].ColorG = float32(green / 256) //- float32(i)
		g.vertices[i].ColorB = float32(blue / 256)
		g.vertices[i].ColorA = 1
	}
	trop := &ebiten.DrawTrianglesOptions{}
	trop.AntiAlias = true
	trop.FillRule = ebiten.FillRuleNonZero

	scr.DrawTriangles(g.vertices, g.indices, g.tr, trop)
}

func (g *Graph) Poly(scr *ebiten.Image, crds []*vec.VecF32, clr clrs.Clr) {
	var path vector.Path
	for i, v := range crds {
		if i == 0 {
			path.MoveTo(v.X+g.x0, -v.Y+g.y0)
			continue
		}
		path.LineTo(v.X+g.x0, -v.Y+g.y0)
	}

	op := &vector.StrokeOptions{}
	op.Width = 1.5
	g.vertices, g.indices = path.AppendVerticesAndIndicesForStroke(g.vertices[:0], g.indices[:0], op)

	red, green, blue, _ := clr.RGBA()
	for i := range g.vertices {
		/*g.vertices[i].DstX += v.X
		g.vertices[i].DstY += v.Y//*/
		g.vertices[i].SrcX = 1
		g.vertices[i].SrcY = 1
		g.vertices[i].ColorR = float32(red / 256)
		g.vertices[i].ColorG = float32(green / 256)
		g.vertices[i].ColorB = float32(blue / 256)
		g.vertices[i].ColorA = 1
	}

	trop := &ebiten.DrawTrianglesOptions{}
	trop.AntiAlias = true
	trop.FillRule = ebiten.FillRuleNonZero

	scr.DrawTriangles(g.vertices, g.indices, g.tr, trop)
}

func (g *Graph) Label(scr *ebiten.Image, text string, size, x, y float32, clr clrs.Clr) {
	g.Font.Set(float64(size), clr)
	g.Font.DrawCenter(scr, text, x+g.x0, -y+g.y0)
}

// ### Shapes ###
func (g *Graph) Arrow(scr *ebiten.Image, x1, y1, x2, y2 float32, clr clrs.Clr) {
	g.Poly(scr, []*vec.VecF32{{x1, y1}, {x2, y2}}, clr)
	w, h := x2-x1, y2-y1
	l := float32(math.Hypot(float64(w), float64(h)))
	if l == 0 {return}
	s, c := h/l, w/l
	var al, aw float32 = 15, 4 // arrow tip length and width
	g.Poly(scr, []*vec.VecF32{
		{ x2 + aw*s - al*c, y2 - aw*c - al*s },
		{ x2, y2 },
		{ x2 - aw*s - al*c, y2 + aw*c - al*s },
	}, clr)
}

func (g *Graph) Coords(scr *ebiten.Image, clr clrs.Clr) {
	g.Arrow(scr, -200, 0, 200, 0, clr)
	g.Arrow(scr, 0, -200, 0, 200, clr)
}

// TODO get rid of
func Draw(scr *ebiten.Image, lls ...*lmnts.Lmnt) {
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

func TestDraw(scr *ebiten.Image) func(*lmnts.Lmnt) {
	i := 0
	return func(el *lmnts.Lmnt) {
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
} //*/
