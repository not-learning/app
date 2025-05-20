package lobby

import (
	"github.com/not-learning/app/clrs"
	"github.com/not-learning/app/frame"
	"github.com/not-learning/app/inter"
	"github.com/not-learning/app/lectures/maths"

	"github.com/hajimehoshi/ebiten/v2"
)

type Lobby struct {
	*frame.Lect

	pow  *maths.Pow
	trig *maths.Trig

	inPow, inTrig bool
	int
}

func Init(x1, y1, x2, y2 float32) *Lobby {
	l := &Lobby{}
	l.Lect = frame.Init(x1, y1, x2, y2)
	l.AddEx(l.sub0, l.shape1, l.anim1, l.xact1, l.zero1)

	l.pow = maths.InitPow(x1, y1, x2, y2)
	l.trig = maths.InitTrig(x1, y1, x2, y2)

	return l
}

func (l *Lobby) sub0(*ebiten.Image) {}

func (l *Lobby) shape1(scr *ebiten.Image) {
	l.inPow = inter.MouseLIn(l.Label(scr, "Степени", 30, 0, 50, clrs.Blue))
	l.inTrig = inter.MouseLIn(l.Label(scr, "Тригонометрия", 30, 0, 0, clrs.Blue))
	//x, y := inter.MousePos()
	//l.Label(scr, strconv.Itoa(x) +" "+ strconv.Itoa(y), 30, 0, 200, clrs.Blue)
}

func (l *Lobby) anim1() bool { return true }

func (l *Lobby) xact1() bool {
	inter.Escape()
	if i, ok := inter.Number(); ok { l.int = i }
	if l.inPow { l.int = 1 }
	if l.inTrig { l.int = 2 }
	return true
}

func (l *Lobby) zero1() {}

func (l *Lobby) Update(scrW, scrH int, ratW, ratH float32) {
	if l.int == 0 { l.Lect.Update(scrW, scrH, ratW, ratH) }
	if l.int == 1 { l.pow.Update(scrW, scrH, ratW, ratH) }
	if l.int == 2 { l.trig.Update(scrW, scrH, ratW, ratH) }
} //*/

func (l *Lobby) Draw(scr *ebiten.Image) {
	scr.Fill(clrs.YCC(20, 128, 128)) // TODO use clrs
	if l.int == 0 { l.Lect.Draw(scr) }
	if l.int == 1 { l.pow.Draw(scr) }
	if l.int == 2 { l.trig.Draw(scr) }
}
