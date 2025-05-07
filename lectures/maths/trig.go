package maths

import (
	"embed"
	"math"

	"github.com/not-learning/app/clrs"
	"github.com/not-learning/app/frame"
	//"github.com/not-learning/app/inter"

	"github.com/hajimehoshi/ebiten/v2"
)

// TODO: prepare for Next, Prev
type Trig struct {
	*frame.Lect
	clrs.Clr

	r, a1, a2, x, y, prevX, prevY float32
	polygon []float32
	n int
	c, s float32

	animP func(float32) []float32
}

func (t *Trig) xy() {
	s, c := math.Sincos(float64(t.a1))
	t.x, t.y = t.r*float32(c), t.r*float32(s)
}

// ## Ex1
func (t *Trig) sub1() func(*ebiten.Image) {
	return t.SubDraw(t.SubWrap(
		`Представь робота, который может вырезать любую фигуру по ее координатам.`,
	))
}

func (t *Trig) shape1(scr *ebiten.Image) {
	t.CirFull(scr, t.x, t.y, 4, clrs.White)
	t.PolyEmp(scr, t.polygon, clrs.Green)
	t.Robot(scr, t.x, t.y, t.r)
}//*/

func (t *Trig) anim1() bool {
	t.polygon = t.animP(7)
	l := len(t.polygon)
	t.x, t.y = t.polygon[l-2], t.polygon[l-1]

	done := (t.prevX == t.x && t.prevY == t.y)
	t.prevX, t.prevY = t.x, t.y
	return done
}

func (t *Trig) xact1() bool { return true }
func (t *Trig) zero1() {
	t.polygon = t.animP(-1)
	t.x, t.y = t.polygon[0], t.polygon[1]
}


// ## Ex2
func (t *Trig) sub2() func(*ebiten.Image) {
	return t.SubDraw(t.SubWrap(
		`Мы хотим вырезать круг. Как задать координаты?`,
	))
}

func (t *Trig) shape2(scr *ebiten.Image) {
	t.Robot(scr, t.x, t.y, t.r+20)
	t.CirFull(scr, t.x, t.y, 4, clrs.White)
}

func (t *Trig) anim2() bool {
	t.a1 += 0.1
	t.xy()
	if t.a1 > 2*2*math.Pi { return true }
	return false
}

func (t *Trig) xact2() bool { return true }
func (t *Trig) zero2() {
	t.a1 = 0
	t.x, t.y = 0, 0
}

// ## Ex3
func (t *Trig) sub3() func(*ebiten.Image) {
	return t.SubDraw(t.SubWrap(
		`Возьмем координатную плоскость.`,
	))
}

func (t *Trig) shape3(scr *ebiten.Image) {
	t.Coords(scr)
}

func (t *Trig) anim3() bool { return true }
func (t *Trig) xact3() bool { return true }
func (t *Trig) zero3() {}

// ## Ex4
func (t *Trig) sub4() func(*ebiten.Image) {
	return t.SubDraw(t.SubWrap(
		`И поместим на нее круг.`,
	))
}

func (t *Trig) shape4(scr *ebiten.Image) {
	t.Coords(scr)
	t.Arc(scr, 0, 0, t.r, 0, t.a1, t.Clr)
}

func (t *Trig) anim4() bool {
	t.a1 += 0.2
	if t.a1 >= 2*2*math.Pi {
		t.a1 = 2*2*math.Pi
		return true
	}
	t.xy()
	return false
}

func (t *Trig) xact4() bool { return true }
func (t *Trig) zero4() { t.a1 = 0 }


// TODO proper tracks
//go:embed tracks/pow
var files embed.FS

func InitTrig(x1, y1, x2, y2 float32) *Trig {
	t := &Trig{}
	t.Lect = frame.Init(x1, y1, x2, y2)
	t.Lect.Tracks.InitFiles("tracks/pow", files)

	t.r = 180

	var a float32 = math.Pi/5
	five1 := frame.Polygon(5, 0, 0, t.r, 0.1)
	five2 := frame.Polygon(5, 0, 0, t.r/2.6, a+0.1)
	star := make([]float32, 24)
	for i := 0; i < len(five1); i += 2 {
		star[2*i], star[2*i+1] = five1[i], five1[i+1]
		star[2*i+2], star[2*i+3] = five2[i], five2[i+1]
	}

	t.animP = frame.AnimPoly(star[:len(star)-2])

	t.AddEx(t.sub1(), t.shape1, t.anim1, t.xact1, t.zero1)
	t.AddEx(t.sub2(), t.shape2, t.anim2, t.xact2, t.zero2)
	t.AddEx(t.sub3(), t.shape3, t.anim3, t.xact3, t.zero3)
	t.AddEx(t.sub4(), t.shape4, t.anim4, t.xact4, t.zero4)

	t.Clr = clrs.Green
	return t
}

/*func (t *Trig) shape1(scr *ebiten.Image) {
	t.Coords(scr, clrs.Green)

	t.Arc(scr, 0, 0, t.r, 0, t.a1, t.Clr)
	//t.Arrow(scr, 0, 0, 1.2*t.x, 1.2*t.y, clrs.White)
	t.PolyEmp(scr, []float32{0, 0, t.x, t.y}, clrs.White)

	//t.PolyEmp(scr, []float32{0, 0, x*1.5, y*1.5}, clrs.White)
	//t.PolyEmp(scr, []float32{0, 0, r*1.5, 0}, clrs.White)

	t.PolyEmp(scr, []float32{t.x, t.y, t.x, 0}, clrs.Blue)
	t.PolyEmp(scr, []float32{t.x, t.y, 0, t.y}, clrs.Blue)

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
