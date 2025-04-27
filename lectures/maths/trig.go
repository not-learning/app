package maths

import (
	"embed"
	"math"

	"github.com/not-learning/app/clrs"
	"github.com/not-learning/app/frame"
	"github.com/not-learning/app/vec"

	"github.com/hajimehoshi/ebiten/v2"
)

type Trig struct {
	*frame.Lect
	clrs.Clr

	r, a1, a2, x, y float32
}

// Надо?
var sub1 = `Как описать вращение математически?
Ну то есть, вот точка движется по кругу. Как превратить это в числа?
Вообще, над этим вопросом интересно подумать самостоятельно...
Но мы не будем.`

/*func (t *Trig) shape1(scr *ebiten.Image) {
	t.Arc(scr, 0, 0, 100, t.a1, t.a2, t.Clr)
}//*/

func (t *Trig) xy() (x, y float32) {
	s, c := math.Sincos(float64(t.a2))
	x, y = t.r*float32(c), t.r*float32(s)
	return
}

func (t *Trig) shape1(scr *ebiten.Image) {
	t.Coords(scr, clrs.Green)

	t.Arrow(scr, 0, 0, 1.2*t.x, 1.2*t.y, clrs.White)

	/*t.Font.Set(20, clrs.Green)
	t.Font.DrawCenter(scr, "►", (r+20), 0)//*/

	/*t.Poly(scr, []*vec.VecF32{{0, 0}, {x*1.5, y*1.5}}, clrs.White)
	t.Poly(scr, []*vec.VecF32{{0, 0}, {r*1.5, 0}}, clrs.White)//*/

	t.Poly(scr, []*vec.VecF32{{t.x, t.y}, {t.x, 0}}, clrs.Blue)
	t.Poly(scr, []*vec.VecF32{{t.x, t.y}, {0, t.y}}, clrs.Blue)

	t.Arc(scr, 0, 0, t.r, t.a1, t.a2, t.Clr)
	t.CirFull(scr, t.x, t.y, 4, clrs.White)
	t.CirFull(scr, t.x, 0, 4, clrs.White)
	t.CirFull(scr, 0, t.y, 4, clrs.White)
	t.Label(scr, "x", 15, t.x+10, -10, clrs.White)
	t.Label(scr, "y", 15, -10, t.y+10, clrs.White)

	/*t.CirFull(scr, x, y, 4, clrs.Black)
	t.CirFull(scr, x, 0, 4, clrs.Black)
	t.CirFull(scr, 0, y, 4, clrs.Black)
	t.CirEmp(scr,  x, y, 4, clrs.White)
	t.CirEmp(scr,  x, 0, 4, clrs.White)
	t.CirEmp(scr,  0, y, 4, clrs.White)//*/
}

func (t *Trig) anim1() {
		t.a2 += 0.02
		t.x, t.y = t.xy()
	if t.a2 == 2*math.Pi { t.a2 = 0 }
}

func (t *Trig) xact1() {
	if t.Lect.MouseR() {
		t.Clr = clrs.Blue
	} else if t.Lect.MouseL() {
		t.Clr = clrs.White
		t.Play()
	} else {
		t.Clr = clrs.Green
	}
	if t.Space() { t.Pause() }
}

var sub2 = `Возьмем координатную плоскость.`

// TODO proper tracks
//go:embed tracks/pow
var files embed.FS

func InitTrig(x1, y1, x2, y2 float32) *Trig {
	t := &Trig{}
	t.Lect = frame.Init(x1, y1, x2, y2)
	t.Lect.Tracks.InitFiles("tracks/pow", files)
	// t.Lect.Tracks.Play()

	t.r, t.a1, t.a2 = 180, 0, 0
	t.x, t.y = t.xy()
	t.AddSubs(sub1, sub2)
	t.AddShapes(t.shape1)

	/*lx, ly := t.NewLabel("x"), t.NewLabel("y")

	lx.SetPos(t.x, 10)
	ly.SetPos(10, t.y)
	t.AddLabels(lx, ly)//*/
	t.AddAnims(t.anim1)
	t.AddXacts(t.xact1)
	t.Clr = clrs.Green
	return t
}
