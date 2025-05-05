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
	prev   *lmnts.Lmnt
	pause  *lmnts.Lmnt
	next   *lmnts.Lmnt

	npl    []*lmnts.Lmnt
	np     *lmnts.Lmnt
	npshow bool
	nph    float32
}

func initBlocks(x1, y1, x2, y2 float32) *blocks {
	b := &blocks{}
	b.nph = 350
	//b.monh = y2 - y1

	b.top = lmnts.New()
	b.top.Name = "top"
	b.top.SetRect(x1, y1, x2, y2)
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
	pl.SetRow()
	pl.SetSize(0, 100)
	b.prev, b.pause, b.next = lmnts.New(), lmnts.New(), lmnts.New()
	b.pause.Name = "pause"
	b.prev.SetSize(125, 75)
	b.pause.SetSize(75, 75)
	b.next.SetSize(125, 75)
	pl.GapsAround(30, b.prev, b.pause, b.next)
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

	b.top.DoAll()

	return b
}

func (b *blocks) numPadShow() {
	if !b.npshow {
		b.npshow = true
		b.top.Add(b.np)
		b.top.DoAll()
	}
}

func (b *blocks) numPadHide() {
	if b.npshow {
		b.npshow = false
		b.top.Del(b.np)
		b.top.DoAll()
	}
}

/*func (b *blocks) numPadShow() {
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
}//*/
