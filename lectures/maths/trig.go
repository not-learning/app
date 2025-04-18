package maths

import (
	"github.com/not-learning/app/frame"
	"github.com/not-learning/app/clrs"
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

type Trig struct {
	*frame.Lect
	a1, a2 float32
}

var sub1 = `Как описать вращение и повороты математически?
Ну то есть, вот точка движется по кругу. Как превратить это в числа?
Вообще, над этим вопросом интересно подумать самостоятельно...
Но мы не будем.`

var sub2 = `Возьмем координатную плоскость.`

func (t *Trig) shape1(scr *ebiten.Image) {
	t.Arc(scr, 0, 0, 100, t.a1, t.a2, clrs.Green)
	//t.Arc(scr, 0, 0, 100, 1, 2, clrs.Green)
}

func (t *Trig) anim1() {
	t.a1 += 0.01
	t.a2 += 0.1
	if t.a1 >= 2*math.Pi { t.a1 = 0 }
	if t.a2 >= 2*math.Pi { t.a2 = 0 }
}

func InitTrig(x1, y1, x2, y2 float32) *Trig {
	t := &Trig{}
	t.Lect = frame.Init(x1, y1, x2, y2)
	t.a1, t.a2 = 0, 0
	t.AddShape(t.shape1)
	t.AddAnim(t.anim1)
	//t.AddAnim(func(){})
	return t
}
