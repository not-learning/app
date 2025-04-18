package maths

import (
	"github.com/not-learning/app/frame"
	"github.com/not-learning/app/clrs"
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

type Trig struct {
	*frame.Lect
}

var sub1 = `Как описать вращение и повороты математически?
Ну то есть, вот точка движется по кругу. Как превратить это в числа?
Вообще, над этим вопросом интересно подумать самостоятельно...
Но мы не будем.`

var sub2 = `Возьмем координатную плоскость.`

func (t *Trig) shape1(scr *ebiten.Image) {
	t.Arc(scr, 225, 225, 100, 0, 2*math.Pi, clrs.Green)
}

//func InitTrig(x1, y1, x2, y2 float32) *Trig {
func InitTrig() *Trig {
	t := &Trig{}
	t.Lect = frame.Init()
	t.AddShape(t.shape1)
	return t
}
