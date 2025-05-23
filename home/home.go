package home

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/not-learning/app/clrs"
	"github.com/not-learning/app/lectures/lobby"
	"github.com/not-learning/app/lectures/maths"
)

type Home struct {
	lobby *lobby.Lobby
	pow   *maths.Pow
	trig  *maths.Trig

	chapter int
}

func Init(x1, y1, x2, y2 float32) *Home {
	h := &Home{}
	h.lobby = lobby.Init(x1, y1, x2, y2)
	h.lobby.ChapterFn(h.Chapter)

	h.pow = maths.InitPow(x1, y1, x2, y2)
	h.pow.ChapterFn(h.Chapter)

	h.trig = maths.InitTrig(x1, y1, x2, y2)
	h.trig.ChapterFn(h.Chapter)

	return h
}

func (h *Home) Chapter(n int) { h.chapter = n }

func (h *Home) Update(scrW, scrH int, ratW, ratH float32) {
	switch h.chapter {
		case 1:
			h.pow.Update(scrW, scrH, ratW, ratH)
		case 2:
			h.trig.Update(scrW, scrH, ratW, ratH)
		default:
			h.lobby.Update(scrW, scrH, ratW, ratH)
	}
}

func (h *Home) Draw(scr *ebiten.Image) {
	scr.Fill(clrs.YCC(20, 128, 128)) // TODO use clrs
	switch h.chapter {
		case 1:
			h.pow.Draw(scr)
		case 2:
			h.trig.Draw(scr)
		default:
			h.lobby.Draw(scr)
	}
}
