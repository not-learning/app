package frame

import (
	"strconv"

	"github.com/not-learning/lmnts"
)

type blocks struct {
	top    *lmnts.Lmnt
	monh   float32
	screen *lmnts.Lmnt
	subs   []*lmnts.Lmnt
	pause  *lmnts.Lmnt

	npl []*lmnts.Lmnt
	np  *lmnts.Lmnt
	nph float32
}

func initBlocks(x1, y1, x2, y2 float32) *blocks {
	b := &blocks{}
	b.nph = 350
	b.monh = y2 - y1

	b.top = lmnts.New()
	b.top.Name = "top"
	b.top.SetRect(x1, y1, x2, y2+b.nph)
	b.screen = lmnts.New()
	b.screen.Name = "screen"
	b.top.Add(b.screen)

	sl := lmnts.New()
	sl.Name = "sl"
	sl.SetSize(0, 40)
	si := lmnts.New()
	si.Name = "si"
	sl.AddL(20, si)
	b.subs = append(b.subs, lmnts.New(), lmnts.New())
	si.Add(b.subs...)
	b.top.Add(sl)

	pl := lmnts.New()
	pl.Name = "pl"
	pl.SetSize(0, 100)
	b.pause = lmnts.New()
	b.pause.Name = "pause"
	b.pause.SetSize(75, 75)
	pl.AddLR(0, 0, b.pause)
	b.top.Add(pl)


	// ## NumPad init
	b.np = lmnts.New()
	b.np.Name = "np"
	b.np.SetSize(0, b.nph)
	nn := 12
	b.npl = make([]*lmnts.Lmnt, nn)
	for i := range nn {
		b.npl[i] = lmnts.New()
		b.npl[i].Name = "np" + strconv.Itoa(i)
		//b.npl[i].SetSize(100, 75)
	}
	in := lmnts.New()
	in.Name = "in"
	b.np.AddTBLR(30, 30, 30, 30, in)
	in.Grid(3, 20, b.npl...)
	b.top.Add(b.np)

	b.top.DoAll()

	return b
}

func (b *blocks) numPadShow() {
	x1, y1, x2, y2 := b.top.Rect()
	if y2-y1 > b.monh+b.nph/2 {
		b.top.SetRect(x1, y1, x2, y2-b.nph)
		b.top.DoAll()
	}
}

func (b *blocks) numPadHide() {
	x1, y1, x2, y2 := b.top.Rect()
	if y2-y1 < b.monh+b.nph/2 {
		b.top.SetRect(x1, y1, x2, y2+b.nph)
		b.top.DoAll()
	}
}
