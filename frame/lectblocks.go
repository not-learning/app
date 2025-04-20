package frame

import "github.com/not-learning/lmnts"

type blocks struct {
	screen *lmnts.Lmnt
	subs   []*lmnts.Lmnt
	pause  *lmnts.Lmnt
}

func InitBlocks(x1, y1, x2, y2 float32) *blocks {
	b := &blocks{}
	top := lmnts.New()
	top.SetRect(x1, y1, x2, y2)

	b.screen = lmnts.New()
	top.Add(b.screen)

	sl := lmnts.New()
	sl.SetSize(0, 40)
	si := lmnts.New()
	sl.AddL(20, si)
	b.subs = append(b.subs, lmnts.New(), lmnts.New())
	si.Add(b.subs...)
	top.Add(sl)

	pl := lmnts.New()
	pl.SetSize(0, 75)
	b.pause = lmnts.New()
	b.pause.SetSize(75, 75)
	pl.AddLR(0, 0, b.pause)
	top.Add(pl)

	top.DoAll()

	return b
}
