package maths

import (
	"math"

	"github.com/not-learning/app/clrs"
	"github.com/not-learning/app/frame"

	"github.com/hajimehoshi/ebiten/v2"
)

type Trig struct {
	*frame.Lect
	a1, a2 float32
	clrs.Clr
	/*r, g, b, a uint8
	n uint32//*/
}

var sub1 = `Как описать вращение и повороты математически?
Ну то есть, вот точка движется по кругу. Как превратить это в числа?
Вообще, над этим вопросом интересно подумать самостоятельно...
Но мы не будем.`

var sub2 = `Возьмем координатную плоскость.`

func (t *Trig) shape1(scr *ebiten.Image) {
	t.Arc(scr, 0, 0, 100, t.a1, t.a2, t.Clr)
}

func (t *Trig) anim1() {
	t.a1 += 0.01
	t.a2 += 0.02
	if t.a1 >= 2*math.Pi { t.a1 = 0 }
	if t.a2 >= 2*math.Pi { t.a2 = 0 }

	/*t.n++
	t.r = uint8(t.n>>2)
	t.g = uint8(t.n>>1)
	t.b = uint8(t.n)
	t.Clr = clrs.RGB(t.r, t.g, t.b)//*/
}

/*func (t *Trig) xact1() {
	if t.Touch() {
		t.Clr = clrs.White
	} else {
		t.Clr = clrs.Blue
	}
}//*/

func (t *Trig) xact2() {
	if t.Lect.Space() || t.Lect.MouseL() {
		t.Clr = clrs.White
	} else {
		t.Clr = clrs.Blue
}
}

func (t *Trig) xact1() {
	if t.Lect.MouseR() {
		t.Clr = clrs.Green
	} else if t.Lect.MouseL() {
		t.Clr = clrs.White
	} else {
		t.Clr = clrs.Blue
	}
}

func InitTrig(x1, y1, x2, y2 float32) *Trig {
	t := &Trig{}
	t.Lect = frame.Init(x1, y1, x2, y2)
	t.a1, t.a2 = 0, 0
	t.AddShape(t.shape1)
	t.AddAnim(t.anim1)
	t.AddXact(t.xact1)
	//t.AddXact(func(){})
	t.Clr = clrs.Blue
	//t.n = uint32(1<<7)
	return t
}
