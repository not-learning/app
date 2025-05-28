package frame

import (
	"strconv"

	"github.com/not-learning/lmnts"
)

type blocks struct {
	x1, y1 float32
	top    *lmnts.Lmnt
	screen *lmnts.Lmnt
	subs   []*lmnts.Lmnt

// Playback controls
	pbc *lmnts.Lmnt
	prev  *lmnts.Lmnt
	pause *lmnts.Lmnt
	next  *lmnts.Lmnt
	pbcshow bool

// Numpad
	np     *lmnts.Lmnt
	npl    []*lmnts.Lmnt
	npshow bool
	nph    float32
}

func initBlocks(x1, y1, x2, y2 float32) *blocks {
	b := &blocks{}
	b.x1, b.y1 = x1, y1
	b.nph = 350

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
	sl.AddLR(20, 20, si)
	b.subs = append(b.subs, lmnts.New(), lmnts.New())
	si.Add(b.subs...)
	b.top.Add(sl)

	// Playback controls
	b.pbc = lmnts.New()
	b.pbc.Name = "pbc"
	b.pbc.SetRow()
	b.pbc.SetSize(0, 100)
	b.prev, b.pause, b.next = lmnts.New(), lmnts.New(), lmnts.New()
	b.pause.Name = "pause"
	b.prev.SetSize(125, 75) // TODO check both options
	b.pause.SetSize(75, 75)
	b.next.SetSize(125, 75)
	b.pbc.GapsAround(0, b.prev, b.pause, b.next)
	//b.top.Add(b.pbc)

	// NumPad
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
	in.Grid(3, 20, b.npl...)
	b.np.AddTBLR(1, 30, 30, 30, in)

	b.top.DoAll()

	return b
}

//var oldW, oldH int
func (b *blocks) update(scrW, scrH int) {
	//if scrW == oldW || scrH == oldH { return } // DEV
	b.top.SetRect(b.x1, b.y1, float32(scrW), float32(scrH))
	b.top.DoAll()
	//oldW, oldH = scrW, scrH
}

// TODO show / hide gets called quite often?
func (b *blocks) playConShow() {
	if !b.pbcshow {
		b.pbcshow = true
		b.top.AddN(-1, b.pbc)
		b.top.DoAll()
	}
}

func (b *blocks) playConHide() {
	if b.pbcshow {
		b.pbcshow = false
		b.top.Del(b.pbc)
		b.top.DoAll()
	}
}

func (b *blocks) numPadShow() {
	if !b.npshow {
		b.npshow = true
		n := -1
		if b.top.IsAdded(b.pbc) { n = -2 }
		b.top.AddN(n, b.np)
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
