package maths

import (//"fmt"
	"embed"
	"math"

	"github.com/not-learning/app/clrs"
	"github.com/not-learning/app/frame"

	"github.com/hajimehoshi/ebiten/v2"
)

type Trig struct {
	*frame.Lect
	clrs.Clr
	animDone bool

	r, a1, a2, x, y, prevX, prevY float32
	polygon []float32
	n int
	c, s float32

	animP func(float32) []float32
}

func (t *Trig) xy() {
	s, c := math.Sincos(float64(t.a2))
	t.x, t.y = t.r*float32(c), t.r*float32(s)
}

var sub1 = `Возьмем робота, который может вырезать любую фигуру по координатам.`

func (t *Trig) shape1(scr *ebiten.Image) {
	t.PolyEmp(scr, t.polygon, clrs.Green)

	if t.animDone {return}

	t.Coords(scr, clrs.Blue)
	t.CirFull(scr, t.x, 0, 4, clrs.Blue)
	t.CirFull(scr, 0, t.y, 4, clrs.Blue)
	t.CirFull(scr, t.x, t.y, 4, clrs.White)
	t.PolyEmp(scr, []float32{t.x, 0, t.x, t.y}, clrs.Blue)
	t.PolyEmp(scr, []float32{0, t.y, t.x, t.y}, clrs.Blue)
}//*/

func (t *Trig) anim1() {
	t.polygon = t.animP(5)
	l := len(t.polygon)
	t.x, t.y = t.polygon[l-2], t.polygon[l-1]

	t.animDone = (t.prevX == t.x && t.prevY == t.y)
	t.prevX, t.prevY = t.x, t.y
}

func (t *Trig) xact1() {
	if t.Space() { t.Pause() }

	if t.Lect.MouseR() {
		t.Clr = clrs.Blue
	} else if t.Lect.MouseL() {
		t.Clr = clrs.White
		t.Play()
	} else {
		t.Clr = clrs.Green
	}
}

var sub2 = `Возьмем координатную плоскость.`

// TODO proper tracks
//go:embed tracks/pow
var files embed.FS

func InitTrig(x1, y1, x2, y2 float32) *Trig {
	t := &Trig{}
	t.Lect = frame.Init(x1, y1, x2, y2)
	t.Lect.Tracks.InitFiles("tracks/pow", files)

	t.animDone = true
	t.a1 = float32(math.Pi/2)

	t.r = 180
	t.n = 100
	t.animP = frame.AnimPoly(frame.Polygon(t.n, 0, 0, t.r, 0))

	t.AddSubs(sub1, sub2)
	t.AddAnims(t.anim1)
	t.AddShapes(t.shape1)
	t.AddXacts(t.xact1)
	t.Clr = clrs.Green
	return t
}

/*func (t *Trig) shape1(scr *ebiten.Image) {
	t.Coords(scr, clrs.Green)

	//t.Arrow(scr, 0, 0, 1.2*t.x, 1.2*t.y, clrs.White)
	t.PolyEmp(scr, []float32{0, 0, t.x, t.y}, clrs.White)

	//t.PolyEmp(scr, []float32{0, 0, x*1.5, y*1.5}, clrs.White)
	//t.PolyEmp(scr, []float32{0, 0, r*1.5, 0}, clrs.White)

	t.PolyEmp(scr, []float32{t.x, t.y, t.x, 0}, clrs.Blue)
	t.PolyEmp(scr, []float32{t.x, t.y, 0, t.y}, clrs.Blue)

	t.Arc(scr, 0, 0, t.r, t.a1, t.a2, t.Clr)
	t.CirFull(scr, t.x, t.y, 4, clrs.White)
	t.CirFull(scr, t.x,   0, 4, clrs.White)
	t.CirFull(scr, 0,   t.y, 4, clrs.White)

	t.Label(scr, "x", 15, t.x, -10, clrs.White)
	t.Label(scr, "y", 15, 10, t.y, clrs.White)
}//*/

/*func (t *Trig) anim1() {
		t.a2 += 0.02
		t.xy()
	if t.a2 >= 4*math.Pi { t.a2 = 2*math.Pi }
	//if t.a2 >= 2.33*math.Pi { t.a2 = 2.33*math.Pi } else { t.a2 += 0.02 }
}//*/
