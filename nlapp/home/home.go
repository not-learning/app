package home

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/not-learning/app/nlapp/clrs"
	"github.com/not-learning/app/nlapp/lectures/lobby"
	"github.com/not-learning/app/nlapp/lectures/maths"
)

type Home struct {
	lobby *lobby.Lobby
	pow   *maths.Pow
	trig  *maths.Trig

	chapter int
}

func Init(x1, y1, x2, y2, scale float32) *Home {
	h := &Home{}
	h.lobby = lobby.Init(x1, y1, x2, y2, scale)
	h.lobby.ChapterFn(h.Chapter)

	h.pow = maths.InitPow(x1, y1, x2, y2, scale)
	h.pow.ChapterFn(h.Chapter)

	h.trig = maths.InitTrig(x1, y1, x2, y2, scale)
	h.trig.ChapterFn(h.Chapter)

	return h
}

func (h *Home) Chapter(n int) { h.chapter = n }

func (h *Home) Update(scrW, scrH int, scale float32) {
	switch h.chapter {
		case 1:
			h.pow.Update(scrW, scrH, scale)
		case 2:
			h.trig.Update(scrW, scrH, scale)
		default:
			h.lobby.Update(scrW, scrH, scale)
	}
}

func (h *Home) Draw(scr *ebiten.Image) {
	scr.Fill(clrs.YCC(20, 128, 128))
	switch h.chapter {
		case 1:
			h.pow.Draw(scr)
		case 2:
			h.trig.Draw(scr)
		default:
			h.lobby.Draw(scr)
	}
}
