package lobby

import (
	"github.com/not-learning/app/clrs"
	"github.com/not-learning/app/frame"
	"github.com/not-learning/app/lectures/maths"

	"github.com/hajimehoshi/ebiten/v2"
)

type Lobby struct {
	*frame.Lect

	pow  *maths.Pow
	trig *maths.Trig

	int
}

func (l *Lobby) sub0(*ebiten.Image) {}

func (l *Lobby) shape1(scr *ebiten.Image) {
	l.Label(scr, "Степени", 30, 0, 300, clrs.Blue)
	l.Label(scr, "Тригонометрия", 30, 0, 250, clrs.Blue)
}

func (l *Lobby) anim1() bool { return true }

func (l *Lobby) xact1() bool {
	frame.Escape()
	if i, ok := frame.Number(); ok { l.int = i }
	return true
}

func (l *Lobby) Update() {
	if l.int == 0 { l.Lect.Update() }
	if l.int == 1 { l.pow.Update() }
	if l.int == 2 { l.trig.Update() }
} //*/

func (l *Lobby) Draw(scr *ebiten.Image) {
	s := scr
	s.Fill(clrs.YCC(20, 128, 128)) // TODO use clrs
	if l.int == 0 { l.Lect.Draw(scr) }
	if l.int == 1 { l.pow.Draw(scr) }
	if l.int == 2 { l.trig.Draw(scr) }
}

func Init(x1, y1, x2, y2 float32) *Lobby {
	l := &Lobby{}
	l.Lect = frame.Init(x1, y1, x2, y2)
	l.AddEx(l.sub0, l.shape1, l.anim1, l.xact1)

	l.pow = maths.InitPow(x1, y1, x2, y2)
	l.trig = maths.InitTrig(x1, y1, x2, y2)

	return l
}
