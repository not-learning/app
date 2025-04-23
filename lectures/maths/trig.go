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
	a1, a2 float32
	clrs.Clr
}

// Надо?
var sub1 = `Как описать вращение математически?
Ну то есть, вот точка движется по кругу. Как превратить это в числа?
Вообще, над этим вопросом интересно подумать самостоятельно...
Но мы не будем.`

/*func (t *Trig) shape1(scr *ebiten.Image) {
	t.Arc(scr, 0, 0, 100, t.a1, t.a2, t.Clr)
}//*/

func (t *Trig) shape1(scr *ebiten.Image) {
	const r float32 = 180
	s, c := math.Sincos(float64(t.a2))
	x, y := r*float32(c), r*float32(s)

	t.Poly(scr, []*vec.VecF32{{-(r+20), 0}, {(r+20), 0}}, clrs.Green)
	t.Poly(scr, []*vec.VecF32{{0, -(r+20)}, {0, (r+20)}}, clrs.Green)
	t.Poly(scr, []*vec.VecF32{{(r+5), -3}, {(r+20), 0}, {(r+5), 3}}, clrs.Green)
	t.Poly(scr, []*vec.VecF32{{-3, (r+5)}, {0, (r+20)}, {3, (r+5)}}, clrs.Green)

	/*t.Font.Set(20, clrs.Green)
	t.Font.DrawCenter(scr, "►", (r+20), 0)//*/

	/*t.Poly(scr, []*vec.VecF32{{0, 0}, {x*1.5, y*1.5}}, clrs.White)
	t.Poly(scr, []*vec.VecF32{{0, 0}, {r*1.5, 0}}, clrs.White)//*/

	t.Poly(scr, []*vec.VecF32{{x, y}, {x, 0}}, clrs.Blue)
	t.Poly(scr, []*vec.VecF32{{x, y}, {0, y}}, clrs.Blue)

	t.Arc(scr, 0, 0, r, t.a1, t.a2, t.Clr)
	t.CirFull(scr, x, y, 3, clrs.White)
	t.CirFull(scr, x, 0, 3, clrs.White)
	t.CirFull(scr, 0, y, 3, clrs.White)
}

func (t *Trig) anim1() {
//	t.a1 += 0.01
	if t.a2 < 2.17*math.Pi {
		t.a2 += 0.02
	} else {
		t.a2 = 2.17*math.Pi
	}
//	if t.a1 >= 2*math.Pi { t.a1 = 0 }
	//if t.a2 >= 4*math.Pi { t.a2 = 2*math.Pi }
}

func (t *Trig) xact1() {
	if t.Lect.MouseR() {
		t.Clr = clrs.Green
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

	t.a1, t.a2 = 0, 0
	t.AddSubs(sub1, sub2)
	t.AddShapes(t.shape1)
	t.AddAnims(t.anim1)
	t.AddXacts(t.xact1)
	t.Clr = clrs.Green
	return t
}
