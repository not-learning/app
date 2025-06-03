package lobby

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/not-learning/app/nlapp/clrs"
	"github.com/not-learning/app/nlapp/frame"
)

type Lobby struct {
	*frame.Frame
	chapter int
}

func Init(x1, y1, x2, y2, scale float32) *Lobby {
	l := &Lobby{}
	l.Frame = frame.Init(x1, y1, x2, y2, scale)
	l.AddEx(l.sub0, l.shape1, l.anim1, l.xact1, l.zero1)
	return l
}

func (l *Lobby) sub0(*ebiten.Image) {}

func (l *Lobby) shape1(scr *ebiten.Image) {
	switch {
		case l.TapIn(l.LabelRect(scr, "Алгебра", 30, 0, 60, clrs.Blue)):
			l.SetChapter(1)
		case l.TapIn(l.LabelRect(scr, "Степени", 30, 0, 30, clrs.Blue)):
			l.SetChapter(2)
		case l.TapIn(l.LabelRect(scr, "Тригонометрия", 30, 0, 0, clrs.Blue)):
			l.SetChapter(3)
	}
}

func (l *Lobby) anim1() bool { return true }

func (l *Lobby) xact1() bool {
	if i, ok := l.Number(); ok {
		l.SetChapter(i)
	}
	return true
}

func (l *Lobby) zero1() {}
